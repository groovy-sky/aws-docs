This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template Layout

A `Layout` defines the placement of elements within a sheet.

For more information, see [Types of layout](../../../quicksight/latest/user/types-of-layout.md) in the _Amazon Quick Suite User Guide_.

This is a union type structure. For this structure to be valid, only one of the attributes can be defined.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configuration" : LayoutConfiguration
}

```

### YAML

```yaml

  Configuration:
    LayoutConfiguration

```

## Properties

`Configuration`

The configuration that determines what the type of layout for a sheet.

_Required_: Yes

_Type_: [LayoutConfiguration](aws-properties-quicksight-template-layoutconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LabelOptions

LayoutConfiguration

All content copied from https://docs.aws.amazon.com/.
