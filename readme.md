**GOSCEN** is a Go library that allows you to define a directed non-cyclic graph of dependencies between executable nodes. Each node is defined by an identifier and a set of dependencies (other nodes). In this way, the library makes available a web server, which, in a new request, executes node dependency in an orderly and unique manner without repetitions starting from the root node.

This library is oriented to the creation of loaders based scorings. In this scheme, the scoring node functions as the root node and the set of loaders works as dependency nodes. However, you can define nodes that are not involved in any dependency path to the root node. In this case, the node will not be executed.

## Example

* First, create your file `config.json`:

```javascript
{
    "id": "fraud_scoring",

    "mode": "passive",

    "entry_point": "/scoring",

    "dependencies": [ "payments", "biometrics" ],

    "loaders": [
        {
            "id": "biometrics",

            "dependencies": [ "payments" ]
        },
        {
            "id": "payments",

            "dependencies": [ "user" ]
        },
        {
            "id": "user",

            "dependencies": []
        }
    ]
}
```

* Then, implement your loaders:

```go
package loader

type Result struct {}

func ExecuteLoader(input ...interface{}) ([]interface{}, apierrors.ApiError) {

    return []interface{}{ &Result{}, }, nil
}
```

* Finally, create and run your scoring:

```go
package main

import "github.com/emikohmann/goscene"

func main() {

    goscene.WithLoaders( goscene.LoadersMapping{

            "biometrics": biometrics.ExecuteLoader,

            "payments":   payments.ExecuteLoader,

            "user":       user.ExecuteLoader,
        },

        scoring.FraudStrongRules,
    )

    goscene.Run()
}
```

* Enjoy your scoring!