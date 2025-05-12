## 실행방법

1. git clone https://github.com/JIWEON-JEONG/coupon-issuance-system.git
2. docker-compose up -d
3. go build 
8. go run .

## 구조

### controller 
- presentation layer
- 클라이언트 요청, 응답에 대한 스펙 구현

### usecase 
- domain or domain_service 들을 facade 하여 요구사항 구현.

### domain
 - 비즈니스, 정책들을 정의 및 구현
 - 도메인 독립적인 의존성만을 가지며 (orm 의존성 제외), 외부 의존성을 최대한 차단.

### model
 - 도메인 및 엔티티 모델

### repository
 - persistence layer
 - 데이터베이스 등과 같은 3rd 파티 상호작용 구현체.



  

