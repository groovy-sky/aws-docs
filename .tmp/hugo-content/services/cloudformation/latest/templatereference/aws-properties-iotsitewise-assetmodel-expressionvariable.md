This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTSiteWise::AssetModel ExpressionVariable

Contains expression variable information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Value" : VariableValue
}

```

### YAML

```yaml

  Name: String
  Value:
    VariableValue

```

## Properties

`Name`

The friendly name of the variable to be used in the expression.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The variable that identifies an asset property from which to use values.

_Required_: Yes

_Type_: [VariableValue](aws-properties-iotsitewise-assetmodel-variablevalue.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnforcedAssetModelInterfaceRelationship

Metric

All content copied from https://docs.aws.amazon.com/.
