This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FSx::S3AccessPointAttachment

An S3 access point attached to an Amazon FSx volume.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FSx::S3AccessPointAttachment",
  "Properties" : {
      "Name" : String,
      "OntapConfiguration" : S3AccessPointOntapConfiguration,
      "OpenZFSConfiguration" : S3AccessPointOpenZFSConfiguration,
      "S3AccessPoint" : S3AccessPoint,
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::FSx::S3AccessPointAttachment
Properties:
  Name: String
  OntapConfiguration:
    S3AccessPointOntapConfiguration
  OpenZFSConfiguration:
    S3AccessPointOpenZFSConfiguration
  S3AccessPoint:
    S3AccessPoint
  Type: String

```

## Properties

`Name`

The name of the S3 access point attachment; also used for the name of the S3 access point.

_Required_: Yes

_Type_: String

_Pattern_: `^(?=[a-z0-9])[a-z0-9-]{1,48}[a-z0-9]$`

_Minimum_: `3`

_Maximum_: `50`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OntapConfiguration`

The ONTAP configuration of the S3 access point attachment.

_Required_: No

_Type_: [S3AccessPointOntapConfiguration](aws-properties-fsx-s3accesspointattachment-s3accesspointontapconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OpenZFSConfiguration`

The OpenZFSConfiguration of the S3 access point attachment.

_Required_: No

_Type_: [S3AccessPointOpenZFSConfiguration](aws-properties-fsx-s3accesspointattachment-s3accesspointopenzfsconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`S3AccessPoint`

The S3 access point configuration of the S3 access point attachment.

_Required_: No

_Type_: [S3AccessPoint](aws-properties-fsx-s3accesspointattachment-s3accesspoint.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Type`

The type of Amazon FSx volume that the S3 access point is attached to.

_Required_: Yes

_Type_: String

_Allowed values_: `OPENZFS | ONTAP`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

`Lifecycle`

The lifecycle status of the S3 access point attachment. The lifecycle can have the following values:

- AVAILABLE - the S3 access point attachment is available for use

- CREATING - Amazon FSx is creating the S3 access point and attachment

- DELETING - Amazon FSx is deleting the S3 access point and attachment

- FAILED - The S3 access point attachment is in a failed state. Delete and detach the S3 access
point attachment, and create a new one.

- UPDATING - Amazon FSx is updating the S3 access point attachment

`S3AccessPoint.Alias`

A DNS alias that is associated with the file system. You can use a DNS alias to access a file system using
user-defined DNS names, in addition to the default DNS name
that Amazon FSx assigns to the file system. For more information, see
[DNS aliases](../../../fsx/latest/windowsguide/managing-dns-aliases.md)
in the _FSx for Windows File Server User Guide_.

`S3AccessPoint.ResourceARN`

The Amazon Resource Name (ARN) for a given resource. ARNs uniquely identify AWS
resources. We require an ARN when you need to specify a resource unambiguously across
all of AWS. For more information, see [Amazon Resource Names (ARNs)](../../../../general/latest/gr/aws-arns-and-namespaces.md) in
the _AWS General Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WindowsConfiguration

FileSystemGID

All content copied from https://docs.aws.amazon.com/.
