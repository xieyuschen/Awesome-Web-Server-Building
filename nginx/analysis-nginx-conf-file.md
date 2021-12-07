# How nginx analyzes the nginx.conf file
Everybody who has once used nginx will know how important of nginx.conf file is. This article will foucs on following questions:  
- how nginx endless loop is managed?
- when this file is loaded
- how nginx analyzes nginx.conf and executes callback functions  


## 1.How nginx endless loop is managed?
As a web server, it must has a endless loop to handle all things needs to be solved by server. Here we just talk about single process so we only need to foucs on function `ngx_single_process_cycle` called at `\src\core\nginx.c:380` defined at `\src\os\unix\ngx_process_cycle.c:279`.  
@todo: Reading it's implement code, we found that there is a typedef `sig_atomic_t` as int type, how it achieves atomic?  

When reading this endless loop to find the way to initialize config, we found some useful things ignored before.  

### 1.1 Set Environment which stored in Config struct(NOT OS ENV)
Nginx loopless function calls method to set environment first- at `\src\os\unix\ngx_process_cycle.c:283`. It set Timezone while environment is NULL which means it's not forked by another process. The environment is `char**` and stored in `ngx_core_conf_t`.  

### 1.2 Call each module's init_process
init_process part in every module is usually NULL so it rarely.
```c
    for (i = 0; cycle->modules[i]; i++) {
        if (cycle->modules[i]->init_process) {
            if (cycle->modules[i]->init_process(cycle) == NGX_ERROR) {
                /* fatal */
                exit(2);
            }
        }
    }
```
### 1.3 Enter endless loop, handle event and timer
Obviously function `ngx_process_events_and_timers` called at `\src\os\unix\ngx_process_cycle.c:300` has already started to handle network event and timer so initialization is finished before call stack steps into `ngx_process_events_and_timers`. If we want to look up how nginx.conf is handled, we must find in front of call stack.
@todo: three condition judges in endless loop.

## 2. When nginx.conf file is loaded by nginx?
