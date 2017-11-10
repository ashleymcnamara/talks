protoc -I$GOPATH/src --go_out=plugins=micro:$GOPATH/src \
	$GOPATH/src/github.com/bketelsen/talks/codemotion/greeting/proto/greeting/greeting.proto
