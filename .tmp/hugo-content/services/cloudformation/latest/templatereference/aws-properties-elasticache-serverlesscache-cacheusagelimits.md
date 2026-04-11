This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ServerlessCache CacheUsageLimits

The usage limits for storage and ElastiCache Processing Units for the cache.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DataStorage" : DataStorage,
  "ECPUPerSecond" : ECPUPerSecond
}

```

### YAML

```yaml

  DataStorage:
    DataStorage
  ECPUPerSecond:
    ECPUPerSecond

```

## Properties

`DataStorage`

The maximum data storage limit in the cache, expressed in Gigabytes.

_Required_: No

_Type_: [DataStorage](aws-properties-elasticache-serverlesscache-datastorage.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ECPUPerSecond`

The number of ElastiCache Processing Units (ECPU) the cache can consume per second.

_Required_: No

_Type_: [ECPUPerSecond](aws-properties-elasticache-serverlesscache-ecpupersecond.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElastiCache::ServerlessCache

DataStorage

All content copied from https://docs.aws.amazon.com/.
