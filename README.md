# http-pprof-switch: A switch to enable/disable go profiling for debugging your applications with just an environment variable.

## Introduction 
A fast and easy way to start a http server for go profiling. You can include
this in all you projects without any impact in your app. You can choose if activate
or deactivate it.

This is a fork of the standard package that adds some extra features: 

1. Zero lines of extra code to run the pprof profiling. Importing the repository is all you need.

2. Automatically run http server with a custom mux to avoid conflicts with your code.

3. Enable/disable or config the package with just few environment variables.


The server is listen in the address `0.0.0.0:1000` by default.


## Setup
To use this module you just need to things: 
1. Import the module
```
import (
       ...
	_"github.com/dacalin/http-pprof-switch"
)
```

2. Set the environment variable to turn it on/off. You can use it in the host or container. 
The environment variables will be check on init.
```
export HTTP_PPROF_SWITCH_ENABLE=1 // ON
export HTTP_PPROF_SWITCH_ENABLE=0 // OFF
```


## Environment varaibles
**HTTP_PPROF_SWITCH_ENABLE**: Enable or disable the server. The profiling will start when and endpoint is reached.

**HTTP_PPROF_SWITCH_PORT**: Default port is 1000. Override it with other port to change the listening port. 

**HTTP_PPROF_SWITCH_HOST**: Default host where ther server listen is 0.0.0.0, but you can modify it by writing this environment variable.


## Endpoints
We used the same endpoints than the original package:

1. `http://<YOUR_HOST>:<HTTP_PPROF_SWITCH_PORT>/debug/pprof/`
2. `http://<YOUR_HOST>:<HTTP_PPROF_SWITCH_PORT>/debug/pprof/cmdline`
3. `http://<YOUR_HOST>:<HTTP_PPROF_SWITCH_PORT>/debug/pprof/profile`
4. `http://<YOUR_HOST>:<HTTP_PPROF_SWITCH_PORT>/debug/pprof/symbol`
5. `http://<YOUR_HOST>:<HTTP_PPROF_SWITCH_PORT>/debug/pprof/trace`





