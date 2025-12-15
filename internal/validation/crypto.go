package validation

import (
	"regexp"
)

var walletRegex = regexp.MustCompile(
	`^(0x[a-fA-F0-9]{40}|bc1[a-z0-9]{25,39})$`,
)

func IsValidCrypto(wallet string) bool{
	return walletRegex.MatchString(wallet)
}