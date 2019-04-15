package finance

type Portfolio struct {
	Records []PortfolioRecord
}

type PortfolioRecord struct {
	Asset        Asset
	Quantity     float64
	DesiredShare float64
}

func (portfolio *Portfolio) NetAssetValue() float64 {
	netAssetValue := 0.0
	for _, record := range portfolio.Records {
		netAssetValue += record.Asset.UnitPrice * record.Quantity
	}
	return netAssetValue
}
