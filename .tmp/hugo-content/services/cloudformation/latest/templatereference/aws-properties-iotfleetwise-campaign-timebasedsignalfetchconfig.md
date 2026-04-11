This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign TimeBasedSignalFetchConfig

Used to configure a frequency-based vehicle signal fetch.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ExecutionFrequencyMs" : Number
}

```

### YAML

```yaml

  ExecutionFrequencyMs: Number

```

## Properties

`ExecutionFrequencyMs`

The frequency with which the signal fetch will be executed.

_Required_: Yes

_Type_: Number

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeBasedCollectionScheme

TimestreamConfig

All content copied from https://docs.aws.amazon.com/.
