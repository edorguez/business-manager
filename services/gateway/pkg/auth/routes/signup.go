package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/edorguez/business-manager/services/gateway/pkg/auth/client"
	"github.com/edorguez/business-manager/services/gateway/pkg/auth/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/shared/types"
	"github.com/edorguez/business-manager/shared/util/file_validator"
)

func SignUp(w http.ResponseWriter, r *http.Request, c *config.Config) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("API Gateway :  Sign Up")

	// We got our body through context, since we saved it in a middleware
	body := r.Context().Value(contracts.SignUpRequest{}).(contracts.SignUpRequest)

	var image []byte

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

		image = fileData.Bytes()
	}

	fmt.Println("API Gateway :  SignUpRequest - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	if err := client.InitAuthServiceClient(c); err != nil {
		json.NewEncoder(w).Encode(&types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		})
		return
	}

	res, err := client.SignUp(body, image, r.Context())

	if err != nil {
		fmt.Println("API Gateway :  Sign Up - ERROR")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Println("API Gateway :  Sign Up - SUCCESS")
	w.WriteHeader(int(res.Status))
	json.NewEncoder(w).Encode(res)
}
