PREFIX ?= /usr
BINDIR ?= $(PREFIX)/bin

all: trt

trt:
	go build

install: all
	install -d $(DESTDIR)$(BINDIR)
	install -m 755 trt $(DESTDIR)$(BINDIR)

uninstall:
	rm -f $(DESTDIR)$(BINDIR)/trt

clean:
	rm -f trt

.PHONY: all install uninstall clean
