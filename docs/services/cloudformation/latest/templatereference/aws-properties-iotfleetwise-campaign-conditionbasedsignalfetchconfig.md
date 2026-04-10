This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign ConditionBasedSignalFetchConfig

Specifies the condition under which a signal fetch occurs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConditionExpression" : String,
  "TriggerMode" : String
}

```

### YAML

```yaml

  ConditionExpression: String
  TriggerMode: String

```

## Properties

`ConditionExpression`

The condition that must be satisfied to trigger a signal fetch.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TriggerMode`

Indicates the mode in which the signal fetch is triggered.

_Required_: Yes

_Type_: String

_Allowed values_: `ALWAYS | RISING_EDGE`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConditionBasedCollectionScheme

DataDestinationConfig

All content copied from https://docs.aws.amazon.com/.
