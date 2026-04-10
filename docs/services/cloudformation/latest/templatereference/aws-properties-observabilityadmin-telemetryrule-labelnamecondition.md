This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ObservabilityAdmin::TelemetryRule LabelNameCondition

Condition that matches based on WAF rule labels, with label names limited to 1024
characters.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LabelName" : String
}

```

### YAML

```yaml

  LabelName: String

```

## Properties

`LabelName`

The label name to match, supporting alphanumeric characters, underscores, hyphens, and
colons.

_Required_: No

_Type_: String

_Pattern_: `^[0-9A-Za-z_\-:]+$`

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Filter

LogDeliveryParameters

All content copied from https://docs.aws.amazon.com/.
