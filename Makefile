GOLANG_CI_LINT := bin/golangci-lint

.PHONY: test
test:
	go test -v ./...

.PHONY: release
release:
	$(eval NEXT_VERSION := $(shell convco version --bump))
	git tag -a v$(NEXT_VERSION) -m "chore(release): v$(NEXT_VERSION)"
	git push origin v$(NEXT_VERSION)
	convco changelog --max-versions 1 > CHANGELOG.md
	gh release create v$(NEXT_VERSION) --title "v$(NEXT_VERSION)" --notes-file CHANGELOG.md

$(GOLANG_CI_LINT):
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.64.7
	@chmod +x $(GOLANG_CI_LINT)
	@echo "golangci-lint installed"

.PHONY: lint
lint: $(GOLANG_CI_LINT)
	$(GOLANG_CI_LINT) run