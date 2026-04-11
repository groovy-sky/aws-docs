This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ReferenceLineDataConfiguration

The data configuration of the reference line.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AxisBinding" : String,
  "DynamicConfiguration" : ReferenceLineDynamicDataConfiguration,
  "SeriesType" : String,
  "StaticConfiguration" : ReferenceLineStaticDataConfiguration
}

```

### YAML

```yaml

  AxisBinding: String
  DynamicConfiguration:
    ReferenceLineDynamicDataConfiguration
  SeriesType: String
  StaticConfiguration:
    ReferenceLineStaticDataConfiguration

```

## Properties

`AxisBinding`

The axis binding type of the reference line. Choose one of the following options:

- `PrimaryY`

- `SecondaryY`

_Required_: No

_Type_: String

_Allowed values_: `PRIMARY_YAXIS | SECONDARY_YAXIS`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DynamicConfiguration`

The dynamic configuration of the reference line data configuration.

_Required_: No

_Type_: [ReferenceLineDynamicDataConfiguration](aws-properties-quicksight-template-referencelinedynamicdataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SeriesType`

The series type of the reference line data configuration. Choose one of the following options:

- `BAR`

- `LINE`

_Required_: No

_Type_: String

_Allowed values_: `BAR | LINE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticConfiguration`

The static data configuration of the reference line data configuration.

_Required_: No

_Type_: [ReferenceLineStaticDataConfiguration](aws-properties-quicksight-template-referencelinestaticdataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReferenceLineCustomLabelConfiguration

ReferenceLineDynamicDataConfiguration

All content copied from https://docs.aws.amazon.com/.
