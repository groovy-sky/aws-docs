This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QuickSight::Dashboard LocalNavigationConfiguration

The navigation configuration for `CustomActionNavigationOperation`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetSheetId" : String
}

```

### YAML

```yaml

  TargetSheetId: String

```

## Properties

`TargetSheetId`

The sheet that is targeted for navigation in the same analysis.

_Required_: Yes

_Type_: String

_Pattern_: `^[\w\-]+$`

_Minimum_: `1`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LoadingAnimation

LongFormatText

All content copied from https://docs.aws.amazon.com/.
