This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGateway::ApiKey StageKey

`StageKey` is a property of the [AWS::ApiGateway::ApiKey](../userguide/aws-resource-apigateway-apikey.md) resource that specifies the stage to associate with the API key. This association allows only clients with the key to make requests to methods in that stage.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RestApiId" : String,
  "StageName" : String
}

```

### YAML

```yaml

  RestApiId: String
  StageName: String

```

## Properties

`RestApiId`

The string identifier of the associated RestApi.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StageName`

The stage name associated with the stage key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [stageKeys](../../../apigateway/latest/api/api-createapikey.md#stageKeys) in the _Amazon API Gateway REST API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGateway::ApiKey

Tag

All content copied from https://docs.aws.amazon.com/.
