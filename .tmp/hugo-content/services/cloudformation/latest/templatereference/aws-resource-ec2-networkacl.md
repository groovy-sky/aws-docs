This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::NetworkAcl

Specifies a network ACL for your VPC.

To add a network ACL entry, see [AWS::EC2::NetworkAclEntry](../userguide/aws-resource-ec2-networkaclentry.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::NetworkAcl",
  "Properties" : {
      "Tags" : [ Tag, ... ],
      "VpcId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EC2::NetworkAcl
Properties:
  Tags:
    - Tag
  VpcId: String

```

## Properties

`Tags`

The tags for the network ACL.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-networkacl-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcId`

The ID of the VPC for the network ACL.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the network ACL.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the network ACL.

## Examples

### Network ACL

The following example creates a network ACL.

#### JSON

```json

"myNetworkAcl" : {
   "Type" : "AWS::EC2::NetworkAcl",
   "Properties" : {
      "VpcId" : { "Ref" : "myVPC" },
      "Tags" : [ { "Key" : "stack", "Value" : "production" } ]
   }
}
```

#### YAML

```yaml

   myNetworkAcl:
      Type: AWS::EC2::NetworkAcl
      Properties:
         VpcId:
           Ref: myVPC
         Tags:
         - Key: stack
           Value: production
```

## See also

- [CreateNetworkAcl](../../../../reference/awsec2/latest/apireference/apireference-query-createnetworkacl.md) in the _Amazon EC2 API_
_Reference_

- [Network ACLs](../../../vpc/latest/userguide/vpc-network-acls.md) in the _Amazon VPC User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
