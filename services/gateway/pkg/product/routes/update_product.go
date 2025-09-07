package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/client"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/util/file_validator"
	"github.com/gorilla/mux"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  UpdateProduct")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value("keyProductUpdate").(contracts.UpdateProductRequest)

	vars := mux.Vars(r)

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

		images = append(images, fileData.Bytes())
	}

	fmt.Println("API Gateway :  UpdateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitProductServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, errUpdate := client.UpdateProduct(vars["id"], body, images, r.Context())

	if errUpdate != nil {
		fmt.Println("API Gateway :  UpdateProduct - ERROR")
		json.NewEncoder(w).Encode(errUpdate)
		return
	}

	fmt.Println("API Gateway :  UpdateProduct - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
