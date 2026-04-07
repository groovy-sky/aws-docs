This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBSecurityGroup

The `AWS::RDS::DBSecurityGroup` resource creates or updates an Amazon RDS
DB security group.

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/vpc-migrate.html) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](https://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Non-VPC2VPC.html) in the _Amazon RDS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::RDS::DBSecurityGroup",
  "Properties" : {
      "DBSecurityGroupIngress" : [ Ingress, ... ],
      "EC2VpcId" : String,
      "GroupDescription" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBSecurityGroup
Properties:
  DBSecurityGroupIngress:
    - Ingress
  EC2VpcId: String
  GroupDescription: String
  Tags:
    - Tag

```

## Properties

`DBSecurityGroupIngress`

Ingress rules to be applied to the DB security group.

_Required_: Yes

_Type_: Array of [Ingress](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-dbsecuritygroup-ingress.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2VpcId`

The identifier of an Amazon virtual private cloud (VPC). This property indicates the VPC that this DB security
group belongs to.

###### Important

This property is included for backwards compatibility and is no longer recommended for providing security information to an RDS
DB instance.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`GroupDescription`

Provides the description of the DB security group.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Metadata assigned to an Amazon RDS resource consisting of a key-value pair.

For more information, see
[Tagging Amazon RDS resources](../../../amazonrds/latest/userguide/user-tagging.md) in the _Amazon RDS User Guide_ or
[Tagging Amazon Aurora and Amazon RDS resources](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.html) in the _Amazon Aurora User Guide_.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-rds-dbsecuritygroup-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the name of the DB security group.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

- [Creating a single VPC security group](#aws-resource-rds-dbsecuritygroup--examples--Creating_a_single_VPC_security_group)

- [Multiple VPC security groups](#aws-resource-rds-dbsecuritygroup--examples--Multiple_VPC_security_groups)

### Creating a single VPC security group

The following example creates a single VPC security group, referred to by
`EC2SecurityGroupName`.

#### JSON

```json

{
    "Resources": {
        "DBinstance": {
            "Type": "AWS::RDS::DBInstance",
            "Properties": {
                "DBSecurityGroups": [
                    {
                        "Ref": "DbSecurityByEC2SecurityGroup"
                    }
                ],
                "AllocatedStorage": "5",
                "DBInstanceClass": "db.t3.small",
                "Engine": "MySQL",
                "MasterUsername": "YourName",
                "MasterUserPassword": "YourPassword"
            },
            "DeletionPolicy": "Snapshot"
        },
        "DbSecurityByEC2SecurityGroup": {
            "Type": "AWS::RDS::DBSecurityGroup",
            "Properties": {
                "GroupDescription": "Ingress for Amazon EC2 security group",
                "DBSecurityGroupIngress": [
                    {
                        "EC2SecurityGroupId": "sg-b0ff1111",
                        "EC2SecurityGroupOwnerId": "111122223333"
                    },
                    {
                        "EC2SecurityGroupId": "sg-ffd722222",
                        "EC2SecurityGroupOwnerId": "111122223333"
                    }
                ]
            }
        }
    }
}
```

#### YAML

```yaml

Resources:
  DBinstance:
    Type: AWS::RDS::DBInstance
    Properties:
      DBSecurityGroups:
        -
          Ref: "DbSecurityByEC2SecurityGroup"
      AllocatedStorage: "5"
      DBInstanceClass: "db.t3.small"
      Engine: "MySQL"
      MasterUsername: "YourName"
      MasterUserPassword: "YourPassword"
    DeletionPolicy: "Snapshot"
  DbSecurityByEC2SecurityGroup:
    Type: AWS::RDS::DBSecurityGroup
    Properties:
      GroupDescription: "Ingress for Amazon EC2 security group"
      DBSecurityGroupIngress:
        -
          EC2SecurityGroupId: "sg-b0ff1111"
          EC2SecurityGroupOwnerId: "111122223333"
        -
          EC2SecurityGroupId: "sg-ffd722222"
          EC2SecurityGroupOwnerId: "111122223333"
```

### Multiple VPC security groups

The following example creates or updates multiple VPC security groups.

#### JSON

```json

"DBSecurityGroup": {
   "Type": "AWS::RDS::DBSecurityGroup",
   "Properties": {
      "EC2VpcId" : { "Ref" : "VpcId" },
      "DBSecurityGroupIngress": [
         {"EC2SecurityGroupName": { "Ref": "WebServerSecurityGroup"}}
      ],
      "GroupDescription": "Frontend Access"
   }
}
```

#### YAML

```yaml

DBSecurityGroup:
  Type: AWS::RDS::DBSecurityGroup
  Properties:
    EC2VpcId:
      Ref: "VpcId"
    DBSecurityGroupIngress:
      -
        EC2SecurityGroupName:
          Ref: "WebServerSecurityGroup"
    GroupDescription: "Frontend Access"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConnectionPoolConfigurationInfoFormat

Ingress
