# Awesome-Web-Server-Building
This project is a guide of building web server framework with some source code analysis of Nginx and Gin. It's useful to learn  how they designed to build a high performance wen server.   

This repository contains two main web server to analysis, one is `Nginx` and another is `Gin`. The reason for chosen those two is that I used Gin for web developing more than 1 year to build a web service and I have read a book about Nginx recently. 


## Source code version 
I chose [v1.7.0](https://github.com/gin-gonic/gin/tree/v1.7.0) for gin.  

## Remark
Ich habe etwas zwei Jahre Deutsche lernen fur meine Hobby. Weil Ich lerne Deutsche, hoffe ich auf Deustsch zu bloggen. Die Blogs moglich viele Felher aber Ich hoffe dass ich gut vor dem Abschluss schreiben kann:)  
Einige Inhalte werde ich noch einmail auf Deutsch schreiben.
# Content
## 1. State machine for dealing http request
When it comes to state machine, I fall into a nightmare for my Compiler course because it consists quite a lot of state machine. Back to my topic here, state machine is quite useful when we deals about a stream based transit protocol like tcp. Http based on every single request which like a packet but tcp is a stream which means there is no clear separation between two closed request.  
State machine which can convert its state as it receives a stream of characters.

## 2. How nginx reads http request from tcp socket?  

 
