This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate InstanceMarketOptions

Specifies the market (purchasing) option for an instance.

`InstanceMarketOptions` is a property of the [AWS::EC2::LaunchTemplate LaunchTemplateData](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MarketType" : String,
  "SpotOptions" : SpotOptions
}

```

### YAML

```yaml

  MarketType: String
  SpotOptions:
    SpotOptions

```

## Properties

`MarketType`

The market type.

_Required_: Conditional

_Type_: String

_Allowed values_: `spot | capacity-block | interruptible-capacity-reservation`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotOptions`

The options for Spot Instances.

_Required_: No

_Type_: [SpotOptions](aws-properties-ec2-launchtemplate-spotoptions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateInstanceMarketOptionsRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplateinstancemarketoptionsrequest.md) in the _Amazon_
_EC2 API Reference_

- [Create a launch template using advanced settings](../../../autoscaling/ec2/userguide/advanced-settings-for-your-launch-template.md) in the _Amazon EC2 Auto Scaling User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

IamInstanceProfile

InstanceRequirements

All content copied from https://docs.aws.amazon.com/.
