This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration FastLaunchSnapshotConfiguration

Configuration settings for creating and managing pre-provisioned snapshots for a
fast-launch enabled Windows AMI.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TargetResourceCount" : Integer
}

```

### YAML

```yaml

  TargetResourceCount: Integer

```

## Properties

`TargetResourceCount`

The number of pre-provisioned snapshots to keep on hand for a fast-launch enabled
Windows AMI.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `10000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FastLaunchLaunchTemplateSpecification

LaunchPermissionConfiguration

All content copied from https://docs.aws.amazon.com/.
