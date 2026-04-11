This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ApiGatewayV2::RoutingRule MatchBasePaths

Represents a `MatchBasePaths` condition.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AnyOf" : [ String, ... ]
}

```

### YAML

```yaml

  AnyOf:
    - String

```

## Properties

`AnyOf`

The string of the case sensitive base path to be matched.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Condition

MatchHeaders

All content copied from https://docs.aws.amazon.com/.
