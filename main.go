package main

import (
	"flag"
	"log"

	"github.com/sgash708/scraping_lawyers/application/usecase"
	"github.com/sgash708/scraping_lawyers/infra/persistence"
	"github.com/sgash708/scraping_lawyers/interfaces/handler"
)

func main() {
	officePersistence := persistence.NewOfficePersistence()
	officeUseCase := usecase.NewUserUseCase(officePersistence)
	officeHandler := handler.NewOfficeHandler(officeUseCase)

	log.Println("処理開始...")

	f := flag.String("flg", "0", "検索フラグ(0:事務所/1:弁護士/2:外国弁護士)")
	flag.Parse()

	// マジックナンバーの追加
	switch *f {
	case "0":
		officeHandler.GetOffice()
	// case "1":
	// 	getLawyers(false)
	// case "2":
	// 	getLawyers(true)
	default:
		log.Fatalln("FLAG指定が抜けているか間違っています。再設定してください。")
	}

	log.Println("処理終了...")
}
