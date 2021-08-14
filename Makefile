GOCMD=go
EXE=fnufetch

${EXE}: *.go */*.go
	${GOCMD} build -o ${EXE}

tidy:
	gofmt -s -w .
	go mod tidy
	go vet

install: ${EXE}
	cp ./${EXE} /usr/bin/${EXE}

userinstall: ${EXE}
	cp ./${EXE} ~/.local/bin/${EXE}

clean:
	rm ${EXE}

