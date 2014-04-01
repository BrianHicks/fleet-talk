.PHONY: image clean upload dev

dev: clean
	go build -o app *.go

app: clean
	GOARCH=amd64 GOOS=linux go build -o app *.go

image: app
	docker build -t brianhicks/hellogo .

upload: image
	docker push brianhicks/hellogo

clean:
	test -f app && rm app || echo "nothing to clean"
