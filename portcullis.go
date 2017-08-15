package portcullis

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/kubex/portcullis-go/keys"
)

// ReqInfo is the structure for deserialised request information
type ReqInfo struct {
	ProjectID   string
	UserID      string
	AppID       string
	VendorID    string
	Username    string
	FirstName   string
	LastName    string
	signature   string
	meta        metadata.MD
	Roles       []string
	Permissions []string
}

// Verify checks that the request signature matches using signature key
func (r *ReqInfo) Verify(sigKey string) bool {
	mac := hmac.New(sha256.New, []byte(sigKey))
	mk := make([]string, len(r.meta))
	i := 0
	for k := range r.meta {
		mk[i] = k
		i++
	}
	sort.Strings(mk)

	m := ""
	for _, v := range mk {
		if strings.HasPrefix(v, keys.GetKeyPrefix()) {
			m = m + v
			b := r.meta[v]
			sort.Strings(b)
			for _, a := range b {
				m = m + a
			}
		}
	}

	mac.Write([]byte(m))
	expectedMAC := hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(r.signature), []byte(expectedMAC))
}

// GlobalAppID is getter for requesting app's Global ID
func (r *ReqInfo) GlobalAppID() string {
	return fmt.Sprintf("%s/%s", r.VendorID, r.AppID)
}

// HasRole check if the user has a specific role
func (r *ReqInfo) HasRole(checkRole string) bool {
	for _, role := range r.Roles {
		if role == checkRole {
			return true
		}
	}
	return false
}

// HasPermission check if the user has a specific permission
func (r *ReqInfo) HasPermission(checkPermission string) bool {
	for _, permission := range r.Permissions {
		if permission == checkPermission {
			return true
		}
	}
	return false
}

// FromContext retrieves request info from given request context
func FromContext(ctx context.Context) ReqInfo {
	md, _ := metadata.FromContext(ctx)
	res := ReqInfo{
		ProjectID:   safeGetMetaValString(keys.GetProjectKey(), md),
		UserID:      safeGetMetaValString(keys.GetUserIDKey(), md),
		Username:    safeGetMetaValString(keys.GetUsernameKey(), md),
		FirstName:   safeGetMetaValString(keys.GetFirstNameKey(), md),
		LastName:    safeGetMetaValString(keys.GetLastNameKey(), md),
		AppID:       safeGetMetaValString(keys.GetAppIDKey(), md),
		VendorID:    safeGetMetaValString(keys.GetAppVendorKey(), md),
		signature:   safeGetMetaValString(keys.GetSignatureKey(), md),
		Roles:       safeGetMetaValStringSlice(keys.GetRolesKey(), md),
		Permissions: safeGetMetaValStringSlice(keys.GetPermissionsKey(), md),
		meta:        md,
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

func safeGetMetaValStringSlice(key string, md metadata.MD) []string {
	result := []string{}
	if md != nil {
		if sliceKeys, hasKey := md[key]; hasKey {
			for _, sliceValue := range sliceKeys {
				result = append(result, sliceValue)
			}
		}
	}
	return result
}
