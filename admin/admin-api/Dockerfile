FROM golang:1.14 as builder

# Build the binary statically.
ENV CGO_ENABLED=0

WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o admin-api .


FROM alpine:3.10.2

RUN apk add -U --no-cache ca-certificates
RUN mkdir -p /var/log/app

WORKDIR /app
COPY --from=builder /app/admin-api .
COPY --from=builder /app/casbin_rbac_model.conf .
COPY --from=builder /app/casbin_rbac_policy.csv .
COPY --from=builder /app/templates ./templates
COPY config.yml .

CMD ["sh","-c","/app/admin-api 2>&1 | tee -a /var/log/app/app.log"]
