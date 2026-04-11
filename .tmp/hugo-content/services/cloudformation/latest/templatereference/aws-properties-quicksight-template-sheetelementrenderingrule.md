This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template SheetElementRenderingRule

The rendering rules of a sheet that uses a free-form layout.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfigurationOverrides" : SheetElementConfigurationOverrides,
  "Expression" : String
}

```

### YAML

```yaml

  ConfigurationOverrides:
    SheetElementConfigurationOverrides
  Expression: String

```

## Properties

`ConfigurationOverrides`

The override configuration of the rendering rules of a sheet.

_Required_: Yes

_Type_: [SheetElementConfigurationOverrides](aws-properties-quicksight-template-sheetelementconfigurationoverrides.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Expression`

The expression of the rendering rules of a sheet.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SheetElementConfigurationOverrides

SheetImage

All content copied from https://docs.aws.amazon.com/.
