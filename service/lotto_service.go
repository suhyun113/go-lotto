package service

import (
	"math/rand"
	"time"
)

type LottoService struct {
	rng *rand.Rand
}

func NewLottoService() *LottoService {
	return &LottoService{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}