package file_validator

func IsValidImage(contentType string) bool {
	allowedTypes := []string{"image/png", "image/jpeg", "image/jpg"}
	for _, allowed := range allowedTypes {
		if contentType == allowed {
			return true
		}
	}
	return false
}
