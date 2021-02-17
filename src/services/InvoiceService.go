package services

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
	"macuka-backend/src/util"
	"net/http"
	"strconv"
)

func CreateInvoice(invoiceDto dto.InvoiceDto) (models.Invoice, error) {
	db := database.GetDatabase()
	date, err := util.ConvertDate(invoiceDto.Date)
	if err != nil {
		return models.Invoice{}, err
	}
	currencyDate, err := util.ConvertDate(invoiceDto.CurrencyDate)
	if err != nil {
		return models.Invoice{}, err
	}
	invoice := models.Invoice{
		InvoiceNumber: invoiceDto.InvoiceNumber,
		Customer:      invoiceDto.Customer,
		Date:          *date,
		CurrencyDate:  *currencyDate,
		CallingNumber: invoiceDto.CallingNumber,
		HasVAT:        invoiceDto.HasVAT,
	}
	db.Create(&invoice)
	log.Print("Added invoice {}", invoice)
	return invoice, nil
}

func getLocations(width float64) []float64 {
	items := make([]float64, 0)
	items = append(items, width*0.05)
	items = append(items, width*0.10)
	items = append(items, width*0.45)
	items = append(items, width*0.55)
	items = append(items, width*0.65)
	items = append(items, width*0.85)
	items = append(items, width*0.95)
	return items
}

func ExportInvoice(id string, writer http.ResponseWriter) string {
	const (
		headerName    = "PODUZEĆE ZA GRAĐENJE, PRIJEVOZ, TRGOVINU I USLUGE"
		headerContact = "Videti 121/f, 52404 Sveti Petar u Šumi\ntel: 052/686 - 440; fax: 052/686 - 550\nmob: 098305698\nžiro-račun: Erste bank 2402006 - 1100072567\ne-mail: macuka@pu.t-com.hr\nIBAN: HR7024020061100072567; SWIFT: ESBCHR22\nMBS: 040140254; POR.BR:1415476;"
		headerOIB     = "OIB: 00645636377"
	)

	db := database.GetDatabase()
	var invoices []models.Invoice
	db.Where("id=", id).Find(&invoices)
	if len(invoices) == 0 {
		return ""
	}

	invoice := invoices[0]
	var customers []models.Customer
	db.Where("id=?", invoice.Customer).Find(&customers)
	customer := customers[0]

	tableHeaders := []string{
		"RB",
		"ROBA/USLUGA",
		"JM",
		"KOL",
		"CIJENA",
		"IZNOS",
	}
	pdf := gofpdf.New("P", gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.AddPage()
	pdf.SetCompression(true)
	pageWidth, pageHeight := pdf.GetPageSize()

	//HEADER
	pdf.SetFont("times", "", 12)
	pdf.SetTextColor(100, 100, 100)
	_, h := pdf.GetFontSize()
	pdf.ImageOptions("/home/mihael/Desktop/481712213.jpg", 0.05*pageWidth, h, 150, 0, false, gofpdf.ImageOptions{
		ReadDpi: true,
	}, 0, "")

	pdf.MultiCell(0, h, headerName, gofpdf.BorderNone, gofpdf.AlignRight, false)
	pdf.SetFont("times", "", 10)
	_, h = pdf.GetFontSize()
	pdf.MultiCell(0, h, headerContact, gofpdf.BorderNone, gofpdf.AlignRight, false)
	pdf.SetFont("times", "b", 10)
	pdf.MultiCell(0, h, headerOIB, gofpdf.BorderNone, gofpdf.AlignRight, false)

	pdf.SetDrawColor(100, 100, 100)
	pdf.SetFillColor(100, 100, 100)
	pdf.Line(0.05*pageWidth, 125, 0.95*pageWidth, 125)

	//INFO
	pdf.SetFont("arial", "B", 12)
	pdf.MultiCell(0, 50, "PRIMATELJ:", gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.SetFont("arial", "", 10)
	_, h = pdf.GetFontSize()
	receiver := customer.Name + "\n" + customer.Address + "\n" + customer.City + "\n" + strconv.Itoa(int(customer.PostalNumber))
	pdf.MultiCell(0, 1.1*h, receiver, gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Invoice info
	pdf.SetFont("arial", "b", 11)
	_, h = pdf.GetFontSize()
	pdf.MultiCell(0, 30, "RAČUN R1 BROJ: "+invoice.InvoiceNumber, gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.SetFont("arial", "", 10)
	w, h := pdf.GetFontSize()
	pdf.MultiCell(0, 1.2*h, fmt.Sprintf("DATUM I VRIJEME:%s\nMJESTO:Sveti Petar u Šumi\nVALUTA PLAĆANJA:%s\nNAČUN PLAĆANJA:TRANSAKCIJSKI RAČUN", invoice.Date, invoice.CurrencyDate), gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Table
	currentY := 350.
	pdf.Line(0.05*pageWidth, currentY, 0.95*pageWidth, currentY)
	locations := getLocations(pageWidth)
	pdf.SetFont("arial", "b", 11)
	for i := 0; i < len(locations)-1; i++ {
		middle := locations[i] + (locations[i+1]-locations[i])/2
		halfLength := int64((len(tableHeaders[i])) / 2)
		pdf.Text(middle-float64(halfLength)*w, currentY-1, tableHeaders[i])
	}

	pdf.SetFont("arial", "", 11)

	//First row
	currentY += h
	for i := 0; i < len(locations)-1; i++ {
		pdf.MoveTo(locations[i], currentY)
		pdf.MultiCell(0, 0, "a", gofpdf.AlignMiddle, gofpdf.AlignMiddle, false)
	}
	currentY += h
	pdf.Line(0.05*pageWidth, currentY, 0.95*pageWidth, currentY)

	for i := 0; i < len(locations); i++ {
		pdf.Line(locations[i], 350, locations[i], currentY)
	}

	pdf.Line(0.05*pageWidth, currentY, 0.95*pageWidth, currentY)

	//Payment
	pdf.MoveTo(0, currentY+30)
	pdf.MultiCell(0, h, "aaab\naaaaaaab\naaaab", gofpdf.BorderNone, gofpdf.AlignRight, false)

	pdf.SetFont("arial", "b", 11)
	pdf.MultiCell(0, h+30, "Poziv na broj: "+invoice.CallingNumber, gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Footer
	_, h = pdf.GetFontSize()
	pdf.MultiCell(0, h, "Porez na dodanu vrijednost se ne obračunava sukladno čl.75.st.3.Zakona o PDV-u.", gofpdf.BorderNone, gofpdf.AlignLeft, false)
	if !invoice.HasVAT {
		pdf.MultiCell(0, h, "Prijenos porezne obveze na investitora – REVERSE CHARGE", gofpdf.BorderNone, gofpdf.AlignLeft, false)
	}

	pdf.SetFont("arial", "", 11)
	w, h = pdf.GetFontSize()
	pdf.MultiCell(0, h, "U slučaju neplaćanja u roku zaračunavamo zakonsku zateznu kamatu.", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	pdf.Text(0.07*pageWidth, 0.8*pageHeight, "Hiacinta Macuka")
	pdf.Text(0.75*pageWidth, 0.8*pageHeight, "Vinko Macuka d.i.g")
	pdf.Text((pageWidth/2)-float64(len("MP"))/2*w, 0.9*pageHeight, "MP")

	err := pdf.Output(writer)
	if err != nil {
		log.Print(err)
		return ""
	}
	return invoice.InvoiceNumber + ".pdf"
}
