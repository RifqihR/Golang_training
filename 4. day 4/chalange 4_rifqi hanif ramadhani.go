package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputvalue := os.Args[1]

	var data = map[int]map[string]string{
		1: {
			"nama":      "Rifqi",
			"alamat":    "Jambi",
			"pekerjaan": "Belum ada",
			"alasan":    "Inging belajar",
		},
		2: {
			"nama":      "Ahmad",
			"alamat":    "Maluku",
			"pekerjaan": "developer",
			"alasan":    "menambah skill",
		},
		3: {
			"nama":      "siti",
			"alamat":    "jambon",
			"pekerjaan": "kasir toko",
			"alasan":    "ingin berganti pekerjaan",
		},
		4: {
			"nama":      "juliah",
			"alamat":    "bandung",
			"pekerjaan": "pedagang",
			"alasan":    "menambah skill",
		},
	}

	input, _ := strconv.Atoi(inputvalue)

	for key, val := range data {
		if key == input {
			fmt.Printf("nama : %s\nalamat : %s\npekerjaan : %s \nalasan : %s\n", val["nama"], val["alamat"], val["pekerjaan"], val["alasan"])
		}
	}

}
