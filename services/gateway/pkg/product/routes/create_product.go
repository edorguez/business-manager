package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	companyClient "github.com/edorguez/business-manager/services/gateway/pkg/company/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	productClient "github.com/edorguez/business-manager/services/gateway/pkg/product/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/contracts"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/file_validator"
)

func CreateProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  CreateProduct")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value("keyProductCreate").(contracts.CreateProductRequest)

	var images [][]byte

	// Get the files
	files := r.MultipartForm.File["files"]
	for _, fileHeader := range files {
		// Open the file
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Read a small portion of the file to determine the content type
		buffer := make([]byte, 512)
		_, err = file.Read(buffer)
		if err != nil && err != io.EOF {
			http.Error(w, "Failed to read file", http.StatusInternalServerError)
			return
		}

		// Reset the file pointer
		file.Seek(0, io.SeekStart)

		// Detect the content type
		contentType := http.DetectContentType(buffer)
		if !file_validator.IsValidImage(contentType) {
			json.NewEncoder(w).Encode(&types.Error{
				Status: http.StatusInternalServerError,
				Error:  "Invalid file type uploaded",
			})
			return
		}

		// Read the entire file
		fileData := bytes.NewBuffer(nil)
		_, err = io.Copy(fileData, file)
		if err != nil {
			http.Error(w, "Failed to read file data", http.StatusInternalServerError)
			return
		}

		images = append(images, fileData.Bytes())
	}

	fmt.Println("API Gateway :  CreateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := productClient.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	if err := companyClient.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	_, errCompany := companyClient.GetCompany(int64(body.CompanyId), r.Context())
	if errCompany != nil {
		fmt.Println("API Gateway :  CreateProduct - ERROR")
		errCompany.Error = "Company not found"
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	res, err := productClient.CreateProduct(body, images, r.Context())
	if err != nil {
		fmt.Println("API Gateway :  CreateProduct - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  CreateProduct - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
