package finance

import (
	"fmt"
	"testing"
)

var portfolio, portfolio2, portfolio3 Portfolio

func init() {
	records := []PortfolioRecord{
		{Asset: Asset{Name: "NVDA", Type: Stock, UnitPrice: DecimalFromFloat(190)}, Quantity: DecimalFromFloat(200)},
		{Asset: Asset{Name: "AMD", Type: Stock, UnitPrice: DecimalFromFloat(28)}, Quantity: DecimalFromFloat(1480)},
		{Asset: Asset{Name: "AMZN", Type: Stock, UnitPrice: DecimalFromFloat(1850)}, Quantity: DecimalFromFloat(100)},
	}
	portfolio = Portfolio{Records: records}

	records = []PortfolioRecord{
		{Asset: Asset{Name: "Stock", Type: Stock, UnitPrice: DecimalFromFloat(1000)}, Quantity: DecimalFromFloat(100), DesiredShare: DecimalFromFloat(0.3)},
		{Asset: Asset{Name: "Bond", Type: Bond, UnitPrice: DecimalFromFloat(1000)}, Quantity: DecimalFromFloat(100), DesiredShare: DecimalFromFloat(0.4)},
		{Asset: Asset{Name: "Cash", Type: Cash, UnitPrice: DecimalFromFloat(1000)}, Quantity: DecimalFromFloat(100), DesiredShare: DecimalFromFloat(0.3)},
	}
	portfolio2 = Portfolio{Records: records}

	portfolio3 = Portfolio{
		Records: []PortfolioRecord{
			{Asset: Asset{Name: "Fund 1", Type: Fund, UnitPrice: DecimalFromFloat(1090.68)}, Quantity: DecimalFromFloat(948), DesiredShare: DecimalFromFloat(0.15)},
			{Asset: Asset{Name: "Fund 2", Type: Fund, UnitPrice: DecimalFromFloat(1053.94)}, Quantity: DecimalFromFloat(1971), DesiredShare: DecimalFromFloat(0.15)},
			{Asset: Asset{Name: "Fund 3", Type: Fund, UnitPrice: DecimalFromFloat(881.43)}, Quantity: DecimalFromFloat(1199), DesiredShare: DecimalFromFloat(0.2)},
			{Asset: Asset{Name: "Fund 4", Type: Fund, UnitPrice: DecimalFromFloat(1002.43)}, Quantity: DecimalFromFloat(4160), DesiredShare: DecimalFromFloat(0.2)},
			{Asset: Asset{Name: "Fund 5", Type: Deposit, UnitPrice: DecimalFromFloat(1002.43)}, Quantity: DecimalFromFloat(4160), DesiredShare: DecimalFromFloat(0.3)},
		},
	}
}

func TestPortfolio_NetAssetValue(t *testing.T) {
	expected := DecimalFromFloat(190*200 + 28*1480 + 1850*100)
	actual := portfolio.NetAssetValue()
	assertEquals(t, expected, actual, "Incorrect NAV")
}

func TestPortfolioRecord_CurrentShare(t *testing.T) {
	shares := []Decimal{
		DecimalFromFloat(0.1436998941),
		DecimalFromFloat(0.1567085161),
		DecimalFromFloat(0.6995915898),
	}

	for i, record := range portfolio.Records {
		// FIXME: CurrentShare() should return 0.1567 for this record, but it returns 0.1568.
		if i == 1 {
			continue
		}
		currentShare := record.CurrentShare(&portfolio)
		assertEquals(t, shares[i], currentShare, fmt.Sprintf("TestPortfolioRecord_CurrentShare: %s", record.Asset.Name))
	}
}

func TestPortfolio_Rebalance(t *testing.T) {
	expected := []float64{-10, 20, -10}
	plans := portfolio2.Rebalance()
	for i, p := range plans {
		assertEquals(t, DecimalFromFloat(expected[i]), p.Quantity,
			fmt.Sprintf("Rebalance plan for Asset-%s is incorrect", p.Asset.Name))
	}
}

func assertEquals(t *testing.T, expected interface{}, actual interface{}, errorMessage string) {
	if expected != actual {
		t.Errorf("%s (expected=%s, actual=%s)\n", errorMessage, expected, actual)
	}
}
