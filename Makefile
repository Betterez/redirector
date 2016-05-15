default:
	@export GOPATH=$(pwd)
	@go install red

run: default
	@bin/red
