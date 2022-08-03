
UID := $(shell id -u)
GID := $(shell id -g)

.PHONY: build
build: meraki/go.mod

meraki/go.mod: spec2.json
	mkdir -p meraki
	docker run -it -u $(UID):$(GID) --rm -v $(PWD):/local openapitools/openapi-generator-cli \
		generate \
		-i /local/spec2.json \
		-g go \
		--git-user-id ShadowApex \
		--git-repo-id meraki-go-sdk \
		--additional-properties=packageName=meraki,structPrefix=true \
		-o /local/meraki

spec2.json:
	wget https://raw.githubusercontent.com/meraki/openapi/master/openapi/spec2.json

.PHONY: clean
clean:
	rm -rf spec2.json meraki
