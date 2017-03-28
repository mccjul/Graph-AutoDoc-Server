FROM scratch
ADD $GOPATH/bin/Graph-AutoDoc-Server /Graph-AutoDoc-Server
CMD ["/Graph-AutoDoc-Server"]