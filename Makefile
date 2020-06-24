xxx:
	@echo "Please select optimal option."

build:
	@go build -o pelican .

clean:
	@rm -f ./pelican

run:
	@go run .

test:
	@go test -v "./..."
