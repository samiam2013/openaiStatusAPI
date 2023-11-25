# OpenAI Status API Golang Client

unofficial Golang support for the openAI API's Status Page's API 

starting with the components endpoint only, this API doesn't require
any key, so it's easy enough to just import a module (like this one) 
and just start using it. 

## Usage

```go
package main

import (
    "fmt"
    osa "github.com/samiam2013/openaiStatusAPI"
)

func main() {
    status, err := osa.GetComponent(osa.API)
    if err != nil {
        log.Fatalf("Failed checking OpenAI Status API: %v", err)
    }
    log.Info("OpenAI API Status: ", status)
}
```

should give you something like 
```
INFO[0000] OpenAI API Status: operational
```
"one of operational, degraded_performance, partial_outage, or major_outage."
- https://status.openai.com/api/v2#status
