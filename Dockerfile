FROM golang:1.23 as builder

RUN useradd -ms /bin/bash secureuser
WORKDIR /home/secureuser/app

COPY go.mod ./
RUN chown -R secureuser:secureuser /home/secureuser/app

USER secureuser

RUN go mod tidy

COPY . ./

RUN go build -buildvcs=false -o securevault

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /home/secureuser/app/securevault .

CMD ["./securevault"]

