CC=$(shell which go)
PYTHON=$(shell which python)
GOBUILD=$(CC) build
GOGENERATE=$(CC) generate
EXECUTABLE=ProtoBufDemo

TMPLIST := data_pb2.py data_pb2.pyc tst.bin ProtoBufDemo data.pb.go

.PHONY := clean all run

all:
	@$(GOGENERATE)
	@$(GOBUILD)

clean:
	@rm -f $(TMPLIST)
	
run: clean all
	@$(PYTHON) writer.py
	@./$(EXECUTABLE)
	