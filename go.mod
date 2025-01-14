module github.com/upbreak/go-todo-app

go 1.23.4

require (
	github.com/caarlos0/env v3.5.0+incompatible // struct에 태그를 사용하여 환경변수를 가져와 사용할 수 있게 해준다
	github.com/go-chi/chi/v5 v5.2.0 // net/http 패키지 타입 정의를 따르며 라우팅 기능을 제공
	github.com/go-playground/validator v9.31.0+incompatible // json에서 unmarshal된 데이터를 검증할 수 있게 도와준다.
	golang.org/x/sync v0.10.0 // main()에서 사용할 run함수를 구현할 때 이용하는 준표준 패키지
)

require (
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/godror/knownpb v0.2.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/planetscale/vtprotobuf v0.6.0 // indirect
	github.com/rs/cors v1.11.1 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8 // indirect
	google.golang.org/protobuf v1.36.1 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
)

require (
	github.com/godror/godror v0.46.0
	github.com/jmoiron/sqlx v1.4.0
)
