# logging

注意：该模块目前还在不停开发中。

## 简介

logging是一个类似python logging模块的日志包，拥有基本的handler、父类继承关系等特性。主要提供日志相关操作。

## 用法

```go
l := logging.GetLogger("x.y.z")
l.Debug("Hello World")
```
