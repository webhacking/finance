package finance

import (
	"fmt"
	"testing"
)

func TestLoadAssetsFromTsv(t *testing.T) {
	assets := LoadAssetsFromTsv("assets.tsv")
	fmt.Println(assets)
}

func TestFetchFundInfo(t *testing.T) {
	value, _ := FetchFundInfo("K55216BY4761")
	fmt.Println(value)
	// expected := 881.62
	// if value != expected {
	// 	assertEquals(t, expected, value, "Incorrect value")
	// }
}
