package main

import (
	"context"
	"flag"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var (
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
