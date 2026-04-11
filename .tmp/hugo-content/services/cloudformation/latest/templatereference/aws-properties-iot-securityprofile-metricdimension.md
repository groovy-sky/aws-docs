This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::SecurityProfile MetricDimension

The dimension of the metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DimensionName" : String,
  "Operator" : String
}

```

### YAML

```yaml

  DimensionName: String
  Operator: String

```

## Properties

`DimensionName`

The name of the dimension.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Operator`

Operators are constructs that perform logical operations. Valid values are
`IN` and `NOT_IN`.

_Required_: No

_Type_: String

_Allowed values_: `IN | NOT_IN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MachineLearningDetectionConfig

MetricsExportConfig

All content copied from https://docs.aws.amazon.com/.
