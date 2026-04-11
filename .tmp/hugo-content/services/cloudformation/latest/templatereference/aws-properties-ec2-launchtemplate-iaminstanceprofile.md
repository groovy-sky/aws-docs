This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate IamInstanceProfile

Specifies an IAM instance profile, which is a container for an IAM role for your
instance. You can use an IAM role to distribute your AWS credentials to
your instances.

If you are creating the launch template for use with an Amazon EC2 Auto Scaling group,
you can specify either the name or the ARN of the instance profile, but not both.

`IamInstanceProfile` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "Name" : String
}

```

### YAML

```yaml

  Arn: String
  Name: String

```

## Properties

`Arn`

The Amazon Resource Name (ARN) of the instance profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the instance profile.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateIamInstanceProfileSpecificationRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplateiaminstanceprofilespecificationrequest.md) in the
_Amazon EC2 API Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HibernationOptions

InstanceMarketOptions

All content copied from https://docs.aws.amazon.com/.
