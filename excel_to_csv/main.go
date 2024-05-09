package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	// Excelファイルを開く
	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// シート名を取得
	sheetName := f.GetSheetName(1)

	// セル結合を解除
	mergedCells, err := f.GetMergeCells(sheetName)
	for _, mergedCell := range mergedCells {
		if err := f.UnmergeCell(sheetName, mergedCell); err != nil {
			fmt.Println(err)
			return
		}
	}

	// 列の数を取得
	maxCol, err := f.GetCols(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 全ての列の値を取得してカンマ区切りの文字列に変換
	rows, err := f.GetRows(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var csvData [][]string
	for _, row := range rows {
		if len(row) > 0 {
			var rowData []string
			for col := 1; col <= maxCol; col++ {
				cellValue, err := f.GetCellValue(sheetName, excelize.ToAlphaString(col)+strconv.Itoa(row[0]))
				if err != nil {
					fmt.Println(err)
					return
				}
				// カンマや改行が含まれている場合は置換する
				cellValue = strings.ReplaceAll(cellValue, "\n", "、")
				cellValue = strings.ReplaceAll(cellValue, ",", "、")
				rowData = append(rowData, cellValue)
			}
			csvData = append(csvData, rowData)
		}
	}

	// 各行のカンマの個数を計算し、最大個数を取得する
	var maxCount int
	for _, row := range csvData {
		count := len(row)
		if count > maxCount {
			maxCount = count
		}
	}

	// 各行が最大個数と同じになるようにカンマを追加する
	for i, row := range csvData {
		count := len(row)
		for count < maxCount {
			csvData[i] = append(csvData[i], "")
			count++
		}
	}

	// CSVファイルに書き込み
	file, err := os.Create("output.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range csvData {
		if err := writer.Write(value); err != nil {
			fmt.Println("error writing record to csv:", err)
			return
		}
	}

	fmt.Println("CSVファイルに変換しました。")
}
