package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {
	/// 从环境变量中获取访问凭证。运行本代码示例之前，请确保已设置环境变量OSS_ACCESS_KEY_ID和OSS_ACCESS_KEY_SECRET。
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("yourEndpoint", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 填写存储空间名称，例如examplebucket。
	bucket, err := client.Bucket("examplebucket")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 指定Object存储类型为低频访问。
	storageType := oss.ObjectStorageClass(oss.StorageIA)

	// 指定Object访问权限为私有。
	objectAcl := oss.ObjectACL(oss.ACLPrivate)

	// 将字符串"Hello OSS"上传至exampledir目录下的exampleobject.txt文件。
	err = bucket.PutObject("exampledir/exampleobject.txt", strings.NewReader("Hello OSS"), storageType, objectAcl)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
