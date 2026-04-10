This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InternetMonitor::Monitor

The `AWS::InternetMonitor::Monitor` resource is an Internet Monitor resource type that contains information about how
you create a monitor in Amazon CloudWatch Internet Monitor. A monitor in Internet Monitor provides visibility into performance and
availability between your applications hosted on AWS and your end users, using a traffic profile that it creates
based on the application resources that you add: Virtual Private Clouds (VPCs), Amazon CloudFront distributions, or WorkSpaces directories.

Internet Monitor also alerts you to internet issues that impact your application in the city-networks (geographies and networks)
where your end users use it. With Internet Monitor, you can quickly pinpoint the locations and providers that are affected, so
that you can address the issue.

For more information, see [Using Amazon CloudWatch Internet Monitor](../../../amazoncloudwatch/latest/monitoring/cloudwatch-internetmonitor.md) in the _Amazon CloudWatch User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::InternetMonitor::Monitor",
  "Properties" : {
      "HealthEventsConfig" : HealthEventsConfig,
      "IncludeLinkedAccounts" : Boolean,
      "InternetMeasurementsLogDelivery" : InternetMeasurementsLogDelivery,
      "LinkedAccountId" : String,
      "MaxCityNetworksToMonitor" : Integer,
      "MonitorName" : String,
      "Resources" : [ String, ... ],
      "ResourcesToAdd" : [ String, ... ],
      "ResourcesToRemove" : [ String, ... ],
      "Status" : String,
      "Tags" : [ Tag, ... ],
      "TrafficPercentageToMonitor" : Integer
    }
}

```

### YAML

```yaml

Type: AWS::InternetMonitor::Monitor
Properties:
  HealthEventsConfig:
    HealthEventsConfig
  IncludeLinkedAccounts: Boolean
  InternetMeasurementsLogDelivery:
    InternetMeasurementsLogDelivery
  LinkedAccountId: String
  MaxCityNetworksToMonitor: Integer
  MonitorName: String
  Resources:
    - String
  ResourcesToAdd:
    - String
  ResourcesToRemove:
    - String
  Status: String
  Tags:
    - Tag
  TrafficPercentageToMonitor: Integer

```

## Properties

`HealthEventsConfig`

A complex type with the configuration information that determines the threshold and other conditions for when Internet Monitor creates a health event
for an overall performance or availability issue, across an application's geographies.

Defines the percentages, for overall performance scores and availability scores for an application, that are the thresholds
for when Internet Monitor creates a health event. You can override the defaults to set a custom threshold for overall performance or availability scores,
or both.

You can also set thresholds for local health scores,, where Internet Monitor creates a health event when scores cross a threshold for one or more city-networks,
in addition to creating an event when an overall score crosses a threshold.

If you don't set a health event threshold, the default value is 95%.

For local thresholds, you also set a minimum percentage of overall traffic that is impacted by an issue before Internet Monitor creates an event.
In addition, you can disable local thresholds, for performance scores, availability scores, or both.

For more information, see [Change health event thresholds](../../../amazoncloudwatch/latest/monitoring/cloudwatch-im-overview.md#IMUpdateThresholdFromOverview) in the Internet Monitor section of the _CloudWatch User Guide_.

_Required_: No

_Type_: [HealthEventsConfig](aws-properties-internetmonitor-monitor-healtheventsconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IncludeLinkedAccounts`

A boolean option that you can set to `TRUE` to include monitors for linked accounts in a list of
monitors, when you've set up cross-account sharing in Internet Monitor. You configure cross-account
sharing by using Amazon CloudWatch Observability Access Manager. For more information, see
[Internet Monitor cross-account\
observability](../../../amazoncloudwatch/latest/monitoring/cwim-cross-account.md) in the Amazon CloudWatch User Guide.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InternetMeasurementsLogDelivery`

Publish internet measurements for a monitor for all city-networks (up to the 500,000 service limit) to another location, such as an Amazon S3 bucket.
Measurements are also published to Amazon CloudWatch Logs for the first 500 (by traffic volume) city-networks (client locations and ASNs, typically
internet service providers or ISPs).

_Required_: No

_Type_: [InternetMeasurementsLogDelivery](aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LinkedAccountId`

The account ID for an account that you've set up cross-account sharing for in Internet Monitor. You configure cross-account
sharing by using Amazon CloudWatch Observability Access Manager. For more information, see
[Internet Monitor cross-account\
observability](../../../amazoncloudwatch/latest/monitoring/cwim-cross-account.md) in the Amazon CloudWatch User Guide.

_Required_: No

_Type_: String

_Pattern_: `^(\d{12})$`

_Minimum_: `12`

_Maximum_: `12`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaxCityNetworksToMonitor`

The maximum number of city-networks to monitor for your resources. A city-network is the location (city) where
clients access your application resources from and the network, such as an internet service provider, that clients
access the resources through.

For more information, see [Choosing a city-network maximum value](../../../amazoncloudwatch/latest/monitoring/imcitynetworksmaximum.md) in _Using Amazon CloudWatch Internet Monitor_.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `500000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MonitorName`

The name of the monitor. A monitor name can contain only alphanumeric characters, dashes (-), periods (.),
and underscores (\_).

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9_.-]+$`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Resources`

The resources that have been added for the monitor, listed by their Amazon Resource Names (ARNs). Use this option to add or
remove resources when making an update.

###### Note

Be aware that if you include content in the `Resources` field when you update a monitor, the `ResourcesToAdd`
and `ResourcesToRemove` fields must be empty.

_Required_: No

_Type_: Array of String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcesToAdd`

The resources to include in a monitor, which you provide as a set of Amazon Resource Names (ARNs). Resources can be Amazon
Virtual Private Cloud VPCs, Network Load Balancers (NLBs), Amazon CloudFront distributions, or Amazon WorkSpaces directories.

You can add a combination of VPCs and CloudFront distributions, or you can add WorkSpaces directories, or you can add NLBs. You can't
add NLBs or WorkSpaces directories together with any other resources.

If you add only VPC resources, at least one VPC must have an Internet Gateway attached to it, to make sure that it has internet connectivity.

###### Note

You can specify this field for a monitor update only if the `Resources` field is empty.

_Required_: No

_Type_: Array of String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcesToRemove`

The resources to remove from a monitor, which you provide as a set of Amazon Resource Names (ARNs)

###### Note

You can specify this field for a monitor update only if the `Resources` field is empty.

_Required_: No

_Type_: Array of String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The status of a monitor. The accepted values that you can specify for `Status` are `ACTIVE` and `INACTIVE`.

_Required_: No

_Type_: String

_Allowed values_: `PENDING | ACTIVE | INACTIVE | ERROR`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags for a monitor, listed as a set of _key:value_ pairs.

_Required_: No

_Type_: Array of [Tag](aws-properties-internetmonitor-monitor-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficPercentageToMonitor`

The percentage of the internet-facing traffic for your application that you want to monitor. You can also, optionally, set
a limit for the number of city-networks (client locations and ASNs, typically internet service providers) that Internet Monitor will monitor traffic for.
The city-networks maximum limit caps the number of city-networks that Internet Monitor monitors for your application, regardless of
the percentage of traffic that you choose to monitor.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ARN of the monitor, such as
`arn:aws:internetmonitor:us-east-1:111122223333:monitor/TestMonitor`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreatedAt`

The time when the monitor was created.

`ModifiedAt`

The last time that the monitor was modified.

`MonitorArn`

The Amazon Resource Name (ARN) of the monitor.

`ProcessingStatus`

The health of data processing for the monitor. For more information, see `ProcessingStatus` under
[MonitorListMember](../../../internet-monitor/latest/api/api-monitorlistmember.md) in the _Amazon CloudWatch Internet Monitor API Reference_.

`ProcessingStatusInfo`

Additional information about the health of the data processing for the monitor.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon CloudWatch Internet Monitor

HealthEventsConfig

All content copied from https://docs.aws.amazon.com/.
