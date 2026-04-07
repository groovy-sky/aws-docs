This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::FlowLog

Specifies a VPC flow log that captures IP traffic for a specified network interface,
subnet, or VPC. To view the log data, use Amazon CloudWatch Logs (CloudWatch Logs) to help
troubleshoot connection issues. For example, you can use a flow log to investigate why
certain traffic isn't reaching an instance, which can help you diagnose overly restrictive
security group rules. For more information, see [VPC Flow Logs](../../../vpc/latest/userguide/flow-logs.md) in the
_Amazon VPC User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::FlowLog",
  "Properties" : {
      "DeliverCrossAccountRole" : String,
      "DeliverLogsPermissionArn" : String,
      "DestinationOptions" : DestinationOptions,
      "LogDestination" : String,
      "LogDestinationType" : String,
      "LogFormat" : String,
      "LogGroupName" : String,
      "MaxAggregationInterval" : Integer,
      "ResourceId" : String,
      "ResourceType" : String,
      "Tags" : [ Tag, ... ],
      "TrafficType" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::FlowLog
Properties:
  DeliverCrossAccountRole: String
  DeliverLogsPermissionArn: String
  DestinationOptions:
    DestinationOptions
  LogDestination: String
  LogDestinationType: String
  LogFormat: String
  LogGroupName: String
  MaxAggregationInterval: Integer
  ResourceId: String
  ResourceType: String
  Tags:
    - Tag
  TrafficType: String

```

## Properties

`DeliverCrossAccountRole`

The ARN of the IAM role that allows the service to publish flow logs across accounts.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DeliverLogsPermissionArn`

The ARN of the IAM role that allows Amazon EC2 to publish flow logs to the log destination.

This parameter is required if the destination type is `cloud-watch-logs`,
or if the destination type is `kinesis-data-firehose` and the delivery stream
and the resources to monitor are in different accounts.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DestinationOptions`

The destination options.

_Required_: No

_Type_: [DestinationOptions](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-flowlog-destinationoptions.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogDestination`

The destination for the flow log data. The meaning of this parameter depends on the destination type.

- If the destination type is `cloud-watch-logs`, specify the ARN of a CloudWatch Logs log group. For example:

arn:aws:logs: _region_: _account\_id_:log-group: _my\_group_

Alternatively, use the `LogGroupName` parameter.

- If the destination type is `s3`, specify the ARN of an S3 bucket. For example:

arn:aws:s3::: _my\_bucket_/ _my\_subfolder_/

The subfolder is optional. Note that you can't use `AWSLogs` as a subfolder name.

- If the destination type is `kinesis-data-firehose`, specify the ARN of a Kinesis Data Firehose delivery stream. For example:

arn:aws:firehose: _region_: _account\_id_:deliverystream: _my\_stream_

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogDestinationType`

The type of destination for the flow log data.

Default: `cloud-watch-logs`

_Required_: No

_Type_: String

_Allowed values_: `cloud-watch-logs | s3 | kinesis-data-firehose`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogFormat`

The fields to include in the flow log record, in the order in which they should appear.
If you omit this parameter, the flow log is created using the default format. If you specify
this parameter, you must include at least one field. For more information about the available fields,
see [Flow log\
records](../../../vpc/latest/userguide/flow-logs.md#flow-log-records) in the _Amazon VPC User Guide_ or [Transit Gateway \
Flow Log records](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-flow-logs.html#flow-log-records) in the _AWS Transit Gateway Guide_.

Specify the fields using the `${field-id}` format, separated by
spaces.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LogGroupName`

The name of a new or existing CloudWatch Logs log group where Amazon EC2 publishes your flow logs.

This parameter is valid only if the destination type is `cloud-watch-logs`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxAggregationInterval`

The maximum interval of time during which a flow of packets is captured and aggregated into a flow log record.
The possible values are 60 seconds (1 minute) or 600 seconds (10 minutes).
This parameter must be 60 seconds for transit gateway resource types.

When a network interface is attached to a [Nitro-based\
instance](https://docs.aws.amazon.com/ec2/latest/instancetypes/ec2-nitro-instances.html), the aggregation interval is always 60 seconds or less, regardless
of the value that you specify.

Default: 600

_Required_: No

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceId`

The ID of the resource to monitor. For example, if the resource type is
`VPC`, specify the ID of the VPC.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceType`

The type of resource to monitor.

_Required_: Yes

_Type_: String

_Allowed values_: `NetworkInterface | Subnet | VPC | TransitGateway | TransitGatewayAttachment | RegionalNatGateway`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags to apply to the flow logs.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-flowlog-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TrafficType`

The type of traffic to monitor (accepted traffic, rejected traffic, or all traffic).
This parameter is not supported for transit gateway resource types. It is required for
the other resource types.

_Required_: No

_Type_: String

_Allowed values_: `ACCEPT | ALL | REJECT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the flow log ID, such as
`fl-123456abc123abc1`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the flow log. For example, `fl-123456abc123abc1`.

## Examples

- [Publish a flow log to CloudWatch Logs to monitor all traffic](#aws-resource-ec2-flowlog--examples--Publish_a_flow_log_to_CloudWatch_Logs_to_monitor_all_traffic)

- [Publish a custom format flow log to CloudWatch Logs for REJECT traffic](#aws-resource-ec2-flowlog--examples--Publish_a_custom_format_flow_log_to_CloudWatch_Logs_for_REJECT_traffic)

- [Publish a custom format flow log to Amazon S3 for ACCEPT traffic](#aws-resource-ec2-flowlog--examples--Publish_a_custom_format_flow_log_to_Amazon_S3_for_ACCEPT_traffic)

### Publish a flow log to CloudWatch Logs to monitor all traffic

The following example creates a flow log for the specified VPC, and captures all
traffic types. Amazon EC2 publishes the logs to the `FlowLogsGroup` log
group.

#### JSON

```json

{
  "MyFlowLog": {
    "Type": "AWS::EC2::FlowLog",
    "Properties": {
      "DeliverLogsPermissionArn": {
        "Fn::GetAtt": [
          "FlowLogRole",
          "Arn"
        ]
      },
      "LogGroupName": "FlowLogsGroup",
      "ResourceId": {
        "Ref": "MyVPC"
      },
      "ResourceType": "VPC",
      "TrafficType": "ALL"
    }
  }
}
```

#### YAML

```yaml

MyFlowLog:
  Type: AWS::EC2::FlowLog
  Properties:
    DeliverLogsPermissionArn: !GetAtt FlowLogRole.Arn
    LogGroupName: FlowLogsGroup
    ResourceId: !Ref MyVPC
    ResourceType: VPC
    TrafficType: ALL
```

### Publish a custom format flow log to CloudWatch Logs for REJECT traffic

The following example creates a flow log for the specified subnet and captures
REJECT traffic. The flow log uses a custom log format (the `LogFormat`
property uses the `${field-id}` format, separated by spaces). Amazon EC2
aggregates the logs over 60 second intervals, and publishes the logs to the
`FlowLogsGroup` log group. The flow log is created with two
tags.

#### JSON

```json

{
  "MyDetailedFlowLogDeliveringToCloudWatchLogs": {
    "Type": "AWS::EC2::FlowLog",
    "Properties": {
      "ResourceId": {
        "Ref": "MySubnet"
      },
      "ResourceType": "Subnet",
      "TrafficType": "REJECT",
      "LogGroupName": "FlowLogsGroup",
      "DeliverLogsPermissionArn": {
        "Fn::GetAtt": [
          "FlowLogRole",
          "Arn"
        ]
      },
      "LogFormat": "${version} ${vpc-id} ${subnet-id} ${instance-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${tcp-flags} ${type} ${pkt-srcaddr} ${pkt-dstaddr}",
      "MaxAggregationInterval": 60,
      "Tags": [
        {
          "Key": "Name",
          "Value": "FlowLogForSubnetA"
        },
        {
          "Key": "Purpose",
          "Value": "RejectTraffic"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

MyDetailedFlowLogDeliveringToCloudWatchLogs:
  Type: AWS::EC2::FlowLog
  Properties:
    ResourceId: !Ref MySubnet
    ResourceType: Subnet
    TrafficType: REJECT
    LogGroupName: FlowLogsGroup
    DeliverLogsPermissionArn: !GetAtt FlowLogRole.Arn
    LogFormat: ${version} ${vpc-id} ${subnet-id} ${instance-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${tcp-flags} ${type} ${pkt-srcaddr} ${pkt-dstaddr}
    MaxAggregationInterval: 60
    Tags:
      - Key: Name
        Value: FlowLogForSubnetA
      - Key: Purpose
        Value: RejectTraffic
```

### Publish a custom format flow log to Amazon S3 for ACCEPT traffic

The following example creates a flow log for the specified subnet and captures
ACCEPT traffic. The flow log uses a custom log format (the `LogFormat`
property uses the `${field-id}` format, separated by spaces). Amazon EC2
aggregates the logs over 60 second intervals, and publishes the logs to an Amazon S3
bucket that's referenced by its ARN, `MyS3Bucket.Arn`. The logs are
published in parquet format in Hive-compatible prefixes partitioned on an hourly basis.
The flow log is created with two tags.

#### JSON

```json

{
  "MyFlowLogDeliveringToS3": {
    "Type": "AWS::EC2::FlowLog",
    "Properties": {
      "ResourceId": {
        "Ref": "MySubnet"
      },
      "ResourceType": "Subnet",
      "TrafficType": "ACCEPT",
      "LogDestination": {
        "Fn::GetAtt": [
          "MyS3Bucket",
          "Arn"
        ]
      },
      "LogDestinationType": "s3",
      "LogFormat": "${version} ${vpc-id} ${subnet-id} ${instance-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${tcp-flags} ${type} ${pkt-srcaddr} ${pkt-dstaddr}",
      "MaxAggregationInterval": 60,
      "DestinationOptions": {
        "FileFormat": "parquet",
        "HiveCompatiblePartitions": true,
        "PerHourPartition": true
      },
      "Tags": [
        {
          "Key": "Name",
          "Value": "FlowLogForSubnetB"
        },
        {
          "Key": "Purpose",
          "Value": "AcceptTraffic"
        }
      ]
    }
  }
}
```

#### YAML

```yaml

MyFlowLogDeliveringToS3:
  Type: AWS::EC2::FlowLog
  Properties:
    ResourceId: !Ref MySubnet
    ResourceType: Subnet
    TrafficType: ACCEPT
    LogDestination: !GetAtt MyS3Bucket.Arn
    LogDestinationType: s3
    LogFormat: ${version} ${vpc-id} ${subnet-id} ${instance-id} ${srcaddr} ${dstaddr} ${srcport} ${dstport} ${protocol} ${tcp-flags} ${type} ${pkt-srcaddr} ${pkt-dstaddr}
    MaxAggregationInterval: 60
    DestinationOptions:
      FileFormat: parquet
      HiveCompatiblePartitions: true
      PerHourPartition: true
    Tags:
      - Key: Name
        Value: FlowLogForSubnetB
      - Key: Purpose
        Value: AcceptTraffic
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::EC2::EnclaveCertificateIamRoleAssociation

DestinationOptions
