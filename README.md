# 如何使用本工具

## 前言

1. 这个工具用于清洗 [Ownthink](https://www.ownthink.com/) 的知识图谱 RDF 数据，将它变成属性图模型。产出结果为一个 vertex.csv 文件和 edge.csv 文件, 分别是清洗后的顶点数据和边数据.

1. 这个工具是用 Golang 编写的，Golang 环境的安装配置参考[文档](docs/golang-install.md)。

## 如何使用

### 先创建 std::hash 的静态库文件

```shell
$ cd hash
$ g++ -c hash.cpp -std=c++11
$ g++ -c bridge.c -std=c++11
$ ar -crs libhash.a hash.o bridge.o
```

### 开始清洗数据

使用 --path 参数指定知识图谱的三元组数据的路径

```shell
$ cd ..
$ go run main.go --path rdf_data.csv
```

这条命令会在当前目录下生成vertex.csv 文件和 edge.csv 文件。

之后, 就可以使用[nebula-importer](https://github.com/vesoft-inc/nebula-importer)来导入到 [Nebula 图数据库](https://github.com/vesoft-inc/nebula) 中啦。

## 附：关于 hash 的说明

Nebula 使用的 hash 是 `MurmurHash2`（和 C++ std::hash 一致），所以清洗工具使用了 `cgo` 来调用 C++ 里的std::hash函数。

```cpp
// https://github.com/vesoft-inc/nebula/blob/master/src/common/base/MurmurHash2.h
uint64_t seed = 0xc70f6907UL;
const uint64_t m = 0xc6a4a7935bd1e995;
const uint32_t r = 47;
uint64_t h = seed ^ (size * m);
```
