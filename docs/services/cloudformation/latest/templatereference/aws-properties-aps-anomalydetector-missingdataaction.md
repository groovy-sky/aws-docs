This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::AnomalyDetector MissingDataAction

Specifies the action to take when data is missing during anomaly detection
evaluation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MarkAsAnomaly" : Boolean,
  "Skip" : Boolean
}

```

### YAML

```yaml

  MarkAsAnomaly: Boolean
  Skip: Boolean

```

## Properties

`MarkAsAnomaly`

Marks missing data points as anomalies.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Skip`

Skips evaluation when data is missing.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Label

RandomCutForestConfiguration

All content copied from https://docs.aws.amazon.com/.
