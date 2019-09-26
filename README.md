# go-owaddress

本项目提供区块链多币种的地址合法性校验底层支持。

## 如何使用

```$xslt
valid, err := owaddress.Verify("btc", address)

valid - bool  : 地址是否合法
err   - error : 币种未注册时会返回错误
```

## 如何添加新币种

```$xslt
1.  在coins目录下新建待添加币种的目录 
2.  在币种目录下定义AddressVerify结构体，继承自address.AddressVerify接口
3.  在该目录下定义 DefaultStruct 和 CoinName = "btc" 用于币种注册
4.  实现AddressVerify的IsValid方法
5.  在owaddress包的register.go的init方法中完成币种注册
6.  在test包下完成测试案例的添加
```