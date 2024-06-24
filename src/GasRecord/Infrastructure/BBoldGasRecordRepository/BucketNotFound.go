package gasrecord_infrastructure_bbold

import "fmt"

type bucketNotFound struct {
	BucketName string
}

func (myError bucketNotFound) Error() string {
	return fmt.Sprintf("Bucket Not Found Error | bucket name: <%v>", myError.BucketName)
}

func newBucketNotFound(bucketName string) bucketNotFound {
	return bucketNotFound{bucketName}
}
