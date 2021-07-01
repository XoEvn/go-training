package jiegouxing

import (
	"fmt"
)

type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

type AWSClient struct{}

func (c *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success,cpu: %f,mem: %f", cpu, mem)
	return nil
}

type AwsClientAdapter struct {
	Client AWSClient
}

func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	fmt.Printf("aws client run success,cpu: %d,mem:%d", cpu, mem)
	return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (c *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aws client run success, cpu： %d, mem: %d", cpu, mem)
	return nil
}

// AliyunClientAdapter 适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
