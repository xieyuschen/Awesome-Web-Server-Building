# How nginx loads the nginx.conf file
Everybody who has once used nginx will know how important of nginx.conf file is. This article will foucs on following questions:  
- how nginx endless loop is managed?
- when this file is loaded
By the way, I will introduce a question for the next article:  
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
Look at `nginx-release-1.21.1\src\core\nginx.c:294` there is a condition judgement. If you have used command `sudo systemctl restart nginx` once, this error output is not strange to you for error on nginx.conf file.  
```c
    if (ngx_test_config) {
        ngx_log_stderr(0, "configuration file %s test failed",
                        init_cycle.conf_file.data);
    }
```
Let's take an eye to variable `ngx_test_config` which defined at `\src\core\ngx_cycle.c:28`. Man who familiar to C language will soon reliaze here must use `extern` keyword, as expected we found the extern declaration at it's conrresponding header file at line 144.  
After finding the key variable, the method to step further are paying attention that who uses it? who modifies it?  

### 2.1 Analyzing command line, set flags and then correspond to flags 
Reading source code then I found that function `ngx_get_options` at `\src\core\nginx.c:211`. To better understand this we first learn about what arguments does nginx start need. Browsing [nginx offical document](https://www.nginx.com/resources/wiki/start/topics/tutorials/commandline/), here I put an example and some options as following.  
```shell
/usr/bin/nginx -t -c ~/mynginx.conf -g "pid /var/run/nginx.pid; worker_processes 2;"
```

<div class="section" id="options">
<h3>Options<a class="headerlink" href="#options" title="Permalink to this headline">¶</a></h3>
<table border="1" class="docutils">
<colgroup>
<col width="15%">
<col width="85%">
</colgroup>
<tbody valign="top">
<tr class="row-odd"><td><code class="docutils literal"><span class="pre">-?,</span> <span class="pre">-h</span></code></td>
<td>Print help.</td>
</tr>
<tr class="row-even"><td><code class="docutils literal"><span class="pre">-v</span></code></td>
<td>Print version.</td>
</tr>
<tr class="row-odd"><td><code class="docutils literal"><span class="pre">-V</span></code></td>
<td>Print NGINX version, compiler version and configure parameters.</td>
</tr>
<tr class="row-even"><td><code class="docutils literal"><span class="pre">-t</span></code></td>
<td>Don’t run, just test the configuration file.
NGINX checks configuration for correct syntax and then try to open files referred in configuration.</td>
</tr>
<tr class="row-odd"><td><code class="docutils literal"><span class="pre">-q</span></code></td>
<td>Suppress non-error messages during configuration testing.</td>
</tr>
<tr class="row-even"><td><code class="docutils literal"><span class="pre">-s</span> <span class="pre">signal</span></code></td>
<td>Send signal to a master process: stop, quit, reopen, reload. (version &gt;= 0.7.53)</td>
</tr>
<tr class="row-odd"><td><code class="docutils literal"><span class="pre">-p</span> <span class="pre">prefix</span></code></td>
<td>Set prefix path (default: <code class="docutils literal"><span class="pre">/usr/local/nginx/</span></code>). (version &gt;= 0.7.53)</td>
</tr>
<tr class="row-even"><td><code class="docutils literal"><span class="pre">-c</span> <span class="pre">filename</span></code></td>
<td>Specify which configuration file NGINX should use instead of the default.</td>
</tr>
<tr class="row-odd"><td><code class="docutils literal"><span class="pre">-g</span> <span class="pre">directives</span></code></td>
<td>Set <a class="reference external" href="https://nginx.org/en/docs/http/ngx_http_core_module.html">global</a> directives. (version &gt;= 0.7.4)</td>
</tr>
</tbody>
</table>
</div>

All options in option table are listed with handling way in function `ngx_get_options`. This function just analyzes options to set flags(like config file path, testing argument and so on) from command line and then do something based on flags value.  We can easily find this design such as `ngx_show_version` to show ngx version to user.  

### 2.2 What does function ngx_process_options do?
Some solutions about flags lives in main scope, the other are stored in scope of `ngx_process_options` function in `\src\core\nginx.c:938`. Firstly it makes sure the config file path as following code. Additional, variable `NGX_CONF_PATH` comes from `./configure`.
```c
    if (ngx_conf_file) {
        cycle->conf_file.len = ngx_strlen(ngx_conf_file);
        cycle->conf_file.data = ngx_conf_file;

    } else {
        ngx_str_set(&cycle->conf_file, NGX_CONF_PATH);
    }
```
Other flags are dealed in `ngx_process_options` scope and stored into cycle data member.  

### 2.3 Step into ngx_init_cycle function again
I finally found that the file loaded may still in `ngx_init_cycle`, so I looked it carefully again. Reading its call stack and I finally found it, the call stack is: `main->ngx_init_cycle->ngx_conf_parse`. In function `ngx_conf_parse` it opens file if the second arguments is not null or @todo: what function does if filename is NULL.  

Attention that ngx_conf_parse is called twice in `ngx_init_cyclye`, the one is in function `ngx_conf_param`, the other is in `ngx_init_cycle` directly. In this function it does two key things, one is that it open file or set flag, the other is that it analyzes configure file and check whether it's valid or not. Here it use **STATE MACHINE** obviously and I will take an eye later @todo.  
All shall we know is that in this stage it analyzes the config file using a state machine, and then executes call back function set by `ngx_command_t` which are interested by a module. Let's see some declaration and implement of callback:  
```c

    //start at \src\core\ngx_conf_file.c:292
    if (cf->handler) {
        //ignore some code
        rv = (*cf->handler)(cf, NULL, cf->handler_conf);
        if (rv == NGX_CONF_OK) {
            continue;
        }
        //ignore some code
    }
```
As code showes above, what's `cf->handler`, let's check it declaration:
```c
typedef char *(*ngx_conf_handler_pt)(ngx_conf_t *cf,
    ngx_command_t *dummy, void *conf);
``` 
We found that this function pointer receives an argument as type `ngx_command_t`, we realize that this handler is possiable to define in `ngx_command_t`. Check `ngx_command_t` and so do I think.
```c
struct ngx_command_s {
    ngx_str_t             name;
    ngx_uint_t            type;
    char               *(*set)(ngx_conf_t *cf, ngx_command_t *cmd, void *conf);
    ngx_uint_t            conf;
    ngx_uint_t            offset;
    void                 *post;
};
```
All things are clear now. To end part1, I answer the question again.  
> Q: When nginx.conf file is loaded by nginx?
> A: It's loaded at `main->ngx_init_cycle->ngx_conf_parse`.  

## 3.How nginx analyzes nginx.conf and executes callback functions?  
Topic `2.3` has already told us the question as title. If you want to know the details of it, please see the next article to learn more about it.  
  

