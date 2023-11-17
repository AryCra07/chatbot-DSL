package consts

// information
const (
	EmptyStr    = ""
	ModeTest    = "test"
	ModeRelease = "release"
)

// layer
const (
	Main       = "Main"
	Config     = "Config"
	Controller = "Controller"
	Service    = "Service"
	Dao        = "Dao"
	Test       = "Test"
)

// message and status
const (
	InvalidRequest = "Request is invalid"
	NotExistId     = "NotExistId"
	SUCCESS        = 0
	FAIL           = -1
	InvalidToken   = -2
)

// database
const (
	DefaultAuth = 0
	AdminAuth   = 1
	DefaultName = "user"
)
