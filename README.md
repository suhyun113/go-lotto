## go-lotto

> Go로 구현한 CLI 기반 로또 게임

구입 금액 입력 -> 로또 자동 발행 -> 당첨/보너스 번호 입력 -> 등수 집계 및 수익률 출력

---

### ✨ 실행 흐름

1. 안내 문구 출력
2. 로또 구입 금액 입력 및 검증 (1,000원 단위, 최소 1,000원)
3. 구입 장 particulars 만큼 로또 자동 발행 후 오름차순 출력
4. 당첨 번호 6개 입력(쉼표 구분) 및 검증
5. 보너스 번호 1개 입력 및 검증
6. 구매 로또 vs 당첨 번호 비교 → 등수별 개수 집계
7. 총 수익률 출력(소수점 첫째 자리까지 반올림)
8. 종료
잘못된 입력이 들어오면
[ERROR]로 시작하는 메시지를 출력한 뒤 해당 단계부터 재입력을 받는다.

### 🧩 고려사항
- 로또 번호 범위: 1~45
- 한 장의 로또 번호 개수: 6개
- 로또 번호는 중복 불가
-보너스 번호는 당첨 번호와 중복 불가
- 로또 1장 가격: 1,000원
- 구입 금액에 맞는 장수만큼 로또가 발행되어야 함

### 🦴 프로젝트 구조
```
go-lotto/
├── constants/
│   └── constants.go           
├── domain/
│   ├── lotto.go                  
│   ├── rank_rules.go          
│   ├── lotto_test.go
│   └── rank_rules_test.go
├── service/
│   ├── lotto_service.go         
│   └── lotto_service_test.go
├── util/
│   ├── parser.go               
│   ├── validator.go           
│   ├── format.go              
│   ├── parser_test.go
│   ├── validator_test.go
│   └── format_test.go
├── view/
│   ├── input_view.go          
│   ├── output_view.go           
│   └── error_view.go            
├── integration_test/
│   └── application_test.go     
├── main.go                     
└── go.mod
```

### 🍀 패키지별 기능
**1) constants**
> 로또 규칙과 등수 데이터를 관리하는 상수 패키지

**2) util**
> 입출력 파싱, 검증, 포맷팅 등 재사용 유틸리티 모음
- parser, validator, format

**3) domain**
> 도메인 모델(Lotto)과 순위 판정 규칙
- Lotto, DeterminRank

**4) service**
> 애플리케이션 핵심 비즈니스 로직

**5) view(IO)
> 입력 처리 & 출력 UI

---

### 🛠️ 사용 환경
- Go Version: go1.25.4
- 모듈 시스템: Go Modules 기반(`go mod tidy` 필요)

---

### ⚙️ 실행 방법
```
go run main.go
```

### 🔎 테스트 실행 방법
- 전체 테스트
```
go test ./...
```
- 통합 테스트
```
go test ./integration_test
```