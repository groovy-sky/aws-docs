This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Timestream::Table RetentionProperties

Retention properties contain the duration for which your time-series data must be stored in the magnetic store
and the memory store.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MagneticStoreRetentionPeriodInDays" : String,
  "MemoryStoreRetentionPeriodInHours" : String
}

```

### YAML

```yaml

  MagneticStoreRetentionPeriodInDays: String
  MemoryStoreRetentionPeriodInHours: String

```

## Properties

`MagneticStoreRetentionPeriodInDays`

The duration for which data must be stored in the magnetic store.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryStoreRetentionPeriodInHours`

The duration for which data must be stored in the memory store.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PartitionKey

S3Configuration
