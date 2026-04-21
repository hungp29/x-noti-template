package telegramtpl

import "strings"

// NormalizeLanguage returns "vi" for Vietnamese; otherwise "en".
func NormalizeLanguage(lang string) string {
	if strings.EqualFold(strings.TrimSpace(lang), "vi") {
		return "vi"
	}
	return "en"
}
