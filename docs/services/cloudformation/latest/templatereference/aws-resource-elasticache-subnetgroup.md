This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::SubnetGroup

Creates a cache subnet group. For more information about cache subnet groups, go to Cache Subnet Groups in the
_Amazon ElastiCache User Guide_ or go to [CreateCacheSubnetGroup](https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_CreateCacheSubnetGroup.html) in the
_Amazon ElastiCache API Reference Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::SubnetGroup",
  "Properties" : {
      "CacheSubnetGroupName" : String,
      "Description" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::SubnetGroup
Properties:
  CacheSubnetGroupName: String
  Description: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`CacheSubnetGroupName`

The name for the cache subnet group. This value is stored as a lowercase
string.

Constraints: Must contain no more than 255 alphanumeric characters or hyphens.

Example: `mysubnetgroup`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description for the cache subnet group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SubnetIds`

The EC2 subnet IDs for the cache subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A tag that can be added to an ElastiCache subnet group. Tags are composed of a Key/Value pair. You can use tags
to categorize and track all your subnet groups. A tag with a null Value is permitted.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-elasticache-subnetgroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the Ref intrinsic function, Ref returns the resource
name.

For more information about using the Ref function, see [Ref](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/intrinsic-function-reference-ref.html).

## Examples

#### JSON

```json

{
    "SubnetGroup": {
        "Type": "AWS::ElastiCache::SubnetGroup",
        "Properties": {
            "Description": "Cache Subnet Group",
            "SubnetIds": [
                {
                    "Ref": "Subnet1"
                },
                {
                    "Ref": "Subnet2"
                }
            ]
        }
    }
}
```

#### YAML

```yaml

SubnetGroup:
  Type: 'AWS::ElastiCache::SubnetGroup'
  Properties:
    Description: Cache Subnet Group
    SubnetIds:
      - !Ref Subnet1
      - !Ref Subnet2
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
