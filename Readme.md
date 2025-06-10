## 1. 程序运行说明
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
