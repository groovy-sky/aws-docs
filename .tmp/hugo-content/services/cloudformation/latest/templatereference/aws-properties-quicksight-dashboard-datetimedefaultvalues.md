This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard DateTimeDefaultValues

The default values of the `DateTimeParameterDeclaration`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DynamicValue" : DynamicDefaultValue,
  "RollingDate" : RollingDateConfiguration,
  "StaticValues" : [ String, ... ]
}

```

### YAML

```yaml

  DynamicValue:
    DynamicDefaultValue
  RollingDate:
    RollingDateConfiguration
  StaticValues:
    - String

```

## Properties

`DynamicValue`

The dynamic value of the `DataTimeDefaultValues`. Different defaults are displayed according to users, groups, and values mapping.

_Required_: No

_Type_: [DynamicDefaultValue](aws-properties-quicksight-dashboard-dynamicdefaultvalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RollingDate`

The rolling date of the `DataTimeDefaultValues`. The date is determined from the dataset based on input expression.

_Required_: No

_Type_: [RollingDateConfiguration](aws-properties-quicksight-dashboard-rollingdateconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StaticValues`

The static values of the `DataTimeDefaultValues`.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DateMeasureField

DateTimeFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
