package finance

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

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

// GetReaderFromFile returns a reader object from a file
func GetReaderFromFile(filename string) io.Reader {
	file, _ := os.Open(filename)
	return bufio.NewReader(file)
}

// LoadAssetsFromTsv loads assets from a .tsv (tab separated values) file
func LoadAssetsFromTsv(filename string) []Asset {
	reader := csv.NewReader(GetReaderFromFile(filename))
	reader.Comma = rune('\t')

	var assets []Asset
	for {
		cols, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatalln(error)
		}

		asset := Asset{
			Name: cols[0],
			Code: cols[1],
			Type: Fund, // TODO: Parse string to make an enum
		}
		assets = append(assets, asset)
	}

	return assets
}
