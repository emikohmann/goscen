### Goscen

**goscen** is a Go library that allows you to define a directed non-cyclic graph of dependencies between executable nodes. Each node is defined by an identifier and a set of dependencies (other nodes). In this way, the library makes available a web server, which, in a new request, executes node dependency in an orderly and unique manner without repetitions starting from the root node.

This library is oriented to the creation of loaders based scorings. In this scheme, the scoring node functions as the root node and the set of loaders works as dependency nodes. However, you can define nodes that are not involved in any dependency path to the root node. In this case, the node will not be executed.

### Example

* First, create your file `config.json`:

```json
{
    "id": "user_scoring",
    "mode": "passive",
    "entry_point": "/user_scoring",
    "dependencies": [
        "user_information"
    ],
    "loaders": [
        {
            "id": "user_information",
            "dependencies": [
                "afip_data",
                "bank_data"
            ]
        },
        {
            "id": "afip_data",
            "dependencies": []
        },
        {
            "id": "bank_data",
            "dependencies": [
                "afip_data"
            ]
        }
    ]
}
```

* Then, implement your loaders:

```go
package loader

type LoaderResult struct {

}

func ExecuteLoader(input ...interface{}) ([]interface{}, apierrors.ApiError) {

    return []interface{}{
        &LoaderResult{},
    }, nil
}
```

* Finally, create yout scoring:

```go
package main

import (
    "github.com/emikohmann/goscene"
)

func main() {
    goscene.WithLoaders(
        goscene.LoadersMapping{
            "afip_data": afip.ExecuteLoader,
            "bank_data": bank.ExecuteLoader,
        },
        scoring.StrongRules
    )

    goscene.Run()
}
```

* Enjoy your scoring!