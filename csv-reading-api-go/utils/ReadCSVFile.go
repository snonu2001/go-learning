package utils

import (
	"encoding/csv"
	"mime/multipart"
)

func ReadCsvData(file multipart.File) ([][]string, error){
	
    reader := csv.NewReader(file)

    // skip first line
    if _, err := reader.Read(); err != nil {
        return [][]string{}, err
    }

    data, err := reader.ReadAll()

    if err != nil {
        return [][]string{}, err
    }

    return data, nil
}