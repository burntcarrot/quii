.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

loc: ## Displays lines of code (LOC).
	@find -name "*.go" -print | xargs wc -l | tail -1 | cut -d ' ' -f 2 > loc.txt
	@cat loc.txt | sed 's/.*/& lines of code/'

# %:
# 	@:
