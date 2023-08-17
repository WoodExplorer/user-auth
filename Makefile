image_name ?=user-auth

builder:
	docker build -t ${image_name} -f build/Dockerfile .

cover:
	go test ./... -cover -coverprofile=coverage.out
	go tool cover -html=coverage.out
