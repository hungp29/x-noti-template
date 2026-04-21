package telegramtpl

import (
	"fmt"
	"strings"

	"github.com/hungp29/x-noti-template/wishn"
)

// wishPriceTelegramTmplData is passed to wish_price_change Telegram body.gotmpl files.
type wishPriceTelegramTmplData struct {
	ProductLinkName   string
	ProductLinkURL    string
	NewPrice          string
	DeltaVsPrevious   string
	DeltaVsLinkLowest string
	Cross             string
	NoteExtra         string // optional "\n"+note when non-empty
}

func buildWishPriceTelegramData(lang string, data map[string]string) (wishPriceTelegramTmplData, error) {
	if err := wishn.ValidateWishPriceNotifyData(data); err != nil {
		return wishPriceTelegramTmplData{}, err
	}
	isLow := strings.EqualFold(strings.TrimSpace(data[wishn.DataIsLowestAmongLinks]), "true")
	var cross string
	switch lang {
	case "vi":
		if isLow {
			cross = "Giá mới là thấp nhất trong tất cả liên kết của sản phẩm này."
		} else {
			cross = fmt.Sprintf(
				"Không phải giá thấp nhất giữa các liên kết. Rẻ nhất: %s — %s. Chênh lệch so với liên kết rẻ nhất: %s.",
				data[wishn.DataCheapestLinkName],
				data[wishn.DataCheapestLinkURL],
				nonEmpty(data[wishn.DataDeltaVsCheapestLink], "—"),
			)
		}
	default:
		if isLow {
			cross = "The new price is the lowest among all links for this product."
		} else {
			cross = fmt.Sprintf(
				"Not the lowest across links. Cheapest: %s — %s. Difference vs cheapest link: %s.",
				data[wishn.DataCheapestLinkName],
				data[wishn.DataCheapestLinkURL],
				nonEmpty(data[wishn.DataDeltaVsCheapestLink], "—"),
			)
		}
	}

	note := strings.TrimSpace(data[wishn.DataPerUnitComparisonNote])
	noteExtra := ""
	if note != "" {
		noteExtra = "\n" + note
	}

	return wishPriceTelegramTmplData{
		ProductLinkName:   data[wishn.DataProductLinkName],
		ProductLinkURL:    data[wishn.DataProductLinkURL],
		NewPrice:          data[wishn.DataNewPrice],
		DeltaVsPrevious:   data[wishn.DataDeltaVsPrevious],
		DeltaVsLinkLowest: data[wishn.DataDeltaVsLinkLowest],
		Cross:             cross,
		NoteExtra:         noteExtra,
	}, nil
}

func nonEmpty(s, fallback string) string {
	if strings.TrimSpace(s) == "" {
		return fallback
	}
	return s
}
