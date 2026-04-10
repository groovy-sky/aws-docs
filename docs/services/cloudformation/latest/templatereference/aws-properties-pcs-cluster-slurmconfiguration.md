This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::Cluster SlurmConfiguration

Additional options related to the Slurm scheduler.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Accounting" : Accounting,
  "AuthKey" : AuthKey,
  "CgroupCustomSettings" : [ CgroupCustomSetting, ... ],
  "JwtAuth" : JwtAuth,
  "ScaleDownIdleTimeInSeconds" : Integer,
  "SlurmCustomSettings" : [ SlurmCustomSetting, ... ],
  "SlurmdbdCustomSettings" : [ SlurmdbdCustomSetting, ... ],
  "SlurmRest" : SlurmRest
}

```

### YAML

```yaml

  Accounting:
    Accounting
  AuthKey:
    AuthKey
  CgroupCustomSettings:
    - CgroupCustomSetting
  JwtAuth:
    JwtAuth
  ScaleDownIdleTimeInSeconds: Integer
  SlurmCustomSettings:
    - SlurmCustomSetting
  SlurmdbdCustomSettings:
    - SlurmdbdCustomSetting
  SlurmRest:
    SlurmRest

```

## Properties

`Accounting`

The accounting configuration includes configurable settings for Slurm accounting.

_Required_: No

_Type_: [Accounting](aws-properties-pcs-cluster-accounting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AuthKey`

The shared Slurm key for authentication, also known as the **cluster**
**secret**.

_Required_: No

_Type_: [AuthKey](aws-properties-pcs-cluster-authkey.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CgroupCustomSettings`

Additional Cgroup-specific configuration that directly maps to Cgroup settings.

_Required_: No

_Type_: Array of [CgroupCustomSetting](aws-properties-pcs-cluster-cgroupcustomsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`JwtAuth`

The JWT authentication configuration for Slurm REST API access.

_Required_: No

_Type_: [JwtAuth](aws-properties-pcs-cluster-jwtauth.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ScaleDownIdleTimeInSeconds`

The time (in seconds) before an idle node is scaled down.

Default: `600`

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlurmCustomSettings`

Additional Slurm-specific configuration that directly maps to Slurm settings.

_Required_: No

_Type_: Array of [SlurmCustomSetting](aws-properties-pcs-cluster-slurmcustomsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlurmdbdCustomSettings`

Additional SlurmDBD-specific configuration that directly maps to SlurmDBD settings.

_Required_: No

_Type_: Array of [SlurmdbdCustomSetting](aws-properties-pcs-cluster-slurmdbdcustomsetting.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SlurmRest`

The Slurm REST API configuration for the cluster.

_Required_: No

_Type_: [SlurmRest](aws-properties-pcs-cluster-slurmrest.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Scheduler

SlurmCustomSetting

All content copied from https://docs.aws.amazon.com/.
