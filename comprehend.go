package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/comprehend"
)

func main() {

	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router))

	router.HandleFunc("/comprehend", getAwsComprehend).Methods("GET")

}

func detectEntity(client *comprehend.Comprehend) (*request.Request, *comprehend.DetectEntitiesOutput) {
	// Example sending a request using the DetectEntitiesRequest method.
	return client.DetectEntitiesRequest(&comprehend.DetectEntitiesInput{LanguageCode: aws.String("en"), Text: aws.String("Amazon.com, Inc. is located in Seattle, WA and was founded July 5th, 1994 by Jeff Bezos, allowing customers to buy everything from books to blenders. Seattle is north of Portland and south of Vancouver, BC. Other notable Seattle - based companies are Starbucks and Boeing.")})
}

func getAwsComprehend(w http.ResponseWriter, r *http.Request){
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	client := comprehend.New(sess, &aws.Config{
		Region: aws.String("us-east-2"),
	})

	req, response := detectEntity(client)

	err := req.Send()

	if err == nil { // resp is now filled
		err := json.NewEncoder(w).Encode(response)
		fmt.Println(response)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	} else {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
	}

	//for _, ent := range response.Entities {
	//	fmt.Printf(aws.StringValue(ent.Text) + "\n")
	//	//fmt.Printf("%v \n", *ent.Text)
	//}

}