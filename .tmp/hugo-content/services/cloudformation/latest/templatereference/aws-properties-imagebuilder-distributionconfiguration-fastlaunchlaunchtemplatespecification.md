This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::DistributionConfiguration FastLaunchLaunchTemplateSpecification

Identifies the launch template that the associated Windows AMI uses for launching an
instance when faster launching is enabled.

###### Note

You can specify either the `launchTemplateName` or the
`launchTemplateId`, but not both.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "LaunchTemplateVersion" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  LaunchTemplateVersion: String

```

## Properties

`LaunchTemplateId`

The ID of the launch template to use for faster launching for a Windows AMI.

_Required_: No

_Type_: String

_Pattern_: `^lt-[a-z0-9-_]{17}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateName`

The name of the launch template to use for faster launching for a Windows AMI.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LaunchTemplateVersion`

The version of the launch template to use for faster launching for a Windows
AMI.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FastLaunchConfiguration

FastLaunchSnapshotConfiguration

All content copied from https://docs.aws.amazon.com/.
