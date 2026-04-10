This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpaces::WorkspacesPool ApplicationSettings

The persistent application settings for users in the pool.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SettingsGroup" : String,
  "Status" : String
}

```

### YAML

```yaml

  SettingsGroup: String
  Status: String

```

## Properties

`SettingsGroup`

The path prefix for the S3 bucket where users’ persistent application settings are
stored.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9_./()!*'-]+$`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Enables or disables persistent application settings for users during their pool
sessions.

_Required_: Yes

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::WorkSpaces::WorkspacesPool

Capacity

All content copied from https://docs.aws.amazon.com/.
