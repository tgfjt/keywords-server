GO ?= go
app := keywords

gqlgen:
	@echo "start gqlgen"
	@cd src && $(GO) run ./scripts/gqlgen.go -v
.PHONY: gqlgen

serve:
	@echo "start devserver"
	@cd app && dev_appserver.py app.yaml
.PHONY: serve

deploy:
	$ appcfg.py update --application=keywords-server --version=1 --oauth2_access_token=$(gcloud auth print-access-token 2> /dev/null) app/app.yaml
.PHONY: deploy
