This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Space EFSFileSystem

A file system, created by you in Amazon EFS, that you assign to a user profile or
space for an Amazon SageMaker AI Domain. Permitted users can access this file system in
Amazon SageMaker AI Studio.

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

The ID of your Amazon EFS file system.

_Required_: Yes

_Type_: String

_Pattern_: `^(fs-[0-9a-f]{8,})$`

_Minimum_: `11`

_Maximum_: `21`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbsStorageSettings

FSxLustreFileSystem

All content copied from https://docs.aws.amazon.com/.
