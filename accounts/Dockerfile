FROM golang

RUN mkdir -p /go/src/github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize

ADD . /go/src/github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts/cmd  -watch github.com/tryu-fullerton-edu/AMSE_CPSC546_Cartelize/accounts
