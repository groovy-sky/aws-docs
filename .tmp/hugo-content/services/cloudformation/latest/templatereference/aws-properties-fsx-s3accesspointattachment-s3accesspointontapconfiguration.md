This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment S3AccessPointOntapConfiguration

Describes the FSx for ONTAP attachment configuration of an S3 access point attachment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FileSystemIdentity" : OntapFileSystemIdentity,
  "VolumeId" : String
}

```

### YAML

```yaml

  FileSystemIdentity:
    OntapFileSystemIdentity
  VolumeId: String

```

## Properties

`FileSystemIdentity`

The file system identity used to authorize file access requests made using the S3 access point.

_Required_: Yes

_Type_: [OntapFileSystemIdentity](aws-properties-fsx-s3accesspointattachment-ontapfilesystemidentity.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`VolumeId`

The ID of the FSx for ONTAP volume that the S3 access point is attached to.

_Required_: Yes

_Type_: String

_Pattern_: `^(fsvol-[0-9a-f]{17,})$`

_Minimum_: `23`

_Maximum_: `23`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

S3AccessPoint

S3AccessPointOpenZFSConfiguration

All content copied from https://docs.aws.amazon.com/.
