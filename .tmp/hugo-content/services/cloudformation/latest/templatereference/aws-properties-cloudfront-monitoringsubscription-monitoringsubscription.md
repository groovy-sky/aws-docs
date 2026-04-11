This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::MonitoringSubscription MonitoringSubscription

A monitoring subscription. This structure contains information about whether
additional CloudWatch metrics are enabled for a given CloudFront distribution.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RealtimeMetricsSubscriptionConfig" : RealtimeMetricsSubscriptionConfig
}

```

### YAML

```yaml

  RealtimeMetricsSubscriptionConfig:
    RealtimeMetricsSubscriptionConfig

```

## Properties

`RealtimeMetricsSubscriptionConfig`

A subscription configuration for additional CloudWatch metrics.

_Required_: No

_Type_: [RealtimeMetricsSubscriptionConfig](aws-properties-cloudfront-monitoringsubscription-realtimemetricssubscriptionconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::CloudFront::MonitoringSubscription

RealtimeMetricsSubscriptionConfig

All content copied from https://docs.aws.amazon.com/.
