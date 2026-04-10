This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::DataMigration Tag

A user-defined key-value pair that describes metadata added to an AWS DMS resource and
that is used by operations such as the following:

- `AddTagsToResource`

- `ListTagsForResource`

- `RemoveTagsFromResource`

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Value: String

```

## Properties

`Key`

A key is the required name of the tag. The string value can be 1-128 Unicode characters
in length and can't be prefixed with "aws:" or "dms:". The string can only contain
only the set of Unicode letters, digits, white-space, '\_', '.', '/', '=', '+', '-' (Java
regular expressions: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

A value is the optional value of the tag. The string value can be 1-256 Unicode
characters in length and can't be prefixed with "aws:" or "dms:". The string can only
contain only the set of Unicode letters, digits, white-space, '\_', '.', '/', '=', '+', '-'
(Java regular expressions: "^(\[\\\p{L}\\\p{Z}\\\p{N}\_.:/=+\\\-\]\*)$").

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceDataSettings

AWS::DMS::DataProvider

All content copied from https://docs.aws.amazon.com/.
