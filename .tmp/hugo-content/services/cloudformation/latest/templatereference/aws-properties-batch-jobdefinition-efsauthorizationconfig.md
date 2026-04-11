This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition EFSAuthorizationConfig

The authorization configuration details for the Amazon EFS file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessPointId" : String,
  "Iam" : String
}

```

### YAML

```yaml

  AccessPointId: String
  Iam: String

```

## Properties

`AccessPointId`

The Amazon EFS access point ID to use. If an access point is specified, the root directory value
specified in the `EFSVolumeConfiguration` must either be omitted or set to
`/` which enforces the path set on the EFS access point. If an access point is used,
transit encryption must be enabled in the `EFSVolumeConfiguration`. For more
information, see [Working\
with Amazon EFS access points](../../../efs/latest/ug/efs-access-points.md) in the _Amazon Elastic File System User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Iam`

Whether or not to use the AWS Batch job IAM role defined in a job definition when mounting the
Amazon EFS file system. If enabled, transit encryption must be enabled in the
`EFSVolumeConfiguration`. If this parameter is omitted, the default value of
`DISABLED` is used. For more information, see [Using Amazon EFS access points](../../../batch/latest/userguide/efs-volumes.md#efs-volume-accesspoints) in
the _AWS Batch User Guide_. EFS IAM authorization requires that
`TransitEncryption` be `ENABLED` and that a `JobRoleArn` is
specified.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EcsTaskProperties

EFSVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
