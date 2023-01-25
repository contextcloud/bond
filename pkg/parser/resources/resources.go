package resources

type Factory func() interface{}
type Resources map[string]Factory

func NewResources() Resources {
	return map[string]Factory{
		"aws_s3_bucket": func() interface{} { return &S3Bucket{} },
	}
}
