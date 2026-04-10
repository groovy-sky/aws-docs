This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoTAnalytics::Datastore CustomerManagedS3

S3-customer-managed; When you choose customer-managed storage, the `retentionPeriod` parameter is ignored. You can't change the choice of Amazon S3 storage after your data store is created.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Bucket" : String,
  "KeyPrefix" : String,
  "RoleArn" : String
}

```

### YAML

```yaml

  Bucket: String
  KeyPrefix: String
  RoleArn: String

```

## Properties

`Bucket`

The name of the Amazon S3 bucket where your data is stored.

_Required_: Yes

_Type_: String

_Pattern_: `[a-zA-Z0-9.\-_]*`

_Minimum_: `3`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KeyPrefix`

(Optional) The prefix used to create the keys of the data store data objects. Each object in an Amazon S3 bucket has a key that is its unique identifier in the bucket. Each object in a bucket has exactly one key. The prefix must end with a forward slash (/).

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9!_.*'()/{}:-]*/`

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The ARN of the role that grants AWS IoT Analytics permission to interact with your Amazon S3 resources.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Column

CustomerManagedS3Storage

All content copied from https://docs.aws.amazon.com/.
