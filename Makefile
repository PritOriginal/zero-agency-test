run-http:
	go run ./cmd/bot-http/ --config=./configs/config.yaml

docker-http:
	docker build -t bot-http -f ./docker/http/Dockerfile .
	docker compose -f ./docker/http/compose.yaml --project-directory . run -it bot-http

build-http:
	go build ./cmd/bot-http/


run-openai:
	go run ./cmd/bot-openai/ --config=./configs/config.yaml

docker-openai:
	docker build -t bot-openai -f ./docker/openai/Dockerfile .
	docker compose -f ./docker/openai/compose.yaml --project-directory . run -it bot-openai

build-openai:
	go build ./cmd/bot-openai/

test:
	go test ./...

test-cover:
	go test ./... -coverprofile cover.test.tmp -coverpkg ./...
	cat cover.test.tmp | grep -v "mocks" > cover.test 
	rm cover.test.tmp 
	go tool cover -func cover.test 

test-integration-http:
	go test -tags=integration ./tests/integration/http -count 1 -timeout 120s

test-integration-openai:
	go test -tags=integration ./tests/integration/openai -count 1 -timeout 120s