This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition RuntimePlatform

An object that represents the compute environment architecture for AWS Batch jobs on
Fargate.

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

The vCPU architecture. The default value is `X86_64`. Valid values are
`X86_64` and `ARM64`.

###### Note

This parameter must be set to `X86_64` for Windows containers.

###### Note

Fargate Spot is not supported on Windows-based containers on
Fargate. A job queue will be blocked if a Windows job is
submitted to a job queue with only Fargate Spot compute environments. However, you can attach
both `FARGATE` and `FARGATE_SPOT` compute environments to the same job
queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingSystemFamily`

The operating system for the compute environment. Valid values are:
`LINUX` (default), `WINDOWS_SERVER_2019_CORE`,
`WINDOWS_SERVER_2019_FULL`, `WINDOWS_SERVER_2022_CORE`, and
`WINDOWS_SERVER_2022_FULL`.

###### Note

The following parameters can’t be set for Windows containers: `linuxParameters`,
`privileged`, `user`, `ulimits`,
`readonlyRootFilesystem`, and `efsVolumeConfiguration`.

###### Note

The AWS Batch Scheduler checks the compute environments that are attached to the job queue before
registering a task definition with Fargate. In this scenario, the job queue is where the job is
submitted. If the job requires a Windows container and the first compute environment is `LINUX`,
the compute environment is skipped and the next compute environment is checked until a Windows-based
compute environment is found.

###### Note

Fargate Spot is not supported on Windows-based containers on Fargate.
A job queue will be blocked if a Windows job is submitted to a job
queue with only Fargate Spot compute environments. However, you can attach both `FARGATE` and
`FARGATE_SPOT` compute environments to the same job queue.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetryStrategy

Secret

All content copied from https://docs.aws.amazon.com/.
