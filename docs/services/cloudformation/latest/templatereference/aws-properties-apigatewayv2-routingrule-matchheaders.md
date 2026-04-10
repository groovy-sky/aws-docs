This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule MatchHeaders

Represents a `MatchHeaders` condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnyOf" : [ MatchHeaderValue, ... ]
}

```

### YAML

```yaml

  AnyOf:
    - MatchHeaderValue

```

## Properties

`AnyOf`

The header name and header value glob to be matched. The matchHeaders condition is matched if any of the header name and header value globs are matched.

_Required_: Yes

_Type_: Array of [MatchHeaderValue](aws-properties-apigatewayv2-routingrule-matchheadervalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MatchBasePaths

MatchHeaderValue

All content copied from https://docs.aws.amazon.com/.
