This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RTBFabric::Link LinkApplicationLogSampling

Describes a link application log sample.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ErrorLog" : Number,
  "FilterLog" : Number
}

```

### YAML

```yaml

  ErrorLog: Number
  FilterLog: Number

```

## Properties

`ErrorLog`

An error log entry.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FilterLog`

A filter log entry.

_Required_: Yes

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HeaderTagAction

LinkAttributes

All content copied from https://docs.aws.amazon.com/.
