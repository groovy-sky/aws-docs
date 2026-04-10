This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign TimeBasedCollectionScheme

Information about a collection scheme that uses a time period to decide how often to
collect data.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "PeriodMs" : Number
}

```

### YAML

```yaml

  PeriodMs: Number

```

## Properties

`PeriodMs`

The time period (in milliseconds) to decide how often to collect data. For example,
if the time period is `60000`, the Edge Agent software collects data once
every minute.

_Required_: Yes

_Type_: Number

_Minimum_: `10000`

_Maximum_: `86400000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TimeBasedSignalFetchConfig

All content copied from https://docs.aws.amazon.com/.
