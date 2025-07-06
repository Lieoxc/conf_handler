## 1. 服务器后端Golang程序运行说明
本程序主要功能
1. clash原始订阅解析，并且把所有节点数据保存到数据库
2. 支持页面组合新订阅，然后配置为新的url，设备端订阅这个新的url，然后返回calsh订阅信息
3. 手动解析clash订阅内容，把节点信息保存到数据库

注意: 需要和页面PHP程序一起配合使用
### 启动程序： 
1. cd /home/xcli/work
2. nohup  ./cfg_handler  & 

### 关闭程序
killall cfg_handler

### 检查程序是否正在运行：

ps -aux | grep cfg_handler 
输出如下：
```
xcli       40045  0.5  0.3 719740 14720 pts/3    Sl   22:35   0:00 ./cfg_handler
xcli       40053  0.0  0.0  10264  2816 pts/3    S+   22:35   0:00 grep --color=auto cfg_handler

```
