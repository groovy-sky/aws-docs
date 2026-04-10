This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::DataRepositoryAssociation AutoImportPolicy

Describes the data repository association's automatic import policy.
The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory
listings up to date by importing changes to your Amazon FSx for Lustre file system
as you modify objects in a linked S3 bucket.

The `AutoImportPolicy` is only supported on Amazon FSx for Lustre file systems
with a data repository association.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Events" : [ String, ... ]
}

```

### YAML

```yaml

  Events:
    - String

```

## Properties

`Events`

The `AutoImportPolicy` can have the following event values:

- `NEW` \- Amazon FSx automatically imports metadata of
files added to the linked S3 bucket that do not currently exist in the FSx
file system.

- `CHANGED` \- Amazon FSx automatically updates file
metadata and invalidates existing file content on the file system as files
change in the data repository.

- `DELETED` \- Amazon FSx automatically deletes files
on the file system as corresponding files are deleted in the data repository.

You can define any combination of event types for your `AutoImportPolicy`.

_Required_: Yes

_Type_: Array of String

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AutoExportPolicy

S3

All content copied from https://docs.aws.amazon.com/.
