version:
	genver.exe

build: version
	go build -o=genver.exe .

install: version
	go install
