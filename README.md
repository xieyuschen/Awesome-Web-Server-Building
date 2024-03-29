# Awesome-Web-Server-Building
This project is a guide of building web server framework with some source code analysis of Nginx and Gin. It's useful to learn  how they designed to build a high performance web server.   

This repository contains two main web server to analysis, one is `Nginx` and another is `Gin`. The reason for chosen those two is that I used Gin for web developing more than 1 year to build a web service and I have read a book about Nginx recently. 


## Source code version 
I chose [v1.7.0](https://github.com/gin-gonic/gin/tree/v1.7.0) for gin, [nginx-1.20.1](https://github.com/nginx/nginx/tree/release-1.21.1) for nginx. 

## Remark
Ich habe etwas zwei Jahre Deutsche lernen fur meine Hobby. Weil Ich lerne Deutsche, hoffe ich auf Deustsch zu bloggen. Die Blogs moglich viele Felher aber Ich hoffe dass ich gut vor dem Abschluss schreiben kann:)  
Einige Inhalte werde ich noch einmail auf Deutsch schreiben.

# Content
## Nginx Part
|Rank|Title|Link|
|--|:--|--|
|0|Questions when I read nginx guide book|[here](./nginx/0.question.md)|
|1|How nginx start up and read config file|[here](./nginx/1.init-and-read-conf.md)|
|2|Steps of configuration file to stored in ngx_cycle_t|[here](./nginx/2.step-of-config-stored.md)|
|3|How nginx loads the nginx.conf file|[here](./nginx/3.load-nginx-conf-file.md)|
|4|How nginx analyzes the nginx.conf file|[here](./nginx/4.nginx-analyzes-conf-file.md)|
|5|Talk about conf_ctx and switch context when execute command callback|[here](./nginx/5.context-switch-when-call-cmd-callback.md)|
|6|How nginx handles http connections|[here](./nginx/6.how-nginx-deal-connection.md)|
|7|How nginx event module works|[here](./nginx/7.how-nginx-event-module-works.md)|
|8|Who manages event module|[here](./nginx/8.who_manages_event_module.md)|
|9|Http module in nginx: receive data from connection and solve it by eleven phases|[here](./nginx/9.http-module-in-nginx-1.md)|
|10|Http module in nginx: give client response and end a request|[here](./nginx/10.http-module-in-nginx-2.md)|
|11|Upstream module|@todo|
|12|Master-Work processes model in nginx|@todo|
|13|Load balance module|@todo|
|14|Limiting flow module|@todo|
|15|Log module|@todo|
|16|Memory management in nginx|@todo|

## Gin Part
|Rank|Title|Link|
|--|:--|--|
|0|Summary about this part|[here](./gin/0.summary.md)|
|1|Package net/http in golang|[here](gin/1.net-http-package-in-go.md)|


# Reference meterials
- Understanding Nginx, Modules Development and Architecture Resolving(Second Edition)  
深入理解Nginx,模块开发与架构解析(第二版), ISBN: `978-7-111-52625-4`  
- Analysis of Nginx Source Code  
Nginx底层设计与源码分析, ISBD: `978-7-111-68274-5`  
- Nginx完全开发指南,使用C,C++,Javascript和Lua  
ISBN:`978-7-121-36436-5`  

