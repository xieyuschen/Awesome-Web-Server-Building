# 1. How nginx start up and read config file
As a complex web server, nginx owns many modules to make sure achieving its aim. For a newer to its source code and architecture it's easy to learn about nginx from top to bottom. Learn about how nginx modules like reading config, load event module can make me better understood.  

Nginx has many excellent designs, for instance event-driven framework, full asynchronized handling, memory pool and so on... It's well to make clear how nginx startup. 

## Nginx startup analysis
### 1. Analyzing command line  
This step will read the most famous file name nginx.conf to initialize nginx, it will create a temperately ngx_cycle_t struct storing values to implement the formal ngx_cycle_t. The temperately struct do many things, like: @todo.  

### 2.Smoothly upgrading or not  
Smoothly upgrading is a common case needed to solve as a web server, so how nginx do for this? Based on my experience, I think to upgrading smoothly for a web server needs to do followings:
>    - Fork a new process and use `exce` system call to restart service process. Stop establishing new connections until `exce` returns by closing listening socket in parent process. By the way, how to notify it when fork finished and notify this is a smoothly upgrading action? Signal or anything else between parent and child process?  
>    - New connections will be handled by child process and old connections are handled by parent process.  
>    - Old process handles all of connection and quit for some flag value was set when fork.  
>    - Send a signal to child process when parent process quit.  

&emsp;&emsp;So let's see how nginx achieves it, it's much more similar than I mentioned above except how it notifies child process it's time to smoothly smoothly upgrading. Nginx use **Environment variables* to make process communication.  
&emsp;&emsp;Additional, sharing listening socket among multiple process will happen **shocking group** which will cause the degradation of system performance. It will be introduced later.  

### 3. Allocate memory for configuration structs
This step is finished by function `ngx_init_cycle`, it is first called at `/src/core/nginx.c:Line292` in function `main`. Let's take a closer look through main and init cycle function. First it use temperately ngx_cycle_t struct to store log, args passed from command line and so on. Then it use this temperately struct to initialize formal ngx_cycle_t struct by calling ngx_init_cycle. 
`ngx_init_cycle` first copy value from old cycle passed as arguments, and then it do the most important things to allocate memory for configuration file. Start at line 232, it allocates memory for all core modules as this:  
```c
for (i = 0; cycle->modules[i]; i++) {
        if (cycle->modules[i]->type != NGX_CORE_MODULE) {
            continue;
        }

        module = cycle->modules[i]->ctx;

        if (module->create_conf) {
            rv = module->create_conf(cycle);
            if (rv == NULL) {
                ngx_destroy_pool(pool);
                return NULL;
            }
            cycle->conf_ctx[cycle->modules[i]->index] = rv;
        }
    }
```
As code shows above, two points about variables initialization in cycle struct should be clear.  
- One is at this time array of cycle->modules is already assigned  
Knowing time when array of cycle->modules is initialized makes us understood it better, read carefully about codes, it quite do initialize cycle->modules in `ngx_init_cycle` at line 225 in `/src/core/ngx_cycle.c` by calling `ngx_cycle_modules`. Just as always it done before, this function allocates memory, check valid and then copy arrays from origin. Additional modules ordered array and their names are stored in `objs/ngx_modules.c` which is generated at `./configure` point as plain text before compiling, so when nginx startup and init, those arrays can be used directly.  
- The other is only **NGX_CORE_MODULE** will be executed to allocate memory  

It's good to know which module belongs to `NGX_CORE_MODULE` so look back of module definition which implements by official authors or triple authors. Each modules needs to own those three structs as following:  
- **ngx_XXX_module** exposed as module interface to stand the whole module which contains the next two structs.  
- **ngx_XXX_module_ctx** contains `create_conf` and `init_conf` used in startup, they're found by module interface.  
- **ngx_XXX_commands** shows the config key it interested.   

Module which not belongs to NGX_CORE_MODULE type will be solved later @todo(show how they are initialized).  

Codes above also refers to how nginx manages its config settings, I will tell it later @todo. By the way, check whether a variable is empty after allocating is necessary. 


### 4. Analyzing nginx.conf by core modules
@todo, Cannot understand source code when reading `src/core/ngx_cycle.c` from line 249 to 295:(   
### 5. Initialize allocated structs according to nginx.conf
This turn still in calling stack of `ngx_init_cycle`, init is similar to create variable to store. Codes are followings:  
```c
    for (i = 0; cycle->modules[i]; i++) {
        if (cycle->modules[i]->type != NGX_CORE_MODULE) {
            continue;
        }

        module = cycle->modules[i]->ctx;

        if (module->init_conf) {
            if (module->init_conf(cycle,
                                  cycle->conf_ctx[cycle->modules[i]->index])
                == NGX_CONF_ERROR)
            {
                environ = senv;
                ngx_destroy_cycle_pools(&conf);
                return NULL;
            }
        }
    }
```

### 6. Open lockfile, create file, etc...
It's unnecessary to learn about those things further, so skip this ... But say something interesting, it use `goto` keyword when faces some errors, how amazing it is!  
### 7. Open listening ports from config file, Init all modules
In line 505 of file `src/core/ngx_cycle.c`, it do open listening ports from config. At this point config has already been initialized by Step5. For more details about what nginx does for network socket flags will be introduced later. @todo  
At line 635 of file `src/core/ngx_cycle.c`, we see function `ngx_init_modules` which init all modules.  
```c
ngx_int_t
ngx_init_modules(ngx_cycle_t *cycle)
{
    ngx_uint_t  i;

    for (i = 0; cycle->modules[i]; i++) {
        if (cycle->modules[i]->init_module) {
            if (cycle->modules[i]->init_module(cycle) != NGX_OK) {
                return NGX_ERROR;
            }
        }
    }

    return NGX_OK;
}
```
Look at the `init_module` function, where its definition? Look back to Step3 each module exposed an interface which is implemented by author. `init_module` functions is defined in this interface which is highly depends on author of it.  
As modules with type **NGX_CORE_MODULE** have been initialized before, modules with such type owns a NULL value on `init_module` function so it can be skipped in this turn. 

### 8. Free useless memory and file descriptor after initializing 
This step is **the end of** function `ngx_init_cycle`. For temperately struct to init, nginx must release used resources which don't be used any more.  


### 9. Choose work mode between master/work and single based on command line arguments
This is not the main theme of this article, I will describe it in this article. @todo