#### Example

> config.json

```json
{
    "id": "scoring_id",
    "type": "progressive|complete",
    "loaders": [
        {
            "id": "loader_1",
            "type": "API",
            "dependencies": []
        },
        {
            "id": "loader_1",
            "type": "API",
            "dependencies": [
                "loader_1"
            ]
        }
    ]
}
```

> main.go

```go
package main

import (
    _ "github.com/emikohmann/goscen"
)

func main() {

}
```

