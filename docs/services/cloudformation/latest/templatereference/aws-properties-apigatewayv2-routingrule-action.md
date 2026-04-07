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

_Type_: [ActionInvokeApi](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-apigatewayv2-routingrule-actioninvokeapi.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ApiGatewayV2::RoutingRule

ActionInvokeApi
