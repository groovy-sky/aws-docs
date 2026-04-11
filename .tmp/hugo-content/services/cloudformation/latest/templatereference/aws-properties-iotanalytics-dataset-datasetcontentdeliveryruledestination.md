This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Dataset DatasetContentDeliveryRuleDestination

The destination to which dataset contents are delivered.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IotEventsDestinationConfiguration" : IotEventsDestinationConfiguration,
  "S3DestinationConfiguration" : S3DestinationConfiguration
}

```

### YAML

```yaml

  IotEventsDestinationConfiguration:
    IotEventsDestinationConfiguration
  S3DestinationConfiguration:
    S3DestinationConfiguration

```

## Properties

`IotEventsDestinationConfiguration`

Configuration information for delivery of dataset contents to AWS IoT Events.

_Required_: No

_Type_: [IotEventsDestinationConfiguration](aws-properties-iotanalytics-dataset-ioteventsdestinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3DestinationConfiguration`

Configuration information for delivery of dataset contents to Amazon S3.

_Required_: No

_Type_: [S3DestinationConfiguration](aws-properties-iotanalytics-dataset-s3destinationconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatasetContentDeliveryRule

DatasetContentVersionValue

All content copied from https://docs.aws.amazon.com/.
