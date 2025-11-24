package domain

import "testing"

func TestNewLotto_SuccessAndSorted(t *testing.T) {
	lotto, err := NewLotto([]int{5, 3, 1, 6, 2, 4})
	if err != nil {
		t.Fatalf("[ERROR] Lotto 생성 실패: %v", err)
	}
	want := []int{1, 2, 3, 4, 5, 6}
	got := lotto.GetNumbers()
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("[ERROR] 번호 정렬 실패: got=%v want=%v", got, want)
		}
	}
}

func TestNewLotto_Fail_Size(t *testing.T) {
	_, err := NewLotto([]int{1, 2, 3})
	if err == nil {
		t.Fatalf("[ERROR] 번호 6개 미만인데 허용됨")
	}
}

func TestNewLotto_Fail_Duplicate(t *testing.T) {
	_, err := NewLotto([]int{1, 1, 2, 3, 4, 5})
	if err == nil {
		t.Fatalf("[ERROR] 번호 중복을 허용함")
	}
}

func TestNewLotto_Fail_Range(t *testing.T) {
	_, err := NewLotto([]int{0, 2, 3, 4, 5, 6})
	if err == nil {
		t.Fatalf("[ERROR] 범위 밖 번호를 허용함")
	}
}

func TestMatchCount(t *testing.T) {
	ticket, _ := NewLotto([]int{1, 2, 3, 7, 8, 9})
	winning, _ := NewLotto([]int{1, 2, 3, 4, 5, 6})

	if ticket.MatchCount(winning) != 3 {
		t.Fatalf("[ERROR] 일치 개수 계산 실패: got=%d want=3", ticket.MatchCount(winning))
	}
}
