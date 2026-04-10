This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMResourceDiscoveryAssociation

An IPAM resource discovery association. An associated resource discovery is a resource discovery that has been associated with an IPAM. IPAM aggregates the resource CIDRs discovered by the associated resource discovery.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMResourceDiscoveryAssociation",
  "Properties" : {
      "IpamId" : String,
      "IpamResourceDiscoveryId" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMResourceDiscoveryAssociation
Properties:
  IpamId: String
  IpamResourceDiscoveryId: String
  Tags:
    - Tag

```

## Properties

`IpamId`

The IPAM ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IpamResourceDiscoveryId`

The resource discovery ID.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

A tag is a label that you assign to an AWS resource. Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-ipamresourcediscoveryassociation-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource discovery ID. For example: `ipam-res-disco-111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpamArn`

The IPAM ARN.

`IpamRegion`

The IPAM home Region.

`IpamResourceDiscoveryAssociationArn`

The resource discovery association ARN.

`IpamResourceDiscoveryAssociationId`

The resource discovery association ID.

`IsDefault`

Defines if the resource discovery is the default. When you create an IPAM, a default resource discovery is created for your IPAM and it's associated with your IPAM.

`OwnerId`

The owner ID.

`ResourceDiscoveryStatus`

The resource discovery status.

- `active` \- Connection or permissions required to read the
results of the resource discovery are intact.

- `not-found` \- Connection or permissions required to read the
results of the resource discovery are broken. This may happen if the owner of the resource discovery stopped sharing it or deleted the resource discovery. Verify the resource discovery still exists and the AWS RAM resource share is still intact.

`State`

The lifecycle state of the association when you associate or disassociate a resource discovery.

- `associate-in-progress` \- Resource discovery is being associated.

- `associate-complete` \- Resource discovery association is complete.

- `associate-failed` \- Resource discovery association has failed.

- `disassociate-in-progress` \- Resource discovery is being disassociated.

- `disassociate-complete` \- Resource discovery disassociation is complete.

- `disassociate-failed ` \- Resource discovery disassociation has failed.

- `isolate-in-progress` \- AWS account that created the resource discovery association has been removed and the resource discovery associatation is being isolated.

- `isolate-complete` \- Resource discovery isolation is complete..

- `restore-in-progress` \- Resource discovery is being restored.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
