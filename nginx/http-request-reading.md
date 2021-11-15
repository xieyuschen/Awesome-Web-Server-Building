# http request message from tcp socket
&emsp;&emsp;As a asynchronized web framework,all operation in nginx like receiving & analyzing data from tcp socket and transfer to event module are all asynchronized. So state machine is worthing to consider as solution. However nginx didn't take this proposal, instead of this way nginx achieves it by setting different callback function at different state.  

&emsp;&emsp;It's clear there is anything member about state machine in structures stand for http request and tcp connection. Google shows [nginx mail man page](http://mailman.nginx.org/pipermail/nginx-devel/2010-May/000238.html) also confirmed it:
> &emsp;&emsp;There is no state machine per se. You need to add/remove your events to/from nginx's event model (which is OS-independent abstraction layer on top of epoll/kqueue/etc).   

&emsp;&emsp;Following content will explain mechanisms how nginx handles network data from tcp to http. Let's step into source code to see more details and additional I have received an advice from my tencent mentor: **The key to read source code is read code from top to bottom, pay attention to relative among variables with containment or constellation relative**.  

# source code 
&emsp;&emsp;First we look at nginx events module for no doubt of closed relation between http request and it. Nginx event module supported multiple method to fit different OS system call like `select`, `epoll`, `kqueue` and so on.  
@todo: functions and file position about things above.  

In `/src/ngx_event.h` there defines a lot of structs refer event. They are shows as following:  
- ngx_event_s  
Event module manages a chunk of events element, basic managed element is an instance of `ngx_event_s`. As a web server, basic event element has high possibility owning a data member about tcp connection and by coincidence there is a `ngx_connection_t` which presents tcp connection, however there is no such type in `nx_event_s`. It confused me until I went back to look for details about this, finally I found `void *data` represent such case. The reason why not use `ngx_connection_t` directly is event maybe a AIO file besides tcp connection.  
- ngx_event_actions_t
- ngx_event_conf_t
- ngx_event_module_t
- ngx_event_aio_s(optional, valid when aio is supported by OS)
- ngx_event_ovlp_t(optional, valid when `IOCP` is supported)

&emsp;&emsp;Then we shall pay attention to how `ngx_event_s` is initialized and managed, in short who managed it at whole framework. In linux it's normal to use `epoll` system call so let we see how epoll is used in nginx. 
- ngx_epoll_module.c  
This file is stored at `/src/event/modules/ngx_epoll_module.c`, it contains some necessary structs as a module like array of ngx_command_t and ngx_event_module_t context.  
Epoll module will be added into module array @todo




