TARGETS = spec.pdf report.pdf

.PHONY: all clean

all: $(TARGETS)

%.pdf: %.tex
	pdflatex $^

%.pdf: %.lyx
	lyx -e pdf $^

clean:
	rm -f *.aux *.log *.out *.lyx# *.lyx~ report.tex $(TARGETS)
