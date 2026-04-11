This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetworkVpcAssociation

Associates a VPC with a service network. When you associate a VPC with the service network,
it enables all the resources within that VPC to be clients and communicate with other services in
the service network. For more information, see [Manage VPC associations](../../../vpc-lattice/latest/ug/service-network-associations.md#service-network-vpc-associations) in the _Amazon VPC Lattice User Guide_.

You can't use this operation if there is a disassociation in progress. If the association
fails, retry by deleting the association and recreating it.

As a result of this operation, the association gets created in the service network account
and the VPC owner account.

If you add a security group to the service network and VPC association, the association must
continue to always have at least one security group. You can add or edit security groups at any
time. However, to remove all security groups, you must first delete the association and recreate
it without security groups.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ServiceNetworkVpcAssociation",
  "Properties" : {
      "DnsOptions" : DnsOptions,
      "PrivateDnsEnabled" : Boolean,
      "SecurityGroupIds" : [ String, ... ],
      "ServiceNetworkIdentifier" : String,
      "Tags" : [ Tag, ... ],
      "VpcIdentifier" : String
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ServiceNetworkVpcAssociation
Properties:
  DnsOptions:
    DnsOptions
  PrivateDnsEnabled: Boolean
  SecurityGroupIds:
    - String
  ServiceNetworkIdentifier: String
  Tags:
    - Tag
  VpcIdentifier: String

```

## Properties

`DnsOptions`

The DNS options for the service network VPC association.

_Required_: No

_Type_: [DnsOptions](aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PrivateDnsEnabled`

Indicates if private DNS is enabled for the service network VPC association.

_Required_: No

_Type_: Boolean

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecurityGroupIds`

The IDs of the security groups. Security groups aren't added by default. You can add a
security group to apply network level controls to control which resources in a VPC are allowed to
access the service network and its services. For more information, see [Control traffic to\
resources using security groups](../../../vpc/latest/userguide/vpc-securitygroups.md) in the _Amazon VPC User_
_Guide_.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceNetworkIdentifier`

The ID or ARN of the service network. You must use an ARN if the resources are in different
accounts.

_Required_: No

_Type_: String

_Pattern_: `^((sn-[0-9a-z]{17})|(arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:servicenetwork/sn-[0-9a-z]{17}))$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

The tags for the association.

_Required_: No

_Type_: Array of [Tag](aws-properties-vpclattice-servicenetworkvpcassociation-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VpcIdentifier`

The ID of the VPC.

_Required_: No

_Type_: String

_Pattern_: `^vpc-(([0-9a-z]{8})|([0-9a-z]{17}))$`

_Minimum_: `5`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the association between the service network and the
VPC.

`CreatedAt`

The date and time that the association was created, specified in ISO-8601 format.

`Id`

The ID of the specified association between the service network and the VPC.

`ServiceNetworkArn`

The Amazon Resource Name (ARN) of the service network.

`ServiceNetworkId`

The ID of the service network.

`ServiceNetworkName`

The name of the service network.

`Status`

The status of the association.

`VpcId`

The ID of the VPC.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DnsOptions

All content copied from https://docs.aws.amazon.com/.
