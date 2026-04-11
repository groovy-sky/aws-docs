# ModifyManagedPrefixList

Modifies the specified managed prefix list.

Adding or removing entries in a prefix list creates a new version of the prefix list.
Changing the name of the prefix list does not affect the version.

If you specify a current version number that does not match the true current version
number, the request fails.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AddEntry.N**

One or more entries to add to the prefix list.

Type: Array of [AddPrefixListEntry](api-addprefixlistentry.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: No

**CurrentVersion**

The current version of the prefix list.

Type: Long

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamPrefixListResolverSyncEnabled**

Indicates whether synchronization with an IPAM prefix list resolver should be enabled for this managed prefix list. When enabled, the prefix list CIDRs are automatically updated based on the associated resolver's CIDR selection rules.

Type: Boolean

Required: No

**MaxEntries**

The maximum number of entries for the prefix list. You cannot modify the entries
of a prefix list and modify the size of a prefix list at the same time.

If any of the resources that reference the prefix list cannot support the new
maximum size, the modify operation fails. Check the state message for the IDs of
the first ten resources that do not support the new maximum size.

Type: Integer

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

**PrefixListName**

A name for the prefix list.

Type: String

Required: No

**RemoveEntry.N**

One or more entries to remove from the prefix list.

Type: Array of [RemovePrefixListEntry](api-removeprefixlistentry.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: No

## Response Elements

The following elements are returned by the service.

**prefixList**

Information about the prefix list.

Type: [ManagedPrefixList](api-managedprefixlist.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example modifies the specified managed prefix list by adding another
entry.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyManagedPrefixList
&PrefixListId=pl-0123123123123aabb
&CurrentVersion=1
&AddEntry.1.Cidr=10.1.0.0/16
&AddEntry.1.Description=Miami office
&AUTHPARAMS
```

#### Sample Response

```

<ModifyManagedPrefixListResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>602f3752-c348-4b14-81e2-example</requestId>
    <prefixList>
        <addressFamily>IPv4</addressFamily>
        <maxEntries>10</maxEntries>
        <ownerId>123456789012</ownerId>
        <prefixListArn>arn:aws:ec2:us-east-1:123456789012:prefix-list/pl-0123123123123aabb</prefixListArn>
        <prefixListId>pl-0123123123123aabb</prefixListId>
        <prefixListName>tgw-attachments</prefixListName>
        <state>modify-in-progress</state>
        <version>1</version>
    </prefixList>
</ModifyManagedPrefixListResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifymanagedprefixlist.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifymanagedprefixlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyLocalGatewayRoute

ModifyNetworkInterfaceAttribute
