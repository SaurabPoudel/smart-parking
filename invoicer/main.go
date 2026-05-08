package main

import (
	"fmt"
	"log"
)

func main() {
	invoicer := NewInvoiceService()
	fmt.Println("Invoice service started")

	// Example usage
	invoices := invoicer.GetAllInvoices()
	log.Printf("Total invoices: %d", len(invoices))
}
