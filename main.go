package main

import (
	"context"
	"flag"
	"log"
	"math/rand"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	stsTypes "github.com/aws/aws-sdk-go-v2/service/sts/types"
)

var (
	role    = flag.String("role", "", "Role to assume")
	message = flag.String("message", "", "Message to send")
	topic   = flag.String("topic", "", "Topic ARN to send message to")
)

func main() {
	ctx := context.Background()

	flag.Parse()

	if *message == "" {
		log.Fatal("Message is required")
	}

	if *topic == "" {
		log.Fatal("Topic is required")
	}

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if *role != "" {
		log.Println("Assuming role: " + *role)

		sourceAccount := sts.NewFromConfig(cfg)

		response, err := sourceAccount.AssumeRole(ctx, &sts.AssumeRoleInput{
			RoleArn:         aws.String(*role),
			RoleSessionName: aws.String("sc_" + strconv.Itoa(10000+rand.Intn(25000))),
		})

		if err != nil {
			log.Fatal(err)
		}

		var assumedRoleCreds *stsTypes.Credentials = response.Credentials
		cfg, err = config.LoadDefaultConfig(ctx, config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(
				*assumedRoleCreds.AccessKeyId,
				*assumedRoleCreds.SecretAccessKey,
				*assumedRoleCreds.SessionToken,
			)))

		if err != nil {
			panic(err)
		}
	}

	cli := sns.NewFromConfig(cfg)

	input := &sns.PublishInput{
		Message:  message,
		TopicArn: topic,
	}

	log.Println("Sending message: " + *message + " to topic: " + *topic)

	out, err := cli.Publish(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message sent: " + *out.MessageId)
}
