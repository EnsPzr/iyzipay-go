package helpers

import (
	"crypto/sha1"
	"encoding/base64"
)

func GenerateHash(apikey, secretKey, randomString string, form string) string {
	willBeHashed := apikey + randomString + secretKey + form

	//SHA1.Create().ComputeHash
	computedHash := computeHash(willBeHashed)
	b64 := base64.StdEncoding.EncodeToString(computedHash)
	return b64
}

func computeHash(text string) []byte {
	c := sha1.New()
	c.Write([]byte(text))
	myBytes := c.Sum(nil)
	return myBytes
}
