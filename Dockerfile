FROM golang:1.19

WORKDIR /app

# This files are needed to install the needed dependencies
COPY go.mod ./
COPY go.sum ./

# Download all dependescies
RUN go mod download

# Copy files
COPY main.go ./main.go
COPY cmd ./cmd

# Build bot
RUN go build -o /bot

# Run bot
CMD [ "/bot" ]