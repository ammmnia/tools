package aws_test

import (
	"context"
	. "github.com/ammmnia/tools/s3/aws"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAWS(t *testing.T) {
	ss, err := NewAWS(Config{
		Bucket:          "ap-northeast-2",
		BucketURL:       "https://flyim.s3.ap-northeast-2.amazonaws.com",
		AccessKeyID:     "AKIAW3MEABHC2EOP5RD2",
		AccessKeySecret: "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",
	})
	assert.Nil(t, err)
	assert.Equal(t, "aws", ss.Engine())
	id, _ := uuid.NewUUID()
	upload, err := ss.InitiateMultipartUpload(context.Background(), id.String())
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "flyim", upload.Bucket)
	t.Logf(" %v", upload)

	object, err := ss.StatObject(context.Background(), "README.md")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "README.md", object.Key)
}
