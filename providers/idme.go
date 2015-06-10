package providers

import (
  "encoding/base64"
  "encoding/json"
  "errors"
  "net/url"
  "strings"
)

type IDmeProvider struct {
  *ProviderData
}

func NewIDmeProvider(p *ProviderData) *IDmeProvider {
  p.ProviderName = "IDme"
  if p.LoginUrl.String() == "" {
    p.LoginUrl = &url.URL{Scheme: "http",
      Host: "idme-authority.dev",
      Path: "/authorize"}
  }
  if p.RedeemUrl.String() == "" {
    p.RedeemUrl = &url.URL{Scheme: "http",
      Host: "idme-authority.dev",
      Path: "/token"}
  }
  if p.ValidateUrl.String() == "" {
    p.ValidateUrl = &url.URL{Scheme: "http",
      Host: "idme-authority.dev",
      Path: "/userinfo"}
  }
  if p.Scope == "" {
    p.Scope = "profile email"
  }
  return &IDmeProvider{ProviderData: p}
}

func (p *IDmeProvider) GetEmailAddress(body []byte, access_token string) (string, error) {

  var email struct {
    Email string `json:"email"`
  }
  params := url.Values{
    "access_token": {access_token},
  }
  endpoint := "http://idme-authority.dev/userinfo" + params.Encode()
  resp, err := http.DefaultClient.get(endpoint)
  if err != nil {
    return "", err
  }
  if email.Email == "" {
    return "", errors.New("missing email")
  }
  return email.Email, nil
}

func (p *IDmeProvider) ValidateToken(access_token string) bool {
  return validateToken(p, access_token, nil)
}
