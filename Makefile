OS:=linux-amd64
APP:= totp-tool
TAG := `git describe --abbrev=0 --tags`
TZ := Asia/Taipei
DATETIME := `TZ=$(TZ) date +%Y/%m/%d.%T`
show-tag:
	echo $(TAG)
build = GOOS=$(1) GOARCH=$(2) go build -ldflags "-X main.version=$(TAG) -X main.name=$(APP) -X main.compileDate=$(DATETIME)($(TZ))" -a -o build/$(APP)$(3) 
tar = cd build && tar -zcvf $(APP)_$(TAG)_$(1)_$(2).tar.gz $(APP)$(3)  && rm $(APP)$(3)

build/linux: 
	$(call build,linux,amd64,)
build/linux_amd64.tar.gz: build/linux
	$(call tar,linux,amd64,)
build/windows: 
	$(call build,windows,amd64,.exe)
build/windows_amd64.tar.gz: build/windows
	$(call tar,windows,amd64,.exe)
build/darwin: 
	$(call build,darwin,amd64,)
build/darwin_amd64.tar.gz: build/darwin
	$(call tar,darwin,amd64,)
clean:
	rm -rf build/
