package constants

const (
	LOTTO_MIN   = 1
	LOTTO_MAX   = 45
	LOTTO_SIZE  = 6
	LOTTO_PRICE = 1000
)

// 등수/상금 테이블
var PRIZE = map[string]int64{
	"FIRST":  2000000000, // 6개
	"SECOND": 30000000,   // 5개 + 보너스
	"THIRD":  1500000,    // 5개
	"FOURTH": 50000,      // 4개
	"FIFTH":  5000,       // 3개
}

// 출력 라벨
var RANK_LABEL = map[string]string{
	"FIFTH":  "3개 일치 (5,000원)",
	"FOURTH": "4개 일치 (50,000원)",
	"THIRD":  "5개 일치 (1,500,000원)",
	"SECOND": "5개 일치, 보너스 볼 일치 (30,000,000원)",
	"FIRST":  "6개 일치 (2,000,000,000원)",
}
