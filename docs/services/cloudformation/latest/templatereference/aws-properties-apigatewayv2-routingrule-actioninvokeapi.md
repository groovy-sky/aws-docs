This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule ActionInvokeApi

Represents an InvokeApi action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApiId" : String,
  "Stage" : String,
  "StripBasePath" : Boolean
}

```

### YAML

```yaml

  ApiId: String
  Stage: String
  StripBasePath: Boolean

```

## Properties

`ApiId`

The API identifier of the target API.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Stage`

The name of the target stage.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StripBasePath`

The strip base path setting. When true, API Gateway strips the incoming matched base path when forwarding the request to the target API.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Action

Condition

All content copied from https://docs.aws.amazon.com/.
