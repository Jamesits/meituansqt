package sqtapi

type Config struct {
	Version      string
	ApiEndpoint  string
	EnterpriseID int64  // "entId"
	AccessKey    string // sometimes called "token"
	SecretKey    string
}
