build:
	go build -o bin/notelify-articles-service
serve-prod: build
	./bin/notelify-articles-service -env=prod

serve-dev: build
	./bin/notelify-articles-service -env=dev

serve-docker: build
	./bin/notelify-articles-service -env=docker_test


test:
	go test -v -tags=myenv ./...
	Env=dev go test -v -tags=myenv ./...


