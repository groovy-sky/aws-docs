This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::GameLift::MatchmakingConfiguration Tag

A label that you can assign to a Amazon GameLift Servers resource.

**Learn more**

[Tagging AWS\
Resources](../../../../general/latest/gr/aws-tagging.md) in the _AWS General Reference_

[AWS Tagging Strategies](https://aws.amazon.com/answers/account-management/aws-tagging-strategies)

**Related actions**

[All APIs by task](../../../../reference/gamelift/latest/developerguide/reference-awssdk.md#reference-awssdk-resources-fleets)

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

The key for a developer-defined key value pair for tagging an AWS resource.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value for a developer-defined key value pair for tagging an AWS resource.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GameProperty

AWS::GameLift::MatchmakingRuleSet

All content copied from https://docs.aws.amazon.com/.
