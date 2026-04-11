This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule Action

Represents a routing rule action. The only supported action is `invokeApi`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InvokeApi" : ActionInvokeApi
}

```

### YAML

```yaml

  InvokeApi:
    ActionInvokeApi

```

## Properties

`InvokeApi`

Represents an InvokeApi action.

_Required_: Yes

_Type_: [ActionInvokeApi](aws-properties-apigatewayv2-routingrule-actioninvokeapi.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ApiGatewayV2::RoutingRule

ActionInvokeApi

All content copied from https://docs.aws.amazon.com/.
