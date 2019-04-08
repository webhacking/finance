package finance

import (
	"fmt"
	"testing"
)

func TestLoadAssetsFromTsv(t *testing.T) {
	assets := LoadAssetsFromTsv("assets.tsv")
	fmt.Println(assets)
}
