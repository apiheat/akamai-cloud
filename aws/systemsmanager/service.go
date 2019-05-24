package systemsmanager

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
)

//GetSSMParamByKey Retrieves value from SSM Param Store by given key.
//Returns string value
func GetSSMParamByKey(key string) string {
	svc := ssm.New(session.New(), aws.NewConfig().WithRegion(os.Getenv("AWS_REGION")))

	keyname := key
	withDecryption := true
	param, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           &keyname,
		WithDecryption: &withDecryption,
	})
	if err != nil {
		log.Println(err)
	}

	value := *param.Parameter.Value

	return value

}
