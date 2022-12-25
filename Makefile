.DEFAULT_GOAL := build
run: 
	@ ./scripts/run.sh $(name)
build:
	@ ./scripts/build.sh $(name)
