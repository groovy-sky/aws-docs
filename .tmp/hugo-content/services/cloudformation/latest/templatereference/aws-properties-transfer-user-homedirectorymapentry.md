This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::User HomeDirectoryMapEntry

Represents an object that contains entries and targets for
`HomeDirectoryMappings` .

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Entry" : String,
  "Target" : String,
  "Type" : String
}

```

### YAML

```yaml

  Entry: String
  Target: String
  Type: String

```

## Properties

`Entry`

Represents an entry for `HomeDirectoryMappings`.

_Required_: Yes

_Type_: String

_Pattern_: `^/.*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Target`

Represents the map target that is used in a `HomeDirectoryMapEntry`.

_Required_: Yes

_Type_: String

_Pattern_: `^/.*$`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of mapping. Set the type to `FILE` if you want the
mapping to point to a file, or `DIRECTORY` for the directory to point to a
directory.

###### Note

By default, home directory mappings have a `Type` of
`DIRECTORY` when you create a Transfer Family server. You would need to
explicitly set `Type` to `FILE` if you want a mapping to have
a file target.

_Required_: No

_Type_: String

_Allowed values_: `FILE | DIRECTORY`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Transfer::User

PosixProfile

All content copied from https://docs.aws.amazon.com/.
