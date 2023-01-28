package resources

type AwsS3Bucket struct {
	BucketName string `hcl:"bucket_name,attr"`
}
