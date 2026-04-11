This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard SetParameterValueConfiguration

The configuration of adding parameters in action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DestinationParameterName" : String,
  "Value" : DestinationParameterValueConfiguration
}

```

### YAML

```yaml

  DestinationParameterName: String
  Value:
    DestinationParameterValueConfiguration

```

## Properties

`DestinationParameterName`

The destination parameter name of the `SetParameterValueConfiguration`.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

Property description not available.

_Required_: Yes

_Type_: [DestinationParameterValueConfiguration](aws-properties-quicksight-dashboard-destinationparametervalueconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SeriesItem

ShapeConditionalFormat

All content copied from https://docs.aws.amazon.com/.
