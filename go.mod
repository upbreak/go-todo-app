module github.com/upbreak/go-todo-app

go 1.23.4

require golang.org/x/sync v0.10.0 // go get -u golang.org/x/sync: main()에서 사용할 run함수를 구현할 때 이용하는 준표준 패키지

require github.com/caarlos0/env v3.5.0+incompatible // go get -u github.com/caarlos0/env: struct에 태그를 사용하여 환경변수를 가져와 사용할 수 있게 해준다
