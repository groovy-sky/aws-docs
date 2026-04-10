This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaigns::Campaign

Contains information about an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ConnectCampaigns::Campaign",
  "Properties" : {
      "ConnectInstanceArn" : String,
      "DialerConfig" : DialerConfig,
      "Name" : String,
      "OutboundCallConfig" : OutboundCallConfig,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ConnectCampaigns::Campaign
Properties:
  ConnectInstanceArn: String
  DialerConfig:
    DialerConfig
  Name: String
  OutboundCallConfig:
    OutboundCallConfig
  Tags:
    - Tag

```

## Properties

`ConnectInstanceArn`

The Amazon Resource Name (ARN) of the Amazon Connect instance.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DialerConfig`

Contains information about the dialer configuration.

_Required_: Yes

_Type_: [DialerConfig](aws-properties-connectcampaigns-campaign-dialerconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the campaign.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `127`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OutboundCallConfig`

Contains information about the outbound call configuration.

_Required_: Yes

_Type_: [OutboundCallConfig](aws-properties-connectcampaigns-campaign-outboundcallconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags used to organize, track, or control access for this resource. For example, { "tags": {"key1":"value1", "key2":"value2"} }.

_Required_: No

_Type_: Array of [Tag](aws-properties-connectcampaigns-campaign-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the campaign name. For example:

`{ "Ref": "myCampaignName" }`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the high-volume outbound campaign.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon Connect Outbound Campaigns

AgentlessDialerConfig

All content copied from https://docs.aws.amazon.com/.
