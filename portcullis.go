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

	"github.com/cubex/portcullis-go/keys"
)

// ReqInfo is the structure for deserialised request information
type ReqInfo struct {
	OrganisationID string
	UserID         string
	AppID          string
	VendorID       string
	Username       string
	FirstName      string
	LastName       string
	signature      string
	meta           metadata.MD
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

// FromContext retrieves request info from given request context
func FromContext(ctx context.Context) ReqInfo {
	md, _ := metadata.FromContext(ctx)
	res := ReqInfo{
		OrganisationID: safeGetMetaValString(keys.GetOrganisationKey(), md),
		UserID:         safeGetMetaValString(keys.GetUserIDKey(), md),
		Username:       safeGetMetaValString(keys.GetUsernameKey(), md),
		FirstName:      safeGetMetaValString(keys.GetFirstNameKey(), md),
		LastName:       safeGetMetaValString(keys.GetLastNameKey(), md),
		AppID:          safeGetMetaValString(keys.GetAppIDKey(), md),
		VendorID:       safeGetMetaValString(keys.GetAppVendorKey(), md),
		signature:      safeGetMetaValString(keys.GetSignatureKey(), md),
		meta:           md,
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
