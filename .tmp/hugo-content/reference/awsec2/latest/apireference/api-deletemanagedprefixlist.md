# DeleteManagedPrefixList

Deletes the specified managed prefix list. You must first remove all references to the prefix list in your resources.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**PrefixListId**

The ID of the prefix list.

Type: String

Required: Yes

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

This example deletes the specified prefix list.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteManagedPrefixList
&PrefixListId=pl-0123123123123aabb
&AUTHPARAMS
```

#### Sample Response

```

<DeleteManagedPrefixListResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>06152571-575a-49aa-af95-example</requestId>
    <prefixList>
        <addressFamily>IPv6</addressFamily>
        <maxEntries>25</maxEntries>
        <ownerId>123456789012</ownerId>
        <prefixListArn>arn:aws:ec2:us-east-1:123456789012:prefix-list/pl-0123123123123aabb</prefixListArn>
        <prefixListId>pl-0123123123123aabb</prefixListId>
        <prefixListName>test-pl</prefixListName>
        <state>delete-in-progress</state>
        <version>1</version>
    </prefixList>
</DeleteManagedPrefixListResponse>'

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletemanagedprefixlist.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletemanagedprefixlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteLocalGatewayVirtualInterfaceGroup

DeleteNatGateway
