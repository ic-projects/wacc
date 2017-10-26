TARGETS = spec.pdf

.PHONY: all clean

all: $(TARGETS)

%.pdf: %.tex
	pdflatex $^

clean:
	rm -f *.aux *.log *.out $(TARGETS)