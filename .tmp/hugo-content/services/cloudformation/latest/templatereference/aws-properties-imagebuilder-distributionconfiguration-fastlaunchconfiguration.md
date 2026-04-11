This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration FastLaunchConfiguration

Define and configure faster launching for output Windows AMIs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "Enabled" : Boolean,
  "LaunchTemplate" : FastLaunchLaunchTemplateSpecification,
  "MaxParallelLaunches" : Integer,
  "SnapshotConfiguration" : FastLaunchSnapshotConfiguration
}

```

### YAML

```yaml

  AccountId: String
  Enabled: Boolean
  LaunchTemplate:
    FastLaunchLaunchTemplateSpecification
  MaxParallelLaunches: Integer
  SnapshotConfiguration:
    FastLaunchSnapshotConfiguration

```

## Properties

`AccountId`

The owner account ID for the fast-launch enabled Windows AMI.

_Required_: No

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Enabled`

A Boolean that represents the current state of faster launching for the Windows AMI.
Set to `true` to start using Windows faster launching, or `false`
to stop using it.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplate`

The launch template that the fast-launch enabled Windows AMI uses when it launches
Windows instances to create pre-provisioned snapshots.

_Required_: No

_Type_: [FastLaunchLaunchTemplateSpecification](aws-properties-imagebuilder-distributionconfiguration-fastlaunchlaunchtemplatespecification.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxParallelLaunches`

The maximum number of parallel instances that are launched for creating
resources.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SnapshotConfiguration`

Configuration settings for managing the number of snapshots that are created from
pre-provisioned instances for the Windows AMI when faster launching is enabled.

_Required_: No

_Type_: [FastLaunchSnapshotConfiguration](aws-properties-imagebuilder-distributionconfiguration-fastlaunchsnapshotconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Distribution

FastLaunchLaunchTemplateSpecification

All content copied from https://docs.aws.amazon.com/.
