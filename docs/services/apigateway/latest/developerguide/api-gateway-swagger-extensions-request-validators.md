# x-amazon-apigateway-request-validators object

Defines the supported request validators for the containing API as a map between a validator name and the associated request validation rules. This extension applies to a REST API.

Property nameTypeDescription

`request_validator_name`

[x-amazon-apigateway-request-validators.requestValidator object](api-gateway-swagger-extensions-request-validators-requestvalidator.md)

Specifies the validation rules consisting of the named validator.
For example:

```nohighlight

    "basic" : {
      "validateRequestBody" : true,
      "validateRequestParameters" : true
    },

```

To apply this validator to a specific method, reference the validator name ( `basic`) as the value of the [x-amazon-apigateway-request-validator property](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-swagger-extensions-request-validator.html) property.

## `x-amazon-apigateway-request-validators` example

The following example shows a set of request validators for an API as a map between a validator name and the associated request validation rules.

OpenAPI 2.0

```nohighlight

{
  "swagger": "2.0",
  ...
  "x-amazon-apigateway-request-validators" : {
    "basic" : {
      "validateRequestBody" : true,
      "validateRequestParameters" : true
    },
    "params-only" : {
      "validateRequestBody" : false,
      "validateRequestParameters" : true
    }
  },
  ...
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

x-amazon-apigateway-request-validator

x-amazon-apigateway-request-validators.requestValidator
