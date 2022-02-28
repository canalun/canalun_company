test:
	cd ./scripts/content_updater && go test -v ./... && cd -

run-scripts:
	. ./local_env_set.sh && cd ./scripts/content_updater && go run *.go

update-mod:
	cd ./scripts/content_updater && go mod tidy  && cd -

local-dev:
	cd ./company_home && python -m SimpleHTTPServer 8000 && cd -