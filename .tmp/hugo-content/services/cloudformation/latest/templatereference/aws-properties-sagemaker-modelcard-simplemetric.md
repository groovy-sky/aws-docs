This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::ModelCard SimpleMetric

A simple metric for evaluating model performance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Notes" : String,
  "Type" : String,
  "Value" : Number,
  "XAxisName" : String,
  "YAxisName" : String
}

```

### YAML

```yaml

  Name: String
  Notes: String
  Type: String
  Value: Number
  XAxisName: String
  YAxisName: String

```

## Properties

`Name`

The name of the metric.

_Required_: Yes

_Type_: String

_Pattern_: `.{1,255}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Notes`

Additional notes about the metric.

_Required_: No

_Type_: String

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of the metric.

_Required_: Yes

_Type_: String

_Allowed values_: `number | string | boolean`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The value of the metric.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`XAxisName`

The name of the X-axis for the metric visualization.

_Required_: No

_Type_: String

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`YAxisName`

The name of the Y-axis for the metric visualization.

_Required_: No

_Type_: String

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SecurityConfig

SourceAlgorithm

All content copied from https://docs.aws.amazon.com/.
