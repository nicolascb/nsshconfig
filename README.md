# OpenSSH config parser
Create, list, edit and delete entries

[OpenSSH Reference](http://man.openbsd.org/cgi-bin/man.cgi/OpenBSD-current/man5/ssh_config.5?query=ssh_config%26sec=5)

## Examples
* [List](#List)
* [Delete](#delete)
* [Edit](#edit)
* [Create](#create)
## List

```go
package main

import (
    "fmt"

    "github.com/nicolascb/nsshconfig"
)

func main() {
    // By default, this filepath is ~/.ssh/config, to change it use: 
    // nsshconfig.SetConfigPath(/path/to/config)
    err := nsshconfig.LoadConfig() 
    if err != nil {
        panic(err)
    }

    fmt.Printf("Found %d entries",nsshconfig.TotalEntries())
    for _, host := range nsshconfig.Hosts() {
       fmt.Printf("#######\n")
       fmt.Printf("Host: %s\n", host.Host)
       fmt.Printf("Options: \n")
       for option, value := range host.Options {
           fmt.Printf("%s=%s\n",option,value)
       }
    }
}
```

## Delete


```go
package main

import (
    "fmt"

    "github.com/nicolascb/nsshconfig"
)

func main() {
    // By default, this filepath is ~/.ssh/config, to change it use: 
    // nsshconfig.SetConfigPath(/path/to/config)
    err := nsshconfig.LoadConfig() 
    if err != nil {
        panic(err)
    }

    // Delete host server01
    err = nsshconfig.Delete("server01")
    if err != nil {
        panic(err)
    }

    // Write file
    err = nsshconfig.WriteConfig()
    if err != nil {
        panic(err)
    }
}
```

## Edit


```go
package main

import (
    "fmt"

    "github.com/nicolascb/nsshconfig"
)

func main() {
    // By default, this filepath is ~/.ssh/config, to change it use: 
    // nsshconfig.SetConfigPath(/path/to/config)
    err := nsshconfig.LoadConfig() 
    if err != nil {
        panic(err)
    }

    // Edit general entry (*)
    host, err := nsshconfig.GetEntryByHost("*")
    if err != nil {
        panic(err)
    }

    // Save
    host.Options["port"] = "5122"
    err = host.Save()
    if err != nil {
        panic(err)
    }

    // Write file
    err = nsshconfig.WriteConfig()
    if err != nil {
        panic(err)
    }
}
```

## Create


```go
package main

import (
    "fmt"

    "github.com/nicolascb/nsshconfig"
)

func main() {
    // By default, this filepath is ~/.ssh/config, to change it use: 
    // nsshconfig.SetConfigPath(/path/to/config)
    err := nsshconfig.LoadConfig() 
    if err != nil {
        panic(err)
    }

    // Options
    options := make(map[string]string)
    options["port"] = "5133"
    options["hostname"] = "gremio.net"
    
    // Add (update if exist)
    err = nsshconfig.New("gremio_server", options)
    if err != nil {
        panic(err)
    }
    
    // Write file
    err = nsshconfig.WriteConfig()
    if err != nil {
        panic(err)
    }
}
```