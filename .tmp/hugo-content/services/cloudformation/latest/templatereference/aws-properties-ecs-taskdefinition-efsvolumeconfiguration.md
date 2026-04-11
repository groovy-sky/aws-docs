This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition EFSVolumeConfiguration

This parameter is specified when you're using an Amazon Elastic File System file
system for task storage. For more information, see [Amazon EFS volumes](../../../amazonecs/latest/developerguide/efs-volumes.md) in
the _Amazon Elastic Container Service Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : AuthorizationConfig,
  "FilesystemId" : String,
  "RootDirectory" : String,
  "TransitEncryption" : String,
  "TransitEncryptionPort" : Integer
}

```

### YAML

```yaml

  AuthorizationConfig:
    AuthorizationConfig
  FilesystemId: String
  RootDirectory: String
  TransitEncryption: String
  TransitEncryptionPort: Integer

```

## Properties

`AuthorizationConfig`

The authorization configuration details for the Amazon EFS file system.

_Required_: No

_Type_: [AuthorizationConfig](aws-properties-ecs-taskdefinition-authorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FilesystemId`

The Amazon EFS file system ID to use.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootDirectory`

The directory within the Amazon EFS file system to mount as the root directory inside
the host. If this parameter is omitted, the root of the Amazon EFS volume will be used.
Specifying `/` will have the same effect as omitting this parameter.

###### Important

If an EFS access point is specified in the `authorizationConfig`, the
root directory parameter must either be omitted or set to `/` which will
enforce the path set on the EFS access point.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitEncryption`

Determines whether to use encryption for Amazon EFS data in transit between the Amazon
ECS host and the Amazon EFS server. Transit encryption must be turned on if Amazon EFS
IAM authorization is used. If this parameter is omitted, the default value of
`DISABLED` is used. For more information, see [Encrypting data in transit](../../../efs/latest/ug/encryption-in-transit.md) in
the _Amazon Elastic File System User Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TransitEncryptionPort`

The port to use when sending encrypted data between the Amazon ECS host and the Amazon
EFS server. If you do not specify a transit encryption port, it will use the port
selection strategy that the Amazon EFS mount helper uses. For more information, see
[EFS mount\
helper](../../../efs/latest/ug/efs-mount-helper.md) in the _Amazon Elastic File System User_
_Guide_.

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DockerVolumeConfiguration

EnvironmentFile

All content copied from https://docs.aws.amazon.com/.
