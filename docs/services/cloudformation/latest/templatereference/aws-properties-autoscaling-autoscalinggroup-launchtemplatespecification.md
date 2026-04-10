This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AutoScaling::AutoScalingGroup LaunchTemplateSpecification

Specifies a launch template to use when provisioning EC2 instances for an Auto Scaling
group.

You must specify the following:

- The ID or the name of the launch template, but not both.

- The version of the launch template.

`LaunchTemplateSpecification` is property of the [AWS::AutoScaling::AutoScalingGroup](../userguide/aws-resource-autoscaling-autoscalinggroup.md) resource. It is also a property of the [AWS::AutoScaling::AutoScalingGroup LaunchTemplate](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplate.md) and [AWS::AutoScaling::AutoScalingGroup LaunchTemplateOverrides](../userguide/aws-properties-autoscaling-autoscalinggroup-launchtemplateoverrides.md) property types.

For information about creating a launch template, see [AWS::EC2::LaunchTemplate](../userguide/aws-resource-ec2-launchtemplate.md) and [Create a launch template for an\
Auto Scaling group](../../../autoscaling/ec2/userguide/create-launch-template.md) in the _Amazon EC2 Auto Scaling User_
_Guide_.

For examples of launch templates, see [Create launch\
templates](../userguide/quickref-ec2-launch-templates.md).

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

You must specify the `LaunchTemplateID` or the `LaunchTemplateName`,
but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`LaunchTemplateName`

The name of the launch template.

You must specify the `LaunchTemplateName` or the `LaunchTemplateID`,
but not both.

_Required_: Conditional

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

`Version`

The version number of the launch template.

Specifying `$Latest` or `$Default` for the template version number
is not supported. However, you can specify `LatestVersionNumber` or
`DefaultVersionNumber` using the `Fn::GetAtt` intrinsic function. For
more information, see [Fn::GetAtt](../userguide/intrinsic-function-reference-getatt.md).

###### Note

For an example of using the `Fn::GetAtt` function, see the [Examples](../userguide/aws-resource-autoscaling-autoscalinggroup.md#aws-resource-autoscaling-autoscalinggroup--examples) section of the `AWS::AutoScaling::AutoScalingGroup`
resource.

_Required_: Yes

_Type_: String

_Update requires_: [Some interruptions](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-some-interrupt)

## See also

- [AWS::EC2::LaunchTemplate](../userguide/aws-resource-ec2-launchtemplate.md)

- [AWS::EC2::LaunchTemplate BlockDeviceMapping](../userguide/aws-properties-ec2-launchtemplate-blockdevicemapping.md)

- [Required\
KMS key policy for use with encrypted volumes](../../../autoscaling/ec2/userguide/key-policy-requirements-ebs-encryption.md) in the _Amazon EC2 Auto_
_Scaling User Guide_

- [Use\
encryption with EBS-backed AMIs](../../../ec2/latest/userguide/amiencryption.md) in the _Amazon EC2 User Guide for_
_Linux Instances_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LaunchTemplateOverrides

LifecycleHookSpecification

All content copied from https://docs.aws.amazon.com/.
