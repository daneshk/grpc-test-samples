FROM golang:1.16-alpine

# Set destination for COPY
WORKDIR /server

#RUN go mod init productinfo/server

# Download Go modules
COPY go.mod .
COPY go.sum .
RUN go mod download

#RUN go get google.golang.org/grpc/internal/channelz@v1.24.0

COPY . .

RUN go build -v -o ./bin/server

#
## Build stage II : Go binaries are self-contained executables.
#FROM scratch
#COPY --from=build /bin/server /bin/server


ENTRYPOINT ["bin/server"]
EXPOSE 50051