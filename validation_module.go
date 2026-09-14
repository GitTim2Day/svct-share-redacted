package main

import (
	"fmt"
	"strings"
)

// SVCT Validation Module — fail-closed constants.
// Any value not in the known set is rejected. No guessing, no sampling, no drift.

const (
	GmailPrimary   = "timnorman730@gmail.com"
	GitHubUser     = "GitTim2Day"
	YahooSecondary = "timothy_h_norman@yahoo.com"
)

// KnownEmails is the canonical set. Only these addresses may be used for sends.
var KnownEmails = map[string]bool{
	GmailPrimary:   true,
	YahooSecondary: true,
}

// KnownGitHubUsers is the canonical set for repo owners/usernames.
var KnownGitHubUsers = map[string]bool{
	GitHubUser: true,
}

// ValidateEmail trims, lowercases, and checks against the known set.
// Returns error if the address is not exactly one of the locked constants.
func ValidateEmail(addr string) error {
	normalized := strings.ToLower(strings.TrimSpace(addr))
	if !KnownEmails[normalized] {
		return fmt.Errorf("unknown email address rejected: %q (not in locked constant set)", addr)
	}
	return nil
}

// ValidateGitHubUser checks against the locked GitHub username set.
func ValidateGitHubUser(user string) error {
	normalized := strings.TrimSpace(user)
	if !KnownGitHubUsers[normalized] {
		return fmt.Errorf("unknown GitHub user rejected: %q (not in locked constant set)", user)
	}
	return nil
}

// MustValidateEmail panics on failure — use only where a bad address is fatal.
func MustValidateEmail(addr string) {
	if err := ValidateEmail(addr); err != nil {
		panic(err)
	}
}

func main() {
	// Self-test: locked constants pass, all known bad variants fail.
	tests := []struct {
		name    string
		addr    string
		wantErr bool
	}{
		{"primary exact", "timnorman730@gmail.com", false},
		{"primary spaced", "  timnorman730@gmail.com  ", false},
		{"primary upper", "TIMNORMAN730@gmail.com", false},
		{"yahoo exact", "timothy_h_norman@yahoo.com", false},
		{"yahoo spaced", " timothy_h_norman@yahoo.com ", false},
		{"bad gett2day", "gett2day@gmail.com", true},
		{"bad gittim2day", "gittim2day@gmail.com", true},
		{"bad timothy s", "timothys_h_norman@yahoo.com", true},
		{"bad timothy prefix", "timothy01775634@gmail.com", true},
		{"empty", "", true},
	}

	passed := 0
	for _, tc := range tests {
		err := ValidateEmail(tc.addr)
		if (err != nil) == tc.wantErr {
			passed++
			fmt.Printf("PASS  %s\n", tc.name)
		} else {
			fmt.Printf("FAIL  %s  err=%v\n", tc.name, err)
		}
	}
	fmt.Printf("\n%d/%d tests passed\n", passed, len(tests))

	// GitHub user check
	if err := ValidateGitHubUser("GitTim2Day"); err != nil {
		fmt.Printf("FAIL  GitHub user: %v\n", err)
	} else {
		fmt.Println("PASS  GitHub user GitTim2Day")
	}
}