This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster Scheduler

The cluster management and job scheduling software associated with the cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Version" : String
}

```

### YAML

```yaml

  Type: String
  Version: String

```

## Properties

`Type`

The software AWS PCS uses to manage cluster scaling and job
scheduling.

_Required_: Yes

_Type_: String

_Allowed values_: `SLURM`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The version of the specified scheduling software that AWS PCS uses to manage
cluster scaling and job scheduling. For more information, see [Slurm versions in AWS PCS](https://docs.aws.amazon.com/pcs/latest/userguide/slurm-versions.html) in the _AWS PCS User Guide_.

Valid Values: `23.11 | 24.05 | 24.11 | 25.05`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Networking

SlurmConfiguration
