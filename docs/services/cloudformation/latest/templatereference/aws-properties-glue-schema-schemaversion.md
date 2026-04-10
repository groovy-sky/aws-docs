This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::Schema SchemaVersion

Specifies the version of a schema.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IsLatest" : Boolean,
  "VersionNumber" : Integer
}

```

### YAML

```yaml

  IsLatest: Boolean
  VersionNumber: Integer

```

## Properties

`IsLatest`

Indicates if this version is the latest version of the schema.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VersionNumber`

The version number of the schema.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Registry

Tag

All content copied from https://docs.aws.amazon.com/.
