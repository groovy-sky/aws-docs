This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ServerlessCache

The resource representing a serverless cache.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::ServerlessCache",
  "Properties" : {
      "CacheUsageLimits" : CacheUsageLimits,
      "DailySnapshotTime" : String,
      "Description" : String,
      "Endpoint" : Endpoint,
      "Engine" : String,
      "FinalSnapshotName" : String,
      "KmsKeyId" : String,
      "MajorEngineVersion" : String,
      "ReaderEndpoint" : Endpoint,
      "SecurityGroupIds" : [ String, ... ],
      "ServerlessCacheName" : String,
      "SnapshotArnsToRestore" : [ String, ... ],
      "SnapshotRetentionLimit" : Integer,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "UserGroupId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::ServerlessCache
Properties:
  CacheUsageLimits:
    CacheUsageLimits
  DailySnapshotTime: String
  Description: String
  Endpoint:
    Endpoint
  Engine: String
  FinalSnapshotName: String
  KmsKeyId: String
  MajorEngineVersion: String
  ReaderEndpoint:
    Endpoint
  SecurityGroupIds:
    - String
  ServerlessCacheName: String
  SnapshotArnsToRestore:
    - String
  SnapshotRetentionLimit: Integer
  SubnetIds:
    - String
  Tags:
    - Tag
  UserGroupId: String

```

## Properties

`CacheUsageLimits`

The cache usage limit for the serverless cache.

_Required_: No

_Type_: [CacheUsageLimits](aws-properties-elasticache-serverlesscache-cacheusagelimits.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DailySnapshotTime`

The daily time that a cache snapshot will be created. Default is NULL, i.e. snapshots will not be created at a
specific time on a daily basis. Available for Valkey, Redis OSS and Serverless Memcached only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the serverless cache.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Endpoint`

Represents the information required for client programs to connect to a cache
node. This value is read-only.

_Required_: No

_Type_: [Endpoint](aws-properties-elasticache-serverlesscache-endpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The engine the serverless cache is compatible with.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FinalSnapshotName`

The name of the final snapshot taken of a cache before the cache is deleted.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyId`

The ID of the AWS Key Management Service (KMS) key that is used to encrypt data at rest in the serverless cache.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MajorEngineVersion`

The version number of the engine the serverless cache is compatible with.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReaderEndpoint`

Represents the information required for client programs to connect to a cache
node. This value is read-only.

_Required_: No

_Type_: [Endpoint](aws-properties-elasticache-serverlesscache-endpoint.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecurityGroupIds`

The IDs of the EC2 security groups associated with the serverless
cache.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServerlessCacheName`

The unique identifier of the serverless cache.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotArnsToRestore`

The ARN of the snapshot from which to restore data into the new cache.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SnapshotRetentionLimit`

The number of days for which ElastiCache retains automatic snapshots before deleting them. Available for Valkey, Redis OSS and Serverless Memcached only. The maximum value allowed is 35 days.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

If no subnet IDs are given and your VPC is in us-west-1, then ElastiCache will select 2 default subnets across AZs in your VPC.
For all other Regions, if no subnet IDs are given then ElastiCache will select 3 default subnets across AZs in your default VPC.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A list of tags to be added to this resource.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticache-serverlesscache-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserGroupId`

The identifier of the user group associated with the serverless cache. Available for Valkey and Redis OSS only. Default is NULL.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ARN`

The Amazon Resource Name (ARN) of the serverless cache.

`CreateTime`

When the serverless cache was created.

`Endpoint.Address`

The DNS hostname of the cache node.

`Endpoint.Port`

The port number that the cache engine is listening on.

`FullEngineVersion`

The name and version number of the engine the serverless cache is compatible with.

`ReaderEndpoint.Address`

The port used by the reader endpoint.

`ReaderEndpoint.Port`

The port used by the endpoint.

`Status`

The current status of the serverless cache. The allowed values are CREATING, AVAILABLE, DELETING, CREATE-FAILED and MODIFYING.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ElastiCache::SecurityGroupIngress

CacheUsageLimits

All content copied from https://docs.aws.amazon.com/.
