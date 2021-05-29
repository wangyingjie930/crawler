package config

const (
	// Service ports
	//ItemSaverPort = 1234
	//WorkerPort0 = 9000

	// ElasticSearch
	ElasticIndex = "dating_profile"

	// RPC Endpoints
	ItemSaverRpc   = "ItemSaverService.Save"
	CrawServiceRpc = "CrawService.Process"

	// Parser name
	ParseCityList = "ParseCityList"
	ParseCity     = "ParseCity"
	ParseProfile  = "ProfileParse"
	NilParser     = "NilParser"

	// Rate limiting
	Qps = 20
)
