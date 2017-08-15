package bibleweb

var bibleAPIKey = "ILhUbZU1gsSMjAXXbr89E7pX2t2lH7kTxOkwxeKM"
var musixMatchKey = "3f8d47fda129f05a4043ef0fce9db474"

const (
	// VersionESV is the key for the ESV
	VersionESV = Version("eng-ESV")
	// VersionNASB is the key for the NASB.
	versionNASB = Version("eng-NASB")
)

// NewAPI returns an APi with the default key
func NewAPI() (*API, error) {
	return &API{
		Key: bibleAPIKey,
	}, nil
}

// API is the main struct to acceess the web bible
type API struct {
	Key string
}

// Version of a bible
type Version string

// ResponseMeta is some weird metadata attached to calls
type ResponseMeta struct {
	Fums          string `json:"fums"`
	FumsTid       string `json:"fums_tid"`
	FumsJsInclude string `json:"fums_js_include"`
	FumsJs        string `json:"fums_js"`
	FumsNoscript  string `json:"fums_noscript"`
}
