# go-msgraph-types

## Introduction

Golang's implementation of [MSGraph API type definitions](https://github.com/microsoftgraph/msgraph-metadata/tree/master).

## Usage

```go

import (
  "github.com/ms-henglu/go-msgraph-types/types"
)

func main() {
  msgraphTypes := DefaultMSGraphSchemaLoader()
  
  // use customized static files
  // msgraphTypes := types.NewMSGraphSchemaLoader(embeddedFiles)
  
  // list available api-versions
  apiVersions := msgraphTypes.ListAPIVersions()  // ["v1.0", "beta"]
  
  // get the resource definition for a specific api-version
  resourceDefinition, err := msgraphTypes.GetResourceDefinition("v1.0", "/applications")
  
  // list resources
  resourceDefinitions, err := msgraphTypes.ListResources("v1.0")  // ["/applications", "/users", ...]
}

```
