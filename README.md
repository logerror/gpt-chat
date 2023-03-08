## 预览
Mac:
![image](https://user-images.githubusercontent.com/13896386/222678871-12582f0d-1995-46ca-a016-2bf348a55758.png)

Windows:
![image](https://user-images.githubusercontent.com/13896386/223722120-7ce80f70-f29e-4d3f-8d0c-8139e31ea403.png)
也可以新建一个bat脚本方便使用`start cmd /k gpt-chat.exe sk-jXXXXXXXXXXXXXXXXX`

Linux:
same as mac



## 如何使用
1. 访问  https://platform.openai.com/account/api-keys 获取你的secret key
2. 下载地址,根据操作系统及cpu架构选择压缩包解压即可 (https://github.com/logerror/gpt-chat/releases)
   3. 在工具的同级目录下创建一个名为`.env`的文件,内容如下
   ```
   # you can find you api-key in: https://platform.openai.com/account/api-keys
   OPEN_AI_SECRET_KEY=""
   BOT_NAME="Bot"
   MODEL="gpt-3.5-turbo-0301"
   ```
4. 你可以用以下两种方式运行gpt-chat
   1. gpt-chat "you secret key"
   2. 使用环境变量声明SecretKey export OPEN_AI_SECRET_KEY="you secret key" && gpt-chat



