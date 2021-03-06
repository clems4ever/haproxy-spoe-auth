package tests

const (
	// ProtectedLdapURL is the URL for the LDAP use case
	ProtectedLdapURL = "http://protected-ldap.example.com:9080/"
	// ProtectedOidcURL is the URL for the OIDC use case
	ProtectedOidcURL = "http://protected-oidc.example.com:9080/"
	// ProtectedOAuth2URL is the URL for the OAuth2 use case
	ProtectedOAuth2URL = "http://protected-oauth2.example.com:9080/"
	// UnprotectedURL is the URL for the unprotected app
	UnprotectedURL = "http://unprotected.example.com:9080/"

	// LogoutOidcURL is the URL used to log out the user
	LogoutOidcURL = "http://auth-oidc.example.com:9080/"
	// LogoutOAuht2URL is the URL used to log out the user
	LogoutOAuht2URL = "http://auth-oauth2.example.com:9080/"
)
