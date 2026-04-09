# x-amazon-apigateway-auth object

Defines an authorization type to be applied for authorization of method invocations in
API Gateway.

Property nameTypeDescription`type``string`Specifies the authorization type. Specify `"NONE"` for
open access. Specify `"AWS_IAM"` to use IAM permissions.
Values are case insensitive.

## x-amazon-apigateway-auth example

The following example sets the authorization type for an API method.

OpenAPI 3.0.1

```nohighlight

{
  "openapi": "3.0.1",
  "info": {
    "title": "openapi3",
    "version": "1.0"
  },
  "paths": {
    "/protected-by-iam": {
      "get": {
        "x-amazon-apigateway-auth": {
          "type": "AWS_IAM"
        }
      }
    }
  }
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-api-key-source

x-amazon-apigateway-authorizer

All content copied from https://docs.aws.amazon.com/.
