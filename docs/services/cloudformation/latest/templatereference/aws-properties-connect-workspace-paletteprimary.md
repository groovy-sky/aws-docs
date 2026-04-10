This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::Workspace PalettePrimary

Contains primary color configuration for a workspace theme.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Active" : String,
  "ContrastText" : String,
  "Default" : String
}

```

### YAML

```yaml

  Active: String
  ContrastText: String
  Default: String

```

## Properties

`Active`

The primary color used for active states.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ContrastText`

The text color that contrasts with the primary color for readability.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Default`

The default primary color used throughout the workspace.

_Required_: No

_Type_: String

_Pattern_: `.*\S.*`

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PaletteNavigation

Tag

All content copied from https://docs.aws.amazon.com/.
