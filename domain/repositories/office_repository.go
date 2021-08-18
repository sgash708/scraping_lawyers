package repositories

import (
	"database/sql"

	"github.com/sgash708/scraping_lawyers/domain"
	// _ "github.com/sgash708/scraping_lawyers/domain"
)

// OfficeRepository officeリポジトリ
type OfficeRepository interface {
	// REF: https://qiita.com/tono-maron/items/345c433b86f74d314c8d#domain%E5%B1%A4
	// データ挿入
	Insert(DB *sql.DB, datas map[int]map[string]string) (int, error)
	// ID数カウント
	GetIDCount(DB *sql.DB, tableName string) (*domain.Office, error)
}
