package access_token

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestExpirationTimeConstant(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "Expiration time should be 24 hours")
}
func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "brand new access token should not be expired")
	assert.EqualValues(t, at.AccessToken, "", "new access token should not have defineds access token id")
	assert.True(t, at.UserID == 0, "new access token should not have an associated user id")
}

func TestGetAccessTokenIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expiring three hours from now should NOT be expired")

}
