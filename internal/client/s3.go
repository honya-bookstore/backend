package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3(ctx context.Context, cfgSrv *config.Server) *s3.Client {
	c, err := awsconfig.LoadDefaultConfig(
		ctx,
		awsconfig.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(cfgSrv.S3AccessKey, cfgSrv.S3SecretKey, ""),
		), awsconfig.WithRegion(cfgSrv.S3RegionName))
	if err != nil {
		log.Printf("failed to load config: %v", err)
		return nil
	}
	client := s3.NewFromConfig(c, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(cfgSrv.S3Endpoint)
	})
	exist, err := client.HeadBucket(
		ctx,
		&s3.HeadBucketInput{Bucket: aws.String(cfgSrv.S3Bucket)},
	)
	if err != nil || exist == nil {
		log.Printf("bucket not exist: %v", err)
	}

	return client
}

func NewS3Presign(s3Client *s3.Client) *s3.PresignClient {
	return s3.NewPresignClient(s3Client)
}
