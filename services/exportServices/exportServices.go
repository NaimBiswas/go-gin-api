package exportServices

import (
	Models "NaimBiswas/go-gin-api/models"
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func ExportToPdf(c *gin.Context, collection *mongo.Collection) {
	var dataArray []bson.M

	response, err := collection.Find(c, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from the database"})
		return
	}
	//var dataArray []bson.M = model
	if err := response.All(c, &dataArray); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch data from the database"})
		return
	}
	//initiate pdf with pdf size and width
	pdf := gofpdf.New("L", "mm", "A3", "")
	pdf.AddPage()

	// Save the PDF to a buffer
	var buf bytes.Buffer
	headers := []string{}
	//style for pdf
	pdf.SetFillColor(27, 229, 252)
	pdf.SetTextColor(64, 64, 64)
	pdf.SetDrawColor(204, 0, 0)
	pdf.SetFont("Arial", "", 12)

	//struct for columns
	userExportType := reflect.TypeOf(Models.UserColumnsForExport{})

	// Generate Pdf headers and header for dataCatch
	for i := 0; i < userExportType.NumField(); i++ {
		field := userExportType.Field(i)
		fieldName := field.Tag.Get("pdfFiled")
		jsonTag := field.Tag.Get("json")
		headers = append(headers, jsonTag)
		// Add the cell to the header row
		pdf.CellFormat(50, 10, fieldName, "1", 0, "CM", true, 0, "")
	}

	//add data as per header
	//const rowsPerPage = 40 // Adjust the number of rows per page as needed
	//rowCounter := 0
	for _, data := range dataArray {
		pdf.Ln(-1)
		for i := 0; i < len(headers); i++ {
			cellValue := data[headers[i]]

			convertHeader := fmt.Sprintf("%v", headers[i])
			convertValue := fmt.Sprintf("%v", cellValue)

			if cellValue == nil || len(convertValue) == 0 || convertValue == "[]" {
				cellValue = "" // Replace nil with an empty string
			} else if convertHeader == "createdAt" || convertHeader == "updatedAt" {
				if val, ok := cellValue.(primitive.DateTime); ok {
					// Convert the primitive.DateTime to a time.Time object
					parsedDate := time.Unix(int64(val)/1000, 0)
					// Format the time.Time object to the desired format "DD/MM/YYYY"
					cellValue = parsedDate.Format("02/01/2006 15:04:05.000")
				}
			}
			pdf.CellFormat(50, 10, fmt.Sprintf("%v", cellValue), "1", 0, "CM", false, 0, "")
		}
		//rowCounter++
		//if rowCounter == rowsPerPage {
		//	pdf.AddPage()
		//	rowCounter = 0
		//	for i := 0; i < len(headers[0:10]); i++ {
		//		if !(headers[i] == "__v" || headers[i] == "_id" || headers[i] == "password") {
		//			pdf.CellFormat(50, 10, headers[i], "1", 0, "CM", true, 0, "")
		//		}
		//	}
		//}
	}

	err = pdf.Output(&buf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate PDF"})
		return
	}

	timestamp := time.Now().Unix()
	formattedTimestamp := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(timestamp), " ", "_"), ":", "_")

	fileName := collection.Name() + "_" + formattedTimestamp
	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+fileName+".pdf")
	// Write the PDF content to the response
	err = ioutil.WriteFile("./files/"+fileName+".pdf", buf.Bytes(), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save PDF to local file"})
		return
	}
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

func ExportToXlsx() {

}

func ExportToCSV() {

}
