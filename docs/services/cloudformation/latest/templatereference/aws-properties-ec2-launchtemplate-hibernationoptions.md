This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate HibernationOptions

Specifies whether your instance is configured for hibernation. This parameter is valid
only if the instance meets the [hibernation\
prerequisites](../../../ec2/latest/userguide/hibernate.md#hibernating-prerequisites). For more information, see [Hibernate Your Instance](../../../ec2/latest/userguide/hibernate.md) in the
_Amazon EC2 User Guide_.

`HibernationOptions` is a property of [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Configured" : Boolean
}

```

### YAML

```yaml

  Configured: Boolean

```

## Properties

`Configured`

If you set this parameter to `true`, the instance is enabled for
hibernation.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnclaveOptions

IamInstanceProfile

All content copied from https://docs.aws.amazon.com/.
