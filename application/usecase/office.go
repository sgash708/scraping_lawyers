package usecase

import (
	"database/sql"

	"github.com/sgash708/scraping_lawyers/domain"
	"github.com/sgash708/scraping_lawyers/domain/repositories"
)

// OfficeUseCase OffceのUseCaseインターフェース
type OfficeUseCase interface {
	Insert(DB *sql.DB, datas map[int]map[string]string) (int, error)
}

// officeUseCase repositoriesのstruct
type officeUseCase struct {
	officeRepository repositories.OfficeRepository
}

// NewUserUseCase 初期化
func NewUserUseCase(officeRepo repositories.OfficeRepository) OfficeUseCase {
	return &officeUseCase{
		officeRepository: officeRepo,
	}
}

// Insert repositoryからInsert関数呼びだす
func (ou *officeUseCase) Insert(DB *sql.DB, datas map[int]map[string]string) (int, error) {
	// 本来は、データの精査をする(Validationなど)

	// domainを通じてinfraで実装した関数を呼出
	count, err := ou.officeRepository.Insert(DB, datas)
	if err != nil {
		return count, err
	}
	return count, nil
}

// GetIDCount repositoryからGetIDCount関数呼びだす
func (ou *officeUseCase) GetIDCount(db *sql.DB, tableName string) (*domain.Office, error) {
	office, err := ou.officeRepository.GetIDCount(db, tableName)
	if err != nil {
		return nil, err
	}
	return office, nil
}
