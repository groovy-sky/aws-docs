This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster SlurmdbdCustomSetting

Additional settings that directly map to SlurmDBD settings.

###### Important

AWS PCS supports a subset of SlurmDBD settings. For more information, see [Configuring custom SlurmDBD settings in AWS PCS](../../../pcs/latest/userguide/slurmdbd-custom-settings.md) in the _AWS PCS User Guide_.

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

AWS PCS supports custom SlurmDBD settings for clusters. For more information,
see [Configuring\
custom SlurmDBD settings in AWS PCS](../../../pcs/latest/userguide/slurmdbd-custom-settings.md)
in the _AWS PCS User Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ParameterValue`

The values for the configured SlurmDBD settings.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SlurmCustomSetting

SlurmRest

All content copied from https://docs.aws.amazon.com/.
