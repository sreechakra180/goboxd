package scanner

import (
	"strings"

	"github.com/user/blackroot/internal/logger"
)

// Scanner checks code for malicious patterns
type Scanner struct{}

func NewScanner() *Scanner {
	return &Scanner{}
}

// ScanCode runs all applicable regex rules against the user code
func (s *Scanner) ScanCode(language string, code string) []string {
	var alerts []string

	// Check common rules first (like rm -rf)
	for _, r := range commonRules {
		if r.Pattern.MatchString(code) {
			logger.WarnLog.Printf("Threat blocked: %s", r.Name)
			alerts = append(alerts, r.Message)
		}
	}

	// Check language specific rules
	lang := strings.ToLower(language)
	if rules, exists := languageRules[lang]; exists {
		for _, r := range rules {
			if r.Pattern.MatchString(code) {
				logger.WarnLog.Printf("Threat blocked in %s: %s", lang, r.Name)
				alerts = append(alerts, r.Message)
			}
		}
	}

	// quick validation before execution: infinite loop heuristics
	// very hacky but works for some script kiddies
	if strings.Contains(code, "while (true)") || strings.Contains(code, "while True:") || strings.Contains(code, "for {") {
		// we don't necessarily block it, maybe just log it or we could add an alert
		// actually, our execution timeout will catch this anyway.
		// logger.InfoLog.Println("Potential infinite loop detected, relying on timeout")
	}

	return alerts
}
