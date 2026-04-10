This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore DatastoreStorage

Where data store data is stored.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManagedS3" : CustomerManagedS3,
  "IotSiteWiseMultiLayerStorage" : IotSiteWiseMultiLayerStorage,
  "ServiceManagedS3" : Json
}

```

### YAML

```yaml

  CustomerManagedS3:
    CustomerManagedS3
  IotSiteWiseMultiLayerStorage:
    IotSiteWiseMultiLayerStorage
  ServiceManagedS3: Json

```

## Properties

`CustomerManagedS3`

Use this to store data store data in an S3 bucket that you manage. The choice of
service-managed or customer-managed S3 storage cannot be changed after creation
of the data store.

_Required_: No

_Type_: [CustomerManagedS3](aws-properties-iotanalytics-datastore-customermanageds3.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IotSiteWiseMultiLayerStorage`

Use this to store data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
You can't change the choice of Amazon S3 storage after your data store is created.

_Required_: No

_Type_: [IotSiteWiseMultiLayerStorage](aws-properties-iotanalytics-datastore-iotsitewisemultilayerstorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceManagedS3`

Use this to store data store data in an S3 bucket managed by the AWS IoT Analytics service.
The choice of service-managed or customer-managed S3 storage cannot be changed after creation
of the data store.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DatastorePartitions

FileFormatConfiguration

All content copied from https://docs.aws.amazon.com/.
