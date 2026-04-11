This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel PropertyType

Contains a property type, which can be one of `attribute`,
`measurement`, `metric`, or `transform`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Attribute" : Attribute,
  "Metric" : Metric,
  "Transform" : Transform,
  "TypeName" : String
}

```

### YAML

```yaml

  Attribute:
    Attribute
  Metric:
    Metric
  Transform:
    Transform
  TypeName: String

```

## Properties

`Attribute`

Specifies an asset attribute property. An attribute generally contains static information,
such as the serial number of an [IIoT](https://en.wikipedia.org/wiki/Internet_of_things) wind turbine.

_Required_: No

_Type_: [Attribute](aws-properties-iotsitewise-assetmodel-attribute.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Metric`

Specifies an asset metric property. A metric contains a mathematical expression that uses
aggregate functions to process all input data points over a time interval and output a single
data point, such as to calculate the average hourly temperature.

_Required_: No

_Type_: [Metric](aws-properties-iotsitewise-assetmodel-metric.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transform`

Specifies an asset transform property. A transform contains a mathematical expression that
maps a property's data points from one form to another, such as a unit conversion from Celsius
to Fahrenheit.

_Required_: No

_Type_: [Transform](aws-properties-iotsitewise-assetmodel-transform.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TypeName`

The type of property type, which can be one of `Attribute`,
`Measurement`, `Metric`, or `Transform`.

_Required_: Yes

_Type_: String

_Allowed values_: `Measurement | Attribute | Transform | Metric`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PropertyPathDefinition

Tag

All content copied from https://docs.aws.amazon.com/.
