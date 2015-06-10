package providers

type Provider interface {
	Data() *ProviderData
	GetEmailAddress(body []byte, access_token string) (string, error)
	Redeem(string, string) ([]byte, string, error)
	ValidateToken(access_token string) bool
	GetLoginURL(redirectURI, finalRedirect string) string
}

func New(provider string, p *ProviderData) Provider {
	switch provider {
	case "myusa":
		return NewMyUsaProvider(p)
	case "linkedin":
		return NewLinkedInProvider(p)
	case "github":
		return NewGitHubProvider(p)
	case "idme":
		return NewIDmeProvider(p)
	default:
		return NewIDmeProvider(p)
	}
}
