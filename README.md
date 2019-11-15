## usage:
```shell
$ cd hash
$ g++ -c hash.cpp -std=c++11
$ g++ -c bridge.c -std=c++11
$ ar -crs libhash.a hash.o bridge.o
$ cd ..
$ go run main.go --path rdf_data.csv
```
This will generate a verte.csv and a edge.csv file in current directory, which is needed for [nebula-importer](https://github.com/vesoft-inc/nebula-importer)
