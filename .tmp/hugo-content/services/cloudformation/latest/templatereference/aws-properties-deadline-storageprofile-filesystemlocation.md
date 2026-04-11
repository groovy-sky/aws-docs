This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Deadline::StorageProfile FileSystemLocation

The details of the file system location for the resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Name" : String,
  "Path" : String,
  "Type" : String
}

```

### YAML

```yaml

  Name: String
  Path: String
  Type: String

```

## Properties

`Name`

The location name.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9A-Za-z ]*$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Path`

The file path.

_Required_: Yes

_Type_: String

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of file.

_Required_: Yes

_Type_: String

_Allowed values_: `SHARED | LOCAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Deadline::StorageProfile

Next

All content copied from https://docs.aws.amazon.com/.
