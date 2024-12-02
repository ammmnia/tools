package kodo_test

import (
	"context"
	. "github.com/ammmnia/tools/s3/kodo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKodo(t *testing.T) {
	k, err := NewKodo(Config{
		Endpoint:        "https://flyim.s3.ap-northeast-2.amazonaws.com",
		Bucket:          "ap-northeast-2",
		BucketURL:       "https://flyim.s3.ap-northeast-2.amazonaws.com",
		AccessKeyID:     "AKIAW3MEABHC2EOP5RD2",
		AccessKeySecret: "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",
	})
	assert.Nil(t, err)
	assert.Equal(t, "kodo", k.Engine())
	upload, err := k.InitiateMultipartUpload(context.TODO(), "xhxh")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "xhxh", upload.UploadID)
}
