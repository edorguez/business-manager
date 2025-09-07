package phone

func RemovePhoneZero(number string) string {
	if len(number) > 0 && number[0] == '0' {
		number = number[1:]
	}

	return number
}
