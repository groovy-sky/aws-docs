# UnassignPrivateIpAddresses

Unassigns the specified secondary private IP addresses or IPv4 Prefix Delegation
prefixes from a network interface.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Ipv4Prefix.N**

The IPv4 prefixes to unassign from the network interface.

Type: Array of strings

Required: No

**NetworkInterfaceId**

The ID of the network interface.

Type: String

Required: Yes

**PrivateIpAddress.N**

The secondary private IP addresses to unassign from the network interface. You can
specify this option multiple times to unassign more than one IP address.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

The following example unassigns two secondary private IP addresses from the
specified network interface.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=UnassignPrivateIpAddresses
&NetworkInterfaceId=eni-197d9972
&PrivateIpAddress.1=10.0.2.60
&PrivateIpAddress.2=10.0.2.65
&AUTHPARAMS
```

#### Sample Response

```

<UnassignPrivateIpAddresses xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <return>true</return>
</UnassignPrivateIpAddresses>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/UnassignPrivateIpAddresses)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/UnassignPrivateIpAddresses)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/unassignprivateipaddresses.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/UnassignPrivateIpAddresses)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/unassignprivateipaddresses.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

UnassignIpv6Addresses

UnassignPrivateNatGatewayAddress
