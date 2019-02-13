package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehend"
	"io/ioutil"
	"log"
)

func main() {
	getAwsComprehend()

}

func getAwsComprehend(){
	sess := getSession()
	client := getClient(sess)
	req, response := detectEntity(client)
	err := req.Send()

	if err == nil { // resp is now filled
		fmt.Println(response)
		getCity(response)
	} else {
		fmt.Println(err)
	}
}

func getCity(response *comprehend.DetectEntitiesOutput){
	for _, ent := range response.Entities {
		fmt.Printf(aws.StringValue(ent.Text) + "\n")
		//fmt.Printf("%v \n", *ent.Text)
	}
}

func getSession() *session.Session{
	return session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
}

func getClient(sess *session.Session) *comprehend.Comprehend{
	return comprehend.New(sess, &aws.Config{
		Region: aws.String("us-east-2"),
	})
}

func getText() string{
	text, err := ioutil.ReadFile("comprehend_text.txt")

	if err != nil{
		log.Fatal(err)
	}
	return string(text)
}

func detectEntity(client *comprehend.Comprehend) (*request.Request, *comprehend.DetectEntitiesOutput) {
	awsText := getText()
	// Example sending a request using the DetectEntitiesRequest method.
	return client.DetectEntitiesRequest(&comprehend.DetectEntitiesInput{LanguageCode: aws.String("en"), Text: aws.String(awsText)})
}