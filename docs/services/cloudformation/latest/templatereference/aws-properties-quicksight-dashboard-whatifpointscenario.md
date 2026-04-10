This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard WhatIfPointScenario

Provides the forecast to meet the target for a particular date.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Date" : String,
  "Value" : Number
}

```

### YAML

```yaml

  Date: String
  Value: Number

```

## Properties

`Date`

The date that you need the forecast results for.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The target value that you want to meet for the provided date.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WaterfallVisual

WhatIfRangeScenario

All content copied from https://docs.aws.amazon.com/.
