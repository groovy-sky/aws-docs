This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Channel ChannelStorage

Where channel data is stored. You may choose one of `serviceManagedS3`,
`customerManagedS3` storage. If not specified, the default is
`serviceManagedS3`. This can't be changed after creation of the channel.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManagedS3" : CustomerManagedS3,
  "ServiceManagedS3" : Json
}

```

### YAML

```yaml

  CustomerManagedS3:
    CustomerManagedS3
  ServiceManagedS3: Json

```

## Properties

`CustomerManagedS3`

Used to store channel data in an S3 bucket that you manage. If customer managed storage is
selected, the `retentionPeriod` parameter is ignored. You can't change the choice
of S3 storage after the data store is created.

_Required_: No

_Type_: [CustomerManagedS3](aws-properties-iotanalytics-channel-customermanageds3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceManagedS3`

Used to store channel data in an S3 bucket managed by AWS IoT Analytics. You can't change the choice
of S3 storage after the data store is created.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IoTAnalytics::Channel

CustomerManagedS3

All content copied from https://docs.aws.amazon.com/.
