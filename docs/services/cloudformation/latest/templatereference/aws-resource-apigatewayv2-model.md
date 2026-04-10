This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::Model

The `AWS::ApiGatewayV2::Model` resource updates data model for a
WebSocket API. For more information, see [Model Selection Expressions](../../../apigateway/latest/developerguide/apigateway-websocket-api-selection-expressions.md#apigateway-websocket-api-model-selection-expressions) in the _API Gateway Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ApiGatewayV2::Model",
  "Properties" : {
      "ApiId" : String,
      "ContentType" : String,
      "Description" : String,
      "Name" : String,
      "Schema" : Json
    }
}

```

### YAML

```yaml

Type: AWS::ApiGatewayV2::Model
Properties:
  ApiId: String
  ContentType: String
  Description: String
  Name: String
  Schema: Json

```

## Properties

`ApiId`

The API identifier.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ContentType`

The content-type for the model, for example, "application/json".

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the model.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the model.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schema`

The schema for the model. For application/json models, this should be JSON schema draft 4 model.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the model ID, such as `abc123`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ModelId`

The model ID.

## Examples

### Model creation example

The following example creates a `model` resource called
`MyModel` for an API called
`MyApi`.

#### JSON

```json

{
    "MyModel": {
        "Type": "AWS::ApiGatewayV2::Model",
        "Properties": {
            "Name": "ModelName",
            "ApiId": {
                "Ref": "MyApi"
            },
            "ContentType": "application/json",
            "Schema": {
                "$schema": "http://json-schema.org/draft-04/schema#",
                "title": "DummySchema",
                "type": "object",
                "properties": {
                    "id": {
                        "type": "string"
                    }
                }
            }
        }
    }
}
```

#### YAML

```yaml

MyModel:
  Type: 'AWS::ApiGatewayV2::Model'
  Properties:
    Name: ModelName
    ApiId: !Ref MyApi
    ContentType: application/json
    Schema:
      $schema: 'http://json-schema.org/draft-04/schema#'
      title: DummySchema
      type: object
      properties:
        id:
          type: string

```

## See also

- [CreateModel](../../../apigatewayv2/latest/api-reference/apis-apiid-models.md#CreateModel) in the _Amazon API Gateway_
_Version 2 API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::IntegrationResponse

AWS::ApiGatewayV2::Route

All content copied from https://docs.aws.amazon.com/.
