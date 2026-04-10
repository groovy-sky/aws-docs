This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::Instance HibernationOptions

Specifies the hibernation options for the instance.

`HibernationOptions` is a property of the [AWS::EC2::Instance](../userguide/aws-properties-ec2-instance.md) resource.

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
omit the `InstanceInterruptionBehavior` parameter (for [`SpotMarketOptions`](../../../../reference/awsec2/latest/apireference/api-spotmarketoptions.md)), or set it to
`hibernate`. When `Configured` is true:

- If you omit `InstanceInterruptionBehavior`, it defaults to
`hibernate`.

- If you set `InstanceInterruptionBehavior` to a value other than
`hibernate`, you'll get an error.

Default: `false`

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EnclaveOptions

InstanceIpv6Address

All content copied from https://docs.aws.amazon.com/.
