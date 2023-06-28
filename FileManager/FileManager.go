package FileManager

import "os"

func LoadFile(fileUrl string) []byte {
	dat, err := os.ReadFile(fileUrl)
	if err != nil {
		panic(err)
	}
	return dat
}
