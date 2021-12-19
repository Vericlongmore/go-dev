#####总结几种 socket 粘包的解包方式：fix length/delimiter based/length field based frame decoder。尝试举例其应用。

怎么处理粘包
方式1: fix length
发送方，每次发送固定长度的数据，并且不超过缓冲区，接受方每次按固定长度区接受数据
方式2: delimiter based
发送方，在数据包添加特殊的分隔符，用来标记数据包边界
方式3: length field based
发送方，在消息数据包头添加包长度信息



Netty提供了4种常用解码器：

LineBasedFrameDecoder - 换行解码器
DelimiterBasedFrameDecoder - 分隔符解码器
FixedLengthFrameDecoder - 定长解码器
LengthFieldBasedFrameDecoder - 消息头定长解码器
