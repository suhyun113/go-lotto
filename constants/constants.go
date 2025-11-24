package constants

const (
	LottoMin = 1
	LottoMax = 45
	LottoSize = 6
	LottoPrice = 1000
)

type Rank int

const (
	RankNone Rank = iota // 꽝
	RankFifth // 3개 
	RankFourth // 4개
	RankThird // 5개
	RankSecond // 5개 + 보너스
	RankFirst // 6개
)

// 등수별 상금 테이블
var Prize = map[Rank]int64{
	RankFirst: 2000000000,
	RankSecond: 30000000,
	RankThird: 1500000,
	RankFourth: 50000,
	RankFifth: 5000,
	RankNone: 0,
}

// 출력 라벨
var RankLabel = map[Rank]string{
	RankFifth:  "3개 일치 (5,000원)",
	RankFourth: "4개 일치 (50,000원)",
	RankThird:  "5개 일치 (1,500,000원)",
	RankSecond: "5개 일치, 보너스 볼 일치 (30,000,000원)",
	RankFirst:  "6개 일치 (2,000,000,000원)",
	RankNone: "꽝",
}

var MatchCount = map[Rank]int{
	RankFirst:  6,
	RankSecond: 5,
	RankThird:  5,
	RankFourth: 4,
	RankFifth:  3,
	RankNone:   0,
}

func (r Rank) String() string {
	switch r {
	case RankFirst:
		return "FIRST"
	case RankSecond:
		return "SECOND"
	case RankThird:
		return "THIRD"
	case RankFourth:
		return "FOURTH"
	case RankFifth:
		return "FIFTH"
	default:
		return "NONE"
	}
}