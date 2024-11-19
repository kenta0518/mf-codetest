FROM mcr.microsoft.com/vscode/devcontainers/go:1.20-bullseye

WORKDIR /app

ENV TZ=Asia/Tokyo

COPY ../go.mod ../go.sum ./
RUN go mod download

RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/go-delve/delve/cmd/dlv@v1.22.0
RUN go install honnef.co/go/tools/cmd/staticcheck@v0.4.0
RUN go install golang.org/x/tools/gopls@latest
RUN go install github.com/golang/mock/mockgen@v1.7.0-rc.1
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY ../ ./

RUN sudo chown -R vscode:vscode /go
