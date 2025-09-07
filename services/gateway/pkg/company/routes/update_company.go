package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/file_validator"
	"github.com/gorilla/mux"
)

func UpdateCompany(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateCompany")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.UpdateCompanyRequest{}).(contracts.UpdateCompanyRequest)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Unable to convert ID",
		})
	}

	var image []byte
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
			json.NewEncoder(w).Encode(&contracts.Error{
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

		image = fileData.Bytes()
	}

	fmt.Println("API Gateway :  UpdateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitCompanyServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errCompany := client.UpdateCompany(int64(id), body, image, r.Context())

	if errCompany != nil {
		fmt.Println("API Gateway :  UpdateCompany - ERROR")
		json.NewEncoder(w).Encode(errCompany)
		return
	}

	fmt.Println("API Gateway :  UpdateCompany - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
