package portcullis

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/fortifi/portcullis-go/keys"
)

// ReqInfo is the structure for deserialised request information
type ReqInfo struct {
	OrganisationID string
	UserID         string
	Username       string
	FirstName      string
	LastName       string
}

// FromContext retrieves request info from given request context
func FromContext(ctx context.Context) ReqInfo {
	md, _ := metadata.FromContext(ctx)
	res := ReqInfo{
		OrganisationID: safeGetMetaValString(keys.GetOrganisationKey(), md),
		UserID:         safeGetMetaValString(keys.GetUserIDKey(), md),
		Username:       safeGetMetaValString(keys.GetUsernameKey(), md),
		FirstName:      safeGetMetaValString(keys.GetFirstNameKey(), md),
		LastName:       safeGetMetaValString(keys.GetLastNameKey(), md),
	}
	return res
}

func safeGetMetaValString(key string, md metadata.MD) string {
	result := ""
	if md != nil {
		if len(md[key]) != 0 {
			result = md[key][0]
		}
	}
	return result
}
