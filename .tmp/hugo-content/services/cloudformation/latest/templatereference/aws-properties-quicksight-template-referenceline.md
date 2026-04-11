This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ReferenceLine

The reference line visual display options.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataConfiguration" : ReferenceLineDataConfiguration,
  "LabelConfiguration" : ReferenceLineLabelConfiguration,
  "Status" : String,
  "StyleConfiguration" : ReferenceLineStyleConfiguration
}

```

### YAML

```yaml

  DataConfiguration:
    ReferenceLineDataConfiguration
  LabelConfiguration:
    ReferenceLineLabelConfiguration
  Status: String
  StyleConfiguration:
    ReferenceLineStyleConfiguration

```

## Properties

`DataConfiguration`

The data configuration of the reference line.

_Required_: Yes

_Type_: [ReferenceLineDataConfiguration](aws-properties-quicksight-template-referencelinedataconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LabelConfiguration`

The label configuration of the reference line.

_Required_: No

_Type_: [ReferenceLineLabelConfiguration](aws-properties-quicksight-template-referencelinelabelconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of the reference line. Choose one of the following options:

- `ENABLE`

- `DISABLE`

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StyleConfiguration`

The style configuration of the reference line.

_Required_: No

_Type_: [ReferenceLineStyleConfiguration](aws-properties-quicksight-template-referencelinestyleconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RangeEndsLabelType

ReferenceLineCustomLabelConfiguration

All content copied from https://docs.aws.amazon.com/.
