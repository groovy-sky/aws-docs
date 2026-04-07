This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaTailor::PlaybackConfiguration LogConfiguration

Defines where AWS Elemental MediaTailor sends logs for the playback configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdsInteractionLog" : AdsInteractionLog,
  "EnabledLoggingStrategies" : [ String, ... ],
  "ManifestServiceInteractionLog" : ManifestServiceInteractionLog,
  "PercentEnabled" : Integer
}

```

### YAML

```yaml

  AdsInteractionLog:
    AdsInteractionLog
  EnabledLoggingStrategies:
    - String
  ManifestServiceInteractionLog:
    ManifestServiceInteractionLog
  PercentEnabled: Integer

```

## Properties

`AdsInteractionLog`

Settings for customizing what events are included in logs for interactions with the ad decision server (ADS).

_Required_: No

_Type_: [AdsInteractionLog](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-adsinteractionlog.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnabledLoggingStrategies`

The method used for collecting logs from AWS Elemental MediaTailor. `LEGACY_CLOUDWATCH` indicates that MediaTailor is sending logs directly to Amazon CloudWatch Logs. `VENDED_LOGS` indicates that MediaTailor is sending logs to CloudWatch, which then vends the logs to your destination of choice. Supported destinations are CloudWatch Logs log group, Amazon S3 bucket, and Amazon Data Firehose stream.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ManifestServiceInteractionLog`

Settings for customizing what events are included in logs for interactions with the origin server.

_Required_: No

_Type_: [ManifestServiceInteractionLog](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-mediatailor-playbackconfiguration-manifestserviceinteractionlog.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PercentEnabled`

The percentage of session logs that MediaTailor sends to your configured log destination. For example, if your playback configuration has 1000 sessions and `percentEnabled` is set to `60`, MediaTailor sends logs for 600 of the sessions to CloudWatch Logs. MediaTailor decides at random which of the playback configuration sessions to send logs for. If you want to view logs for a specific session, you can use the [debug log mode](https://docs.aws.amazon.com/mediatailor/latest/ug/debug-log-mode.html).

Valid values: `0` \- `100`

_Required_: Yes

_Type_: Integer

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

LivePreRollConfiguration

ManifestProcessingRules
