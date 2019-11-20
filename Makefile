PREFIX=/usr/local
BINDIR=${PREFIX}/bin
DESTDIR=
BLDDIR = build
BLDFLAGS=
EXT=
ifeq (${GOOS},windows)
    EXT=.exe
endif

# APPS = conduit nsqlookupd nsqadmin nsq_to_nsq nsq_to_file nsq_to_http nsq_tail nsq_stat to_nsq
APPS = conduit
all: $(APPS)

$(BLDDIR)/conduit:        $(wildcard cmd/conduit/*.go internal/app/conduit/*.go)
# $(BLDDIR)/nsqlookupd:  $(wildcard apps/nsqlookupd/*.go nsqlookupd/*.go nsq/*.go internal/*/*.go)

$(BLDDIR)/%:
	@mkdir -p $(dir $@)
	go build ${BLDFLAGS} -o $@ ./cmd/$*

$(APPS): %: $(BLDDIR)/%

clean:
	rm -fr $(BLDDIR)

.PHONY: install clean all
.PHONY: $(APPS)

install: $(APPS)
	install -m 755 -d ${DESTDIR}${BINDIR}
	for APP in $^ ; do install -m 755 ${BLDDIR}/$$APP ${DESTDIR}${BINDIR}/$$APP${EXT} ; done