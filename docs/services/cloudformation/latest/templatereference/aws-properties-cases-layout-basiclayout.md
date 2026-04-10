This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Layout BasicLayout

Content specific to `BasicLayout` type. It configures fields in the top panel
and More Info tab of agent application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MoreInfo" : LayoutSections,
  "TopPanel" : LayoutSections
}

```

### YAML

```yaml

  MoreInfo:
    LayoutSections
  TopPanel:
    LayoutSections

```

## Properties

`MoreInfo`

This represents sections in a tab of the page layout.

_Required_: No

_Type_: [LayoutSections](aws-properties-cases-layout-layoutsections.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TopPanel`

This represents sections in a panel of the page layout.

_Required_: No

_Type_: [LayoutSections](aws-properties-cases-layout-layoutsections.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Cases::Layout

FieldGroup

All content copied from https://docs.aws.amazon.com/.
