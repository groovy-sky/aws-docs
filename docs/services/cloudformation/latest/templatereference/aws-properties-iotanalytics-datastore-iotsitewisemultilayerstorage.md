This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore IotSiteWiseMultiLayerStorage

Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.
You can't change the choice of Amazon S3 storage after your data store is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CustomerManagedS3Storage" : CustomerManagedS3Storage
}

```

### YAML

```yaml

  CustomerManagedS3Storage:
    CustomerManagedS3Storage

```

## Properties

`CustomerManagedS3Storage`

Stores data used by AWS IoT SiteWise in an Amazon S3 bucket that you manage.

_Required_: No

_Type_: [CustomerManagedS3Storage](aws-properties-iotanalytics-datastore-customermanageds3storage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FileFormatConfiguration

ParquetConfiguration

All content copied from https://docs.aws.amazon.com/.
