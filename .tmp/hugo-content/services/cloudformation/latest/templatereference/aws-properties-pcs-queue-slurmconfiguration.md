This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Queue SlurmConfiguration

Additional options related to the Slurm scheduler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SlurmCustomSettings" : [ SlurmCustomSetting, ... ]
}

```

### YAML

```yaml

  SlurmCustomSettings:
    - SlurmCustomSetting

```

## Properties

`SlurmCustomSettings`

Additional Slurm-specific configuration that directly maps to Slurm settings.

_Required_: No

_Type_: Array of [SlurmCustomSetting](aws-properties-pcs-queue-slurmcustomsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ErrorInfo

SlurmCustomSetting

All content copied from https://docs.aws.amazon.com/.
