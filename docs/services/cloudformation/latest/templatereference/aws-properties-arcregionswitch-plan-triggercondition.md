This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan TriggerCondition

Defines a condition that must be met for a trigger to fire.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AssociatedAlarmName" : String,
  "Condition" : String
}

```

### YAML

```yaml

  AssociatedAlarmName: String
  Condition: String

```

## Properties

`AssociatedAlarmName`

The name of the CloudWatch alarm associated with the condition.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Condition`

The condition that must be met. Valid values include `green` and `red`.

_Required_: Yes

_Type_: String

_Allowed values_: `red | green`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Trigger

Workflow

All content copied from https://docs.aws.amazon.com/.
