This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ConnectCampaignsV2::Campaign CommunicationLimitsConfig

Contains the communication limits configuration for an outbound campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AllChannelsSubtypes" : CommunicationLimits,
  "InstanceLimitsHandling" : String
}

```

### YAML

```yaml

  AllChannelsSubtypes:
    CommunicationLimits
  InstanceLimitsHandling: String

```

## Properties

`AllChannelsSubtypes`

The CommunicationLimits that apply to all channel subtypes defined in an outbound campaign.

_Required_: No

_Type_: [CommunicationLimits](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-connectcampaignsv2-campaign-communicationlimits.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InstanceLimitsHandling`

Opt-in or Opt-out from instance-level limits.

_Required_: No

_Type_: String

_Allowed values_: `OPT_IN | OPT_OUT`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CommunicationLimits

CommunicationTimeConfig
