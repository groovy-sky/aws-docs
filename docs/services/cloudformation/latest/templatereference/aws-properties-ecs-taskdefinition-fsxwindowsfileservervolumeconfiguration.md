This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition FSxWindowsFileServerVolumeConfiguration

This parameter is specified when you're using [Amazon FSx for Windows File\
Server](../../../fsx/latest/windowsguide/what-is.md) file system for task storage.

For more information and the input format, see [Amazon FSx for Windows File\
Server volumes](../../../amazonecs/latest/developerguide/wfsx-volumes.md) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AuthorizationConfig" : FSxAuthorizationConfig,
  "FileSystemId" : String,
  "RootDirectory" : String
}

```

### YAML

```yaml

  AuthorizationConfig:
    FSxAuthorizationConfig
  FileSystemId: String
  RootDirectory: String

```

## Properties

`AuthorizationConfig`

The authorization configuration details for the Amazon FSx for Windows File Server
file system.

_Required_: No

_Type_: [FSxAuthorizationConfig](aws-properties-ecs-taskdefinition-fsxauthorizationconfig.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FileSystemId`

The Amazon FSx for Windows File Server file system ID to use.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RootDirectory`

The directory within the Amazon FSx for Windows File Server file system to mount as
the root directory inside the host.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FSxAuthorizationConfig

HealthCheck

All content copied from https://docs.aws.amazon.com/.
