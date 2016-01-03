package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const (
	usageMsg = "Usage:\n\tmain region_name command"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(usageMsg)
		os.Exit(-1)
	}

	region := os.Args[1]
	if region == "" {
		fmt.Println(usageMsg)
		os.Exit(-1)
	}

	command := os.Args[2]

	switch command {
	case "ls":
		ls(region)
		break
	case "create":
		fmt.Println("Under construction...")
		break
	default:
		fmt.Println("Unknown command")
	}
}

func ls(region string) {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String(region)})

	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("> Number of reservation sets: ", len(resp.Reservations))

	for idx, res := range resp.Reservations {
		fmt.Println("  > Number of instances: ", len(res.Instances))
		for _, inst := range resp.Reservations[idx].Instances {
			fmt.Println("    - Instance ID: ", *inst.InstanceId)
		}
	}
}
