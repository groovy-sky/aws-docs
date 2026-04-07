This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Pinpoint::ApplicationSettings

Specifies the settings for an Amazon Pinpoint application. In Amazon Pinpoint, an
_application_ (also referred to as an _app_ or
_project_) is a collection of related settings, customer
information, segments, and campaigns, and other types of Amazon Pinpoint
resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Pinpoint::ApplicationSettings",
  "Properties" : {
      "ApplicationId" : String,
      "CampaignHook" : CampaignHook,
      "CloudWatchMetricsEnabled" : Boolean,
      "Limits" : Limits,
      "QuietTime" : QuietTime
    }
}

```

### YAML

```yaml

Type: AWS::Pinpoint::ApplicationSettings
Properties:
  ApplicationId: String
  CampaignHook:
    CampaignHook
  CloudWatchMetricsEnabled: Boolean
  Limits:
    Limits
  QuietTime:
    QuietTime

```

## Properties

`ApplicationId`

The unique identifier for the Amazon Pinpoint application.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`CampaignHook`

The settings for the Lambda function to use by default as a code hook for campaigns in
the application. To override these settings for a specific campaign, use the Campaign
resource to define custom Lambda function settings for the campaign.

_Required_: No

_Type_: [CampaignHook](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-applicationsettings-campaignhook.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CloudWatchMetricsEnabled`

Property description not available.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Limits`

The default sending limits for campaigns in the application. To override these limits
for a specific campaign, use the Campaign resource to define custom limits for the
campaign.

_Required_: No

_Type_: [Limits](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-applicationsettings-limits.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QuietTime`

The default quiet time for campaigns in the application. Quiet time is a specific time
range when campaigns don't send messages to endpoints, if all the following conditions
are met:

\- The `EndpointDemographic.Timezone` property of the endpoint is set to a
valid value.

\- The current time in the endpoint's time zone is later than or equal to the time
specified by the `QuietTime.Start` property for the application (or a
campaign that has custom quiet time settings).

\- The current time in the endpoint's time zone is earlier than or equal to the time
specified by the `QuietTime.End` property for the application (or a campaign
that has custom quiet time settings).

If any of the preceding conditions isn't met, the endpoint will receive messages from
a campaign, even if quiet time is enabled.

To override the default quiet time settings for a specific campaign, use the Campaign
resource to define a custom quiet time for the campaign.

_Required_: No

_Type_: [QuietTime](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-pinpoint-applicationsettings-quiettime.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the unique identifier ( `ApplicationId`) for
the Amazon Pinpoint application that you're specifying the settings for.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Pinpoint::App

CampaignHook
