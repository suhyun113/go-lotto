package domain

import "github.com/suhyun113/go-lotto/constants"

func DetermineRank(matchCount int, bonusMatch bool) constants.Rank {
	switch matchCount {
	case 6:
		return constants.RankFirst
	case 5:
		if bonusMatch {
			return constants.RankSecond
		}
		return constants.RankThird
	case 4:
		return constants.RankFourth
	case 3:
		return constants.RankFifth
	default:
		return constants.RankNone
	}
}
