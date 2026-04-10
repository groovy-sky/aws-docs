This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::SecurityGroup

The `AWS::ElastiCache::SecurityGroup` resource creates a cache security group. For more information
about cache security groups, go to [CacheSecurityGroups](../../../amazonelasticache/latest/dg/vpcs.md) in the _Amazon ElastiCache User Guide_ or go to [CreateCacheSecurityGroup](../../../../reference/amazonelasticache/latest/apireference/api-createcachesecuritygroup.md) in the _Amazon ElastiCache API Reference Guide_.

For more information, see [CreateCacheSubnetGroup](../../../../reference/amazonelasticache/latest/apireference/api-createcachesubnetgroup.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::SecurityGroup",
  "Properties" : {
      "Description" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::SecurityGroup
Properties:
  Description: String
  Tags:
    - Tag

```

## Properties

`Description`

A description for the cache security group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A tag that can be added to an ElastiCache security group. Tags are composed of a Key/Value pair. You can use
tags to categorize and track all your security groups. A tag with a null Value is permitted.

_Required_: No

_Type_: Array of [Tag](aws-properties-elasticache-securitygroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
