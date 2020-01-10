package jwt

// signKey for signing algorithm
var signKey = "Myriad-Dreamin"

// GetSignKey instance
func GetSignKey() string {
	return signKey
}

// SetSignKey instance
func SetSignKey(key string) {
	signKey = key
}
