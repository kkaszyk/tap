build:
	go build -o bin/pat main.go trace_format.go trace_processor.go trace_parser.go

run : build
	./bin/pat traces/sample_trace

clean:
	/bin/rm -rf bin
