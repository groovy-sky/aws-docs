This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityLake::DataLake LifecycleConfiguration

Provides lifecycle details of Amazon Security Lake object. To manage your data so that it is stored cost effectively, you can configure retention settings for the data. You can specify your preferred Amazon S3 storage class and the
time period for Amazon S3 objects to stay in that storage class before they transition to a
different storage class or expire. For more information about Amazon S3 Lifecycle
configurations, see [Managing your storage\
lifecycle](../../../s3/latest/userguide/object-lifecycle-mgmt.md) in the _Amazon Simple Storage Service User Guide_.

In Security Lake, you specify retention settings at the Region level. For example, you
might choose to transition all S3 objects in a specific AWS Region to the `S3
                Standard-IA` storage class 30 days after they're written to the data lake.
The default Amazon S3 storage class is S3 Standard.

###### Important

Security Lake doesn't support Amazon S3 Object Lock. When the data lake buckets are created, S3 Object Lock is disabled by
default. Enabling S3 Object Lock with default retention mode interrupts the delivery of normalized log data to the data lake.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Expiration" : Expiration,
  "Transitions" : [ Transitions, ... ]
}

```

### YAML

```yaml

  Expiration:
    Expiration
  Transitions:
    - Transitions

```

## Properties

`Expiration`

Provides data expiration details of the Amazon Security Lake object.

_Required_: No

_Type_: [Expiration](aws-properties-securitylake-datalake-expiration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Transitions`

Provides data storage transition details of Amazon Security Lake object. By configuring these settings, you can specify your preferred Amazon S3 storage class and the time period for S3 objects to stay in that storage class before they transition to a different storage class.

_Required_: No

_Type_: [Array](aws-properties-securitylake-datalake-transitions.md) of [Transitions](aws-properties-securitylake-datalake-transitions.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Expiration

ReplicationConfiguration

All content copied from https://docs.aws.amazon.com/.
