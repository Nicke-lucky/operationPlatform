package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/yunify/qingstor-sdk-go/config"
	qs "github.com/yunify/qingstor-sdk-go/service"
	"os"
)

//private String zone = "sh1a";
//private String bucketName = "gdkjetcpark";
//private String accessKey="QGYIKYOMPKJLFWVURGBG",
//private String accessSecret="QLPAaRF1legVvjbA8nfz2bN2EiuKRvD9f8HKZISX");
//
//http://gdkjetcpark.sh1a.qingstor.com/

//可以在不同区域创建 存储空间(Bucket)
const (
	Zone         = "sh1a" //地区
	BucketName   = "gdkjetcpark"
	AccessKey    = "QGYIKYOMPKJLFWVURGBG"
	AccessSecret = "QLPAaRF1legVvjbA8nfz2bN2EiuKRvD9f8HKZISX"
)

func QingStorInit() {
	//发起请求前首先建立需要初始化服务:
	//初始化了一个 QingStor Service
	//configuration, _ := config.New("ACCESS_KEY_ID", "SECRET_ACCESS_KEY")
	configuration, _ := config.New(AccessKey, AccessSecret)
	qsService, _ := qs.Init(configuration)

	//{
	//获取账户下的 Bucket 列表
	qsOutput, _ := qsService.ListBuckets(nil)
	// Print the HTTP status code.
	// Example: 200
	fmt.Println(qs.IntValue(qsOutput.StatusCode))
	// Print the bucket count.
	// Example: 5
	fmt.Println(qs.IntValue(qsOutput.Count))

	//}

	//初始化并创建 Bucket, 需要指定 Bucket[桶] 名称和所在 Zone:
	//bucket, _ := qsService.Bucket("test-bucket", "pek3a")
	//bucket, _ := qsService.Bucket(	BucketName, Zone)
	//putBucketOutput, _ := bucket.Put()

	//取 Bucket 中存储的 Object 列表
	bOutput, err := bucket.ListObjects(nil)
	// Print the HTTP status code.
	// Example: 200
	fmt.Println(qs.IntValue(bOutput.StatusCode))
	// Print the key count.
	// Example: 0
	fmt.Println(len(bOutput.Keys))

	//创建一个 Object
	//例如上传一张屏幕截图:
	// Open file
	fname := ""
	f, err := os.Open("./version/" + fname)
	if err != nil {
		logrus.Print(err)
	}

	defer func() {
		_ = f.Close()
	}()

	// Put object
	oOutput, err := bucket.PutObject(fname, &service.PutObjectInput{Body: f})

	// 所有 >= 400 的 HTTP 返回码都被视作错误
	if err != nil {
		// Example: QingStor Error: StatusCode 403, Code "permission_denied"...
		fmt.Println(err)
	} else {
		// Print the HTTP status code.
		// Example: 201
		fmt.Println(qs.IntValue(oOutput.StatusCode))
	}

}

//StatusCode
