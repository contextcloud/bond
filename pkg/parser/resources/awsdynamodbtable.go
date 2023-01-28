package resources

type AwsDynamodbTable struct {
	TableName string `hcl:"table_name,attr"`
}
