dev:
	./air run main.go


test-api:
	APP_ENV=test go test -v  ./.../.../test/... 

mock-build:
	mockgen -package project_test \
	-destination api/project/test/mock_redis_test.go \
	github.com/birdglove2/nitad-backend/redis RedisStorage