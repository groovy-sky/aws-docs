This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::SpotFleet FleetLaunchTemplateSpecification

Specifies the launch template to be used by the Spot Fleet request for configuring Amazon EC2 instances.

You must specify the following:

- The ID or the name of the launch template, but not both.

- The version of the launch template.

`FleetLaunchTemplateSpecification` is a property of the [AWS::EC2::SpotFleet](../userguide/aws-resource-ec2-spotfleet.md) resource.

For information about creating a launch template, see
[AWS::EC2::LaunchTemplate](../userguide/aws-resource-ec2-launchtemplate.md) and
[Create a launch template](../../../ec2/latest/userguide/ec2-launch-templates.md#create-launch-template)
in the _Amazon EC2 User Guide_.

For examples of launch templates, see [Examples](../userguide/aws-resource-ec2-launchtemplate.md#aws-resource-ec2-launchtemplate--examples).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LaunchTemplateId" : String,
  "LaunchTemplateName" : String,
  "Version" : String
}

```

### YAML

```yaml

  LaunchTemplateId: String
  LaunchTemplateName: String
  Version: String

```

## Properties

`LaunchTemplateId`

The ID of the launch template.

You must specify the `LaunchTemplateId` or the `LaunchTemplateName`, but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LaunchTemplateName`

The name of the launch template.

You must specify the `LaunchTemplateName` or the `LaunchTemplateId`, but not both.

_Required_: Conditional

_Type_: String

_Pattern_: `[a-zA-Z0-9\(\)\.\-/_]+`

_Minimum_: `3`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Version`

The version number of the launch template.

Specifying `$Latest` or `$Default` for the template version number
is not supported. However, you can specify
`LatestVersionNumber` or `DefaultVersionNumber` using the
`Fn::GetAtt` intrinsic function. For more information, see [Fn::GetAtt](../userguide/aws-resource-ec2-launchtemplate.md#aws-resource-ec2-launchtemplate-return-values-fn--getatt).

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EbsBlockDevice

GroupIdentifier

All content copied from https://docs.aws.amazon.com/.
