# ManagedPrefixList

Describes a managed prefix list.

## Contents

**addressFamily**

The IP address version.

Type: String

Required: No

**ipamPrefixListResolverSyncEnabled**

Indicates whether synchronization with an IPAM prefix list resolver is enabled for this managed prefix list. When enabled, the prefix list CIDRs are automatically updated based on the resolver's CIDR selection rules.

Type: Boolean

Required: No

**ipamPrefixListResolverTargetId**

The ID of the IPAM prefix list resolver target associated with this managed prefix list. When set, this prefix list becomes an IPAM managed prefix list.

An IPAM-managed prefix list is a customer-managed prefix list that has been associated with an IPAM prefix list resolver target. When a prefix list becomes IPAM managed, its CIDRs are automatically synchronized based on the IPAM prefix list resolver's CIDR selection rules, and direct CIDR modifications are restricted.

Type: String

Required: No

**maxEntries**

The maximum number of entries for the prefix list.

Type: Integer

Required: No

**ownerId**

The ID of the owner of the prefix list.

Type: String

Required: No

**prefixListArn**

The Amazon Resource Name (ARN) for the prefix list.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**prefixListId**

The ID of the prefix list.

Type: String

Required: No

**prefixListName**

The name of the prefix list.

Type: String

Required: No

**state**

The current state of the prefix list.

Type: String

Valid Values: `create-in-progress | create-complete | create-failed | modify-in-progress | modify-complete | modify-failed | restore-in-progress | restore-complete | restore-failed | delete-in-progress | delete-complete | delete-failed`

Required: No

**stateMessage**

The state message.

Type: String

Required: No

**TagSet.N**

The tags for the prefix list.

Type: Array of [Tag](api-tag.md) objects

Required: No

**version**

The version of the prefix list.

Type: Long

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/managedprefixlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/managedprefixlist.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/managedprefixlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

MaintenanceDetails

MediaAcceleratorInfo
