package finance

// Asset represents investment vehicles such as stocks, bonds, funds, deposits, options, futures, and such.
type Asset struct {
	Name      string
	Code      string
	Type      AssetType
	UnitPrice float64
}

// AssetType is a type of an investment vehicle
type AssetType int

const (
	Unknown AssetType = 0
	Cash    AssetType = 1
	Bond    AssetType = 2
	Stock   AssetType = 3
	Fund    AssetType = 4
	Deposit AssetType = 5
)

// LoadAssetsFromTsv loads assets from a .tsv (tab separated values) file
func LoadAssetsFromTsv(filename string) {

}
