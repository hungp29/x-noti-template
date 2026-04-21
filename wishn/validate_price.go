package wishn

import (
	"fmt"
	"strings"
)

// ValidateWishPriceNotifyData checks required keys for wish price-change notifications.
func ValidateWishPriceNotifyData(data map[string]string) error {
	if data == nil {
		data = map[string]string{}
	}
	req := []string{
		DataProductLinkURL,
		DataNewPrice,
		DataDeltaVsPrevious,
		DataDeltaVsLinkLowest,
		DataIsLowestAmongLinks,
	}
	for _, k := range req {
		if strings.TrimSpace(data[k]) == "" {
			return fmt.Errorf("missing %q", k)
		}
	}
	if !strings.EqualFold(strings.TrimSpace(data[DataIsLowestAmongLinks]), "true") {
		for _, k := range []string{DataDeltaVsCheapestLink, DataCheapestLinkName, DataCheapestLinkURL} {
			if strings.TrimSpace(data[k]) == "" {
				return fmt.Errorf("missing %q for non-lowest case", k)
			}
		}
	}
	return nil
}
