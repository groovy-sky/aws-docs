This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::DataRepositoryAssociation AutoExportPolicy

Describes a data repository association's automatic export policy. The
`AutoExportPolicy` defines the types of updated objects on the
file system that will be automatically exported to the data repository.
As you create, modify, or delete files, Amazon FSx for Lustre
automatically exports the defined changes asynchronously once your application finishes
modifying the file.

The `AutoExportPolicy` is only supported on Amazon FSx for Lustre file systems
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

The `AutoExportPolicy` can have the following event values:

- `NEW` \- New files and directories are automatically exported
to the data repository as they are added to the file system.

- `CHANGED` \- Changes to files and directories on the
file system are automatically exported to the data repository.

- `DELETED` \- Files and directories are automatically deleted
on the data repository when they are deleted on the file system.

You can define any combination of event types for your `AutoExportPolicy`.

_Required_: Yes

_Type_: Array of String

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::FSx::DataRepositoryAssociation

AutoImportPolicy

All content copied from https://docs.aws.amazon.com/.
