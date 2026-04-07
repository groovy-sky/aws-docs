This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::IntelligentPromptRouter RoutingCriteria

Routing criteria for a prompt router.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ResponseQualityDifference" : Number
}

```

### YAML

```yaml

  ResponseQualityDifference: Number

```

## Properties

`ResponseQualityDifference`

The criteria's response quality difference.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PromptRouterTargetModel

Tag
