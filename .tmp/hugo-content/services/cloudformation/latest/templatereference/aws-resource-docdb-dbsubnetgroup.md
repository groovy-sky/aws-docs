This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DocDB::DBSubnetGroup

The `AWS::DocDB::DBSubnetGroup` Amazon DocumentDB (with MongoDB compatibility) resource describes a DBSubnetGroup.
subnet groups must contain at least one subnet in at least two Availability Zones in the AWS Region.
For more information, see [DBSubnetGroup](../../../documentdb/latest/developerguide/api-dbsubnetgroup.md) in the _Amazon DocumentDB Developer Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DocDB::DBSubnetGroup",
  "Properties" : {
      "DBSubnetGroupDescription" : String,
      "DBSubnetGroupName" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DocDB::DBSubnetGroup
Properties:
  DBSubnetGroupDescription: String
  DBSubnetGroupName: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`DBSubnetGroupDescription`

The description for the subnet group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSubnetGroupName`

The name for the subnet group. This value is stored as a lowercase string.

Constraints: Must contain no more than 255 letters, numbers, periods, underscores,
spaces, or hyphens. Must not be default.

Example: `mySubnetgroup`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The Amazon EC2 subnet IDs for the subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

The tags to be assigned to the subnet group.

_Required_: No

_Type_: Array of [Tag](aws-properties-docdb-dbsubnetgroup-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DBSubnetGroup's name, such as `default`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

#### JSON

```json

{
   "AWSTemplateFormatVersion" : "2010-09-09",
   "Resources" : {
      "myDBSubnetGroup" : {
         "Type" : "AWS::DocDB::DBSubnetGroup",
         "Properties" : {
            "DBSubnetGroupDescription" : "description",
            "SubnetIds" : [ "subnet-7b5b4112", "subnet-7b5b4115" ],
            "Tags" : [ {"Key" : "String", "Value" : "String"} ]
         }
      }
   }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: "2010-09-09"
Resources:
   myDBSubnetGroup:
      Type: "AWS::DocDB::DBSubnetGroup"
      Properties:
         DBSubnetGroupDescription: "description"
         SubnetIds:
            - "subnet-7b5b4112"
            - "subnet-7b5b4115"
         Tags:
            -
               Key: "String"
               Value: "String"
```

## See also

- [Subnet](../../../documentdb/latest/developerguide/api-subnet.md)

- [DBSubnetGroup](../../../documentdb/latest/developerguide/api-dbsubnetgroup.md)

- [CreateDBSubnetGroup](../../../documentdb/latest/developerguide/api-createdbsubnetgroup.md)

- [DeleteDBSubnetGroup](../../../documentdb/latest/developerguide/api-deletedbsubnetgroup.md)

- [DescribeDBSubnetGroups](../../../documentdb/latest/developerguide/api-describedbsubnetgroups.md)

- [ModifyDBSubnetGroup](../../../documentdb/latest/developerguide/api-modifydbsubnetgroup.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
