// Package license helps manage commercial licenses and check if they are valid for the version of  used.
package license

import _d "github.com/bamzi/pdfext/internal/license"

// SetMeteredKeyPersistentCache sets the metered License API Key persistent cache.
// Default value 'true', set to `false` will report the usage immediately to license server,
// this can be used when there's no access to persistent data storage.
func SetMeteredKeyPersistentCache(val bool) { _d.SetMeteredKeyPersistentCache(val) }

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey() *LicenseKey { return _d.MakeUnlicensedKey() }

// SetLicenseKey sets and validates the license key.
func SetLicenseKey(content string, customerName string) error {
	return _d.SetLicenseKey(content, customerName)
}

// LicenseKey represents a loaded license key.
type LicenseKey = _d.LicenseKey

// SetMeteredKey sets the metered API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey(apiKey string) error { return _d.SetMeteredKey(apiKey) }

const (
	LicenseTierUnlicensed = _d.LicenseTierUnlicensed
	LicenseTierCommunity  = _d.LicenseTierCommunity
	LicenseTierIndividual = _d.LicenseTierIndividual
	LicenseTierBusiness   = _d.LicenseTierBusiness
)

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState() (_d.MeteredStatus, error) { return _d.GetMeteredState() }

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey() *LicenseKey                { return _d.GetLicenseKey() }
func SetMeteredKeyUsageLogVerboseMode(val bool) { _d.SetMeteredKeyUsageLogVerboseMode(val) }
