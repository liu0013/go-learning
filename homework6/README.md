# 总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。
~~~
fix length: 简单的自定义协议
delimiter based: SMTP
length field based frame: GRPC
~~~

# 实现一个从 socket connection 中解码出 goim 协议的解码器。