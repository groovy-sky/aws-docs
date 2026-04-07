This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::Campaign WriteTreatmentResource

Specifies the settings for a campaign treatment. A
_treatment_ is a variation of a campaign that's used for
A/B testing of a campaign.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomDeliveryConfiguration" : CustomDeliveryConfiguration,
  "MessageConfiguration" : MessageConfiguration,
  "Schedule" : Schedule,
  "SizePercent" : Integer,
  "TemplateConfiguration" : TemplateConfiguration,
  "TreatmentDescription" : String,
  "TreatmentName" : String
}

```

### YAML

```yaml

  CustomDeliveryConfiguration:
    CustomDeliveryConfiguration
  MessageConfiguration:
    MessageConfiguration
  Schedule:
    Schedule
  SizePercent: Integer
  TemplateConfiguration:
    TemplateConfiguration
  TreatmentDescription: String
  TreatmentName: String

```

## Properties

`CustomDeliveryConfiguration`

The delivery configuration settings for sending the treatment through a custom
channel. This object is required if the `MessageConfiguration` object
for the treatment specifies a `CustomMessage` object.

_Required_: No

_Type_: [CustomDeliveryConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-customdeliveryconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MessageConfiguration`

The message configuration settings for the treatment.

_Required_: No

_Type_: [MessageConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-messageconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Schedule`

The schedule settings for the treatment.

_Required_: No

_Type_: [Schedule](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-schedule.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SizePercent`

The allocated percentage of users (segment members) to send the treatment
to.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TemplateConfiguration`

The message template to use for the treatment.

_Required_: No

_Type_: [TemplateConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-campaign-templateconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatmentDescription`

A custom description of the treatment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TreatmentName`

A custom name for the treatment.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TemplateConfiguration

AWS::Pinpoint::EmailChannel
