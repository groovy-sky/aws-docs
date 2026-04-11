This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup SlurmCustomSetting

Additional settings that directly map to Slurm settings.

###### Important

AWS PCS supports a subset of Slurm settings. For more information, see [Configuring custom Slurm settings in AWS PCS](../../../pcs/latest/userguide/slurm-custom-settings.md) in the _AWS PCS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ParameterName" : String,
  "ParameterValue" : String
}

```

### YAML

```yaml

  ParameterName: String
  ParameterValue: String

```

## Properties

`ParameterName`

AWS PCS supports custom Slurm settings for clusters,
compute node groups, and queues. For more information,
see [Configuring\
custom Slurm settings in AWS PCS](../../../pcs/latest/userguide/slurm-custom-settings.md)
in the _AWS PCS User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The values for the configured Slurm settings.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlurmConfiguration

SpotOptions

All content copied from https://docs.aws.amazon.com/.
