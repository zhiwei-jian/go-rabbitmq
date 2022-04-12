package utils

import (
	"encoding/base64"
	"strings"
)

func Base64Decode(str string) string {
	var strs []string = strings.Split(str, "\"")
	str = strs[1]
	// reader := strings.NewReader(str)
	// decoder := base64.NewDecoder(base64.RawStdEncoding, reader)

	// buf := make([]byte, 1024)

	// dst := ""
	// for {
	// 	n, err := decoder.Read(buf)
	// 	dst += string(buf[:n])
	// 	if n < 0 || err != nil {
	// 		break
	// 	}
	// }

	rawDecodedText, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return string(rawDecodedText)
}
