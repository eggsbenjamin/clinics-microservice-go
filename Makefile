.SILENT:

COLOUR=\033[0;93m
RESET=\033[m

install:
	printf "\n${COLOUR}Installing Dependencies...\n${RESET}"
	@glide install

build_test_deps:
	@printf "\n${COLOUR}Building Test Dependencies...\n${RESET}"
	@go build -v -o bin/ginkgo ./vendor/github.com/onsi/ginkgo/ginkgo
	@chmod -R 0777 bin

run_system_tests:
	@printf "\n${COLOUR}Running System Tests...\n${RESET}"
	@bin/ginkgo systemtest

run_unit_tests:
	@printf "\n${COLOUR}Running Unit Tests...\n${RESET}"
	@bin/ginkgo services utils mappers handlers

run_all_tests:
	@make run_unit_tests && make run_system_tests
