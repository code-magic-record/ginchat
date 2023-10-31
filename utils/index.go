package utils

func GetPhoneFromToken(token string) string {
	var phone string
	phone, err := RDB.Get(RDB.Context(), token).Result()
	if err != nil {
		phone = ""
	}
	return phone
}
