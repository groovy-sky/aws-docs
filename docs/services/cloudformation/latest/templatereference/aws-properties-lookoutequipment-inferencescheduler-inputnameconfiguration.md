This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LookoutEquipment::InferenceScheduler InputNameConfiguration

The `InputNameConfiguration` property type specifies Property description not available. for an [AWS::LookoutEquipment::InferenceScheduler](aws-resource-lookoutequipment-inferencescheduler.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ComponentTimestampDelimiter" : String,
  "TimestampFormat" : String
}

```

### YAML

```yaml

  ComponentTimestampDelimiter: String
  TimestampFormat: String

```

## Properties

`ComponentTimestampDelimiter`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(\-|\_|\s)?$`

_Minimum_: `0`

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimestampFormat`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^EPOCH|yyyy-MM-dd-HH-mm-ss|yyyyMMddHHmmss$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataOutputConfiguration

S3InputConfiguration

All content copied from https://docs.aws.amazon.com/.
