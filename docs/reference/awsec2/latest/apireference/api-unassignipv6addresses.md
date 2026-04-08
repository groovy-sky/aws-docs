# UnassignIpv6Addresses

Unassigns the specified IPv6 addresses or Prefix Delegation prefixes from a network
interface.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Ipv6Addresses.N**

The IPv6 addresses to unassign from the network interface.

Type: Array of strings

Required: No

**Ipv6Prefix.N**

The IPv6 prefixes to unassign from the network interface.

Type: Array of strings

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**networkInterfaceId**

The ID of the network interface.

Type: String

**requestId**

The ID of the request.

Type: String

**unassignedIpv6Addresses**

The IPv6 addresses that have been unassigned from the network interface.

Type: Array of strings

**unassignedIpv6PrefixSet**

The IPv6 prefixes that have been unassigned from the network interface.

Type: Array of strings

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

The following example unassigns two IPv6 addresses from the specified network
interface.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UnassignIpv6Addresses
&NetworkInterfaceId=eni-197d9972
&Ipv6Addresses.1=2001:db8:1234:1a00::123
&Ipv6Addresses.2=2001:db8:1234:1a00::456
&AUTHPARAMS
```

#### Sample Response

```

<UnassignIpv6AddressesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>94d446d7-fc8e-4918-94f9-example</requestId>
    <networkInterfaceId>eni-197d9972</networkInterfaceId>
    <unassignedIpv6Addresses>
        <item>2001:db8:1234:1a00::123</item>
        <item>2001:db8:1234:1a00::456</item>
    </unassignedIpv6Addresses>
</UnassignIpv6AddressesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/unassignipv6addresses.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/unassignipv6addresses.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

TerminateInstances

UnassignPrivateIpAddresses
