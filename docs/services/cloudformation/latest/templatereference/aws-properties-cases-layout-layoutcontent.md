This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Layout LayoutContent

Object to store union of different versions of layout content.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Basic" : BasicLayout
}

```

### YAML

```yaml

  Basic:
    BasicLayout

```

## Properties

`Basic`

Content specific to `BasicLayout` type. It configures fields in the top panel
and More Info tab of agent application.

_Required_: Yes

_Type_: [BasicLayout](aws-properties-cases-layout-basiclayout.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FieldItem

LayoutSections

All content copied from https://docs.aws.amazon.com/.
