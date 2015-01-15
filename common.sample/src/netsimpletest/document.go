//document.go
/*
一：传统Socket编程步骤
1.建立Socket：使用Socket()函数。
2.绑定Socket：使用bind()函数。
3.监听：使用listen()函数；或者链接，使用connect()函数。
4.接受连接：使用accept()函数。
5.接收：使用receive()函数，或者发送：使用Send()函数。
*/
/*
GO语言Socket编程
1.无论使用什么协议建立什么形式的连接，都只需要使用net.Dial()即可。
Dial() 函数原型：
func Dial(net,addr string)(Conn,error)
*/
/*
常见协议的调用方式
tcp连接：conn,err := net.Dial("tcp","192.168.0.10:2100")
udp连接：conn,err :=net.Dial("udp","192.168.0.10:2100")
ICMP连接(使用协议名称)：conn,err := net.Dial("ip4:icmp","www.baidu.com")
ICMP连接(使用协议编号)：net.Dial("ip4:1","10.0.0.3")
*/
/*
Dial()函数支持的网络协议：
"tcp"、"tcp4"(仅限ipv4)、"tcp6"(仅限ipv6)
、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ipv6"
*/
package netsimpletest
