This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::DataLake Expiration

Provides data expiration details of the Amazon Security Lake object. You can specify your preferred Amazon S3 storage class and the time period for S3 objects to stay in that storage class before they expire. For more information about Amazon S3 Lifecycle
configurations, see [Managing your storage\
lifecycle](../../../s3/latest/userguide/object-lifecycle-mgmt.md) in the _Amazon Simple Storage Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Days" : Integer
}

```

### YAML

```yaml

  Days: Integer

```

## Properties

`Days`

The number of days before data expires in the Amazon Security Lake object.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EncryptionConfiguration

LifecycleConfiguration

All content copied from https://docs.aws.amazon.com/.
