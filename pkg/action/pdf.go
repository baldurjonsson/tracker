package action

import (
	"fmt"
	"strconv"

	"github.com/baldurjonsson/tracker/pkg/store"
	"github.com/go-pdf/fpdf"
	"github.com/urfave/cli/v2"
	"golang.org/x/text/encoding/charmap"
)

var encoder = charmap.ISO8859_1.NewEncoder()

func convertToISO88591(s string) string {
	str, _ := encoder.String(s)
	return str
}

func getTimesheet(c *cli.Context) (*store.Timesheet, error) {
	s := c.Context.Value("store").(*store.Store)
	id, err := strconv.Atoi(c.Args().First())
	if err != nil {
		return nil, err
	}
	for _, t := range s.Timesheets.Timesheets {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, fmt.Errorf("timesheet with id %d not found", id)
}
func RenderTimesheetPDF(c *cli.Context) error {
	timesheet, err := getTimesheet(c)
	if err != nil {
		return err
	}

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "B", 12)
		pdf.Cell(170, 10, convertToISO88591(timesheet.Name))
		pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(0, 10, fmt.Sprint(pdf.PageNo()), "", 0, "R", false, 0, "")
		pdf.Ln(-1)
		pdf.SetFont("Arial", "", 12)
		pdf.Cell(20, 6, "Nafn:")
		pdf.Cell(140, 6, convertToISO88591(timesheet.Profile.Name))
		pdf.Ln(-1)
		pdf.Cell(20, 6, "Netfang:")
		pdf.Cell(140, 6, convertToISO88591(timesheet.Profile.Email))
		pdf.Cell(0, 10, "")
		pdf.Ln(-1)
	})
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(30, 7, "Project", "1", 0, "", false, 0, "")
	pdf.CellFormat(15, 7, "Hours", "1", 0, "", false, 0, "")
	pdf.CellFormat(25, 7, "Date", "1", 0, "", false, 0, "")
	pdf.CellFormat(110, 7, "Notes", "1", 0, "", false, 0, "")
	pdf.Ln(-1)
	pdf.SetFont("Arial", "", 10)
	for _, e := range timesheet.Entries {
		pdf.CellFormat(30, 6, convertToISO88591(e.Project), "1", 0, "", false, 0, "")
		pdf.CellFormat(15, 6, fmt.Sprintf("%.1f", e.Hours), "1", 0, "", false, 0, "")
		pdf.CellFormat(25, 6, e.Date.Format("2006-01-02"), "1", 0, "", false, 0, "")
		pdf.CellFormat(110, 6, convertToISO88591(e.Notes), "1", 0, "", false, 0, "")
		pdf.Ln(-1)
	}

	pdf.OutputFileAndClose(c.String("filename"))
	return nil
}
