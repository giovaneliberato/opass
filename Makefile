

build:
	go build -o cli_opass ./cmd/opass/

clean:
	rm cli_opass

install: build
	mkdir -p $(HOME)/.opass
	cp cli_opass /usr/local/bin/opass

uninstall:
	rm /usr/local/bin/opass
	rm -rf $(HOME)/.opass
