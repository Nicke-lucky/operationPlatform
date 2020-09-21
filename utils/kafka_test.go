package utils

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestKafkaProduction(t *testing.T) {

}

func TestConsumer(t *testing.T) {
	for {

		logrus.Print("消费数据")
		Consumer()

	}

}

//
func TestProducer(t *testing.T) {
	//logrus.Print("生成数据")
	Producer("topic1", "09876")

}

//UseKafka
func TestUseKafka(t *testing.T) {

	UseKafka()
	//2.12-2.5.0
	//	consumer groups require Version to be >= V0_10_2_0
}

//ConsumerGroup()
func TestConsumerGroup(t *testing.T) {
	ConsumerGroup()
}

//ConsumerGroup()
func TestConsumerGroups(t *testing.T) {
	/*jsonData := []byte(`{"head":{
	    "topic":"billExitDataCollectTopic",
	    "index":"325",
	    "topicreply":"SG_GATEWAY_mytopic_test",
	    "id":"3202999999110320200729162439000000c6",
	    "topictime":"2020-07-29 16:24:39",
	    "lane_id":"1103",
	    "parking_id":"2002009998",
	    "company_id":"3202999999",
	    "source_type":"ddd"},
	 "data":{"bill_id":"3202999999110320200729162439000000c6","programver":"1.0.0.071713.RC",
	"programstarttime":"2020-07-29 16:22:09","company_id":"3202999999","parking_id":"2002009998","channel_id":"",
	"lane_id":"1103","record_id":"32029999991103202007291624390000","record_no":"15","lane_record_no":"15","etc_terminal_id":"01320002da0f",
	"etc_termtrad_no":"000000c6",
	    "obu_serial":"3301031706220961","obu_id":"621325aa",
	    "obu_issuer":"d5e3bdad33010001","obu_plate":"浙B663DP","obu_plate_color":"0","obu_expire_date":"20280121","card_id":"1512230200031046",
	    "card_network":"3301","card_trade_no":"0739","card_issuer":"d5e3bdadd5e3bdad","card_type":"23","card_expired":"20260706",
	"reset_money":"1999236042",
	"money":"500","tac":"7e8a2262","trade_time":"2020-07-29 16:24:39","trade_type":"0","lane_key":"KIEKSUNJ8Z4A4HLP","vehicle":"1","obu_status":"2200",
	    "plate_num":"浙BY003V","plate_color":"0","algorithm_type":"","black_ver":"202007291624","notify":"","entry_time":"2020-07-29 15:54:39",
	"duration":"30","bill_description":"this is desc!","file0019":"aa290000000000aa28f02e0000000000000000000000000000000000000000000000000000000000000000",
	"file0015":"d5e3bdadd5e3bdad1711210115122302000310462016070620260706d5e342593030335600000000000000","obuinfo":"",
	"vehicileinfo":"d5e34236363344500000000000000100000000000000000000000500b8a5481204c38b1085d274388b4af831323134373433001200000000000000",
	"feebackinfo":"","costtime":"343","psamversion":"0","card_version":"17","obu_type":"0","key_type":"0","before_money":"0","devicetype":"0"}}`)*/
	//ProcessMessage("billDataCollectTopic", jsonData)
	//zdzBillExitDataCollectTopic
	zdzjsonData := []byte(`{"head":{
"topic":"zdzBillExitDataCollectTopic", 
"index":"803", "topicreply":"172.18.70.21messageNoticeTopic", "id":"320107001111022020072911264400000175", "topictime":"2020-07-29 11:26:48", 
"lane_id":"1102", "parking_id":"3201000016", "company_id":"3201070011", 
"source_type":"zdz"}, 
"data":{"card_network":"3201","obu_expire_date":"20220215","trade_time":"2020-07-29 11:26:44","plate_num":"苏D54321",
"obu_plate":"苏D54321","bill_description":"苏D54321|2020-07-29 11:26:43出场|红公馆|伟龙泊时捷ETC测试|停车1小时0分钟|10.00元",
"obu_plate_color":"0","black_ver":"202007291126","obu_issuer":"bdadcbd532010001",
"notify":"http://office.mobcb.com:11142/api/v3/external/etc/bill/exit/notify?mallId=BGM3OICRjHRm7ige","vehicle":"1",
"duration":"60","card_issuer":"bdadcbd532010001","before_money":"10000","record_no":"220734","tac":"230b19c8",
"reset_money":"9000","etc_terminal_id":"01320002da13","bill_id":"320107001111022020072911264400000175",
"lane_key":"3MX6SQFE0YZUKY04","entry_time":"2020-07-29 10:26:43","lane_record_no":"220734","plate_color":"0","company_id":"3201070011",
"card_expired":"20250110","card_trade_no":"0269","lane_id":"1102","card_type":"23","card_id":"1512230200031048","algorithm_type":"9",
"record_id":"777629924269142","money":"1000","etc_termtrad_no":"00000175","sign_data":"7AE2C4D545ADCD7F","obu_serial":"2101192702238269",
"parking_id":"9801000016","trade_type":"9","obu_id":"621325aa","obu_status":"205"}}`)
	ProcessMessage("zdzBillExitDataCollectTopic", zdzjsonData)
}

//GetRandStr
func TestGetRandStr(t *testing.T) {
	logrus.Print(GetRandStr(4))
}
