This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition Volume

The data volume configuration for tasks launched using this task definition.
Specifying a volume configuration in a task definition is optional. The volume
configuration may contain multiple volumes but only one volume configured at launch is
supported. Each volume defined in the volume configuration may only specify a
`name` and one of either `configuredAtLaunch`,
`dockerVolumeConfiguration`, `efsVolumeConfiguration`,
`fsxWindowsFileServerVolumeConfiguration`, or `host`. If an
empty volume configuration is specified, by default Amazon ECS uses a host volume. For
more information, see [Using data volumes in\
tasks](https://docs.aws.amazon.com/AmazonECS/latest/developerguide/using_data_volumes.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ConfiguredAtLaunch" : Boolean,
  "DockerVolumeConfiguration" : DockerVolumeConfiguration,
  "EFSVolumeConfiguration" : EFSVolumeConfiguration,
  "FSxWindowsFileServerVolumeConfiguration" : FSxWindowsFileServerVolumeConfiguration,
  "Host" : HostVolumeProperties,
  "Name" : String
}

```

### YAML

```yaml

  ConfiguredAtLaunch: Boolean
  DockerVolumeConfiguration:
    DockerVolumeConfiguration
  EFSVolumeConfiguration:
    EFSVolumeConfiguration
  FSxWindowsFileServerVolumeConfiguration:
    FSxWindowsFileServerVolumeConfiguration
  Host:
    HostVolumeProperties
  Name: String

```

## Properties

`ConfiguredAtLaunch`

Indicates whether the volume should be configured at launch time. This is used to
create Amazon EBS volumes for standalone tasks or tasks created as part of a service.
Each task definition revision may only have one volume configured at launch in the
volume configuration.

To configure a volume at launch time, use this task definition revision and specify a
`volumeConfigurations` object when calling the
`CreateService`, `UpdateService`, `RunTask` or
`StartTask` APIs.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DockerVolumeConfiguration`

This parameter is specified when you use Docker volumes.

Windows containers only support the use of the `local` driver. To use bind
mounts, specify the `host` parameter instead.

###### Note

Docker volumes aren't supported by tasks run on AWS Fargate.

_Required_: No

_Type_: [DockerVolumeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskdefinition-dockervolumeconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EFSVolumeConfiguration`

This parameter is specified when you use an Amazon Elastic File System file system for
task storage.

_Required_: No

_Type_: [EFSVolumeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskdefinition-efsvolumeconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FSxWindowsFileServerVolumeConfiguration`

This parameter is specified when you use Amazon FSx for Windows File Server file
system for task storage.

_Required_: No

_Type_: [FSxWindowsFileServerVolumeConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskdefinition-fsxwindowsfileservervolumeconfiguration.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Host`

This parameter is specified when you use bind mount host volumes. The contents of the
`host` parameter determine whether your bind mount host volume persists
on the host container instance and where it's stored. If the `host` parameter
is empty, then the Docker daemon assigns a host path for your data volume. However, the
data isn't guaranteed to persist after the containers that are associated with it stop
running.

Windows containers can mount whole directories on the same drive as
`$env:ProgramData`. Windows containers can't mount directories on a
different drive, and mount point can't be across drives. For example, you can mount
`C:\my\path:C:\my\path` and `D:\:D:\`, but not
`D:\my\path:C:\my\path` or `D:\:C:\my\path`.

_Required_: No

_Type_: [HostVolumeProperties](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ecs-taskdefinition-hostvolumeproperties.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the volume. Up to 255 letters (uppercase and lowercase), numbers,
underscores, and hyphens are allowed.

When using a volume configured at launch, the `name` is required and must
also be specified as the volume name in the `ServiceVolumeConfiguration` or
`TaskVolumeConfiguration` parameter when creating your service or
standalone task.

For all other types of volumes, this name is referenced in the
`sourceVolume` parameter of the `mountPoints` object in the
container definition.

When a volume is using the `efsVolumeConfiguration`, the name is
required.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Ulimit

VolumeFrom
