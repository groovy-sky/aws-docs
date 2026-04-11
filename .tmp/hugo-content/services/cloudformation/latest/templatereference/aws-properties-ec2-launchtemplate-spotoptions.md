This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::LaunchTemplate SpotOptions

Specifies options for Spot Instances.

`SpotOptions` is a property of [AWS::EC2::LaunchTemplate InstanceMarketOptions](../userguide/aws-properties-ec2-launchtemplate-launchtemplatedata-instancemarketoptions.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockDurationMinutes" : Integer,
  "InstanceInterruptionBehavior" : String,
  "MaxPrice" : String,
  "SpotInstanceType" : String,
  "ValidUntil" : String
}

```

### YAML

```yaml

  BlockDurationMinutes: Integer
  InstanceInterruptionBehavior: String
  MaxPrice: String
  SpotInstanceType: String
  ValidUntil: String

```

## Properties

`BlockDurationMinutes`

Deprecated.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceInterruptionBehavior`

The behavior when a Spot Instance is interrupted. The default is
`terminate`.

_Required_: No

_Type_: String

_Allowed values_: `hibernate | stop | terminate`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxPrice`

The maximum hourly price you're willing to pay for a Spot Instance. We do not
recommend using this parameter because it can lead to increased interruptions. If you do
not specify this parameter, you will pay the current Spot price. If you do specify this
parameter, it must be more than USD $0.001. Specifying a value below USD $0.001 will
result in an `InvalidParameterValue` error message when the launch template
is used to launch an instance.

###### Important

If you specify a maximum price, your Spot Instances will be interrupted more
frequently than if you do not specify this parameter.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SpotInstanceType`

The Spot Instance request type.

If you are using Spot Instances with an Auto Scaling group, use `one-time`
requests, as the Amazon EC2 Auto Scaling service handles requesting new Spot Instances
whenever the group is below its desired capacity.

_Required_: No

_Type_: String

_Allowed values_: `one-time | persistent`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidUntil`

The end date of the request, in UTC format
( _YYYY-MM-DD_ T _HH:MM:SS_ Z). Supported only for
persistent requests.

- For a persistent request, the request remains active until the
`ValidUntil` date and time is reached. Otherwise, the request
remains active until you cancel it.

- For a one-time request, `ValidUntil` is not supported. The request
remains active until all instances launch or you cancel the request.

Default: 7 days from the current date

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [LaunchTemplateSpotMarketOptionsRequest](../../../../reference/awsec2/latest/apireference/api-launchtemplatespotmarketoptionsrequest.md) in the _Amazon EC2_
_API Reference_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Reference

Tag

All content copied from https://docs.aws.amazon.com/.
