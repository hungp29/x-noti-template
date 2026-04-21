package wishn

// TemplateWishPriceChange is used when a wish product link price changes after a scan.
const TemplateWishPriceChange = "wish_price_change"

// Data keys for wish price-change templates (mail + Telegram).
const (
	DataProductLinkName              = "product_link_name"
	DataProductLinkURL               = "product_link_url"
	DataNewPrice                     = "new_price"
	DataDeltaVsPrevious              = "delta_vs_previous"
	DataDeltaVsLinkLowest            = "delta_vs_link_lowest"
	DataIsLowestAmongLinks           = "is_lowest_among_links" // "true" or "false"
	DataDeltaVsCheapestLink         = "delta_vs_cheapest_link" // empty when lowest among links
	DataCheapestLinkName             = "cheapest_link_name"
	DataCheapestLinkURL              = "cheapest_link_url"
	DataUnit                         = "unit"
	DataPerUnitComparisonNote      = "per_unit_comparison_note" // optional; empty when N/A
	DataCrossLinkLowestSummaryLine = "cross_link_lowest_summary" // optional helper line
)
