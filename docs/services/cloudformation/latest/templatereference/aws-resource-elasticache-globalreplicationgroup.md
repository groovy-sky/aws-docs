This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::GlobalReplicationGroup

Consists of a primary cluster that accepts writes and an associated secondary cluster
that resides in a different Amazon region. The secondary cluster accepts only reads. The
primary cluster automatically replicates updates to the secondary cluster.

- The **GlobalReplicationGroupIdSuffix** represents
the name of the Global datastore, which is what you use to associate a secondary
cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::GlobalReplicationGroup",
  "Properties" : {
      "AutomaticFailoverEnabled" : Boolean,
      "CacheNodeType" : String,
      "CacheParameterGroupName" : String,
      "Engine" : String,
      "EngineVersion" : String,
      "GlobalNodeGroupCount" : Integer,
      "GlobalReplicationGroupDescription" : String,
      "GlobalReplicationGroupIdSuffix" : String,
      "Members" : [ GlobalReplicationGroupMember, ... ],
      "RegionalConfigurations" : [ RegionalConfiguration, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::GlobalReplicationGroup
Properties:
  AutomaticFailoverEnabled: Boolean
  CacheNodeType: String
  CacheParameterGroupName: String
  Engine: String
  EngineVersion: String
  GlobalNodeGroupCount: Integer
  GlobalReplicationGroupDescription: String
  GlobalReplicationGroupIdSuffix: String
  Members:
    - GlobalReplicationGroupMember
  RegionalConfigurations:
    - RegionalConfiguration

```

## Properties

`AutomaticFailoverEnabled`

Specifies whether a read-only replica is automatically promoted to read/write primary if the existing primary
fails.

`AutomaticFailoverEnabled` must be enabled for Valkey or Redis OSS (cluster mode enabled) replication
groups.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheNodeType`

The cache node type of the Global datastore

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CacheParameterGroupName`

The name of the cache parameter group to use with the Global datastore. It must be compatible with the major
engine version used by the Global datastore.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The ElastiCache engine. For Valkey or Redis OSS only.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The Elasticache Valkey or Redis OSS engine version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalNodeGroupCount`

The number of node groups that comprise the Global Datastore.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalReplicationGroupDescription`

The optional description of the Global datastore

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GlobalReplicationGroupIdSuffix`

The suffix name of a Global Datastore. The suffix guarantees uniqueness of the Global Datastore name across
multiple regions.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Members`

The replication groups that comprise the Global datastore.

_Required_: Yes

_Type_: Array of [GlobalReplicationGroupMember](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-globalreplicationgroup-globalreplicationgroupmember.html)

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegionalConfigurations`

The Regions that comprise the Global Datastore.

_Required_: No

_Type_: Array of [RegionalConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-globalreplicationgroup-regionalconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`GlobalReplicationGroupId`

The ID used to associate a secondary cluster to the Global Replication Group.

`Status`

The status of the Global Datastore. Can be `Creating`, `Modifying`,
`Available`, `Deleting` or `Primary-Only`. Primary-only status indicates the global
datastore contains only a primary cluster. Either all secondary clusters are deleted or not successfully
created.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

GlobalReplicationGroupMember
