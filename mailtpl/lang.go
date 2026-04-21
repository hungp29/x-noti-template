package mailtpl

import "strings"

// NormalizeLanguage returns "vi" when lang is Vietnamese; otherwise "en".
func NormalizeLanguage(lang string) string {
	if strings.EqualFold(strings.TrimSpace(lang), "vi") {
		return "vi"
	}
	return "en"
}
