# x-amazon-apigateway-integration.responseParameters object

Specifies mappings from integration method response parameters to method response
parameters. You can map `header`, `body`, or static values to the `header` type of the method
response. Supported only for REST APIs.

Property nameTypeDescription`method.response.header.<param-name>``string`

The named parameter value can be derived from the
`header` and `body` types of the
integration response parameters.

## `x-amazon-apigateway-integration.responseParameters` example

The following example maps `body` and `header` parameters of
the integration response to two `header` parameters of the method
response.

```nohighlight

"responseParameters" : {
    "method.response.header.Location" : "integration.response.body.redirect.url",
    "method.response.header.x-user-id" : "integration.response.header.x-userid"
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-integration.responseTemplates

x-amazon-apigateway-integration.tlsConfig

All content copied from https://docs.aws.amazon.com/.
