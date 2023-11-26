package re_capcha

import (
	"encoding/csv"
	"fmt"
	"os"
)

// WriteLocalFile ローカルファイルへ結果を保存
func WriteLocalFile(path string, records [][]string) {

	// ファイルを開く（存在しない場合は新規作成、存在する場合は追記モードで開く）
	file, err := os.OpenFile(fmt.Sprintf("%s/%s", path, "result.csv"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(fmt.Sprintf("ファイルのオープンに失敗しました: %v", err))
	}
	defer file.Close()

	// CSVライターを作成
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// データをCSVファイルに追記
	for _, record := range records {
		if err := writer.Write(record); err != nil {
			fmt.Println(fmt.Sprintf("CSVへの書き込みに失敗しました: %v", err))
		}
	}
}
