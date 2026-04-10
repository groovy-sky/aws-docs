This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CleanRooms::ConfiguredTable DifferentialPrivacyColumn

Specifies the name of the column that contains the unique identifier of your users, whose privacy you want to protect.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String
}

```

### YAML

```yaml

  Name: String

```

## Properties

`Name`

The name of the column, such as user\_id, that contains the unique identifier of your users, whose privacy you want to protect. If you want to turn on differential privacy for two or more tables in a collaboration, you must configure the same column as the user identifier column in both analysis rules.

_Required_: Yes

_Type_: String

_Pattern_: `[a-z0-9_](([a-z0-9_ ]+-)*([a-z0-9_ ]+))?`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DifferentialPrivacy

GlueTableReference

All content copied from https://docs.aws.amazon.com/.
