This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Template DefaultFilterControlConfiguration

The default configuration for all dependent controls of the filter.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ControlOptions" : DefaultFilterControlOptions,
  "Title" : String
}

```

### YAML

```yaml

  ControlOptions:
    DefaultFilterControlOptions
  Title: String

```

## Properties

`ControlOptions`

The control option for the `DefaultFilterControlConfiguration`.

_Required_: Yes

_Type_: [DefaultFilterControlOptions](aws-properties-quicksight-template-defaultfiltercontroloptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Title`

The title of the `DefaultFilterControlConfiguration`. This title is shared by all controls that are tied to this filter.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DefaultDateTimePickerControlOptions

DefaultFilterControlOptions

All content copied from https://docs.aws.amazon.com/.
