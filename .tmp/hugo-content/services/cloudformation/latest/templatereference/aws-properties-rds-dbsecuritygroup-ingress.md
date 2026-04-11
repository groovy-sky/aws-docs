This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RDS::DBSecurityGroup Ingress

The `Ingress` property type specifies an individual ingress rule within an
`AWS::RDS::DBSecurityGroup` resource.

###### Note

EC2-Classic was retired on August 15, 2022. If you haven't migrated from EC2-Classic to a VPC, we recommend that
you migrate as soon as possible. For more information, see [Migrate from EC2-Classic to a VPC](../../../ec2/latest/userguide/vpc-migrate.md) in the
_Amazon EC2 User Guide_, the blog [EC2-Classic Networking is Retiring – \
Here’s How to Prepare](https://aws.amazon.com/blogs/aws/ec2-classic-is-retiring-heres-how-to-prepare), and [Moving a DB instance not in a VPC \
into a VPC](../../../amazonrds/latest/userguide/user-vpc-non-vpc2vpc.md) in the _Amazon RDS User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CIDRIP" : String,
  "EC2SecurityGroupId" : String,
  "EC2SecurityGroupName" : String,
  "EC2SecurityGroupOwnerId" : String
}

```

### YAML

```yaml

  CIDRIP: String
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

## Examples

### Specifying an ingress rule

The following example specifies two security group ingress rules.

#### JSON

```json

"DBSecurityGroupIngress":[
   {
      "EC2SecurityGroupId":"sg-b0ff1111",
      "EC2SecurityGroupOwnerId":"111122223333"
   },
   {
      "EC2SecurityGroupId":"sg-ffd722222",
      "EC2SecurityGroupOwnerId":"111122223333"
   }
]
```

#### YAML

```yaml

DBSecurityGroupIngress:
  - EC2SecurityGroupId: sg-b0ff1111
    EC2SecurityGroupOwnerId: '111122223333'
  - EC2SecurityGroupId: sg-ffd722222
    EC2SecurityGroupOwnerId: '111122223333'
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::RDS::DBSecurityGroup

Tag

All content copied from https://docs.aws.amazon.com/.
