package telegramtpl

import (
	"fmt"
	"strings"

	"github.com/hungp29/x-noti-template/wishn"
)

// Render builds the Telegram message body for templateID and language ("vi" or "en").
func Render(templateID, lang string, data map[string]string) (text string, err error) {
	if data == nil {
		data = map[string]string{}
	}
	lang = NormalizeLanguage(lang)
	tid := strings.TrimSpace(templateID)
	switch tid {
	case wishn.TemplateWishPriceChange:
		tmplData, err := buildWishPriceTelegramData(lang, data)
		if err != nil {
			return "", err
		}
		tpl, err := getTelegramTemplate(tid, lang)
		if err != nil {
			return "", err
		}
		return execTelegramTpl(tpl, tmplData)
	default:
		return "", fmt.Errorf("unknown telegram template_id %q", templateID)
	}
}
