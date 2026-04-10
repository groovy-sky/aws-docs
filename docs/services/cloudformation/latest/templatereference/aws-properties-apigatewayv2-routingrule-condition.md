This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule Condition

Represents a condition. Conditions can contain up to two `matchHeaders` conditions and one `matchBasePaths` conditions. API Gateway evaluates header conditions and base path conditions together. You can only use AND between header and base path conditions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MatchBasePaths" : MatchBasePaths,
  "MatchHeaders" : MatchHeaders
}

```

### YAML

```yaml

  MatchBasePaths:
    MatchBasePaths
  MatchHeaders:
    MatchHeaders

```

## Properties

`MatchBasePaths`

The base path to be matched.

_Required_: No

_Type_: [MatchBasePaths](aws-properties-apigatewayv2-routingrule-matchbasepaths.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MatchHeaders`

The headers to be matched.

_Required_: No

_Type_: [MatchHeaders](aws-properties-apigatewayv2-routingrule-matchheaders.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ActionInvokeApi

MatchBasePaths

All content copied from https://docs.aws.amazon.com/.
