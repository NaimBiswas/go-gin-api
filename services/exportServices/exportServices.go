package exportServices

import (
	Models "NaimBiswas/go-gin-api/models"
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	fileName := generateFileName(collection.Name(), ".pdf")

	c.Header("Content-Type", "application/pdf")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	// Write the PDF content to the response
	err = ioutil.WriteFile("./files/"+fileName, buf.Bytes(), 0644)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "message": "failed to save PDF to local file"})
		return
	}
	c.Data(http.StatusOK, "application/pdf", buf.Bytes())
}

func ExportToXlsx() {

}

func ExportToCSV(c *gin.Context, collection *mongo.Collection) {
	var result []bson.M
	sortOptions := options.Find()
	sortOptions.SetSort(bson.D{{"modifiedAt", -1}})
	cur, err := collection.Find(c, bson.M{}, sortOptions)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err, "failed to fetch data from the database")
		c.Abort()
		return
	}

	defer cur.Close(c)

	var data [][]string
	var headers map[string]bool

	// Read the MongoDB cursor and extract data
	if err := cur.All(c, &result); err != nil {
		errorResponse(c, http.StatusForbidden, err, "Something went wrong")
		return
	}

	// Extract headers from the first document
	if len(result) > 0 {
		headers = make(map[string]bool)
		for key := range result[len(result)-1] {
			if !(key == "_id" || key == "__v" || key == "password") {
				headers[key] = true
			}
		}
	}
	// Create the CSV header row
	var headerRow []string
	for key := range headers {
		headerRow = append(headerRow, key)
	}
	sort.Strings(headerRow)
	data = append(data, headerRow)

	// Extract values from each document
	for _, doc := range result {
		var values []string
		for key := range doc {
			for i := 0; i < len(headerRow); i++ {
				if key == headerRow[i] {
					if doc[key] == nil {
						values = append(values, fmt.Sprintf("%v", ""))
					} else {
						values = append(values, fmt.Sprintf("%v", doc[headerRow[i]]))
					}
				}
			}
		}
		data = append(data, values)
	}
	// Convert data to CSV
	var csvContent strings.Builder
	writer := csv.NewWriter(&csvContent)
	fileName := generateFileName(collection.Name(), ".csv")
	writer.WriteAll(data)
	writer.Flush()

	csvData := csvContent.String()

	// Set the response headers
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Data(http.StatusOK, "text/csv", []byte(csvData))
}

func generateFileName(name string, extension string) string {
	timestamp := time.Now().Unix()
	formattedTimestamp := strings.ReplaceAll(strings.ReplaceAll(fmt.Sprint(timestamp), " ", "_"), ":", "_")

	fileName := name + "_" + formattedTimestamp
	return fileName + extension
}

func errorResponse(c *gin.Context, statusCode int, error error, message string) {
	c.JSON(statusCode, gin.H{"error": error.Error(), "message": message})
	c.Abort()
}
