This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DMS::ReplicationSubnetGroup

The `AWS::DMS::ReplicationSubnetGroup` resource creates an AWS DMS
replication subnet group. Subnet groups must contain at least two subnets in two
different Availability Zones in the same AWS Region.

###### Note

Resource creation fails if the `dms-vpc-role` AWS Identity and Access Management (IAM)
role doesn't already exist. For more information, see
[Creating the IAM Roles to Use With the AWS CLI and AWS DMS API](https://docs.aws.amazon.com/dms/latest/userguide/CHAP_Security.APIRole.html)
in the _AWS Database Migration Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DMS::ReplicationSubnetGroup",
  "Properties" : {
      "ReplicationSubnetGroupDescription" : String,
      "ReplicationSubnetGroupIdentifier" : String,
      "SubnetIds" : [ String, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DMS::ReplicationSubnetGroup
Properties:
  ReplicationSubnetGroupDescription: String
  ReplicationSubnetGroupIdentifier: String
  SubnetIds:
    - String
  Tags:
    - Tag

```

## Properties

`ReplicationSubnetGroupDescription`

The description for the subnet group.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ReplicationSubnetGroupIdentifier`

The identifier for the replication subnet group. If you don't specify a name, CloudFormation
generates a unique ID and uses that ID for the identifier.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SubnetIds`

One or more subnet IDs to be assigned to the subnet group.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

One or more tags to be assigned to the subnet group.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-dms-replicationsubnetgroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the replication subnet group,
such as `mystack-myrepsubnetgroup-0a12bc456789de0fg`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

## Examples

#### JSON

```json

{
    "AWSTemplateFormatVersion": "2010-09-09",
    "Resources": {
        "myReplicationSubnetGroup": {
            "Type": "AWS::DMS::ReplicationSubnetGroup",
            "Properties": {
                "ReplicationSubnetGroupIdentifier": "identifier",
                "ReplicationSubnetGroupDescription": "description",
                "SubnetIds": [
                    "subnet-7b5b4112",
                    "subnet-7b5b4115"
                ],
                "Tags": [
                    {
                        "Key": "String",
                        "Value": "String"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  myReplicationSubnetGroup:
    Type: AWS::DMS::ReplicationSubnetGroup
    Properties:
      ReplicationSubnetGroupDescription: description
      ReplicationSubnetGroupIdentifier: identifier
      SubnetIds:
        - subnet-7b5b4112
        - subnet-7b5b4115
      Tags:
        - Key: String
          Value: String
```

## See also

- [CreateReplicationSubnetGroup](https://docs.aws.amazon.com/dms/latest/APIReference/API_CreateReplicationSubnetGroup.html)
in the _AWS Database Migration Service API Reference_

- [Managing AWS resources as a single unit with AWS CloudFormation stacks](../userguide/stacks.md)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
