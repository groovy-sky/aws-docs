This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::SecurityGroupIngress

The AWS::ElastiCache::SecurityGroupIngress type authorizes ingress to a cache security group from hosts in
specified Amazon EC2 security groups. For more information about ElastiCache security group ingress, go to [AuthorizeCacheSecurityGroupIngress](../../../../reference/amazonelasticache/latest/apireference/api-authorizecachesecuritygroupingress.md) in the _Amazon ElastiCache API Reference Guide_.

###### Note

Updates are not supported.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ElastiCache::SecurityGroupIngress",
  "Properties" : {
      "CacheSecurityGroupName" : String,
      "EC2SecurityGroupName" : String,
      "EC2SecurityGroupOwnerId" : String
    }
}

```

### YAML

```yaml

Type: AWS::ElastiCache::SecurityGroupIngress
Properties:
  CacheSecurityGroupName: String
  EC2SecurityGroupName: String
  EC2SecurityGroupOwnerId: String

```

## Properties

`CacheSecurityGroupName`

The name of the Cache Security Group to authorize.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2SecurityGroupName`

Name of the EC2 Security Group to include in the authorization.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EC2SecurityGroupOwnerId`

Specifies the Amazon Account ID of the owner of the EC2 security group specified in the EC2SecurityGroupName
property. The Amazon access key ID is not an acceptable value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When the logical ID of this resource is provided to the `Ref` intrinsic function, `Ref`
returns the resource name.

For more information about using the `Ref` function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::ElastiCache::ServerlessCache

All content copied from https://docs.aws.amazon.com/.
