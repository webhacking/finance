package finance

// RebalancePlan indicates how much of the asset must be purchased or sold
type RebalancePlan struct {
	Asset    Asset
	Quantity float64
}

// NOTE: Not sure if taking *Portfolio as a parameter, but couldn't think of a better way
func (record *PortfolioRecord) CurrentShare(portfolio *Portfolio) float64 {
	nav := portfolio.NetAssetValue()
	return record.Asset.UnitPrice * record.Quantity / nav
}

// Rebalance a given portfolio according to the desired shares of each asset
// TODO: Need an option to determine whether we want to allow non-discrete quantities
func (portfolio *Portfolio) Rebalance() []RebalancePlan {
	plans := make([]RebalancePlan, len(portfolio.Records))
	nav := portfolio.NetAssetValue()
	for i, record := range portfolio.Records {
		unitPrice := record.Asset.UnitPrice
		currentValue := unitPrice * record.Quantity
		targetValue := record.DesiredShare * nav

		plans[i].Asset = record.Asset
		plans[i].Quantity = (targetValue - currentValue) / unitPrice
	}
	return plans
}
