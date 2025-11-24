package domain

import (
	"testing"

	"github.com/suhyun113/go-lotto/constants"
)

func TestDetermineRank(t *testing.T) {
	tests := []struct {
		match int
		bonus bool
		want  constants.Rank
	}{
		{6, false, constants.RankFirst},
		{5, true, constants.RankSecond},
		{5, false, constants.RankThird},
		{4, false, constants.RankFourth},
		{3, false, constants.RankFifth},
		{2, false, constants.RankNone},
	}

	for _, tt := range tests {
		got := DetermineRank(tt.match, tt.bonus)
		if got != tt.want {
			t.Fatalf("[ERROR] 등수 판정 실패: match=%d bonus=%v got=%v want=%v",
				tt.match, tt.bonus, got, tt.want)
		}
	}
}
