This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBSubnetGroup

The `AWS::RDS::DBSubnetGroup` resource creates a database subnet group.
Subnet groups must contain at least two subnets in two different Availability Zones in
the same region.

For more information, see [Working with DB subnet groups](../../../amazonrds/latest/userguide/user-vpc-workingwithrdsinstanceinavpc.md#USER_VPC.Subnets) in the _Amazon RDS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBSubnetGroup",
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

Type: AWS::RDS::DBSubnetGroup
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

The description for the DB subnet group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSubnetGroupName`

The name for the DB subnet group. This value is stored as a lowercase string.

Constraints:

- Must contain no more than 255 letters, numbers, periods, underscores, spaces, or hyphens.

- Must not be default.

- First character must be a letter.

Example: `mydbsubnetgroup`

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

The EC2 Subnet IDs for the DB subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to assign to the DB subnet group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-dbsubnetgroup-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB subnet group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create a DB subnet group

The following example creates a DB subnet group with two subnet IDs and a set of tags.

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "myDBSubnetGroup": {
            "Type": "AWS::RDS::DBSubnetGroup",
            "Properties": {
                "DBSubnetGroupDescription": "Description of subnet group",
                "SubnetIds": [
                    "subnet-7b5b4112",
                    "subnet-7b5b4115"
                ],
                "Tags": [
                    {
                        "Key": "mykey",
                        "Value": "myvalue"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

---
AWSTemplateFormatVersion: "2010-09-09"
Resources:
  myDBSubnetGroup:
    Type: AWS::RDS::DBSubnetGroup
    Properties:
      DBSubnetGroupDescription: Description of subnet group
      SubnetIds:
        - subnet-7b5b4112
        - subnet-7b5b4115
      Tags:
        - Key: mykey
          Value: myvalue
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
