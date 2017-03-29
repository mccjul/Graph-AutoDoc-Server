FROM golang:onbuild

# Make the source code path
# RUN mkdir -p /go/src/github.com/mccjul/Graph-AutoDoc-Server

# # Add all source code
# ADD . /go/src/github.com/mccjul/Graph-AutoDoc-Server

# # Run the Go installer
# RUN go install /go/src/github.com/mccjul/Graph-AutoDoc-Server

# # Indicate the binary as our entrypoint
# ENTRYPOINT /go/bin/Graph-AutoDoc-Server

# # Expose your port
EXPOSE 1323