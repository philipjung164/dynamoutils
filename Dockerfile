FROM golang:1.10

WORKDIR /go/src/git.projectbro.com/isd/dynamoutils
ADD . .
CMD ["make", "test_ci"]
