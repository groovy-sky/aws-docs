This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster SlurmRest

The Slurm REST API configuration includes settings for enabling and configuring the Slurm REST API. It's a property of the [ClusterSlurmConfiguration](../userguide/aws-properties-pcs-cluster-slurmconfiguration.md) object.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

The default value for `mode` is `NONE`. A value of `STANDARD` means the Slurm REST API is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `STANDARD | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlurmdbdCustomSetting

AWS::PCS::ComputeNodeGroup

All content copied from https://docs.aws.amazon.com/.
