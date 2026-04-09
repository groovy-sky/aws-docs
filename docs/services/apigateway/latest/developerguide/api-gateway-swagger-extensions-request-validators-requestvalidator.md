# x-amazon-apigateway-request-validators.requestValidator object

Specifies the validation rules of a request validator as part of the
[x-amazon-apigateway-request-validators object](api-gateway-swagger-extensions-request-validators.md) map definition.

Property nameTypeDescription

`validateRequestBody`

`Boolean`

Specifies whether to validate the request body ( `true`) or not ( `false`).

`validateRequestParameters`

`Boolean`

Specifies whether to validate the required request parameters ( `true`) or not ( `false`).

## `x-amazon-apigateway-request-validators.requestValidator` example

The following example shows a parameter-only request validator:

```nohighlight

"params-only": {
    "validateRequestBody" : false,
    "validateRequestParameters" : true
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-request-validators

x-amazon-apigateway-security-policy

All content copied from https://docs.aws.amazon.com/.
