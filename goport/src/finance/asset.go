package finance

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/korean"
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

// e.g., http://www.fundnuri.com/fr/fr_ovw.asp?FUND_CD=K55232BU5747
// TODO: Move this function elsewhere
func FetchFundInfo(code string) (float64, error) {
	url := fmt.Sprintf("http://www.fundnuri.com/fr/fr_ovw.asp?FUND_CD=%s", code)
	headers := make(map[string]string)
	params := make(map[string]string)

	resp, error := Fetch(url, headers, params)
	if error != nil {
		log.Fatal(error)
	}
	defer resp.Body.Close()

	reader := ForceConvertEncoding(resp.Body, korean.EUCKR)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}

	// FIXME: Refactoring needed
	var value string
	r, _ := regexp.Compile("[0-9,.]+")
	elements := doc.Find("td span")
	elements.Each(func(i int, s *goquery.Selection) {
		text := strings.Trim(s.Text(), " ")
		if text == "기준가" {
			node := s.Parent().Parent().Next().Children()
			value = r.FindString(node.Text())
		}
	})

	return strconv.ParseFloat(value, 64)
}
