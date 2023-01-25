package resources

type S3Bucket struct {
	BucketName string `hcl:"bucket_name,attr"`
}
