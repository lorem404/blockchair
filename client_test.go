package blockchair

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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

func TestSetRateLimitFunc(t *testing.T) {
	// rl := RateLimit{}
	ratefunc := func(rl RateLimit) {}
	expect := func(c *Client) { c.RateLimitFunc = ratefunc }
	actual := SetRateLimitFunc(ratefunc)
	assert.Equal(t, expect, actual, "SetRateLimitFunc(ratefunc) should return func(c *Client) { c.RateLimitFunc = ratefunc }")
}

// No implementation is provided for this function.
// Since there is no logic to execute, no tests are needed at the moment.
func TestDefaultRateLimitFunc(t *testing.T) {}

func TestPercentageLeft(t *testing.T) {
	rl := RateLimit{Remaining: 2, Limit: 2}
	expect := 100
	actual := rl.PercentageLeft()
	assert.Equal(t, expect, actual, "PercentageLeft() should return 100")
}

func TestWaitTime(t *testing.T) {
	rl := RateLimit{
		Period: 10,
		Limit:  5,
	}
	expect := time.Duration(2) * time.Second
	actual := rl.WaitTime()
	assert.Equal(t, expect, actual, "WaitTime() should return 2s")
}

func TestWaitTimeRemaining(t *testing.T) {
	t.Run("WaitTimeRemaining() should return 2s", func(t *testing.T) {
		rl := RateLimit{
			Remaining: 3,
			Period:    6,
		}
		expect := time.Duration(2) * time.Second
		actual := rl.WaitTimeRemaining()
		assert.Equal(t, expect, actual, "WaitTimeRemaining() should return 2s")
	})
	t.Run("WaitTimeRemaining() should return 6s", func(t *testing.T) {
		rl := RateLimit{
			Remaining: 1,
			Period:    6,
		}
		expect := time.Duration(6) * time.Second
		actual := rl.WaitTimeRemaining()
		assert.Equal(t, expect, actual, "WaitTimeRemaining() should return 6s")
	})
}

func TestRateLimitStrategySleep(t *testing.T)      {}
func TestRateLimitStrategyConcurrent(t *testing.T) {}
func TestParseRate(t *testing.T)                   {}
func TestLoadResponse(t *testing.T)                {}

func TestNew(t *testing.T) {
	expect := &Client{client: &http.Client{}, RateLimitFunc: defaultRateLimitFunc}
	actual := New()
	assert.Equal(t, expect, actual, "New() should return a Client struct")
}

func TestSetClient(t *testing.T) {
	t.Run("set valid client", func(t *testing.T) {
		c := &Client{}
		hc := &http.Client{}
		c.SetClient(hc)

		if c.client != hc {
			t.Errorf("expected client to be set, got %v", c.client)
		}
	})
}
