package main

import (
	"fmt"
	"io"
	"log"
	"time"
	"net/smtp"
	"net/http"
	"strings"
	"strconv"
)

func main () {
	
	for {
        // Perform the task
        fmt.Println("Running task at:", time.Now())
        
	getPrice("GOOG", ":", "NASDAQ", 200.0)
	getPrice("AAPL", ":", "NASDAQ", 140.7)
	getPrice("XMR", "-", "CAD", 100.0)

	// Sleep for a minute
        time.Sleep(1 * time.Minute)
    	
	}	
}

//get price of stock using google finanace NASDAQ, I guess you could change it to whatever you like tho.
func getPrice(stock, seperator, exchange string, myPrice float64) {
	// URL of the page to scrape

	//in the google finance url there is a : between the exchange, nasdaq, and the stock, aapl
	//however, for crypto that : is changed to a -
	//so make the seperator var one of those

	url := "https://www.google.com/finance/quote/" + stock + seperator + exchange // Replace with your desired URL

	// Fetch the HTML source code
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch the URL: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body into a string
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the response body: %v", err)
	}
	htmlContent := string(bodyBytes)

	// Define the start and end markers
	startTag := `<div class="YMlKec fxKbKc">`
	endTag := `</div>`

	// Find the starting position of the start tag
	startIndex := strings.Index(htmlContent, startTag)
	if startIndex == -1 {
		fmt.Println("Start tag not found!")
		return
	}

	// Move the startIndex past the start tag
	startIndex += len(startTag)

	// Find the position of the end tag after the start tag
	endIndex := strings.Index(htmlContent[startIndex:], endTag)
	if endIndex == -1 {
		fmt.Println("End tag not found!")
		return
	}

	// Extract the content between the tags
	extractedText := htmlContent[startIndex : startIndex+endIndex]
	isPriceCompatible(stock, extractedText, myPrice)
	//fmt.Printf("Type: %T", extractedText)
	fmt.Println("Extracted Data:", stock, "on", exchange, "is", extractedText)
	fmt.Println("You want", stock, "for:", myPrice)
}

//f is the float which is now an int.
func stringToFloat(numberString string) float64 {

        // Convert string to float64
        number, err := strconv.ParseFloat(numberString, 64)
        if err != nil {
                log.Fatalf("Error converting string to float64: %v", err)
        }

        // Print the result
        return number

}

//comapre the current price of stock to the price i'd like to buy it for.
func isPriceCompatible(stock, extractedText string, myPrice float64) {
	
	price := stringToFloat(strings.ReplaceAll(extractedText, "$", ""))
	//fmt.Printf("Type: %T, Value: %v", price, price)	
	//if my price is greater than or equal to the price of the stock. send an email.
	if price <= myPrice {
		fmt.Println("stock price is:", price, "I want it for:", myPrice)
		msg := "Buy" + stock + "now" 
		email("botEmailHere@Email.com", "AppPasswordGoesHere", msg, "YourEmailHere@Email.com")

	}
}

func email(from, password, msg, victim string) {

    toList := []string{victim}

    // mail change the address as smtp.mail.yahoo.com
    host := "smtp.gmail.com"

    // Its the default port of smtp server
    port := "587"

    // strings need to be converted into slice bytes
    body := []byte(msg)

    auth := smtp.PlainAuth("", from, password, host)

    err := smtp.SendMail(host+":"+port, auth, from, toList, body)

    // handling the errors
    check(err)

    fmt.Println("Successfully sent mail to all user in toList")
}

func check(err error) {
	
	if err != nil {
                fmt.Println("Error:", err)
		
        }
	

}
