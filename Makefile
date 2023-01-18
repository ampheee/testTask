complier=go	build
runnerPath=build/webPageComparer
bin=webPageComparer
buildPath=build/

all: run build

run: build
	./$(runnerPath)
	xdg-open "http://localhost:8080/when/"

build: 
	mkdir build
	$(complier) .
	mv $(bin) $(buildPath)