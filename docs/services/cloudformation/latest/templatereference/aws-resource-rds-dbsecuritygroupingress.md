This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBSecurityGroupIngress

The `AWS::RDS::DBSecurityGroupIngress` resource enables ingress to a DB
security group using one of two forms of authorization. First, you can add EC2 or VPC
security groups to the DB security group if the application using the database is
running on EC2 or VPC instances. Second, IP ranges are available if the application
accessing your database is running on the Internet.

This type supports updates. For more information about updating stacks, see [AWS \
CloudFormation Stacks Updates](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks.html).

For details about the settings for DB security group ingress, see [AuthorizeDBSecurityGroupIngress](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_AuthorizeDBSecurityGroupIngress.html).

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
  "Type" : "AWS::RDS::DBSecurityGroupIngress",
  "Properties" : {
      "CIDRIP" : String,
      "DBSecurityGroupName" : String,
      "EC2SecurityGroupId" : String,
      "EC2SecurityGroupName" : String,
      "EC2SecurityGroupOwnerId" : String
    }
}

```

### YAML

```yaml

Type: AWS::RDS::DBSecurityGroupIngress
Properties:
  CIDRIP: String
  DBSecurityGroupName: String
  EC2SecurityGroupId: String
  EC2SecurityGroupName: String
  EC2SecurityGroupOwnerId: String

```

## Properties

`CIDRIP`

The IP range to authorize.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DBSecurityGroupName`

The name of the DB security group to add authorization to.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2SecurityGroupId`

Id of the EC2 security group to authorize.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2SecurityGroupName`

Name of the EC2 security group to authorize.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName`
or `EC2SecurityGroupId` must be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2SecurityGroupOwnerId`

AWS account number of the owner of the EC2 security group
specified in the `EC2SecurityGroupName` parameter.
The AWS access key ID isn't an acceptable value.
For VPC DB security groups, `EC2SecurityGroupId` must be provided.
Otherwise, `EC2SecurityGroupOwnerId` and either `EC2SecurityGroupName` or `EC2SecurityGroupId` must be provided.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the DB security group that this ingress rule is associated
with.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Enable ingress to a DB security group

The following example creates a DB security group and allows ingress to it from a
specified VPC security group.

#### JSON

```json

{
  "Resources": {
    "MyDBSecurityGroupIngress": {
      "Type": "AWS::RDS::DBSecurityGroupIngress",
      "Properties": {
        "DBSecurityGroupName": {
          "Ref": "MyDBSecurityGroup"
        },
        "EC2SecurityGroupId": {
          "Ref": "MyVPCSecurityGroup"
        }
      }
    },
    "MyDBSecurityGroup": {
      "Type": "AWS::RDS::DBSecurityGroup",
      "Properties": {
        "GroupDescription": "My database security group"
      }
    },
    "MyVPCSecurityGroup": {
      "Type": "AWS::EC2::SecurityGroup",
      "Properties": {
        "GroupDescription": "My VPC security group",
        "VpcId": "vpc-12345678"
      }
    }
  }
}
```

#### YAML

```yaml

Resources:
  MyDBSecurityGroupIngress:
    Type: AWS::RDS::DBSecurityGroupIngress
    Properties:
      DBSecurityGroupName:
        Ref: MyDBSecurityGroup
      EC2SecurityGroupId:
        Ref: MyVPCSecurityGroup

  MyDBSecurityGroup:
    Type: AWS::RDS::DBSecurityGroup
    Properties:
      GroupDescription: My database security group

  MyVPCSecurityGroup:
    Type: AWS::EC2::SecurityGroup
    Properties:
      GroupDescription: My VPC security group
      VpcId: vpc-12345678
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::RDS::DBShardGroup
