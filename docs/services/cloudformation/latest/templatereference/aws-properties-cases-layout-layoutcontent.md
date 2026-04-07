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

_Type_: [BasicLayout](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-cases-layout-basiclayout.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FieldItem

LayoutSections
