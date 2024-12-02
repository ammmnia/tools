package kodo_test

import (
	. "github.com/ammmnia/tools/s3/kodo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKodo(t *testing.T) {
	k, err := NewKodo(Config{
		Endpoint:        "https://flyim.s3.ap-northeast-2.amazonaws.com",
		Bucket:          "flyim",
		BucketURL:       "https://flyim.s3.ap-northeast-2.amazonaws.com",
		AccessKeyID:     "AKIAW3MEABHC2EOP5RD2",
		AccessKeySecret: "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",
	})
	assert.Nil(t, err)
	assert.Equal(t, "kodo", k.Engine())

}
