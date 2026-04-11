This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InternetMonitor::Monitor LocalHealthEventsConfig

Configuration information that determines the threshold and other conditions for when Internet Monitor creates a health event
for a local performance or availability issue, when scores cross a threshold for one or more city-networks.

Defines the percentages, for performance scores or availability scores, that are the local thresholds
for when Amazon CloudWatch Internet Monitor creates a health event. Also defines whether a local threshold is enabled or disabled, and the minimum percentage
of overall traffic that must be impacted by an issue before Internet Monitor creates an event when a threshold is crossed for a local health score.

If you don't set a local health event threshold, the default value is 60%.

For more information, see [Change health event thresholds](../../../amazoncloudwatch/latest/monitoring/cloudwatch-im-overview.md#IMUpdateThresholdFromOverview) in the Internet Monitor section of the _Amazon CloudWatch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthScoreThreshold" : Number,
  "MinTrafficImpact" : Number,
  "Status" : String
}

```

### YAML

```yaml

  HealthScoreThreshold: Number
  MinTrafficImpact: Number
  Status: String

```

## Properties

`HealthScoreThreshold`

The health event threshold percentage set for a local health score.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MinTrafficImpact`

The minimum percentage of overall traffic for an application that must be impacted by an issue before Internet Monitor creates an event when a
threshold is crossed for a local health score.

If you don't set a minimum traffic impact threshold, the default value is 0.01%.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of whether Internet Monitor creates a health event based on a threshold percentage set for a local health score. The status can be `ENABLED`
or `DISABLED`.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InternetMeasurementsLogDelivery

S3Config

All content copied from https://docs.aws.amazon.com/.
