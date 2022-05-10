.PHONY: test clean

test:
	 go test ./...

clean:
	git clean -fXd