This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppStream::Stack StorageConnector

A connector that enables persistent storage for users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConnectorType" : String,
  "Domains" : [ String, ... ],
  "ResourceIdentifier" : String
}

```

### YAML

```yaml

  ConnectorType: String
  Domains:
    - String
  ResourceIdentifier: String

```

## Properties

`ConnectorType`

The type of storage connector.

_Required_: Yes

_Type_: String

_Allowed values_: `HOMEFOLDERS | GOOGLE_DRIVE | ONE_DRIVE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Domains`

The names of the domains for the account.

_Required_: No

_Type_: Array of String

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceIdentifier`

The ARN of the storage connector.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ApplicationSettings

StreamingExperienceSettings
