package gasrecord_infrastructure_bbold

import "fmt"

type BucketNotFound struct {
	BucketName string
}

func (myError BucketNotFound) Error() string {
	return fmt.Sprintf("Bucket Not Found Error | bucket name: <%v>", myError.BucketName)
}

func NewBucketNotFound(bucketName string) BucketNotFound {
	return BucketNotFound{bucketName}
}
