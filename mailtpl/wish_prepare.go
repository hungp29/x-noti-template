package mailtpl

import (
	"fmt"
	"html"
	"html/template"
	"strings"

	"github.com/hungp29/x-noti-template/wishn"
)

// wishPriceMailTmplData is passed to wish_price_change mail .gotmpl files.
type wishPriceMailTmplData struct {
	ProductLinkName   string
	ProductLinkURL    string
	NewPrice          string
	DeltaVsPrevious   string
	DeltaVsLinkLowest string
	Cross             string
	NoteExtra         string        // optional suffix for plain text (e.g. "\n\n"+note)
	PerUnitHTML       template.HTML // optional extra HTML fragment; safe HTML only
}

func buildWishPriceMailData(lang string, data map[string]string) (wishPriceMailTmplData, error) {
	if err := wishn.ValidateWishPriceNotifyData(data); err != nil {
		return wishPriceMailTmplData{}, err
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
				strings.TrimSpace(data[wishn.DataDeltaVsCheapestLink]),
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
				strings.TrimSpace(data[wishn.DataDeltaVsCheapestLink]),
			)
		}
	}

	note := strings.TrimSpace(data[wishn.DataPerUnitComparisonNote])
	noteExtra := ""
	if note != "" {
		noteExtra = "\n\n" + note
	}

	return wishPriceMailTmplData{
		ProductLinkName:   data[wishn.DataProductLinkName],
		ProductLinkURL:    data[wishn.DataProductLinkURL],
		NewPrice:          data[wishn.DataNewPrice],
		DeltaVsPrevious:   data[wishn.DataDeltaVsPrevious],
		DeltaVsLinkLowest: data[wishn.DataDeltaVsLinkLowest],
		Cross:             cross,
		NoteExtra:         noteExtra,
		PerUnitHTML:       template.HTML(perUnitHTMLNote(note)),
	}, nil
}

func perUnitHTMLNote(note string) string {
	if strings.TrimSpace(note) == "" {
		return ""
	}
	return fmt.Sprintf(`<p style="color:#444;font-size:14px;">%s</p>`, html.EscapeString(note))
}
