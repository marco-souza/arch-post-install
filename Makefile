ALL: clean run

run: src/main.go src/utils.go
	go run src/main.go src/utils.go

build: src/main.go src/utils.go
	go build -o post-install src/main.go src/utils.go

clean:
	rm -rf /tmp/zshrc/ ~/.oh-my-zsh/

release: build
	zip post-install.zip data/apps.json post-install
	git add .
	git tag -a $v -m "Release $v"
	git commit -m "REL : Release v$v"
	git push