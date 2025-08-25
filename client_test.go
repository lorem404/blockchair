package blockchair

import (
	"testing"

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
func TestSetRateLimitFunc(t *testing.T)            {}
func TestUserAgent(t *testing.T)                   {}
func TestDefaultRateLimitFunc(t *testing.T)        {}
func TestPercentageLeft(t *testing.T)              {}
func TestWaitTime(t *testing.T)                    {}
func TestWaitTimeRemaining(t *testing.T)           {}
func TestRateLimitStrategySleep(t *testing.T)      {}
func TestRateLimitStrategyConcurrent(t *testing.T) {}
func TestParseRate(t *testing.T)                   {}
func TestLoadResponse(t *testing.T)                {}
func TestNew(t *testing.T)                         {}
func TestSetClient(t *testing.T)                   {}
