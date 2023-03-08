package schedulers

import (
	"github.com/jasonlvhit/gocron"
	"github.com/lee-lou2/api/platform/kakao"
	"log"
)

// refreshKaKaOToken 카카오 토큰 재발급
func refreshKaKaOToken() {
	if err := kakao.RefreshKaKaoToken(); err != nil {
		log.Println(err)
	} else {
		log.Printf("토큰 재발급 완료")
	}
}

// RunSchedulers 스케쥴러 실행
func RunSchedulers() {
	// 최초 토큰 생성
	gocron.Every(60 * 10).Second().Do(refreshKaKaOToken)
	<-gocron.Start()
}
