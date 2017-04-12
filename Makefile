TEST_TARGET=./test/...

test:
 
  @go test $$TEST_TARGET -cover -race -timeout=5s
