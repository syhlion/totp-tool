OS:=linux-amd64
NAME:= totp-tool
TAG := `git describe --abbrev=0 --tags`
TZ := Asia/Taipei
DATETIME := `TZ=$(TZ) date +%Y/%m/%d.%T`
show-tag:
	echo $(TAG)
build = GOOS=$(1) GOARCH=$(2) go build -ldflags "-X main.version=$(TAG) -X main.name=$(NAME) -X main.compileDate=$(DATETIME)($(TZ))" -a -o build/$(NAME)$(3) 
tar = cp *.env.example ./build && cp test/conn-test/conn-test.env.example ./build &&cd build && tar -zcvf $(GUSHER)_$(TAG)_$(1)_$(2).tar.gz $(JWTGENERATE)$(3) $(CONNTEST)$(3) $(GUSHER)$(3) *.env.example  test/ && rm $(JWTGENERATE)$(3) $(CONNTEST)$(3) $(GUSHER)$(3)  *.env.example  && rm -rf test/

build/linux: 
	go test
	$(call build,linux,amd64,)
build/windows: 
	go test
	$(call build,windows,amd64,.exe)
build/darwin: 
	go test
	$(call build,darwin,amd64,)
clean:
	rm -rf build/
