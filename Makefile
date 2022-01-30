test:
	go test -v ./...

run-scripts:
	. ./local_env_set.sh && cd ./scripts/content_updater && go run *.go

update-mod:
	cd ./scripts/content_updater && go mod tidy  && cd -