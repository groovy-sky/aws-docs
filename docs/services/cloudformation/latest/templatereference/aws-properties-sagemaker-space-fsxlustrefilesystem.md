This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space FSxLustreFileSystem

A custom file system in Amazon FSx for Lustre.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemId" : String
}

```

### YAML

```yaml

  FileSystemId: String

```

## Properties

`FileSystemId`

Amazon FSx for Lustre file system ID.

_Required_: Yes

_Type_: String

_Pattern_: `^(fs-[0-9a-f]{8,})$`

_Minimum_: `11`

_Maximum_: `21`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EFSFileSystem

JupyterServerAppSettings

All content copied from https://docs.aws.amazon.com/.
