FROM alpine

LABEL author="JitenP@Outlook.Com"

WORKDIR /app

COPY app .

ENTRYPOINT [ "./app","--count=100"]

# TO build go lang application for alpine based images 
# CGO_ENABLED=0 go build -a -installsuffix cgo -o app main.go

# sha256:490a52f647f2e27378f946935f06cb086e91e32cfec2874cd8ea78c2cfa36c40
# sha256:6be6c1b286265ab481a88a4d7a2dd53aea828391a70e0c6c975f7b98726ba724
# sha256:6be6c1b286265ab481a88a4d7a2dd53aea828391a70e0c6c975f7b98726ba724