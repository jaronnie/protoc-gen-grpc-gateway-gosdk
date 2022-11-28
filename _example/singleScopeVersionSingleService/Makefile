.PHONY: gensdk
gensdk:
	@export PATH=`pwd`/tools/proto/bin:"${PATH}" && sh script/generate.sh && sh script/goimports.sh all

.PHONY: gensdk.debug
gensdk.debug:
	@export PATH=`pwd`/tools/proto/bin:"${PATH}" && sh script/generate.sh debug

.PHONY: gensdk.debug.fmt
gensdk.debug.fmt:
	@export PATH=`pwd`/tools/proto/bin:"${PATH}" && sh script/generate.sh debug && sh script/goimports.sh all bfsdk
