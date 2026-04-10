This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::VpcLattice::ServiceNetworkServiceAssociation

Associates the specified service with the specified service network. For more information,
see [Manage service associations](../../../vpc-lattice/latest/ug/service-network-associations.md#service-network-service-associations) in the _Amazon VPC Lattice User Guide_.

You can't use this operation if the service and service network are already associated or if
there is a disassociation or deletion in progress. If the association fails, you can retry the
operation by deleting the association and recreating it.

You cannot associate a service and service network that are shared with a caller. The caller
must own either the service or the service network.

As a result of this operation, the association is created in the service network account and
the association owner account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::VpcLattice::ServiceNetworkServiceAssociation",
  "Properties" : {
      "DnsEntry" : DnsEntry,
      "ServiceIdentifier" : String,
      "ServiceNetworkIdentifier" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::VpcLattice::ServiceNetworkServiceAssociation
Properties:
  DnsEntry:
    DnsEntry
  ServiceIdentifier: String
  ServiceNetworkIdentifier: String
  Tags:
    - Tag

```

## Properties

`DnsEntry`

The DNS information of the service.

_Required_: No

_Type_: [DnsEntry](aws-properties-vpclattice-servicenetworkserviceassociation-dnsentry.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ServiceIdentifier`

The ID or ARN of the service.

_Required_: No

_Type_: String

_Pattern_: `^((svc-[0-9a-z]{17})|(arn:[a-z0-9\-]+:vpc-lattice:[a-zA-Z0-9\-]+:\d{12}:service/svc-[0-9a-z]{17}))$`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Type_: Array of [Tag](aws-properties-vpclattice-servicenetworkserviceassociation-tag.md)

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the association.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the association between the service network and the
service.

`CreatedAt`

The date and time that the association was created, specified in ISO-8601 format.

`DnsEntry.DomainName`

The domain name of the service.

`DnsEntry.HostedZoneId`

The ID of the hosted zone.

`Id`

The ID of the of the association between the service network and the service.

`ServiceArn`

The Amazon Resource Name (ARN) of the service.

`ServiceId`

The ID of the service.

`ServiceName`

The name of the service.

`ServiceNetworkArn`

The Amazon Resource Name (ARN) of the service network

`ServiceNetworkId`

The ID of the service network.

`ServiceNetworkName`

The name of the service network.

`Status`

The status of the association between the service network and the service.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

DnsEntry

All content copied from https://docs.aws.amazon.com/.
