# x-amazon-apigateway-integration.requestTemplates object

Specifies mapping templates for a request payload of the specified MIME types.

Property nameTypeDescription`MIME type``string`

An example of the MIME type is `application/json`. For information about creating a
mapping template, see [Mapping template transformations for REST APIs in API Gateway](models-mappings.md).

## x-amazon-apigateway-integration.requestTemplates example

The following example sets mapping templates for a request payload of the
`application/json` and `application/xml` MIME types.

```nohighlight

"requestTemplates" : {
    "application/json" : "#set ($root=$input.path('$')) { \"stage\": \"$root.name\", \"user-id\": \"$root.key\" }",
    "application/xml" : "#set ($root=$input.path('$')) <stage>$root.name</stage> "
}

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

x-amazon-apigateway-integrations

x-amazon-apigateway-integration.requestParameters

All content copied from https://docs.aws.amazon.com/.
