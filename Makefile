PACKAGES=$(shell go list ./...)

test_unit:
	@echo "--> go test "
	@go test -race $(PACKAGES)
