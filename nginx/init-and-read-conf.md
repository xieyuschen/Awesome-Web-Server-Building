# How nginx start up and read config file
As a complex web server, nginx owns many modules to make sure achieving its aim. For a newer to its source code and architecture it's easy to learn about nginx from top to bottom. Learn about how nginx modules like reading config, load event module can make me better understood.  

Nginx has many excellent designs, for instance event-driven framework, full asynchronized handling, memory pool and so on... It's well to make clear how nginx startup. 

## Nginx startup analysis
- Analyzing command line  
This step will read the most famous file name nginx.conf to initialize nginx, it will create a temperately ngx_cycle_t struct storing values to implement the formal ngx_cycle_t. The temperately struct do many things, like: @todo.  

- Smoothly upgrading or not  
Smoothly upgrading is a common case needed to solve as a web server, so how nginx do for this? Based on my experience, I think to upgrading smoothly for a web server needs to do followings:  
    - Fork a new process and use `exce` system call to restart service process. Stop establishing new connections until `exce` returns. By the way, how to notify it? Signal or anything else between parent and child process?  
    - New connections will be handled by child process and old connections are handled by parent process.  
    - Old process handles all of connection and quit for some flag value was set when fork.  
    - Send a signal to child process when parent process quit.  
    