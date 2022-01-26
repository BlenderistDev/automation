gen-automation-core-mock:
	@mockgen -source=interfaces/interfaces.go -destination=testing/interfaces/interfaces.go

test:
	@go test ./... -race -count=1 -cover -coverprofile=coverage.txt && go tool cover -func=coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm coverage.txt

format:
	@go fmt ./...

gen-datamapper-mock:
	@mockgen -source=datamapper/dataMapper.go -destination=testing/datamapper/dataMapper.go
