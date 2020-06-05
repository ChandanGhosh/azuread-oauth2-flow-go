package config

const (
	ClientID      = "ba086ea1-c1bd-4107-bea2-36e43b13a972"
	AuthURL       = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	RedirectURL   = "https://login.microsoftonline.com/common/oauth2/nativeclient"
	TokenURL      = "https://login.microsoftonline.com/common/oauth2/v2.0/token"
	DriveURL      = "https://graph.microsoft.com/v1.0/me/drive"
	MeURL         = "https://graph.microsoft.com/v1.0/me"
	ItemsByIDURL  = "https://graph.microsoft.com/v1.0/me/drive/items/"
	ItemByPathURL = "https://graph.microsoft.com/v1.0/me/drive/root:/"
	DrivesByIDURL = "https://graph.microsoft.com/v1.0/drives/"

	Scopes        = "&scope=User.Read Files.ReadWrite Files.ReadWrite.All Sites.Read.All Sites.ReadWrite.All offline_access"
	DeviceCodeURL = "https://login.microsoftonline.com/common/oauth2/v2.0/devicecode"
	GrantType     = "&grant_type=urn:ietf:params:oauth:grant-type:device_code"
)

// Config ..
type Config struct {
}
