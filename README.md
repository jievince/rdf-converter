## usage:
```shell
cd hash
g++ -c hash.cpp -std=c++11
g++ -c bridge.c -std=c++11
ar -crs libhash.a hash.o bridge.o
go run main.go --path rdf_data.csv
```
