version:
	genver.exe

gen:
	go generate

#build: gen
#	go build -o=genver.exe .

#install: gen
#	go install

build: version
	go build -o=genver.exe .

install: version
	go install
