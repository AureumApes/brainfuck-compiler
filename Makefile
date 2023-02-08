build:
	GOOS=linux go build -o $(PROGRAMNAME)-linux .
	GOOS=windows go build -o $(PROGRAMNAME)-windows.exe .
