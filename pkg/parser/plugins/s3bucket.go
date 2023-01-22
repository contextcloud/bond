package plugins

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
)

type S3Bucket struct {
	Foo  string  `hcl:"foo"`
	Foo2 *string `hcl:"foo2"`
}

func S3BucketDecoder(b hcl.Body) (interface{}, error) {
	var obj S3Bucket
	if err := gohcl.DecodeBody(b, nil, &obj); err != nil {
		return nil, err
	}
	return &obj, nil
}
