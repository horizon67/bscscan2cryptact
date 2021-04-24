package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jszwec/csvutil"
)

type Transaction struct {
	CreatedAt string  `csv:"CreatedAt"`
	Action    string  `csv:"Action"`
	Source    string  `csv:"Source"`
	Base      string  `csv:"Base"`
	Volume    float64 `csv:"Volume"`
	Price     float64 `csv:"Price"`
	Counter   string  `csv:"Counter"`
	Fee       float64 `csv:"Fee"`
	FeeCcy    string  `csv:"FeeCcy"`
}

func NewTransaction() *Transaction {
	t := new(Transaction)
	t.Action = "SELL"
	t.Source = "PancakeSwap"
	t.FeeCcy = "USD"
	return t
}

func main() {
	transaction := NewTransaction()
	var transactions []Transaction

	flag.Parse()
	doc, err := goquery.NewDocument("https://bscscan.com/tx/" + flag.Args()[0])
	if err != nil {
		panic(err)
	}
	// Volume, Price
	var values []string
	selection := doc.Find("ul#wrapperContent > li > div.media-body > span.mr-1.d-inline-block")
	selection.Each(func(index int, s *goquery.Selection) {
		values = append(values, s.Text())
	})
	transaction.Volume, err = strconv.ParseFloat(values[1], 64)
	if err != nil {
		panic(err)
	}
	totalPrice, err := strconv.ParseFloat(values[len(values)-3], 64)
	if err != nil {
		panic(err)
	}
	transaction.Price = totalPrice / transaction.Volume

	// CreatedAt
	selection = doc.Find("span#clock").Parent()
	rep := regexp.MustCompile(`\((.+?)\)`)
	created_at := rep.FindStringSubmatch(selection.Text())[1]
	t, _ := time.Parse("Jan-2-2006 03:04:05 PM +MST", created_at)
	transaction.CreatedAt = t.In(time.FixedZone("Asia/Tokyo", 9*60*60)).Format("2006/1/2 15:04:05")

	// Base, Counter
	var symbols []string
	selection = doc.Find("ul#wrapperContent > li")
	selection.Each(func(index int, s *goquery.Selection) {
		symbols = append(symbols, s.Find("a.d-inline-block").Text())
	})
	transaction.Base = symbols[0]
	transaction.Counter = symbols[1]

	// Fee
	selection = doc.Find("span#ContentPlaceHolder1_spanTxFee")
	rep = regexp.MustCompile(`\(\$(.+?)\)`)
	fee := rep.FindStringSubmatch(selection.Text())[1]
	transaction.Fee, err = strconv.ParseFloat(fee, 64)
	if err != nil {
		panic(err)
	}
	transactions = append(transactions, *transaction)
	b, err := csvutil.Marshal(transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
