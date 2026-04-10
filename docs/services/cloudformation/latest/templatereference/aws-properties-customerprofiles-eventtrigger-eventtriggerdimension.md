This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CustomerProfiles::EventTrigger EventTriggerDimension

A specific event dimension to be assessed.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ObjectAttributes" : [ ObjectAttribute, ... ]
}

```

### YAML

```yaml

  ObjectAttributes:
    - ObjectAttribute

```

## Properties

`ObjectAttributes`

A list of object attributes to be evaluated.

_Required_: Yes

_Type_: Array of [ObjectAttribute](aws-properties-customerprofiles-eventtrigger-objectattribute.md)

_Minimum_: `1`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventTriggerCondition

EventTriggerLimits

All content copied from https://docs.aws.amazon.com/.
