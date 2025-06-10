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


=============注意：首次打开面板浏览器将提示不安全=================

 请选择以下其中一种方式解决不安全提醒
 1、下载证书，地址：https://dg2.bt.cn/ssl/baota_root.pfx，双击安装,密码【www.bt.cn】
 2、点击【高级】-【继续访问】或【接受风险并继续】访问
 教程：https://www.bt.cn/bbs/thread-117246-1-1.html
 mac用户请下载使用此证书：https://dg2.bt.cn/ssl/mac.crt

========================面板账户登录信息==========================

 【云服务器】请在安全组放行 18869 端口
 外网面板地址: https://111.173.117.113:18081/d1efb521
 内网面板地址: https://100.100.110.247:18081/d1efb521
 username: qqpssfcb
 password: 7a5836bb

 浏览器访问以下链接，添加宝塔客服
 https://www.bt.cn/new/wechat_customer
