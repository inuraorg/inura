COMPOSEFLAGS=-d
ITESTS_L2_HOST=http://localhost:9545
BEDROCK_TAGS_REMOTE?=origin
OP_STACK_GO_BUILDER?=us-docker.pkg.dev/oplabs-tools-artifacts/images/op-stack-go:latest

# Requires at least Python v3.9; specify a minor version below if needed
PYTHON?=python3

build: build-go build-ts
.PHONY: build

build-go: submodules inura-node inura-proposer inura-batcher
.PHONY: build-go

lint-go:
	golangci-lint run -E goimports,sqlclosecheck,bodyclose,asciicheck,misspell,errorlint --timeout 5m -e "errors.As" -e "errors.Is" ./...
.PHONY: lint-go

build-ts: submodules
	if [ -n "$$NVM_DIR" ]; then \
		. $$NVM_DIR/nvm.sh && nvm use; \
	fi
	pnpm install
	pnpm build
.PHONY: build-ts

ci-builder:
	docker build -t ci-builder -f ops/docker/ci-builder/Dockerfile .

golang-docker:
	# We don't use a buildx builder here, and just load directly into regular docker, for convenience.
	GIT_COMMIT=$$(git rev-parse HEAD) \
	GIT_DATE=$$(git show -s --format='%ct') \
	IMAGE_TAGS=$$(git rev-parse HEAD),latest \
	docker buildx bake \
			--progress plain \
			--load \
			-f docker-bake.hcl \
			inura-node inura-batcher inura-proposer inura-challenger
.PHONY: golang-docker

submodules:
	git submodule update --init --recursive
.PHONY: submodules

inura-bindings:
	make -C ./inura-bindings
.PHONY: inura-bindings

inura-node:
	make -C ./inura-node inura-node
.PHONY: inura-node

generate-mocks-inura-node:
	make -C ./inura-node generate-mocks
.PHONY: generate-mocks-inura-node

generate-mocks-inura-service:
	make -C ./inura-service generate-mocks
.PHONY: generate-mocks-inura-service

inura-batcher:
	make -C ./inura-batcher inura-batcher
.PHONY: inura-batcher

inura-proposer:
	make -C ./inura-proposer inura-proposer
.PHONY: inura-proposer

inura-challenger:
	make -C ./inura-challenger inura-challenger
.PHONY: inura-challenger

inura-program:
	make -C ./inura-program inura-program
.PHONY: inura-program

cannon:
	make -C ./cannon cannon
.PHONY: cannon

cannon-prestate: inura-program cannon
	./cannon/bin/cannon load-elf --path inura-program/bin/inura-program-client.elf --out inura-program/bin/prestate.json --meta inura-program/bin/meta.json
	./cannon/bin/cannon run --proof-at '=0' --stop-at '=1' --input inura-program/bin/prestate.json --meta inura-program/bin/meta.json --proof-fmt 'inura-program/bin/%d.json' --output ""
	mv inura-program/bin/0.json inura-program/bin/prestate-proof.json

mod-tidy:
	# Below GOPRIVATE line allows mod-tidy to be run immediately after
	# releasing new versions. This bypasses the Go modules proxy, which
	# can take a while to index new versions.
	#
	# See https://proxy.golang.org/ for more info.
	export GOPRIVATE="github.com/ethereum-optimism" && go mod tidy
.PHONY: mod-tidy

clean:
	rm -rf ./bin
.PHONY: clean

nuke: clean devnet-clean
	git clean -Xdf
.PHONY: nuke

pre-devnet:
	@if ! [ -x "$(command -v geth)" ]; then \
		make install-geth; \
	fi
	@if [ ! -e inura-program/bin ]; then \
		make cannon-prestate; \
	fi
.PHONY: pre-devnet

devnet-up: pre-devnet
	./ops/scripts/newer-file.sh .devnet/allocs-l1.json ./packages/contracts-bedrock \
		|| make devnet-allocs
	PYTHONPATH=./bedrock-devnet $(PYTHON) ./bedrock-devnet/main.py --monorepo-dir=.
.PHONY: devnet-up

# alias for devnet-up
devnet-up-deploy: devnet-up

devnet-test: pre-devnet
	PYTHONPATH=./bedrock-devnet $(PYTHON) ./bedrock-devnet/main.py --monorepo-dir=. --test
.PHONY: devnet-test

devnet-down:
	@(cd ./ops-bedrock && GENESIS_TIMESTAMP=$(shell date +%s) docker compose stop)
.PHONY: devnet-down

devnet-clean:
	rm -rf ./packages/contracts-bedrock/deployments/devnetL1
	rm -rf ./.devnet
	cd ./ops-bedrock && docker compose down
	docker image ls 'ops-bedrock*' --format='{{.Repository}}' | xargs -r docker rmi
	docker volume ls --filter name=ops-bedrock --format='{{.Name}}' | xargs -r docker volume rm
.PHONY: devnet-clean

devnet-allocs: pre-devnet
	PYTHONPATH=./bedrock-devnet $(PYTHON) ./bedrock-devnet/main.py --monorepo-dir=. --allocs

devnet-logs:
	@(cd ./ops-bedrock && docker compose logs -f)
	.PHONY: devnet-logs

test-unit:
	make -C ./inura-node test
	make -C ./inura-proposer test
	make -C ./inura-batcher test
	make -C ./inura-e2e test
	pnpm test
.PHONY: test-unit

test-integration:
	bash ./ops-bedrock/test-integration.sh \
		./packages/contracts-bedrock/deployments/devnetL1
.PHONY: test-integration

# Remove the baseline-commit to generate a base reading & show all issues
semgrep:
	$(eval DEV_REF := $(shell git rev-parse develop))
	SEMGREP_REPO_NAME=inuraorg/inura semgrep ci --baseline-commit=$(DEV_REF)
.PHONY: semgrep

clean-node-modules:
	rm -rf node_modules
	rm -rf packages/**/node_modules

tag-bedrock-go-modules:
	./ops/scripts/tag-bedrock-go-modules.sh $(BEDROCK_TAGS_REMOTE) $(VERSION)
.PHONY: tag-bedrock-go-modules

update-op-geth:
	./ops/scripts/update-op-geth.py
.PHONY: update-op-geth

bedrock-markdown-links:
	docker run --init -it -v `pwd`:/input lycheeverse/lychee --verbose --no-progress --exclude-loopback \
		--exclude twitter.com --exclude explorer.optimism.io --exclude linux-mips.org \
		--exclude-mail /input/README.md "/input/specs/**/*.md"

install-geth:
	./ops/scripts/geth-version-checker.sh && \
	 	(echo "Geth versions match, not installing geth..."; true) || \
 		(echo "Versions do not match, installing geth!"; \
 			go install -v github.com/ethereum/go-ethereum/cmd/geth@$(shell cat .gethrc); \
 			echo "Installed geth!"; true)
.PHONY: install-geth
