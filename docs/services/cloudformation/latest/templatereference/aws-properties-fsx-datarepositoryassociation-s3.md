This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::DataRepositoryAssociation S3

The configuration for an Amazon S3 data repository linked to an
Amazon FSx Lustre file system with a data repository association.
The configuration defines which file events (new, changed, or
deleted files or directories) are automatically imported from
the linked data repository to the file system or automatically
exported from the file system to the data repository.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AutoExportPolicy" : AutoExportPolicy,
  "AutoImportPolicy" : AutoImportPolicy
}

```

### YAML

```yaml

  AutoExportPolicy:
    AutoExportPolicy
  AutoImportPolicy:
    AutoImportPolicy

```

## Properties

`AutoExportPolicy`

Describes a data repository association's automatic export policy. The
`AutoExportPolicy` defines the types of updated objects on the
file system that will be automatically exported to the data repository.
As you create, modify, or delete files, Amazon FSx for Lustre
automatically exports the defined changes asynchronously once your application finishes
modifying the file.

The `AutoExportPolicy` is only supported on Amazon FSx for Lustre file systems
with a data repository association.

_Required_: No

_Type_: [AutoExportPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-datarepositoryassociation-autoexportpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoImportPolicy`

Describes the data repository association's automatic import policy.
The AutoImportPolicy defines how Amazon FSx keeps your file metadata and directory
listings up to date by importing changes to your Amazon FSx for Lustre file system
as you modify objects in a linked S3 bucket.

The `AutoImportPolicy` is only supported on Amazon FSx for Lustre file systems
with a data repository association.

_Required_: No

_Type_: [AutoImportPolicy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-datarepositoryassociation-autoimportpolicy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AutoImportPolicy

Tag
