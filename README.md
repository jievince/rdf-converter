# 如何使用本工具

## 介绍

  这个工具用于清洗 [Ownthink](https://www.ownthink.com/) 的知识图谱 RDF 数据，将它变成属性图模型。产出结果为一个 vertex.csv 文件和 edge.csv 文件, 分别是清洗后的顶点数据和边数据。目前只对数据进行了简单去重。
  
  你也可以直接去 [kaggle](https://www.kaggle.com/littlewey/nebula-ownthink-property-graph) 下载完全去重后的数据。

## 如何使用

使用 --path 参数指定知识图谱的三元组数据的路径

```shell
$ go build
$ ./rdf-converter --path rdf_data.csv
```

这条命令会在当前目录下生成vertex.csv 文件和 edge.csv 文件。

之后, 就可以使用[nebula-importer](https://github.com/vesoft-inc/nebula-importer)来导入到 [Nebula 图数据库](https://github.com/vesoft-inc/nebula) 中啦。

说明：nebula-importer要求使用一个yaml配置文件去导入数据，你可以直接使用我写好的这个rdf-import.yaml。
