package finance

import (
	"fmt"
	"log"
	"testing"
)

func TestLoadAssetsFromTsv(t *testing.T) {
	assets := LoadAssetsFromTsv("assets.tsv")

	for _, asset := range assets {
		value, error := FetchFundInfo(asset.Code)
		if error != nil {
			log.Println(error)
		}
		fmt.Printf("%s: %f\n", asset.Code, value)
	}
}

func TestFetchFundInfo(t *testing.T) {
	value, _ := FetchFundInfo("K55216BY4761")
	fmt.Println(value)
	// expected := 881.62
	// if value != expected {
	// 	assertEquals(t, expected, value, "Incorrect value")
	// }
}
