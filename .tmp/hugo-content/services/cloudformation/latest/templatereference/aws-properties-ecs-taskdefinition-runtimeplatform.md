This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECS::TaskDefinition RuntimePlatform

Information about the platform for the Amazon ECS service or task.

For more information about `RuntimePlatform`, see [RuntimePlatform](../../../amazonecs/latest/developerguide/task-definition-parameters.md#runtime-platform) in the _Amazon Elastic Container Service Developer_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CpuArchitecture" : String,
  "OperatingSystemFamily" : String
}

```

### YAML

```yaml

  CpuArchitecture: String
  OperatingSystemFamily: String

```

## Properties

`CpuArchitecture`

The CPU architecture.

You can run your Linux tasks on an ARM-based platform by setting the value to
`ARM64`. This option is available for tasks that run on Linux Amazon EC2
instance, Amazon ECS Managed Instances, or Linux containers on Fargate.

_Required_: No

_Type_: String

_Allowed values_: `X86_64 | ARM64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`OperatingSystemFamily`

The operating system.

_Required_: No

_Type_: String

_Allowed values_: `WINDOWS_SERVER_2019_FULL | WINDOWS_SERVER_2019_CORE | WINDOWS_SERVER_2016_FULL | WINDOWS_SERVER_2004_CORE | WINDOWS_SERVER_2022_CORE | WINDOWS_SERVER_2022_FULL | WINDOWS_SERVER_2025_CORE | WINDOWS_SERVER_2025_FULL | WINDOWS_SERVER_20H2_CORE | LINUX`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestartPolicy

S3FilesVolumeConfiguration

All content copied from https://docs.aws.amazon.com/.
