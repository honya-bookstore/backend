package objectstorages3

import (
	"errors"
	"net"

	"backend/internal/domain"

	"github.com/aws/smithy-go"
	"github.com/hashicorp/go-multierror"

	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
)

func ToDomainErrorFromS3(err error) error {
	if err == nil {
		return nil
	}

	var apiErr smithy.APIError
	if errors.As(err, &apiErr) {
		switch apiErr.ErrorCode() {

		case "AccessDenied":
			return multierror.Append(domain.ErrForbidden, errors.New("s3 access denied"), err)

		case "NoSuchBucket", "NoSuchKey":
			return multierror.Append(domain.ErrNotFound, errors.New("s3 resource not found"), err)

		case "InvalidBucketName", "InvalidObjectState":
			return multierror.Append(domain.ErrInvalid, errors.New("invalid s3 request"), err)

		case "BucketAlreadyExists", "BucketAlreadyOwnedByYou":
			return multierror.Append(domain.ErrExists, errors.New("bucket conflict"), err)

		case "ServiceUnavailable", "SlowDown":
			return multierror.Append(domain.ErrUnavailable, errors.New("s3 service unavailable"), err)

		case "RequestTimeout":
			return multierror.Append(domain.ErrTimeout, errors.New("s3 request timeout"), err)

		default:
			return multierror.Append(domain.ErrServiceError, errors.New("unknown s3 error"), err)
		}
	}

	var httpErr *awshttp.ResponseError

	if errors.As(err, &httpErr) {
		switch httpErr.HTTPStatusCode() {
		case 400:
			return multierror.Append(domain.ErrInvalid, errors.New("s3 bad request"), err)
		case 403:
			return multierror.Append(domain.ErrForbidden, errors.New("s3 forbidden"), err)
		case 404:
			return multierror.Append(domain.ErrNotFound, errors.New("s3 not found"), err)
		case 409:
			return multierror.Append(domain.ErrConflict, errors.New("s3 conflict"), err)
		case 408:
			return multierror.Append(domain.ErrTimeout, errors.New("s3 timeout"), err)
		case 503:
			return multierror.Append(domain.ErrUnavailable, errors.New("s3 unavailable"), err)
		default:
			return multierror.Append(domain.ErrServiceError, errors.New("s3 http error"), err)
		}
	}
	var netErr net.Error
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return multierror.Append(domain.ErrTimeout, errors.New("s3 network timeout"), err)
		}
		return multierror.Append(domain.ErrUnavailable, errors.New("s3 network failure"), err)
	}

	return multierror.Append(domain.ErrUnknown, errors.New("unexpected s3 error"), err)
}
