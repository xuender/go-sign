build:
	go build -o helloword main.go
	SECRET_KEY=secret_key gosign -e=SECRET_KEY helloword

run:
	SECRET_KEY=secret_key ./helloword
