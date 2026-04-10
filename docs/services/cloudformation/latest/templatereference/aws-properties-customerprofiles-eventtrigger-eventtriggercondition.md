This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger EventTriggerCondition

Specifies the circumstances under which the event should trigger the
destination.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EventTriggerDimensions" : [ EventTriggerDimension, ... ],
  "LogicalOperator" : String
}

```

### YAML

```yaml

  EventTriggerDimensions:
    - EventTriggerDimension
  LogicalOperator: String

```

## Properties

`EventTriggerDimensions`

A list of dimensions to be evaluated for the event.

_Required_: Yes

_Type_: Array of [EventTriggerDimension](aws-properties-customerprofiles-eventtrigger-eventtriggerdimension.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogicalOperator`

The operator used to combine multiple dimensions.

_Required_: Yes

_Type_: String

_Allowed values_: `ANY | ALL | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CustomerProfiles::EventTrigger

EventTriggerDimension

All content copied from https://docs.aws.amazon.com/.
