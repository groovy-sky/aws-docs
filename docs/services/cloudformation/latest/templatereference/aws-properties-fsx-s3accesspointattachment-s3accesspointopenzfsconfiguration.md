This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment S3AccessPointOpenZFSConfiguration

Describes the FSx for OpenZFS attachment configuration of an S3 access point attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemIdentity" : OpenZFSFileSystemIdentity,
  "VolumeId" : String
}

```

### YAML

```yaml

  FileSystemIdentity:
    OpenZFSFileSystemIdentity
  VolumeId: String

```

## Properties

`FileSystemIdentity`

The file system identity used to authorize file access requests made using the S3 access point.

_Required_: Yes

_Type_: [OpenZFSFileSystemIdentity](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-fsx-s3accesspointattachment-openzfsfilesystemidentity.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeId`

The ID of the FSx for OpenZFS volume that the S3 access point is attached to.

_Required_: Yes

_Type_: String

_Pattern_: `^(fsvol-[0-9a-f]{17,})$`

_Minimum_: `23`

_Maximum_: `23`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3AccessPointOntapConfiguration

S3AccessPointVpcConfiguration
