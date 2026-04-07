This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::EventBridgeRuleTemplate EventBridgeRuleTemplateTarget

The target to which to send matching events.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String
}

```

### YAML

```yaml

  Arn: String

```

## Properties

`Arn`

Target ARNs must be either an SNS topic or CloudWatch log group.

_Required_: Yes

_Type_: String

_Pattern_: `^arn.+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::MediaLive::EventBridgeRuleTemplate

AWS::MediaLive::EventBridgeRuleTemplateGroup
