# Agenda

> 课程《服务计算》第七周作业：用 Go 实现命令行 Agenda


## 运行

```
go get github.com/lishicheng/Agenda
$GOPATH/bin/Agenda
```

## 注意事项

- 由于使用了 cobra 需要科学上网才能编译运行

## 配置文件

默认使用 ``$HOME/.agenda-go.yaml``

如果找不到的话，将使用如下的默认设置

```yaml
# 工作目录。其他配置如果使用相对路径，则相对该工作目录
cwd: .
# log 的路径。如果环境变量里有 DEBUG 变量，则 log 会输出到 stderr
log: /dev/null
# 用户数据的路径。JSON 格式
user_data: data/userinfo.json
# 会议数据的路径。JSON 格式
meeting_data: data/meetinginfo.json

```

## 项目分工

- 李仕成(master)
  1.项目框架和架构设计
  2.用户实体、服务、UI
  3.需求实现：用户注册、登录、登出
  4.接入持续集成
  5.README编写
- 黄梓锋(contributor)
  1.命令设计
  2.实现用户查询、删除、会议创建、增删参与者、查询、会议取消、退出、清空
  3.会议实体、服务、UI
  4.代码测试


## 任务目标

 - 熟悉 go 命令行工具管理项目
 - 综合使用 go 的函数、数据结构与接口，编写一个简单命令行应用 agenda
 - 使用面向对象的思想设计程序，使得程序具有良好的结构命令，并能方便修改、扩展新的命令不会影响其他命令的代码
 - 项目部署在 Github 上，合适多人协作，特别是代码归并
 - 支持日志（原则上不使用debug调试程序）


## 具体运行测试示例
- 登录
	- Command: `user_login -u=username -p=password`
- 登出
	- Command: `logout`
若有用户登录则打印相应登出信息，若无用户登录则提示出错
- 注册
	- Command: `user_register -u=username -p=password -m=mail -c=cellphone`
用户名具有唯一性，不允许与现有用户名相同

- 删除账户
	- Command: `delete_accout`
	如无用户登录则返回出错信息
	如有用户登录则删除用户作为发起者的会议，同时将用户从作为参与者的会议中删除
	删除登录信息
- 查询用户
	- Command: `user_query -u=username`

- 创建会议
	- Command: `create_meeting -t=title -p=participator -s=starttime -e=endtime`
		只允许已登录用户进行此操作
		会议标题具有唯一性
		参与者必须为注册用户
		参与者数量 >= 1
		开始时间与结束时间应符合逻辑，采用24小时制
		不允许发起者或参与者在此时间段内有其他会议
- 会议查询
	- Command: `query_meeting -s=starttime -e=endtime`
- 取消会议
	- Command: `delete_meeting -t=title`
	仅允许已登录用户删除自己发起的会议
- 增加会议参与者
	- Command: `participator_add -t=title -p=participator`
- 删除会议参与者
	- Command: `participator_remove -t=title -p=participator`
	仅允许已登录用户对自己发起的仍存在的会议进行删除操作, 否则返回出错信息
- 退出会议
	- Command: `meeting_quit -t=title`
    仅允许已登录用户退出自己参加的会议
	若因此导致会议参与人数为0则删除会议
- 清空会议
	- Command: `clear_meeting`
删除已登录用户发起的所有会议
