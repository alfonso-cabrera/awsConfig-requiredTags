package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/configservice"

	"github.com/iamthemuffinman/logsip"
)

var log = logsip.New(os.Stdout)

func main() {

	// Create a new aws session.
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-1"),
	})
	if err != nil {
		log.Fatal("Couldn't create intial AWS session. Error was: %s", err.Error())
		return
	}

	// Create session with AWS Config.
	svc := configservice.New(sess)

	// Define params for AWS Config rule.
	params := &configservice.PutConfigRuleInput{
		ConfigRule: &configservice.ConfigRule{ // Required
			Source: &configservice.Source{ // Required
					Owner: aws.String("AWS"),
					SourceIdentifier: aws.String("REQUIRED_TAGS"),
			},
			ConfigRuleName:            aws.String("required-tags"),
			InputParameters:           aws.String("{\"tag1Key\":\"Environment\",\"tag2Key\":\"Project\",\"tag3Key\":\"AlertGroup\"}"),
			Description:               aws.String("Check for compliance with AWS tagging structure. Checks EC2 & RDS"),
			Scope: &configservice.Scope{
					ComplianceResourceTypes: []*string{
							aws.String("AWS::RDS::DBInstance"),
							aws.String("AWS::EC2::Instance"),
					},
			},
	},
}
	// Pass params to create the Config rule
	resp, err := svc.PutConfigRule(params)
	if err != nil {
	    log.Fatal("Config rule creation failed! Error: %s", err.Error())
	    return
	}

	// Pretty-print the response data.
	// The PutConfigRule API call does not return output if successful.
	// TODO: Add call to describe-config-rules so we can verify the rule creation & status.
	fmt.Println(resp)

}
