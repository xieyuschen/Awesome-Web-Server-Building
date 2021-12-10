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

## Gin Part
@todo: to be finished


 
