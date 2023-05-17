docs:
	pandoc -s introduction.md 01-problem.md -o space-invaders.pdf \
	 --pdf-engine=xelatex \
	--table-of-contents \
	--number-sections

install-ugo:
	go get -u github.com/jromero/ugo/cmd/ugo

test: install-ugo
	ugo run -v introduction.md 01-problem.md
