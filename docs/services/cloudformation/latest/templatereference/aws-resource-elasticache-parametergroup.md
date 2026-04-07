This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ParameterGroup

The `AWS::ElastiCache::ParameterGroup` type creates a new cache parameter group. Cache parameter
groups control the parameters for a cache cluster.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::ParameterGroup",
  "Properties" : {
      "CacheParameterGroupFamily" : String,
      "Description" : String,
      "Properties" : {Key: Value, ...},
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::ParameterGroup
Properties:
  CacheParameterGroupFamily: String
  Description: String
  Properties:
    Key: Value
  Tags:
    - Tag

```

## Properties

`CacheParameterGroupFamily`

The name of the cache parameter group family that this cache parameter group is compatible with.

Valid values are: `valkey8` \| `valkey7` \| `memcached1.4` \| `memcached1.5` \| `memcached1.6` \|
`redis2.6` \| `redis2.8` \| `redis3.2` \| `redis4.0` \| `redis5.0`
\| `redis6.x` \| `redis7`

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description for this cache parameter group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Properties`

A comma-delimited list of parameter name/value pairs.

For example:

```

"Properties" : {
   "cas_disabled" : "1",
   "chunk_size_growth_factor" : "1.02"
}

```

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z0-9]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A tag that can be added to an ElastiCache parameter group. Tags are composed of a Key/Value pair. You can use
tags to categorize and track all your parameter groups. A tag with a null Value is permitted.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-parametergroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

### Fn::GetAtt

`CacheParameterGroupName`

A user-specified name for the cache parameter group. This value is stored as a lowercase string.

## Examples

#### JSON

```json

{
    "MyParameterGroup": {
        "Type": "AWS::ElastiCache::ParameterGroup",
        "Properties": {
            "Description": "MyNewParameterGroup",
            "CacheParameterGroupFamily": "memcached1.4",
            "Properties": {
                "cas_disabled": "1",
                "chunk_size_growth_factor": "1.02"
            }
        }
    }
}
```

#### YAML

```yaml

MyParameterGroup:
  Type: 'AWS::ElastiCache::ParameterGroup'
  Properties:
    Description: MyNewParameterGroup
    CacheParameterGroupFamily: memcached1.4
    Properties:
      cas_disabled: '1'
      chunk_size_growth_factor: '1.02'
```

## See also

- [CreateCacheParameterGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheParameterGroup.html) in the _Amazon ElastiCache API Reference Guide_

- [ModifyCacheParameterGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_ModifyCacheParameterGroup.html) in the _Amazon ElastiCache API Reference Guide_

- [CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ReshardingConfiguration

Tag
