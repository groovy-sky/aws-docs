---
title: "CloudWatch metrics for your VPCs"
---

# CloudWatch metrics for your VPCs

Amazon VPC publishes data about your VPCs to Amazon CloudWatch. You can retrieve statistics about your VPCs
as an ordered set of time-series data, known as _metrics_. Think of a metric
as a variable to monitor and the data as the value of that variable over time. For more
information, see the [Amazon CloudWatch User Guide](../../../amazoncloudwatch/latest/monitoring.md).

###### Contents

- [NAU metrics and dimensions](#nau-cloudwatch)

- [Enable or disable NAU monitoring](#nau-monitoring-enable)

- [NAU CloudWatch alarm example](#nau-cloudwatch-alarm-example)

## NAU metrics and dimensions

[Network Address Usage](network-address-usage.md)
(NAU) is a metric applied to resources in your virtual network to help you plan for and
monitor the size of your VPC. There is no cost to monitor NAU. Monitoring NAU is helpful
because if you exhaust the NAU or peered NAU quotas for your VPC, you can't launch new
EC2 instances or provision new resources, such as Network Load Balancers, VPC endpoints, Lambda
functions, transit gateway attachments, and NAT gateways.

If you've enabled Network Address Usage monitoring for a VPC, Amazon VPC sends metrics
related to NAU to Amazon CloudWatch. The size of a VPC is measured by the number of Network Address
Usage (NAU) units that the VPC contains.

You can use these metrics to understand the rate of your VPC growth, forecast when
your VPC will reach its size limit, or create alarms when size thresholds are crossed.

The `AWS/EC2` namespace includes the following metrics for monitoring NAU.

MetricDescription`NetworkAddressUsage`

The NAU count per VPC.

**Reporting criteria**

- Every 24 hours.

**Dimensions**

- Name: `Per-VPC Metrics`, Value: The VPC
ID.

`NetworkAddressUsagePeered`The NAU count for the VPC and all VPCs that it's peered
with.

**Reporting**
**criteria**

- Every 24 hours.

**Dimensions**

- Name: `Per-VPC Metrics`, Value: The VPC
ID.

The `AWS/Usage` namespace includes the following metrics for monitoring NAU.

MetricDescription`ResourceCount`

The NAU count per VPC.

**Reporting criteria**

- Every 24 hours.

**Dimensions**

- Name: `Service`, Value: `EC2`

- Name: `Type`, Value: `Resource`

- Name: `Resource`, Value: The VPC ID.

- Name: `Class`, Value: `NetworkAddressUsage`

`ResourceCount`

The NAU count for the VPC and all VPCs that it's peered with.

**Reporting criteria**

- Every 24 hours.

**Dimensions**

- Name: `Service`, Value: `EC2`

- Name: `Type`, Value: `Resource`

- Name: `Resource`, Value: The VPC ID.

- Name: `Class`, Value: `NetworkAddressUsagePeered`

`ResourceCount`

A combined view of NAU usage across VPCs.

**Reporting criteria**

- Every 24 hours.

**Dimensions**

- Name: `Service`, Value: `EC2`

- Name: `Type`, Value: `Resource`

- Name: `Resource`, Value: `VPC`

- Name: `Class`, Value: `NetworkAddressUsage`

`ResourceCount`

A combined view of NAU usage across peered VPCs.

**Reporting criteria**

- Every 24 hours.

**Dimensions**

- Name: `Service`, Value: `EC2`

- Name: `Type`, Value: `Resource`

- Name: `Resource`, Value: `VPC`

- Name: `Class`, Value: `NetworkAddressUsagePeered`

## Enable or disable NAU monitoring

To view NAU metrics in CloudWatch, you must first enable monitoring on each VPC to monitor.

###### To enable or disable monitoring NAU

1. Open the Amazon VPC console at [https://console.aws.amazon.com/vpc/](https://console.aws.amazon.com/vpc).

2. In the navigation pane, choose **Your VPCs**.

3. Select the check box for the VPC.

4. Select **Actions**, **Edit VPC settings**.

5. Do one of the following:

- To enable monitoring, select **Network mapping units metrics settings**,
**Enable network address usage metrics**.

- To disable monitoring, clear **Network mapping units metrics settings**,
**Enable network address usage metrics**.

###### To enable or disable monitoring using the command line

- [modify-vpc-attribute](../../../cli/latest/reference/ec2/modify-vpc-attribute.md)
(AWS CLI)

- [Edit-EC2VpcAttribute](../../../powershell/latest/reference/items/edit-ec2vpcattribute.md)
(AWS Tools for Windows PowerShell)

## NAU CloudWatch alarm example

You can use the following AWS CLI command and example `.json` to create an
Amazon CloudWatch alarm and SNS notification that tracks NAU utilization of the VPC with 50,000
NAUs as the threshold. This sample requires you to first create an Amazon SNS topic. For more
information, see [Getting started with Amazon SNS](../../../sns/latest/dg/sns-getting-started.md) in
the _Amazon Simple Notification Service Developer Guide_.

```nohighlight

aws cloudwatch put-metric-alarm --cli-input-json file://nau-alarm.json
```

The following is an example of `nau-alarm.json`.

```json

{
    "Namespace": "AWS/EC2",
    "MetricName": "NetworkAddressUsage",
    "Dimensions": [{
        "Name": "Per-VPC Metrics",
        "Value": "vpc-0123456798"
    }],
    "AlarmActions": ["arn:aws:sns:us-west-1:123456789012:my_sns_topic"],
    "ComparisonOperator": "GreaterThanThreshold",
    "Period": 86400,
    "EvaluationPeriods": 1,
    "Threshold": 50000,
    "AlarmDescription": "Tracks NAU utilization of the VPC with 50k NAUs as the threshold",
    "AlarmName": "VPC NAU Utilization",
    "Statistic": "Maximum"
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Troubleshoot

Billing and usage reports

All content copied from https://docs.aws.amazon.com/.
