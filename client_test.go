package blockchair

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestGetSupportedCrypto(t *testing.T) {
	expect := []string{"bitcoin", "bitcoin-cash", "litecoin", "dogecoin", "dash", "groestlcoin", "zcash", "ecash", "bitcoin/testnet"}
	actual := GetSupportedCrypto()
	assert.Equal(t, expect, actual, "GetSupportedCrypto() should return the correct list of supported cryptocurrencies")
}

func TestGetSupportedCryptoEth(t *testing.T) {
	expect := []string{"ethereum/testnet", "ethereum"}
	actual := GetSupportedCryptoEth()
	assert.Equal(t, expect, actual, "GetSupportedCryptoEth() should return the correct list")

}

func TestGetSupportedCryptoMultichaint(t *testing.T) {
	expect := []string{"bitcoin", "bitcoin-cash", "litecoin", "dash", "groestlcoin", "zcash", "ethereum"}
	actual := GetSupportedCryptoMultichain()
	assert.Equal(t, expect, actual, "GetSupportedCryptoMultichain() should return the correct list")

}

// func TestSetRateLimitFunc(t *testing.T) {
func TestUserAgent(t *testing.T) {
	t.Run("c.userAgent() should return blockchair-api-golang-0.1.3 foo", func(t *testing.T) {
		c := Client{UserAgent: "foo"}

		expect := UserAgent + " " + "foo"
		actual := c.userAgent()
		assert.Equal(t, expect, actual, "c.userAgent() should return blockchair-api-golang-0.1.3 foo")
	})

	t.Run("c.userAgent() should return blockchair-api-golang-0.1.3", func(t *testing.T) {
		c := Client{UserAgent: ""}

		expect := UserAgent
		actual := c.userAgent()
		assert.Equal(t, expect, actual, "c.userAgent() should return blockchair-api-golang-0.1.3")
	})
}
func TestSetRateLimitFunc(t *testing.T) {}

func TestDefaultRateLimitFunc(t *testing.T) {}

func TestPercentageLeft(t *testing.T) {
	rl := RateLimit{Remaining: 2, Limit: 2}
	expect := 100
	actual := rl.PercentageLeft()
	assert.Equal(t, expect, actual, "rl.PercentageLeft() should return 100")
}

func TestWaitTime(t *testing.T)                    {}
func TestWaitTimeRemaining(t *testing.T)           {}
func TestRateLimitStrategySleep(t *testing.T)      {}
func TestRateLimitStrategyConcurrent(t *testing.T) {}
func TestParseRate(t *testing.T)                   {}
func TestLoadResponse(t *testing.T)                {}
func TestNew(t *testing.T) {
	expect := &Client{client: &http.Client{}, RateLimitFunc: defaultRateLimitFunc}
	actual := New()
	assert.Equal(t, expect, actual, "New() should return a Client struct")
}
func TestSetClient(t *testing.T) {}
