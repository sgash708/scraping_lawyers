package persistence

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/sgash708/scraping_lawyers/domain/model"
	"github.com/sgash708/scraping_lawyers/domain/repository"
)

// count
var count string

// OfficePersistence データベース構造体
type OfficePersistence struct{}

// NewOfficePersistence
func NewOfficePersistence() repository.OfficeRepository {
	/*
	 * ①OfficeRepositoryはGetIDCountとInsert関数を持っている。
	 * ②officePersistenceをレシーバに持つGetIDCount関数とInsert関数を実装する。
	 * ③officePersistenceがinterfaceのOfficeRepositoryを満たしていることになる。
	 * ④interfaceの性質上、（今回に関してはレシーバの）型は不問。
	 * ⑤返り値はOfficeRepositoryというinterfaceなのでそれを満たしているので返り値として満たされる。
	 */
	return &OfficePersistence{}
}

// getOfficeQuery オフィス追加クエリ
func getOfficeQuery(data map[string]string) string {
	return fmt.Sprintf(`
INSERT INTO nichibenren.offices (
	id,
	corporate_name,
	bar_association,
	name,
	office_bar_association,
	postcode,
	address,
	phone_number,
	fax_number,
	corp_kanji,
	office_kanji
)
VALUES (
	%s,
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	'%s',
	%s,
	%s
)
ON DUPLICATE KEY UPDATE
	name = VALUES(name),
	office_kanji = VALUES(office_kanji),
	address = VALUES(address),
	postcode = VALUES(postcode),
	phone_number = VALUES(phone_number),
	fax_number = VALUES(fax_number);
`,
		data["届出番号"], data["法人名"], data["弁護士会"], data["事務所名"], data["事務所弁護士会"],
		data["郵便番号"], data["住所"], data["電話番号"], data["FAX番号"], data["法人名漢字フラグ"],
		data["事務所名漢字フラグ"])
}

// Insert データ追加
func (d *OfficePersistence) Insert(db *sql.DB, datas map[int]map[string]string) (int, error) {
	recordNum := 0
	for _, data := range datas {
		if data["届出番号"] == "" {
			continue
		}
		_, err := db.Exec(getOfficeQuery(data))
		if err != nil {
			return recordNum, err
		}
		recordNum++
	}

	return recordNum, nil
}

// GetIDCount IDの数カウントする
func (d *OfficePersistence) GetIDCount(db *sql.DB, tableName string) (*model.Office, error) {
	office := model.Office{}
	if err := db.QueryRow(fmt.Sprintf("SELECT COUNT(id) as count FROM nichibenren.%s", tableName)).Scan(&count); err != nil {
		return nil, err
	}

	cnt, err := strconv.Atoi(count)
	if err != nil {
		return nil, err
	}
	office.CNT = cnt

	return &office, nil
}
