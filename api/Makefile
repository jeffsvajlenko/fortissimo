all: go cs python
go: go/fortissimo.pb.go
cs: cs/bin/debug/FortissimoApi.0.0.1.nupkg
python:

go/fortissimo.pb.go: fortissimo.proto
	@protoc --go_out=plugins=grpc:go/fortissimo --go_opt=paths=source_relative  fortissimo.proto
	@cd go && go mod tidy

cs/Fortissimo.cs:
	@protoc --csharp_out=cs fortissimo.proto

cs/bin/debug/FortissimoApi.0.0.1.nupkg: cs/Fortissimo.cs
	@cd cs/ && dotnet build && dotnet pack

python/fortissimo_pb2.py:
	@protoc --python_out=python fortissimo.proto

clean:
	@rm -rf go/fortissimo/fortissimo.pb.go cs/Fortissimo.cs cs/bin/ cs/obj/ python/