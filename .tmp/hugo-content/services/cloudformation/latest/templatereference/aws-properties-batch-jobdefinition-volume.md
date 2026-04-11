This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Batch::JobDefinition Volume

A data volume that's used in a job's container properties.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EfsVolumeConfiguration" : EFSVolumeConfiguration,
  "Host" : Host,
  "Name" : String
}

```

### YAML

```yaml

  EfsVolumeConfiguration:
    EFSVolumeConfiguration
  Host:
    Host
  Name: String

```

## Properties

`EfsVolumeConfiguration`

This parameter is specified when you're using an Amazon Elastic File System file system for job storage. Jobs
that are running on Fargate resources must specify a `platformVersion` of at least
`1.4.0`.

_Required_: No

_Type_: [EFSVolumeConfiguration](aws-properties-batch-jobdefinition-efsvolumeconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Host`

The contents of the `host` parameter determine whether your data volume persists
on the host container instance and where it's stored. If the host parameter is empty, then the
Docker daemon assigns a host path for your data volume. However, the data isn't guaranteed to
persist after the containers that are associated with it stop running.

###### Note

This parameter isn't applicable to jobs that are running on Fargate resources and
shouldn't be provided.

_Required_: No

_Type_: [Host](aws-properties-batch-jobdefinition-host.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the volume. It can be up to 255 characters long. It can contain uppercase and lowercase letters,
numbers, hyphens (-), and underscores (\_). This name is referenced in the
`sourceVolume` parameter of container definition `mountPoints`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Ulimit

AWS::Batch::JobQueue

All content copied from https://docs.aws.amazon.com/.
