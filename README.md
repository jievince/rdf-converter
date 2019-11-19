因为nebula内部使用的hash是std::hash, 而std::hash所采用的哈希算法并不固定, 它是根据系统来决定的, 我们需要保证这个数据清洗工具生成的点的vid的hash算法和nebula内部保持一致. 所以我们需要想办法在我们这个Golang写的清洗工具里调用C++里的std::hash函数. 查阅资料可知, Golang官方有个cgo特性, 可以在Golang代码中调用C函数, 所以我们先用C包装C++ std::hash函数, 然后用cgo调用这个hash函数.


## 用法

### 先创建std::hash的静态库文件
```shell
$ cd hash
$ g++ -c hash.cpp -std=c++11
$ g++ -c bridge.c -std=c++11
$ ar -crs libhash.a hash.o bridge.o
```

### 开始清洗数据

使用--path参数指定知识图谱的三元组数据的路径,

```shell
$ cd ..
$ go run main.go --path rdf_data.csv
```

这条命令会在当前目录下生成一个vertex.csv文件和edge.csv文件, 分别是清洗后的顶点数据和边数据.

最后, 我们就可以愉快的使用[nebula-importer](https://github.com/vesoft-inc/nebula-importer)来导入vertex.csv和edge.csv这两个数据文件啦
