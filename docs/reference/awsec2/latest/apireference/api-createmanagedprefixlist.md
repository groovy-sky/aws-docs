# CreateManagedPrefixList

Creates a managed prefix list. You can specify entries for the prefix list.
Each entry consists of a CIDR block and an optional description.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AddressFamily**

The IP address type.

Valid Values: `IPv4` \| `IPv6`

Type: String

Required: Yes

**ClientToken**

Unique, case-sensitive identifier you provide to ensure the idempotency of the
request. For more information, see [Ensuring\
idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Constraints: Up to 255 UTF-8 characters in length.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Entry.N**

One or more entries for the prefix list.

Type: Array of [AddPrefixListEntry](api-addprefixlistentry.md) objects

Array Members: Minimum number of 0 items. Maximum number of 100 items.

Required: No

**MaxEntries**

The maximum number of entries for the prefix list.

Type: Integer

Required: Yes

**PrefixListName**

A name for the prefix list.

Constraints: Up to 255 characters in length. The name cannot start with `com.amazonaws`.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the prefix list during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

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

This example creates a managed prefix list with a maximum of 10 entries, and
adds 2 entries. The prefix list support IPv4 CIDR blocks.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateManagedPrefixList
&PrefixListName=tgw-attachments
&Entry.1.Cidr=10.0.0.0/16
&Entry.1.Description=vpc-a
&Entry.2.Cidr=10.2.0.0/16
&Entry.2.Description=vpc-b
&MaxEntries=10
&AddressFamily=IPv4
&AUTHPARAMS
```

#### Sample Response

```

<CreateManagedPrefixListResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>2989de6e-5305-49c7-809a-example</requestId>
    <prefixList>
            <addressFamily>IPv4</addressFamily>
            <maxEntries>10</maxEntries>
            <ownerId>123456789012</ownerId>
            <prefixListArn>arn:aws:ec2:us-east-1:123456789012:prefix-list/pl-0123123123123abcd</prefixListArn>
            <prefixListId>pl-0123123123123abcd</prefixListId>
            <prefixListName>tgw-attachments</prefixListName>
            <state>create-in-progress</state>
            <tagSet/>
            <version>1</version>
    </prefixList>
</CreateManagedPrefixListResponse>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateManagedPrefixList)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateManagedPrefixList)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createmanagedprefixlist.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateManagedPrefixList)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createmanagedprefixlist.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateMacSystemIntegrityProtectionModificationTask

CreateNatGateway
