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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UnassignIpv6Addresses)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/UnassignIpv6Addresses)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TerminateInstances

UnassignPrivateIpAddresses
