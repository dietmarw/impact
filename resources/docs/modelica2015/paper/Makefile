# SVG to PDF conversion
SVG2PDF   = rsvg-convert
IMAGEDIRS = fig

.PHONY: all impact.md

all: clean impact.pdf

abstract.pdf: images abstract.tex impact.bib
	pdflatex abstract.tex
	biber abstract
	pdflatex abstract.tex

impact.pdf: images impact.tex impact.bib
	pdflatex impact.tex
	biber impact
	pdflatex impact.tex

impact.md: impact.pdf
	pandoc impact.tex -o impact.md --biblio impact.bib

clean:
	rm -f *.aux *.bbl *.blg *-blx.bib *.log *.out *.run.xml *.toc *.bcf *~
	rm -f *.nlo *.ist *.nls *.ilg

distclean: clean
	rm -f impact.pdf abstract.pdf
	rm -rf auto/ _region_*
	$(foreach file, $(PDFs), $(shell rm -f "$(file)"))

# Pattern rule for converting SVG to PDF
%.pdf:  %.svg
	test -L "$<" || $(SVG2PDF)  -f pdf -o "$@" "$<"

# Build a list of SVG files to convert to PDFs
PDFs:=	$(foreach dir, $(IMAGEDIRS), $(patsubst %.svg,%.pdf,$(wildcard $(dir)/*.svg)))

# Make a rule to build the PDFs
images: $(PDFs)
