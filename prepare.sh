set -e
set -x

go install github.com/evilmartians/lefthook@latest
go install github.com/git-chglog/git-chglog/cmd/git-chglog@latest
go install github.com/vektra/mockery/v2@latest
pip install MarkdownPP

lefthook install
