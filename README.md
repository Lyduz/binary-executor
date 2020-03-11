# Binary Exector
### Shell Executor Written in Go

## Notes
- The Go binary is available in this repository. You can download it and deploy in your machine.
- Make sure to add the appropriate permissions to executable. E.g "chmod 775 binary-executor"
- The flags needed are --path and --commands
- --path is the default and first command to be executed. It requires you to assign the default directory you want your coommands to be executed
- --commands are set of commands separated by ";" symbol. See example for details.

## Run It
```sh
    // again, add necessary permissions:

    sudo chown $USER:$USER binary-executor
    sudo chmod 775 binary-executor


    // run it with help argument
    
    ./binary-executor -h


    // pass commands
    
    ./binary-executor --path=/var/www/api/api.lyduz.com --commands="pwd; whoami; cd /var/www; uptime"

```

## TODO
- windows deployment issue