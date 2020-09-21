package utils

import (
	"context"
	"errors"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"operationPlatform/types"
	"os"
	"os/signal"

	"strconv"
	"strings"
	"syscall"
	"time"

	"sync"
)

const (
	kafkaConn1 = "localhost:9092" //本机
	kafkaConn2 = "127.0.0.1:9093"
	kafkaConn3 = "127.0.0.1:9094"
	topic      = "topic1"
)

//代理
var brokerAddrs = []string{kafkaConn1, kafkaConn2, kafkaConn3}

//生产数据
func Producer(msgType string, value string) {
	config := sarama.NewConfig()

	// 等待服务器 所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll

	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	//producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	if err != nil {
		panic(err)
	}

	defer producer.Close()

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		//Topic: "test",//包含了消息的主题
		Partition: int32(10),                   //
		Key:       sarama.StringEncoder("key"), //
	}

	log.Println("msgType = ", msgType, ",value = ", value)
	msg.Topic = msgType
	//将字符串转换为字节数组
	msg.Value = sarama.ByteEncoder(value)

	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		log.Println("Send message Fail")
	}
	log.Printf("Partition = %d, offset=%d\n", partition, offset)
}

var (
	wg sync.WaitGroup
)

//消费者
func Consumer() {
	// 根据给定的代理地址和配置创建一个消费者
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	//conf := &sarama.Config{}
	ConsumerGroup, grouperr := sarama.NewConsumerGroup([]string{"localhost:9092"}, "12324", nil)
	if grouperr != nil {
		panic(grouperr)
	}
	log.Println(ConsumerGroup)
	//ConsumerGroup.Consume()
	//Partitions(topic):该方法返回了该topic的所有分区id
	partitionList, err := consumer.Partitions(topic)
	if err != nil {
		panic(err)
	}

	for partition := range partitionList {
		//ConsumePartition方法根据主题，分区和给定的偏移量创建创建了相应的分区消费者
		//如果该分区消费者已经消费了该信息将会返回error
		//sarama.OffsetNewest:表明了为最新消息
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()
		wg.Add(1)
		go func(sarama.PartitionConsumer) {
			defer wg.Done()
			//Messages()该方法返回一个消费消息类型的只读通道，由代理产生
			for msg := range pc.Messages() {
				log.Printf("%s---Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
			}
		}(pc)
	}
	wg.Wait()
	consumer.Close()
}

type Kafka struct {
	brokers []string
	topics  []string
	//OffsetNewest int64 = -1
	//OffsetOldest int64 = -2
	startOffset       int64
	version           string
	ready             chan bool
	group             string
	channelBufferSize int
}

func NewKafka() *Kafka {
	return &Kafka{
		brokers:           brokers,
		topics:            topics,
		group:             group,
		channelBufferSize: 2,
		ready:             make(chan bool),
		version:           "1.1.1",
	}
}

var brokers = []string{"172.18.70.21:9092"}
var topics = []string{types.DdkafkaTopic, types.ZdzkafkaTopic}
var group = "39"

func (p *Kafka) Init() func() {
	log.Println("++++++++++++++++++++++++++++++++++kafka init...")

	version, err := sarama.ParseKafkaVersion(p.version)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++Error parsing Kafka version:  ", err)
	}
	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange // 分区分配策略
	config.Consumer.Offsets.Initial = -2                                   // 未找到组消费位移的时候从哪边开始消费
	config.ChannelBufferSize = p.channelBufferSize                         // channel长度

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(p.brokers, p.group, config)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++Error creating consumer group client: ", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			//util.HandlePanic("client.Consume panic", log.StandardLogger())
		}()
		for {
			if err := client.Consume(ctx, p.topics, p); err != nil {
				log.Println("++++++++++++++++++++++++++++++++Error from consumer: ", err)
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				log.Println("+++++++++++++++++++++++++ctx.Err():", ctx.Err())
				return
			}
			p.ready = make(chan bool)
		}
	}()
	<-p.ready
	log.Infoln("+++++++++++++++++++++++++++++Sarama consumer up and running!...")
	// 保证在系统退出时，通道里面的消息被消费
	return func() {
		log.Println("+++++++++++++++++++++++++++++kafka close")
		cancel()
		wg.Wait()
		if err = client.Close(); err != nil {
			log.Println("++++++++++++++++++++++++++++++++++Error closing client: ", err)
		}
	}
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (p *Kafka) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(p.ready)
	return nil
}

//清理
// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (p *Kafka) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

//消费主张
// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
func (p *Kafka) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	// 具体消费消息
	for message := range claim.Messages() {
		msg := string(message.Value)
		log.Println("+++++++++++++++++++++++++++++++msg:", msg)
		time.Sleep(time.Second)
		//run.Run(msg)
		// 更新位移
		session.MarkMessage(message, "")
	}
	return nil
}

func UseKafka() {
	k := NewKafka()
	f := k.Init()

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Warnln("terminating: via signal")
	}
	f()
}

type consumerGroupHandler struct {
	name string //groupname
}

func (consumerGroupHandler) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (consumerGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

//消费主张
func (h consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		log.Printf("++++++++++++++%s group Message topic:%q partition:%d offset:%d  value:%s\n", h.name, msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		//消息处理
		log.Println("消息已接受，正处理中+++++++++++++++++++++++++++++++消息已接受，正处理中+++++++++++")
		err := ProcessMessage(msg.Topic, msg.Value)
		if err != nil {
			log.Println("+++++++++++++++++++++++【ProcessMessage  error】++++++++++++++++++++++")
		}
		log.Println("消息处理完成+++++++++++++++++++++++++++++++消息处理完成++++++++++++++++++++++++++++++++++++++++++++")
		// 手动确认消息
		sess.MarkMessage(msg, "")
	}
	return nil
}

//处理消息
func ProcessMessage(topic string, msg []byte) error {
	log.Println("++++++++++++++++++++++++【topic,msg的值】 :", topic, string(msg[1:10]))
	var (
		Totalstr     string
		Parkingid    string
		Card_network string
	)
	switch topic {
	case types.DdkafkaTopic:

	case types.ZdzkafkaTopic:

	case "topic1":
		log.Println(string(msg))
		return nil
	case "sun":
		log.Println("++++++ topic:sun 获取的值：", string(msg))
		return nil
	}
	log.Println("+++++++++结算金额Totalstr:", Totalstr, "+++++++停车场id——Parkingid:", Parkingid, "++++++卡网络号Card_network：", Card_network)
	//把数据更新到redis
	//conn := RedisConn //初始化redis
	conn := Pool.Get()
	defer conn.Close()
	//1、获取redis中数据
	rhgeterr, value := RedisHGet(&conn, "jiesstatistical", Parkingid)
	if rhgeterr != nil {
		log.Println("+++++++++++++++++++++++【rhgeterr】：", rhgeterr)
		return rhgeterr
	}
	//该停车场为第一次出现
	if value == nil {
		log.Println("+++++++++++++++++++【该停车场为第一次出现】++++++++++++++++")

		rhseterr := RedisHSet(&conn, "jiesstatistical", Parkingid, Totalstr+"|"+strconv.Itoa(1))
		if rhseterr != nil {
			log.Println("+++++++++++++++++++【该停车场为第一次出现】++++++++++++++++++++【RedisHSet error】：", rhseterr)
			return rhseterr
		}
	}
	rhgeterr1, value1 := RedisHGet(&conn, "jiesstatistical", Parkingid)
	if rhgeterr1 != nil {
		log.Println("+++++++++++++++++++++++【rhgeterr】：", rhgeterr)
		return rhgeterr1
	}
	vstr := string(value1.([]uint8))
	log.Println("The hget value is ：", vstr)

	if !StringExist(vstr, "|") {
		return errors.New("+++++++++++++++[kafka  326 hang  error ] ++++++++++++++++++++++++")
	}
	vs := strings.Split(vstr, `"`)
	v := strings.Split(vs[1], `|`)
	zje, _ := strconv.Atoi(v[0])
	zts, _ := strconv.Atoi(v[1])

	//2、处理数据
	total, _ := strconv.Atoi(Totalstr)
	zts = zts + 1
	zje = zje + total
	log.Println("zje:", zje, "total:", total, "zje + total:", zje, "zts", zts)
	//根据消息 更新redis
	//3、hset redis

	// key:"jiesstatistical"  item: 停车场id  value："金额｜总条数"
	rhseterr := RedisHSet(&conn, "jiesstatistical", Parkingid, strconv.Itoa(zje)+"|"+strconv.Itoa(zts))
	if rhseterr != nil {
		log.Println("++++++++++++++++++++++++++++++++++rhseterr+++++++++++++++++++++++:", rhseterr)
	}

	log.Println("++++++++++++++++++++停车场总金额、总笔数更新到redis 成功！+++++++++++++++++++++")
	switch Card_network {
	//省内结算总金额
	case "3201":
		//1、查询数据getredis
		geterr, getvalue := RedisGet(&conn, "snjiesuantotal")
		if geterr != nil {
			log.Println(geterr, "geterr++++++++++++++++++++")
			return errors.New("+++++++++++++++[kafka  362 hang  error ] ++++++++++++++++++++++++")
		}
		if getvalue == nil {
			log.Println("结算总金额、总笔数 get redis  属于第一次")
			setRedis := RedisSet(&conn, "snjiesuantotal", Totalstr+"|"+strconv.Itoa(1))
			if setRedis != nil {
				log.Println(setRedis, "setRedis++++++++++++++++++++")
				return errors.New("+++++++++++++++[kafka  369 hang  error ] ++++++++++++++++++++++++")
			}
		}

		getvstr := string(getvalue.([]uint8))
		log.Println("++++++++++++++++++++++++++++++++++The  get redis value is :", getvstr)

		if !StringExist(getvstr, "|") {
			return errors.New("+++++++++++++++[kafka  370 hang  error ] ++++++++++++++++++++++++")
		}
		//\"3000|3\" 去掉 " 号
		vst := strings.Split(getvstr, `"`)
		getjsv := strings.Split(vst[1], `|`)

		//处理数据 处理 结算总金额、总笔数
		jszje, _ := strconv.Atoi(getjsv[0])
		jszts, _ := strconv.Atoi(getjsv[1])

		//2、处理数据
		jszts = jszts + 1
		jszje = jszje + total
		log.Println("jszje:", getjsv[0], "total:", total, "jszje + total:", jszje)
		//3、更新到redis
		setredis := RedisSet(&conn, "snjiesuantotal", strconv.Itoa(jszje)+"|"+strconv.Itoa(jszts))
		if setredis != nil {
			log.Println("++++++++++++++++++++++++++++++++++The  set redis value is :", setredis)
			return errors.New("+++++++++++++++[kafka  395 hang  error ] ++++++++++++++++++++++++")
		}

	default:
		//省外结算总金额
		//1、查询数据getredis
		geterr, getvalue := RedisGet(&conn, "swjiesuantotal")
		if geterr != nil {
			log.Println("++++++++++++++++++++++++++++++++++The  get redis value is :", geterr)
			return geterr
		}
		if getvalue == nil {
			log.Println("结算总金额、总笔数 get redis  属于第一次")
			setRedis := RedisSet(&conn, "swjiesuantotal", Totalstr+"|"+strconv.Itoa(1))
			if setRedis != nil {
				log.Println("++++++++++++++++++++++++++++++++++The  set redis value is :", setRedis)
				return setRedis
			}
		}

		getvstr := string(getvalue.([]uint8))
		log.Println("The  get redis value is ", getvstr)

		if !StringExist(getvstr, "|") {
			return errors.New("+++++++++++++++[kafka 419  hang  error ] ++++++++++++++++++++++++")
		}
		//\"3000|3\" 去掉 " 号
		vst := strings.Split(getvstr, `"`)
		getjsv := strings.Split(vst[1], `|`)

		//处理数据 处理 结算总金额、总笔数
		jszje, _ := strconv.Atoi(getjsv[0])
		jszts, _ := strconv.Atoi(getjsv[1])

		//2、处理数据
		jszts = jszts + 1
		jszje = jszje + total
		log.Println("jszje:", getjsv[0], "total:", total, "jszje + total:", jszje, "jszts", jszts)
		//3、更新到redis
		setredis := RedisSet(&conn, "swjiesuantotal", strconv.Itoa(jszje)+"|"+strconv.Itoa(jszts))
		if setredis != nil {
			log.Println("++++++++++++++++++++++++++++++++++The  set redis value is :", setredis)
			return setredis
		}
	}
	return nil
}

func handleErrors(group *sarama.ConsumerGroup, wg *sync.WaitGroup) {
	wg.Done()
	for err := range (*group).Errors() {
		log.Println("ERROR", err)
	}
}

//消费
func consume(group *sarama.ConsumerGroup, wg *sync.WaitGroup, name string) error {
	log.Println(name+" group "+"start ok ++++++++++++++++++++消费 kafka consume +++++++++++++++++++++++name:", name)
	wg.Done()
	ctx := context.Background()
	for {
		//c1 ：group  topics := []string{"zdzBillExitDataCollectTopic", "topic1","sun", "billDataCollectTopic"}
		ddtopic := types.DdkafkaTopic
		zdztopic := types.ZdzkafkaTopic
		topics := []string{zdztopic, "topic1", "sun", ddtopic}
		handler := consumerGroupHandler{name: name}

		log.Println("+++++++++++++++++++topics:", topics)
		err := (*group).Consume(ctx, topics, handler)
		if err != nil {
			log.Println("++++++++++++++++++【(*group).Consume  error】  +++++++++++++++++++++", err)
			//panic(err)
			return err
		}
		log.Println(name+" group "+"start ok", "+++++++++++++++++++++++[kafka ok]+++++++++++++++++++++++++")

	}
}

//main 调用 消费kafka
func ConsumerGroup() error {
	var wg sync.WaitGroup
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = false
	config.Version = sarama.V0_10_2_0
	client, err := sarama.NewClient([]string{types.KafkaIpa, types.KafkaIpb, types.KafkaIpc, "172.18.70.21:9092"}, config)
	defer client.Close()
	if err != nil {
		log.Println("++++++++++++++++++++++++++sarama.NewClient 执行出错: ", err)
		return err
		//panic(err)
	}

	group1, err := sarama.NewConsumerGroupFromClient("c1", client)
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++sarama.NewConsumerGroupFromClient 执行出错: ", err)
		return err
		//panic(err)
	}
	//group2, err := sarama.NewConsumerGroupFromClient("c2", client)
	//if err != nil {
	//	panic(err)
	//}
	//group3, err := sarama.NewConsumerGroupFromClient("c3", client)
	//if err != nil {
	//	panic(err)
	//}
	defer group1.Close()
	//defer group2.Close()
	//defer group3.Close()
	wg.Add(1)
	//处理kafka
	cerr := consume(&group1, &wg, "c1")
	if cerr != nil {
		return cerr
	}
	//go consume(&group2,&wg,"c2")
	//go consume(&group3,&wg,"c3")
	wg.Wait()
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	select {
	case <-signals:
	}
	return nil
}
