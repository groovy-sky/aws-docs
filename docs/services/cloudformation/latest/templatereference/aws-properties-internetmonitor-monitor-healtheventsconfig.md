This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InternetMonitor::Monitor HealthEventsConfig

Define the health event threshold percentages for the performance score and availability score for your application's
monitor. Amazon CloudWatch Internet Monitor creates a health event when there's an internet issue that
affects your application end users where a health score percentage is at or below a set threshold.

If you don't set a health event threshold, the default value is 95%.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AvailabilityLocalHealthEventsConfig" : LocalHealthEventsConfig,
  "AvailabilityScoreThreshold" : Number,
  "PerformanceLocalHealthEventsConfig" : LocalHealthEventsConfig,
  "PerformanceScoreThreshold" : Number
}

```

### YAML

```yaml

  AvailabilityLocalHealthEventsConfig:
    LocalHealthEventsConfig
  AvailabilityScoreThreshold: Number
  PerformanceLocalHealthEventsConfig:
    LocalHealthEventsConfig
  PerformanceScoreThreshold: Number

```

## Properties

`AvailabilityLocalHealthEventsConfig`

The configuration that determines the threshold and other conditions for when Internet Monitor creates a health event for a local availability issue.

_Required_: No

_Type_: [LocalHealthEventsConfig](aws-properties-internetmonitor-monitor-localhealtheventsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AvailabilityScoreThreshold`

The health event threshold percentage set for availability scores. When the overall availability
score is at or below this percentage, Internet Monitor creates a health event.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceLocalHealthEventsConfig`

The configuration that determines the threshold and other conditions for when Internet Monitor creates a health event for a local performance issue.

_Required_: No

_Type_: [LocalHealthEventsConfig](aws-properties-internetmonitor-monitor-localhealtheventsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PerformanceScoreThreshold`

The health event threshold percentage set for performance scores. When the overall performance
score is at or below this percentage, Internet Monitor creates a health event.

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::InternetMonitor::Monitor

InternetMeasurementsLogDelivery

All content copied from https://docs.aws.amazon.com/.
