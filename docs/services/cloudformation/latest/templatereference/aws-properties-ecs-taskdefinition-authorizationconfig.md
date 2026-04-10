This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition AuthorizationConfig

The authorization configuration details for the Amazon EFS file system.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessPointId" : String,
  "IAM" : String
}

```

### YAML

```yaml

  AccessPointId: String
  IAM: String

```

## Properties

`AccessPointId`

The Amazon EFS access point ID to use. If an access point is specified, the root
directory value specified in the `EFSVolumeConfiguration` must either be
omitted or set to `/` which will enforce the path set on the EFS access
point. If an access point is used, transit encryption must be on in the
`EFSVolumeConfiguration`. For more information, see [Working with Amazon\
EFS access points](../../../efs/latest/ug/efs-access-points.md) in the _Amazon Elastic File System User_
_Guide_.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IAM`

Determines whether to use the Amazon ECS task role defined in a task definition when
mounting the Amazon EFS file system. If it is turned on, transit encryption must be
turned on in the `EFSVolumeConfiguration`. If this parameter is omitted, the
default value of `DISABLED` is used. For more information, see [Using\
Amazon EFS access points](../../../amazonecs/latest/developerguide/efs-volumes.md#efs-volume-accesspoints) in the _Amazon Elastic Container Service_
_Developer Guide_.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ECS::TaskDefinition

ContainerDefinition

All content copied from https://docs.aws.amazon.com/.
