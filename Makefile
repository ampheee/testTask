name=dateComparer
complier=go	build
buildPath=build/
appPath=app/
app=main.go
all: run build

run: build
	./$(buildPath)$(name)
	xdg-open "http://localhost:8080/when/"

build: 
	mkdir build
	$(complier) $(appPath)$(app)
	mv main $(name) && mv $(name) $(buildPath)

rebuild: clean build
rerun: clean run

clean:
	rm -rf $(buildPath) main

