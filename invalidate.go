package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront"
	"time"
)

func invalidate(distId string, invPath string) error {
	sess, err := session.NewSession()
	if err != nil {
		return err
	}

	svc := cloudfront.New(sess)
	input := &cloudfront.CreateInvalidationInput{
		DistributionId: &distId,
		InvalidationBatch: &cloudfront.InvalidationBatch{
			CallerReference: aws.String(time.Now().Format("2006-02-01 15:04:05")),
			Paths: &cloudfront.Paths{
				Quantity: aws.Int64(1),
				Items:    []*string{aws.String(invPath)},
			},
		},
	}

	_, err = svc.CreateInvalidation(input)
	if err != nil {
		return err
	}

	return nil
}
