GO ?= go
app := keywords

gqlgen:
	@echo "start gqlgen"
	@cd src && $(GO) run ./scripts/gqlgen.go -v
.PHONY: gqlgen

serve:
	@echo "start server"
	@cd app $(GO) run ./main.go
.PHONY: serve

devserver:
	@echo "start devserver"
	@cd app && dev_appserver.py app.yaml
.PHONY: devserver

deploy:
	cd app && gcloud app deploy --project keywords-server
.PHONY: deploy
