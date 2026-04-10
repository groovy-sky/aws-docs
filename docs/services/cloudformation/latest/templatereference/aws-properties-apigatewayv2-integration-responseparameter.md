This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Integration ResponseParameter

Supported only for HTTP APIs. You use response parameters to transform the HTTP response from a backend
integration before returning the response to clients. Specify a key-value map from a selection key to response
parameters. The selection key must be a valid HTTP status code within the range of 200-599. Response parameters are a key-value map. The key
must match the pattern `<action>:<header>.<location>` or
`overwrite.statuscode`. The action can be `append`, `overwrite` or
`remove`. The value can be a static value, or map to response data, stage variables, or context
variables that are evaluated at runtime. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destination" : String,
  "Source" : String
}

```

### YAML

```yaml

  Destination: String
  Source: String

```

## Properties

`Destination`

Specifies the location of the response to modify, and how to modify it. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Source`

Specifies the data to update the parameter with. To learn more, see [Transforming API requests and responses](../../../apigateway/latest/developerguide/http-api-parameter-mapping.md).

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::Integration

ResponseParameterMap

All content copied from https://docs.aws.amazon.com/.
