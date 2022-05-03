build:
	go build -o helloword \
	-ldflags "-X 'github.com/xuender/go-sign.Safe=strong'" \
	main.go

	SECRET_KEY=secret_key sign -e=SECRET_KEY helloword

run:
	SECRET_KEY=secret_key ./helloword
