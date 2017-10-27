.PHONY = clean

src: main.go wacc.go
	go build

wacc.go: grammar/wacc.peg
	peg grammar/wacc.peg
	mv grammar/wacc.peg.go wacc.go

clean:
	rm wacc.go src
