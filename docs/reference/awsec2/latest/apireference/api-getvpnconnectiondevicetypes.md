# GetVpnConnectionDeviceTypes

Obtain a list of customer gateway devices for which sample configuration
files can be provided. The request has no additional parameters. You can also see the
list of device types with sample configuration files available under [Your customer gateway\
device](../../../../services/vpn/latest/s2svpn/your-cgw.md) in the _AWS Site-to-Site VPN User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**MaxResults**

The maximum number of results returned by `GetVpnConnectionDeviceTypes` in
paginated output. When this parameter is used, `GetVpnConnectionDeviceTypes`
only returns `MaxResults` results in a single page along with a
`NextToken` response element. The remaining results of the initial
request can be seen by sending another `GetVpnConnectionDeviceTypes` request
with the returned `NextToken` value. This value can be between 200 and 1000.
If this parameter is not used, then `GetVpnConnectionDeviceTypes` returns all
results.

Type: Integer

Valid Range: Minimum value of 200. Maximum value of 1000.

Required: No

**NextToken**

The `NextToken` value returned from a previous paginated
`GetVpnConnectionDeviceTypes` request where `MaxResults` was
used and the results exceeded the value of that parameter. Pagination continues from the
end of the previous results that returned the `NextToken` value. This value
is null when there are no more results to return.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**nextToken**

The `NextToken` value to include in a future
`GetVpnConnectionDeviceTypes` request. When the results of a
`GetVpnConnectionDeviceTypes` request exceed `MaxResults`,
this value can be used to retrieve the next page of results. This value is null when
there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

**vpnConnectionDeviceTypeSet**

List of customer gateway devices that have a sample configuration file available for
use.

Type: Array of [VpnConnectionDeviceType](api-vpnconnectiondevicetype.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetVpnConnectionDeviceTypes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetVpnConnectionDeviceTypes)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetVpnConnectionDeviceTypes)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/getvpnconnectiondevicetypes.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetVpnConnectionDeviceSampleConfiguration

GetVpnTunnelReplacementStatus
