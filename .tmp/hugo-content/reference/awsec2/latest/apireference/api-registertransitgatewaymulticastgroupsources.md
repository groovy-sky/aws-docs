# RegisterTransitGatewayMulticastGroupSources

Registers sources (network interfaces) with the specified transit gateway multicast group.

A multicast source is a network interface attached to a supported instance that sends
multicast traffic. For more information about supported instances, see [Multicast\
on transit gateways](../../../../services/vpc/latest/tgw/tgw-multicast-overview.md) in the _AWS Transit Gateways Guide_.

After you add the source, use [SearchTransitGatewayMulticastGroups](api-searchtransitgatewaymulticastgroups.md) to verify that the source was added to the multicast
group.

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

The group sources' network interface IDs to register with the transit gateway multicast group.

Type: Array of strings

Required: Yes

**TransitGatewayMulticastDomainId**

The ID of the transit gateway multicast domain.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**registeredMulticastGroupSources**

Information about the transit gateway multicast group sources.

Type: [TransitGatewayMulticastRegisteredGroupSources](api-transitgatewaymulticastregisteredgroupsources.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example registers the network interface as a group source
`eni-07f290fc3cEXAMPLE` with the multicast domain
`tgw-mcast-domain-0c4905cef7EXAMPLE`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RegisterTransitGatewayMulticastGroupSources
&TransitGatewayMulticastDomainId=tgw-mcast-domain-0c4905cef7EXAMPLE
&NetworkInterfaceIds=eni-07f290fc3cEXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<RegisterTransitGatewayMulticastGroupSourcesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>b66c84ed-eb8a-4e6d-8d79-2347fEXAMPLE</requestId>
    <registeredMulticastGroupSources>
        <groupIpAddress>224.0.1.0</groupIpAddress>
        <registeredNetworkInterfaceIds>
            <item>eni-07f290fc3cEXAMPLE</item>
        </registeredNetworkInterfaceIds>
        <transitGatewayMulticastDomainId>tgw-mcast-domain-0c4905cef7EXAMPLE</transitGatewayMulticastDomainId>
    </registeredMulticastGroupSources>
</RegisterTransitGatewayMulticastGroupSourcesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/registertransitgatewaymulticastgroupsources.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RegisterTransitGatewayMulticastGroupMembers

RejectCapacityReservationBillingOwnership
