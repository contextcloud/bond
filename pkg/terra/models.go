package terra

type BackendType string

const (
	BackendTypeS3    BackendType = "s3"
	BackendTypeLocal BackendType = "local"
)

type Backend struct {
	Type    BackendType
	Options interface{}
}

type BackendS3 struct {
	Bucket string `hcl:"bucket"`
	Region string `hcl:"region"`
}

type BackendLocal struct {
	Path string `hcl:"path"`
}
