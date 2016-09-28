package portcullis_test

import (
	"fmt"
	"testing"

	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"

	"github.com/fortifi/portcullis-go"
	"github.com/fortifi/portcullis-go/keys"
)

const (
	testOrg      = "this-is-a-test-org-id"
	testUserID   = "this-is-a-test-user-id"
	testUsername = "this-is-a-test-username"

	testAppID  = "test-app-id"
	testVendor = "test-vendor"
)

// TestAuthDataExtraction tests for valid transaction of portcullis meta data values
func TestAuthDataExtraction(t *testing.T) {
	metamap := map[string]string{}
	metamap[keys.GetOrganisationKey()] = testOrg
	metamap[keys.GetUserIDKey()] = testUserID
	metamap[keys.GetUsernameKey()] = testUsername
	metamap[keys.GetAppIDKey()] = testAppID
	metamap[keys.GetAppVendorKey()] = testVendor

	meta := metadata.New(metamap)
	ctx := metadata.NewContext(context.Background(), meta)
	in := portcullis.FromContext(ctx)

	if in.GlobalAppID() != fmt.Sprintf("%s/%s", testVendor, testAppID) {
		t.Error("Global app ID does not contain expected value")
	}

	if in.OrganisationID != testOrg {
		t.Error("Organisation does not contain expected value")
	}

	if in.Username != testUsername {
		t.Error("Username does not contain expected value")
	}

	if in.UserID != testUserID {
		t.Error("userID does not contain expected value")
	}
}

// TestAuthDataExtractionWithMissingFields tests for valid extraction of portcullis meta with missing values
func TestAuthDataExtractionWithMissingFields(t *testing.T) {
	metamap := map[string]string{}
	metamap[keys.GetUsernameKey()] = testUsername
	meta := metadata.New(metamap)
	ctx := metadata.NewContext(context.Background(), meta)

	org := portcullis.FromContext(ctx).OrganisationID
	username := portcullis.FromContext(ctx).Username
	userID := portcullis.FromContext(ctx).UserID

	if username != testUsername {
		t.Error("Username does not contain expected value")
	}

	if org != "" {
		t.Error("Organisation does not contain expected value")
	}

	if userID != "" {
		t.Error("userID does not contain expected value")
	}
}

// TestExtractionWithInvalidContext tests extraction result with context contains no metadata
func TestExtractionWithInvalidContext(t *testing.T) {
	ctx := context.TODO()
	org := portcullis.FromContext(ctx).OrganisationID

	if org != "" {
		t.Error("Organisation does not contain expected value")
	}
}
