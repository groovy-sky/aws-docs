This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::DataLake Transitions

Provides transition lifecycle details of the Amazon Security Lake object. For more information about Amazon S3 Lifecycle
configurations, see [Managing your storage\
lifecycle](../../../s3/latest/userguide/object-lifecycle-mgmt.md) in the _Amazon Simple Storage Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Days" : Integer,
  "StorageClass" : String
}

```

### YAML

```yaml

  Days: Integer
  StorageClass: String

```

## Properties

`Days`

The number of days before data transitions to a different S3 Storage Class in the Amazon Security Lake object.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StorageClass`

The list of storage classes that you can choose from based on the data access,
resiliency, and cost requirements of your workloads.
The default storage class is **S3 Standard**.
For information about other storage classes, see
[Setting the storage class of an object](../../../s3/latest/userguide/sc-howtoset.md)
in the _Amazon S3 User Guide_.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::SecurityLake::Subscriber

All content copied from https://docs.aws.amazon.com/.
