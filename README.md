# eft-bot
基于[LagrangeGo](github.com/LagrangeDev/LagrangeGo)实现的逃离塔科夫机器人

A bot implementation based on [LagrangeGo](github.com/LagrangeDev/LagrangeGo) for video game Escape From Tarkov

## Usage
### Run Bot
首先你需要下载Golang编译环境 [website](https://golang.org/)

接着下载源代码并编译为二进制文件.

接着你需要将源代码中的配置文件`etc/config copy.yaml`重命名为`etc/config.yaml`并将其移动到相对于二进制文件的`etc/config.yaml`路径.

最后修改配置文件中的相应配置信息,执行二进制文件,扫码登录QQ`etc/qrcode.png`

First, you need to download the Golang compilation environment from [website](https://golang.org/).

Next, download the source code and compile it into a binary file.

Then, rename the configuration file etc/config copy.yaml from the source code to etc/config.yaml and move it to the etc/config.yaml path relative to the binary file.

Finally, modify the corresponding configuration information in the configuration file, execute the binary file, and scan the QR code to log in to QQ etc/qrcode.png.

### Interact With Bot
在bot订阅的群聊中输入合法的命令即可获得bot的回应.

一条合法的命令由前缀和指令两部分组成.

前缀默认为`eft`,你可以在配置文件中修改

To receive a response from the bot in a group chat where the bot is subscribed, you need to input a valid command.

A valid command consists of two parts: a prefix and an instruction.

The default prefix is eft, but you can modify it in the configuration file.

### Current Supported Command

| command        | meaning                                                   |
| -------------- | --------------------------------------------------------- |
| 帮助           | 查看当前可用的指令                                        |
| help           | Display current supported command                         |
| 汇率           | 展示游戏中美元,欧元汇率                                   |
| exchange       | Display the US dollar and euro exchange rates in the game |
| x87 {num}      | Calculate item's price in euro through 口关's theory      |
| Other Input... | Display other item's flea market information              |
