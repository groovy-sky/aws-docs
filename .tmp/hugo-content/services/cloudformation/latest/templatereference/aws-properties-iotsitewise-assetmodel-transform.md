This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel Transform

Contains an asset transform property. A transform is a one-to-one mapping of a property's
data points from one form to another. For example, you can use a transform to convert a
Celsius data stream to Fahrenheit by applying the transformation expression to each data point
of the Celsius stream. A transform can only have a data type of `DOUBLE` and
consume properties with data types of `INTEGER` or `DOUBLE`.

For more information, see [Transforms](../../../iot-sitewise/latest/userguide/asset-properties.md#transforms) in the _AWS IoT SiteWise User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expression" : String,
  "Variables" : [ ExpressionVariable, ... ]
}

```

### YAML

```yaml

  Expression: String
  Variables:
    - ExpressionVariable

```

## Properties

`Expression`

The mathematical expression that defines the transformation function. You can specify up
to 10 variables per expression. You can specify up to 10 functions per
expression.

For more information, see [Quotas](../../../iot-sitewise/latest/userguide/quotas.md) in the _AWS IoT SiteWise User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Variables`

The list of variables used in the expression.

_Required_: Yes

_Type_: Array of [ExpressionVariable](aws-properties-iotsitewise-assetmodel-expressionvariable.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TumblingWindow

All content copied from https://docs.aws.amazon.com/.
