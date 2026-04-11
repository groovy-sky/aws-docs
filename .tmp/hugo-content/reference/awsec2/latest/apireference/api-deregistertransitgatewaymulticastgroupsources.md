# DeregisterTransitGatewayMulticastGroupSources

Deregisters the specified sources (network interfaces) from the transit gateway multicast group.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**GroupIpAddress**

The IP address assigned to the transit gateway multicast group.

Type: String

Required: No

**NetworkInterfaceIds.N**

The IDs of the group sources' network interfaces.

Type: Array of strings

Required: No

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**deregisteredMulticastGroupSources**

Information about the deregistered group sources.

Type: [TransitGatewayMulticastDeregisteredGroupSources](api-transitgatewaymulticastderegisteredgroupsources.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example deregisters the network interface as a group source
`eni-07f290fc3cEXAMPLE` from the multicast domain
`tgw-mcast-domain-0c4905cef7EXAMPLE`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeregisterTransitGatewayMulticastGroupSources
&TransitGatewayMulticastDomainId=tgw-mcast-domain-0c4905cef7EXAMPLE
&NetworkInterfaceIds=eni-07f290fc3cEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DeregisterTransitGatewayMulticastGroupSourcesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>1ca916e8-a4b5-4ff8-9fc3-3052dEXAMPLE</requestId>
    <deregisteredMulticastGroupSources>
        <deregisteredNetworkInterfaceIds>
            <item>eni-07f290fc3cEXAMPLE</item>
        </deregisteredNetworkInterfaceIds>
        <groupIpAddress>224.0.1.0</groupIpAddress>
        <transitGatewayMulticastDomainId>tgw-mcast-domain-0c4905cef7EXAMPLE</transitGatewayMulticastDomainId>
    </deregisteredMulticastGroupSources>
</DeregisterTransitGatewayMulticastGroupSourcesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deregistertransitgatewaymulticastgroupsources.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeregisterTransitGatewayMulticastGroupMembers

DescribeAccountAttributes
