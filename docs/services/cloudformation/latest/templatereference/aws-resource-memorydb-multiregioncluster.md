This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MemoryDB::MultiRegionCluster

Represents a multi-Region cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::MemoryDB::MultiRegionCluster",
  "Properties" : {
      "Description" : String,
      "Engine" : String,
      "EngineVersion" : String,
      "MultiRegionClusterNameSuffix" : String,
      "MultiRegionParameterGroupName" : String,
      "NodeType" : String,
      "NumShards" : Integer,
      "Tags" : [ Tag, ... ],
      "TLSEnabled" : Boolean,
      "UpdateStrategy" : String
    }
}

```

### YAML

```yaml

Type: AWS::MemoryDB::MultiRegionCluster
Properties:
  Description: String
  Engine: String
  EngineVersion: String
  MultiRegionClusterNameSuffix: String
  MultiRegionParameterGroupName: String
  NodeType: String
  NumShards: Integer
  Tags:
    - Tag
  TLSEnabled: Boolean
  UpdateStrategy: String

```

## Properties

`Description`

The description of the multi-Region cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Engine`

The name of the engine used by the multi-Region cluster.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EngineVersion`

The version of the engine used by the multi-Region cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiRegionClusterNameSuffix`

A suffix to be added to the Multi-Region cluster name. Amazon MemoryDB automatically applies a prefix to the Multi-Region cluster Name when it is created. Each Amazon Region has its own prefix. For instance, a Multi-Region cluster Name created in the US-West-1 region will begin with "virxk", along with the suffix name you provide. The suffix guarantees uniqueness of the Multi-Region cluster name across multiple regions.

_Required_: No

_Type_: String

_Pattern_: `[a-z][a-z0-9\-]*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MultiRegionParameterGroupName`

The name of the multi-Region parameter group associated with the cluster.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`NodeType`

The node type used by the multi-Region cluster.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NumShards`

The number of shards in the multi-Region cluster.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to be applied to the multi-Region cluster.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-memorydb-multiregioncluster-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TLSEnabled`

Indiciates if the multi-Region cluster is TLS enabled.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`UpdateStrategy`

The strategy to use for the update operation. Supported values are "coordinated" or "uncoordinated".

_Required_: No

_Type_: String

_Allowed values_: `COORDINATED | UNCOORDINATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`ARN`

The Amazon Resource Name (ARN) of the multi-Region cluster.

`MultiRegionClusterName`

The name of the multi-Region cluster.

`Status`

The current status of the multi-Region cluster.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
