package finance

type Portfolio struct {
	Records []PortfolioRecord
}

type PortfolioRecord struct {
	Asset        Asset
	Quantity     Decimal
	DesiredShare Decimal
}

func (portfolio *Portfolio) NetAssetValue() Decimal {
	netAssetValue := Decimal(0)
	for _, record := range portfolio.Records {
		netAssetValue += record.Asset.UnitPrice.Mul(record.Quantity)
	}
	return netAssetValue
}
