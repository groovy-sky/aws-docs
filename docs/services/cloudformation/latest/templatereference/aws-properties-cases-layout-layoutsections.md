This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cases::Layout LayoutSections

Ordered list containing different kinds of sections that can be added. A LayoutSections
object can only contain one section.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Sections" : [ Section, ... ]
}

```

### YAML

```yaml

  Sections:
    - Section

```

## Properties

`Sections`

Ordered list containing different kinds of sections that can be added. A LayoutSections
object can only contain one section.

_Required_: No

_Type_: Array of [Section](aws-properties-cases-layout-section.md)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LayoutContent

Section

All content copied from https://docs.aws.amazon.com/.
