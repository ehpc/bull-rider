PROTOBUFDIR=$(CURDIR)/protobuf
PROTOSDIR=$(PROTOBUFDIR)/protos
PROTOBUFGODIR=$(PROTOBUFDIR)/go
MODNAME=github.com/ehpc/bull-rider

PROTOC=protoc

all: clean protobuf

clean:
	@rm -rf $(PROTOBUFGODIR)

protobuf: protobuf_go

protobuf_go:
	@mkdir $(PROTOBUFGODIR)
	@for f in $$(find $(PROTOSDIR) -type f -printf '%P\n' -name *.proto); do \
		$(PROTOC) -I$(PROTOSDIR)/ --go_out=paths=source_relative:$(PROTOBUFGODIR)/ $(PROTOSDIR)/$$f; \
	done
