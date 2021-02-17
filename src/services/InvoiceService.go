package services

import (
	"github.com/jung-kurt/gofpdf"
	"log"
	"macuka-backend/src/database"
	"macuka-backend/src/dto"
	"macuka-backend/src/models"
	"macuka-backend/src/util"
)

func CreateInvoice(invoiceDto dto.InvoiceDto) (models.Invoice, error) {
	db := database.GetDatabase()
	date, err := util.ConvertDate(invoiceDto.Date)
	if err != nil {
		return models.Invoice{}, err
	}
	invoice := models.Invoice{
		Customer: invoiceDto.Customer,
		Amount:   invoiceDto.Amount,
		Date:     *date,
	}
	db.Create(&invoice)
	log.Print("Added invoice {}", invoice)
	return invoice, nil
}

func ExportInvoice() {
	const (
		headerName    = "PODUZEĆE ZA GRAĐENJE, PRIJEVOZ, TRGOVINU I USLUGE"
		headerContact = "Videti 121/f, 52404 Sveti Petar u Šumi\ntel: 052/686 - 440; fax: 052/686 - 550\nmob: 098305698\nžiro-račun: Erste bank 2402006 - 1100072567\ne-mail: macuka@pu.t-com.hr\nIBAN: HR7024020061100072567; SWIFT: ESBCHR22\nMBS: 040140254; POR.BR:1415476;"
		headerOIB     = "OIB: 00645636377"
	)
	pdf := gofpdf.New("P", gofpdf.UnitPoint, gofpdf.PageSizeA4, "")
	pdf.AddPage()
	pageWidth, _ := pdf.GetPageSize()
	pdf.SetLeftMargin(0.05 * pageWidth)
	pdf.SetRightMargin(0.05 * pageWidth)

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
	pdf.MultiCell(0, 1.1*h, "aaaa\nh", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	//Invoice info
	pdf.SetFont("arial", "b", 11)
	_, h = pdf.GetFontSize()
	pdf.MultiCell(0, 30, "RAČUN R1 BROJ:a", gofpdf.BorderNone, gofpdf.AlignLeft, false)
	pdf.SetFont("arial", "", 10)
	_, h = pdf.GetFontSize()
	pdf.MultiCell(0, 1.2*h, "DATUM I VRIJEME:A\nMJESTO:a\nVALUTA PLAĆANJA:1\nNAČUN PLAĆANJA:TRANSAKCIJSKI RAČUN", gofpdf.BorderNone, gofpdf.AlignLeft, false)

	err := pdf.OutputFileAndClose("file.pdf")
	if err != nil {
		log.Print(err)
	}
}
