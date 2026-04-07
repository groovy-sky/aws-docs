This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTFleetWise::Campaign

Creates an orchestration of data collection rules. The AWS IoT FleetWise Edge Agent software
running in vehicles uses campaigns to decide how to collect and transfer data to the
cloud. You create campaigns in the cloud. After you or your team approve campaigns,
AWS IoT FleetWise automatically deploys them to vehicles.

For more information, see [Campaigns](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/campaigns.html) in the _AWS IoT FleetWise Developer Guide_.

###### Important

Access to certain AWS IoT FleetWise features is currently gated. For more information, see [AWS Region and feature availability](https://docs.aws.amazon.com/iot-fleetwise/latest/developerguide/fleetwise-regions.html) in the _AWS IoT FleetWise Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoTFleetWise::Campaign",
  "Properties" : {
      "Action" : String,
      "CollectionScheme" : CollectionScheme,
      "Compression" : String,
      "DataDestinationConfigs" : [ DataDestinationConfig, ... ],
      "DataExtraDimensions" : [ String, ... ],
      "DataPartitions" : [ DataPartition, ... ],
      "Description" : String,
      "DiagnosticsMode" : String,
      "ExpiryTime" : String,
      "Name" : String,
      "PostTriggerCollectionDuration" : Number,
      "Priority" : Integer,
      "SignalCatalogArn" : String,
      "SignalsToCollect" : [ SignalInformation, ... ],
      "SignalsToFetch" : [ SignalFetchInformation, ... ],
      "SpoolingMode" : String,
      "StartTime" : String,
      "Tags" : [ Tag, ... ],
      "TargetArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoTFleetWise::Campaign
Properties:
  Action: String
  CollectionScheme:
    CollectionScheme
  Compression: String
  DataDestinationConfigs:
    - DataDestinationConfig
  DataExtraDimensions:
    - String
  DataPartitions:
    - DataPartition
  Description: String
  DiagnosticsMode: String
  ExpiryTime: String
  Name: String
  PostTriggerCollectionDuration: Number
  Priority: Integer
  SignalCatalogArn: String
  SignalsToCollect:
    - SignalInformation
  SignalsToFetch:
    - SignalFetchInformation
  SpoolingMode: String
  StartTime: String
  Tags:
    - Tag
  TargetArn: String

```

## Properties

`Action`

Specifies how to update a campaign. The action can be one of the following:

- `APPROVE` \- To approve delivering a data collection scheme to
vehicles.

- `SUSPEND` \- To suspend collecting signal data. The campaign is deleted from vehicles and all vehicles in the suspended campaign will stop sending data.

- `RESUME` \- To reactivate the `SUSPEND` campaign. The campaign is redeployed to all vehicles and the vehicles will resume sending data.

- `UPDATE` \- To update a campaign.

_Required_: No

_Type_: String

_Allowed values_: `APPROVE | SUSPEND | RESUME | UPDATE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CollectionScheme`

The data collection scheme associated with the campaign. You can specify a scheme
that collects data based on time or an event.

_Required_: Yes

_Type_: [CollectionScheme](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-collectionscheme.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Compression`

Whether to compress signals before transmitting data to AWS IoT FleetWise. If you
don't want to compress the signals, use `OFF`. If it's not specified,
`SNAPPY` is used.

Default: `SNAPPY`

_Required_: No

_Type_: String

_Allowed values_: `OFF | SNAPPY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataDestinationConfigs`

The destination where the campaign sends data. You can choose to send data to be stored in Amazon S3 or Amazon Timestream.

Amazon S3 optimizes the cost of data storage and provides additional mechanisms to use vehicle data, such as data lakes, centralized data storage, data processing pipelines, and analytics. AWS IoT FleetWise supports at-least-once file delivery to S3. Your vehicle data is stored on multiple AWS IoT FleetWise servers for redundancy and high availability.

You can use Amazon Timestream to access and analyze time series data, and Timestream to query
vehicle data so that you can identify trends and patterns.

_Required_: No

_Type_: Array of [DataDestinationConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-datadestinationconfig.html)

_Minimum_: `1`

_Maximum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataExtraDimensions`

A list of vehicle attributes to associate with a campaign.

Enrich the data with specified vehicle attributes. For example, add `make` and `model` to the campaign, and AWS IoT FleetWise will associate the data with those attributes as dimensions in Amazon Timestream. You can then query the data against `make` and `model`.

Default: An empty array

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 0`

_Maximum_: `150 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DataPartitions`

The data partitions associated with the signals collected from the vehicle.

_Required_: No

_Type_: Array of [DataPartition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-datapartition.html)

_Minimum_: `0`

_Maximum_: `20`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the campaign.

_Required_: No

_Type_: String

_Pattern_: `^[^\u0000-\u001F\u007F]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DiagnosticsMode`

Option for a vehicle to send diagnostic trouble codes to AWS IoT FleetWise. If you
want to send diagnostic trouble codes, use `SEND_ACTIVE_DTCS`. If it's not
specified, `OFF` is used.

Default: `OFF`

_Required_: No

_Type_: String

_Allowed values_: `OFF | SEND_ACTIVE_DTCS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExpiryTime`

The time the campaign expires, in seconds since epoch (January 1, 1970 at
midnight UTC time). Vehicle data isn't collected after the campaign expires.

Default: 253402214400 (December 31, 9999, 00:00:00 UTC)

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of a campaign.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z\d\-_:]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PostTriggerCollectionDuration`

How long (in milliseconds) to collect raw data after a triggering event
initiates the collection. If it's not specified, `0` is used.

Default: `0`

_Required_: No

_Type_: Number

_Minimum_: `0`

_Maximum_: `4294967295`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Priority`

A number indicating the priority of one campaign over another campaign for
a certain vehicle or fleet. A campaign with the lowest value is deployed to vehicles
before any other campaigns. If it's not specified, `0` is used.

Default: `0`

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignalCatalogArn`

The Amazon Resource Name (ARN) of the signal catalog associated with the campaign.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignalsToCollect`

A list of information about signals to collect.

_Required_: No

_Type_: Array of [SignalInformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-signalinformation.html)

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SignalsToFetch`

A list of information about signals to fetch.

_Required_: No

_Type_: Array of [SignalFetchInformation](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-signalfetchinformation.html)

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SpoolingMode`

Whether to store collected data after a vehicle lost a connection with the
cloud. After a connection is re-established, the data is automatically forwarded to
AWS IoT FleetWise. If you want to store collected data when a vehicle loses connection with the
cloud, use `TO_DISK`. If it's not specified, `OFF` is used.

Default: `OFF`

_Required_: No

_Type_: String

_Allowed values_: `OFF | TO_DISK`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StartTime`

The time, in milliseconds, to deliver a campaign after it was approved. If
it's not specified, `0` is used.

Default: `0`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata that can be used to manage the campaign.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iotfleetwise-campaign-tag.html)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetArn`

The Amazon Resource Name (ARN) of a vehicle or fleet to which the campaign is deployed.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the campaign.

`CreationTime`

The time the campaign was created in seconds since epoch (January 1, 1970 at midnight
UTC time).

`LastModificationTime`

The last time the campaign was modified.

`Status`

The state of the campaign. The status can be one of: `CREATING`,
`WAITING_FOR_APPROVAL`, `RUNNING`, and `SUSPENDED`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS IoT FleetWise

CollectionScheme
