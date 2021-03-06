//document.go
/*棋牌游戏项目-玩家基本操作
(1)登陆游戏
(2)查看房间列表
(3)创建房间
(4)加入房间
(5)进行游戏
(6)房间内聊天
(7)游戏完成，退出房间
(8)退出登陆
*/

/*
一：用户登陆流程详细设计
(1)用户唯一ID
(2)用户名，用于显示
(3)玩家等级
(4)经验值
实际的游戏设计当然要比这个复杂得多，比如还有社交关系、道具和技能等。鉴于这些细节 并不影响架构，这里我们都一并略去。
*/
/*功能划分
1.玩家会话管理系统，用于管理每一位登陆的玩家，包括玩家信息和玩家状态
2.大厅管理
3.房间管理系统，创建、管理和销毁每一个房间
4.游戏会话管理系统，管理房间内的所有动作，包括游戏进程和房间内聊天
5.聊天管理系统，用于接收管理员的广播信息

*/
//中央服务器
/*
功能：
1.在线玩家的状态管理。
2.服务器管理。
3.聊天系统。
*/
package cg
