package sdkms

const (
	SdkmsBaseEndpoint = "https://sdkms.fortanix.com"

	SessionAuthEndpoint = SdkmsBaseEndpoint + "/sys/v1/session/auth"

	GetAccountsEndpoint   = SdkmsBaseEndpoint + "/sys/v1/accounts"
	SelectAccountEndpoint = SdkmsBaseEndpoint + "/sys/v1/session/select_account"

	GetGroupsEndpoint = SdkmsBaseEndpoint + "/sys/v1/groups"

	PostKeyEndpoint   = SdkmsBaseEndpoint + "/crypto/v1/keys"
	ExportKeyEndpoint = SdkmsBaseEndpoint + "/crypto/v1/keys/export"
)
