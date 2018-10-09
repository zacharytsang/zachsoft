修改配置文件

进入自己安装idea路径的bin目录下，(不知道自己安装在哪儿的同学，请右击桌面的idea图标，选择属性，在目标这个地方就可以找到了)，将刚刚下载好的JetbrainsCrack.jar复制到此目录下；

还是在bin目录下，找到idea.exe.vmoptions和idea64.exe.vmoptions（比如你要激活pycharm,找到pycharm.exe.vmoptions和pycharm64.exe.vmoptions），用记事本打开它们，在两个文件最后分别加上：

-javaagent:D:\Develop\IntelliJ IDEA 15.0.6\bin\JetbrainsCrack-2.6.2.jar

3. 修改Activation Code

重启idea，进入Register(在Help中有，或者进入项目之前选择工程的界面也有)

选择Activation Code，粘贴如下代码，记得修改下面的用户名部分和邮箱部分：


ThisCrackLicenseId-{
"licenseId":"ThisCrackLicenseId",
"licenseeName":"你想要的用户名",
"assigneeName":"",
"assigneeEmail":"随便填一个邮箱(我填的:idea@163.com)",
"licenseRestriction":"For This Crack, Only Test! Please support genuine!!!",
"checkConcurrentUse":false,
"products":[
{"code":"II","paidUpTo":"2099-12-31"},
{"code":"DM","paidUpTo":"2099-12-31"},
{"code":"AC","paidUpTo":"2099-12-31"},
{"code":"RS0","paidUpTo":"2099-12-31"},
{"code":"WS","paidUpTo":"2099-12-31"},
{"code":"DPN","paidUpTo":"2099-12-31"},
{"code":"RC","paidUpTo":"2099-12-31"},
{"code":"PS","paidUpTo":"2099-12-31"},
{"code":"DC","paidUpTo":"2099-12-31"},
{"code":"RM","paidUpTo":"2099-12-31"},
{"code":"CL","paidUpTo":"2099-12-31"},
{"code":"PC","paidUpTo":"2099-12-31"}
],
"hash":"2911276/0",
"gracePeriodDays":7,
"autoProlongated":false}

