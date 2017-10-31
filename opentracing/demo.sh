docker run -d -p 5775:5775/udp -p 6831:6831/udp -p 16686:16686 jaegertracing/all-in-one:latest

cd $GOPATH/src/github.com/jaegertracing/jaeger

make install

cd $GOPATH/src/github.com/jaegertracing/jaeger/examples/hotrod

go run main.go all
