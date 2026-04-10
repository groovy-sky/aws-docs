This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::M2::Environment FsxStorageConfiguration

###### Important

AWS Mainframe Modernization Service (Managed Runtime Environment experience) will no longer be open to new customers starting on November 7, 2025. If you would like to use the service, please sign up prior to November 7, 2025. For capabilities similar to AWS Mainframe Modernization Service (Managed Runtime Environment experience) explore AWS Mainframe Modernization Service (Self-Managed Experience). Existing customers can
continue to use the service as normal. For more information, see
[AWS Mainframe Modernization availability change](../../../m2/latest/userguide/mainframe-modernization-availability-change.md).

Defines the storage configuration for an Amazon FSx file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemId" : String,
  "MountPoint" : String
}

```

### YAML

```yaml

  FileSystemId: String
  MountPoint: String

```

## Properties

`FileSystemId`

The file system identifier.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,200}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MountPoint`

The mount point for the file system.

_Required_: Yes

_Type_: String

_Pattern_: `^\S{1,200}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EfsStorageConfiguration

HighAvailabilityConfig

All content copied from https://docs.aws.amazon.com/.
