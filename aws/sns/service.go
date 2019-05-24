package sns

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

//SendSnsMessageJSON sends given JSON message into provided topic ARN.
//
//Params
//*message: contains JSON to be send to SNS
//*topicArn: sns resurce to which we will send
//*messageParams: is a map of MessageAttributeValue
//
//Returns nil on success
func SendSnsMessageJSON(message, topicArn string, messageParams map[string]*sns.MessageAttributeValue) error {
	svc := sns.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

	snsMessage := SNSMessage{
		Default: message,
	}
	snsMessageBytes, _ := json.Marshal(snsMessage)
	snsMessageStr := string(snsMessageBytes)

	// Logs the message we will be sending to SNS.
	log.Println(snsMessageStr)

	publishOutput, err := svc.Publish(&sns.PublishInput{
		// MessageAttributes: map[string]*sns.MessageAttributeValue{
		// 	"resource": &sns.MessageAttributeValue{
		// 		DataType:    aws.String("String"),
		// 		StringValue: aws.String("slack"),
		// 	},
		// },
		MessageAttributes: messageParams,
		TopicArn:          aws.String(topicArn),
		Message:           aws.String(snsMessageStr),
		MessageStructure:  aws.String("json"),
	})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(publishOutput.GoString())

	return nil
}
