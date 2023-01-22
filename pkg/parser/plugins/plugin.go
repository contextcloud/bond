package plugins

import (
	"github.com/hashicorp/hcl/v2"
)

type Decoder func(hcl.Body) (interface{}, error)

type Plugins map[string]Decoder

func NewPlugins() Plugins {
	return map[string]Decoder{
		"s3_bucket": S3BucketDecoder,
	}
}
