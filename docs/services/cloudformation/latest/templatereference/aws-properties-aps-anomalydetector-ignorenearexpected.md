This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::AnomalyDetector IgnoreNearExpected

Configuration for threshold settings that determine when values near expected values
should be ignored during anomaly detection.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Amount" : Number,
  "Ratio" : Number
}

```

### YAML

```yaml

  Amount: Number
  Ratio: Number

```

## Properties

`Amount`

The absolute amount by which values can differ from expected values before being
considered anomalous.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Ratio`

The ratio by which values can differ from expected values before being considered
anomalous.

_Required_: No

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AnomalyDetectorConfiguration

Label

All content copied from https://docs.aws.amazon.com/.
