This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DataPathColor

The color map that determines the color options for a particular element.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Color" : String,
  "Element" : DataPathValue,
  "TimeGranularity" : String
}

```

### YAML

```yaml

  Color: String
  Element:
    DataPathValue
  TimeGranularity: String

```

## Properties

`Color`

The color that needs to be applied to the element.

_Required_: Yes

_Type_: String

_Pattern_: `^#[A-F0-9]{6}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Element`

The element that the color needs to be applied to.

_Required_: Yes

_Type_: [DataPathValue](aws-properties-quicksight-dashboard-datapathvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeGranularity`

The time granularity of the field that the color needs to be applied to.

_Required_: No

_Type_: String

_Allowed values_: `YEAR | QUARTER | MONTH | WEEK | DAY | HOUR | MINUTE | SECOND | MILLISECOND`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataLabelType

DataPathLabelType

All content copied from https://docs.aws.amazon.com/.
