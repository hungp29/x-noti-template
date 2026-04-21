package telegramtpl

import (
	"testing"

	"github.com/hungp29/x-noti-template/wishn"
)

func TestRenderWishPriceChangeEN(t *testing.T) {
	data := map[string]string{
		wishn.DataProductLinkName:        "Coffee",
		wishn.DataProductLinkURL:         "https://shop.example/p/1",
		wishn.DataNewPrice:               "VND 99000",
		wishn.DataDeltaVsPrevious:        "-1000 VND",
		wishn.DataDeltaVsLinkLowest:      "0 VND",
		wishn.DataIsLowestAmongLinks:     "true",
		wishn.DataDeltaVsCheapestLink:    "",
		wishn.DataCheapestLinkName:      "-",
		wishn.DataCheapestLinkURL:       "-",
		wishn.DataUnit:                  "pack",
		wishn.DataPerUnitComparisonNote: "",
	}
	s, err := Render(wishn.TemplateWishPriceChange, "en", data)
	if err != nil {
		t.Fatal(err)
	}
	if s == "" || len(s) < 20 {
		t.Fatalf("unexpectedly short: %q", s)
	}
}
