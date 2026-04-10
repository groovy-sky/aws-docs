This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template ImageMenuOption

The menu options for the interactions of an image.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityStatus" : String
}

```

### YAML

```yaml

  AvailabilityStatus: String

```

## Properties

`AvailabilityStatus`

The availability status of the image menu. If the value of this property is set to `ENABLED`, dashboard readers can interact with the image menu.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImageInteractionOptions

InnerFilter

All content copied from https://docs.aws.amazon.com/.
