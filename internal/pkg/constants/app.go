package constants

const (
	SERVER_ADDRESS = "127.0.0.1:8080"
	EMAIL_REGEX    = "^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$"
	MOBILE_REGEX   = "^([+]\\d{2})?\\d{10}$"
)

const (
	USER      = "user"
	ADMIN     = "admin"
	ORGANIZER = "organizer"
)

const (
	ACTIVE_STATUS   = "active"
	INACTIVE_STATUS = "inactive"
	BANNED_STATUS   = "banned"
)

type TokenKeyType string

const TokenKey TokenKeyType = "token-key"
