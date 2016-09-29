package keys

import "strings"

const (
	keyprefix = "portc-"

	organisationKey = "organisation"
	usernameKey     = "username"
	userIDKey       = "userid"
	appIDkey        = "appid"
	vendorKey       = "appvendor"
	sigKey          = "signature"

	firstNameKey = "first-name"
	lastNameKey  = "last-name"
)

// GetKeyPrefix returns portcullis key prefix
func GetKeyPrefix() string {
	return keyprefix
}

// GetSignatureKey retrieves the key used for portcullis verification signature
func GetSignatureKey() string {
	return keyprefix + sigKey
}

// GetAppIDKey retrieves the key used for App ID
func GetAppIDKey() string {
	return keyprefix + appIDkey
}

// GetAppVendorKey retrieves the key used for app vendor
func GetAppVendorKey() string {
	return keyprefix + vendorKey
}

// GetOrganisationKey retrieves the key used for organisation
func GetOrganisationKey() string {
	return keyprefix + organisationKey
}

// GetUsernameKey retrieves the key used for username
func GetUsernameKey() string {
	return keyprefix + usernameKey
}

// GetUserIDKey retrieves the key used for user ID
func GetUserIDKey() string {
	return keyprefix + userIDKey
}

// GetFirstNameKey retrieves the first name of the user make the request
func GetFirstNameKey() string {
	return keyprefix + firstNameKey
}

// GetLastNameKey retrieves the last name of the user making the request
func GetLastNameKey() string {
	return keyprefix + lastNameKey
}

// GetGenericKeyForString retrieves key for given generic value
func GetGenericKeyForString(in string) string {
	key := strings.Replace(in, " ", "-", -1)
	key = strings.ToLower(key)
	return keyprefix + key
}
