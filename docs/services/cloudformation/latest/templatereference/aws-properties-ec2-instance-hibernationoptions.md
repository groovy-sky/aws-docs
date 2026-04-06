This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance HibernationOptions

Specifies the hibernation options for the instance.

`HibernationOptions` is a property of the [AWS::EC2::Instance](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance.html) resource.

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

Set to `true` to enable your instance for hibernation.

For Spot Instances, if you set `Configured` to `true`, either
omit the `InstanceInterruptionBehavior` parameter (for [`SpotMarketOptions`](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_SpotMarketOptions.html)), or set it to
`hibernate`. When `Configured` is true:

- If you omit `InstanceInterruptionBehavior`, it defaults to
`hibernate`.

- If you set `InstanceInterruptionBehavior` to a value other than
`hibernate`, you'll get an error.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EnclaveOptions

InstanceIpv6Address
