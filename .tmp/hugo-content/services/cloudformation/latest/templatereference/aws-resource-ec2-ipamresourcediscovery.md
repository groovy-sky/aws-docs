This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EC2::IPAMResourceDiscovery

A resource discovery is an IPAM component that enables IPAM to manage and monitor resources that belong to the owning account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EC2::IPAMResourceDiscovery",
  "Properties" : {
      "Description" : String,
      "OperatingRegions" : [ IpamOperatingRegion, ... ],
      "OrganizationalUnitExclusions" : [ IpamResourceDiscoveryOrganizationalUnitExclusion, ... ],
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::EC2::IPAMResourceDiscovery
Properties:
  Description: String
  OperatingRegions:
    - IpamOperatingRegion
  OrganizationalUnitExclusions:
    - IpamResourceDiscoveryOrganizationalUnitExclusion
  Tags:
    - Tag

```

## Properties

`Description`

The resource discovery description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatingRegions`

The operating Regions for the resource discovery. Operating Regions are AWS Regions where the IPAM is allowed to manage IP address CIDRs. IPAM only discovers and monitors resources in the AWS Regions you select as operating Regions.

_Required_: No

_Type_: Array of [IpamOperatingRegion](aws-properties-ec2-ipamresourcediscovery-ipamoperatingregion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OrganizationalUnitExclusions`

If your IPAM is integrated with AWS Organizations,
you can exclude an [organizational unit (OU)](../../../organizations/latest/userguide/orgs-getting-started-concepts.md#organizationalunit) from being managed by IPAM. When you exclude an OU, IPAM will not manage the IP addresses in accounts in that OU. For more information, see [Exclude organizational units from IPAM](../../../vpc/latest/ipam/exclude-ous.md) in the _Amazon Virtual Private Cloud IP Address Manager User Guide_.

_Required_: No

_Type_: Array of [IpamResourceDiscoveryOrganizationalUnitExclusion](aws-properties-ec2-ipamresourcediscovery-ipamresourcediscoveryorganizationalunitexclusion.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A tag is a label that you assign to an AWS resource. Each tag consists of a key and an optional value. You can use tags to search and filter your resources or track your AWS costs.

_Required_: No

_Type_: Array of [Tag](aws-properties-ec2-ipamresourcediscovery-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource discovery ID. For example: `ipam-res-disco-111122223333`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IpamResourceDiscoveryArn`

The resource discovery ARN.

`IpamResourceDiscoveryId`

The resource discovery ID.

`IpamResourceDiscoveryRegion`

The resource discovery Region.

`IsDefault`

Defines if the resource discovery is the default. The default resource discovery is the resource discovery automatically created when you create an IPAM.

`OwnerId`

The owner ID.

`State`

The resource discovery's state.

- `create-in-progress` \- Resource discovery is being created.

- `create-complete` \- Resource discovery creation is complete.

- `create-failed` \- Resource discovery creation has failed.

- `modify-in-progress` \- Resource discovery is being modified.

- `modify-complete` \- Resource discovery modification is complete.

- `modify-failed` \- Resource discovery modification has failed.

- `delete-in-progress` \- Resource discovery is being deleted.

- `delete-complete` \- Resource discovery deletion is complete.

- `delete-failed` \- Resource discovery deletion has failed.

- `isolate-in-progress` \- AWS account that created the resource discovery has been removed and the resource discovery is being isolated.

- `isolate-complete` \- Resource discovery isolation is complete.

- `restore-in-progress` \- AWS account that created the resource discovery and was isolated has been restored.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

IpamOperatingRegion

All content copied from https://docs.aws.amazon.com/.
