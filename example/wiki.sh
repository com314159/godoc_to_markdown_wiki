go run $GOPATH/src/github.com/com314159/godoc_to_markdown_wiki/main.go -toml="$GOPATH//src/github.com/com314159/godoc_to_markdown_wiki/example/wiki.toml"
cd wiki
git add --all && git commit -m "Update Wiki" && git push origin master