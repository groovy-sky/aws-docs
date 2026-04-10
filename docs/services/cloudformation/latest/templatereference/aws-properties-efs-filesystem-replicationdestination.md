This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EFS::FileSystem ReplicationDestination

Describes the destination file system in the replication configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityZoneName" : String,
  "FileSystemId" : String,
  "KmsKeyId" : String,
  "Region" : String,
  "RoleArn" : String,
  "Status" : String,
  "StatusMessage" : String
}

```

### YAML

```yaml

  AvailabilityZoneName: String
  FileSystemId: String
  KmsKeyId: String
  Region: String
  RoleArn: String
  Status: String
  StatusMessage: String

```

## Properties

`AvailabilityZoneName`

For One Zone file systems, the replication configuration must specify the Availability Zone in which the destination file system is located.

Use the format `us-east-1a`
to specify the Availability Zone. For
more information about One Zone file systems, see [EFS file system types](../../../efs/latest/ug/storage-classes.md) in the _Amazon EFS User Guide_.

###### Note

One Zone file system type is not available in all Availability Zones in AWS Regions where
Amazon EFS is available.

_Required_: No

_Type_: String

_Pattern_: `.+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FileSystemId`

The ID of the destination Amazon EFS file system.

_Required_: No

_Type_: String

_Pattern_: `^(arn:aws[-a-z]*:elasticfilesystem:[0-9a-z-:]+:file-system/fs-[0-9a-f]{8,40}|fs-[0-9a-f]{8,40})$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of an AWS KMS key used to protect the encrypted file system.

_Required_: No

_Type_: String

_Pattern_: `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|mrk-[0-9a-f]{32}|alias/[a-zA-Z0-9/_-]+|(arn:aws[-a-z]*:kms:[a-z0-9-]+:\d{12}:((key/[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})|(key/mrk-[0-9a-f]{32})|(alias/[a-zA-Z0-9/_-]+))))$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The AWS Region in which the destination file system is located.

###### Note

For One Zone file systems, the replication configuration must specify the AWS Region in which the destination
file system is located.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-((iso[a-z]{0,1}-)|(gov-)){0,1}[a-z]+-{0,1}[0-9]{0,1}$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The Amazon Resource Name (ARN) of the current source file system in the
replication configuration.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Describes the status of the replication configuration. For more information
about replication status, see [Viewing\
replication details](../../../efs/latest/ug/awsbackup.md#restoring-backup-efsmonitoring-replication-status.html) in the _Amazon EFS User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | ENABLING | DELETING | ERROR | PAUSED | PAUSING`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StatusMessage`

Message that provides details about the `PAUSED` or `ERRROR` state
of the replication destination configuration. For more information
about replication status messages, see [Viewing\
replication details](../../../efs/latest/ug/awsbackup.md#restoring-backup-efsmonitoring-replication-status.html) in the _Amazon EFS User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ReplicationConfiguration

AWS::EFS::MountTarget

All content copied from https://docs.aws.amazon.com/.
