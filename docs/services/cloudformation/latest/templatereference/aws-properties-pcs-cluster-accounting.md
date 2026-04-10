This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster Accounting

The accounting configuration includes configurable settings for Slurm accounting.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DefaultPurgeTimeInDays" : Integer,
  "Mode" : String
}

```

### YAML

```yaml

  DefaultPurgeTimeInDays: Integer
  Mode: String

```

## Properties

`DefaultPurgeTimeInDays`

The default value for all purge settings for `slurmdbd.conf`. For more
information, see the [slurmdbd.conf documentation at SchedMD](https://slurm.schedmd.com/slurmdbd.conf.html).

The default value for `defaultPurgeTimeInDays` is `-1`.

A value of `-1` means there is no purge time and records persist as long as
the cluster exists.

###### Important

`0` isn't a valid value.

_Required_: No

_Type_: Integer

_Minimum_: `-1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The default value for `mode` is `NONE`. A value of
`STANDARD` means Slurm accounting is enabled.

_Required_: Yes

_Type_: String

_Allowed values_: `STANDARD | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCS::Cluster

AuthKey

All content copied from https://docs.aws.amazon.com/.
