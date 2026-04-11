This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::AnomalyDetector Dimension

A dimension is a name/value pair that is part of the identity of a metric. Because
dimensions are part of the unique identifier for a metric, whenever you add a unique
name/value pair to one of your metrics, you are creating a new variation of that metric.
For example, many Amazon EC2 metrics publish `InstanceId` as a
dimension name, and the actual instance ID as the value for that dimension.

You can assign up to 30 dimensions to a metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : String
}

```

### YAML

```yaml

  Name: String
  Value: String

```

## Properties

`Name`

The name of the dimension.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The value of the dimension. Dimension values must contain only ASCII characters and
must include at least one non-whitespace character. ASCII control characters are not
supported as part of dimension values.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuration

Metric

All content copied from https://docs.aws.amazon.com/.
