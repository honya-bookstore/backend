package objectstorages3

import (
	"context"
	"time"

	"backend/config"
	"backend/internal/application"
	"backend/internal/delivery/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type Media struct {
	s3Client        *s3.Client
	s3PresignClient *s3.PresignClient
	cfgSrv          *config.Server
}

func ProvideMedia(
	s3Client *s3.Client,
	s3PresignClient *s3.PresignClient,
	cfgSrv *config.Server,
) *Media {
	return &Media{
		s3Client:        s3Client,
		s3PresignClient: s3PresignClient,
		cfgSrv:          cfgSrv,
	}
}

var _ application.MediaObjectStorage = (*Media)(nil)

func (p *Media) GetUploadImageURL(ctx context.Context) (*http.UploadImageURLResponseDTO, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	idStr := id.String()
	url, err := p.s3PresignClient.PresignPutObject(
		ctx,
		&s3.PutObjectInput{
			Bucket: aws.String(p.cfgSrv.S3Bucket),
			Key:    aws.String(S3MediaFolderTemp + idStr),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	return &http.UploadImageURLResponseDTO{
		URL: url.URL,
		Key: idStr,
	}, nil
}

func (p *Media) GetDeleteImageURL(ctx context.Context, imageID uuid.UUID) (*http.DeleteImageURLResponseDTO, error) {
	url, err := p.s3PresignClient.PresignDeleteObject(
		ctx,
		&s3.DeleteObjectInput{
			Bucket: aws.String(p.cfgSrv.S3Bucket),
			Key:    aws.String(S3MediaFolder + imageID.String()),
		},
		s3.WithPresignExpires(10*time.Minute),
	)
	if err != nil {
		return nil, ToDomainErrorFromS3(err)
	}
	return &http.DeleteImageURLResponseDTO{
		URL: url.URL,
	}, nil
}

func (p *Media) PersistImageFromTemp(ctx context.Context, key string) error {
	_, err := p.s3Client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(p.cfgSrv.S3Bucket),
		CopySource: aws.String(p.cfgSrv.S3Bucket + "/" + S3MediaFolderTemp + key),
		Key:        aws.String(S3MediaFolder + key),
	})
	if err != nil {
		return ToDomainErrorFromS3(err)
	}
	_, err = p.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(p.cfgSrv.S3Bucket),
		Key:    aws.String(S3MediaFolderTemp + key),
	})
	if err != nil {
		return ToDomainErrorFromS3(err)
	}
	return nil
}
