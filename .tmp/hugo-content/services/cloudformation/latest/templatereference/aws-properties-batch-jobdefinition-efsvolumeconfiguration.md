This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EFSVolumeConfiguration

This is used when you're using an Amazon Elastic File System file system for job storage. For more
information, see [Amazon EFS\
Volumes](../../../batch/latest/userguide/efs-volumes.md) in the _AWS Batch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : EFSAuthorizationConfig,
  "FileSystemId" : String,
  "RootDirectory" : String,
  "TransitEncryption" : String,
  "TransitEncryptionPort" : Integer
}

```

### YAML

```yaml

  AuthorizationConfig:
    EFSAuthorizationConfig
  FileSystemId: String
  RootDirectory: String
  TransitEncryption: String
  TransitEncryptionPort: Integer

```

## Properties

`AuthorizationConfig`

The authorization configuration details for the Amazon EFS file system.

_Required_: No

_Type_: [EFSAuthorizationConfig](aws-properties-batch-jobdefinition-efsauthorizationconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemId`

The Amazon EFS file system ID to use.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RootDirectory`

The directory within the Amazon EFS file system to mount as the root directory inside the host.
If this parameter is omitted, the root of the Amazon EFS volume is used instead. Specifying
`/` has the same effect as omitting this parameter. The maximum length is 4,096
characters.

###### Important

If an EFS access point is specified in the `authorizationConfig`, the root
directory parameter must either be omitted or set to `/`, which enforces the path set
on the Amazon EFS access point.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryption`

Determines whether to enable encryption for Amazon EFS data in transit between the Amazon ECS host and
the Amazon EFS server. Transit encryption must be enabled if Amazon EFS IAM authorization is used. If
this parameter is omitted, the default value of `DISABLED` is used. For more
information, see [Encrypting data in transit](../../../efs/latest/ug/encryption-in-transit.md) in the _Amazon Elastic File System User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TransitEncryptionPort`

The port to use when sending encrypted data between the Amazon ECS host and the Amazon EFS server. If
you don't specify a transit encryption port, it uses the port selection strategy that the Amazon EFS
mount helper uses. The value must be between 0 and 65,535. For more information, see [EFS mount helper](../../../efs/latest/ug/efs-mount-helper.md) in the
_Amazon Elastic File System User Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EFSAuthorizationConfig

EksContainer

All content copied from https://docs.aws.amazon.com/.
