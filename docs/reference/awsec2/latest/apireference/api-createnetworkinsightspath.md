# CreateNetworkInsightsPath

Creates a path to analyze for reachability.

Reachability Analyzer enables you to analyze and debug network reachability between
two resources in your virtual private cloud (VPC). For more information, see the
[Reachability Analyzer Guide](../../../../services/vpc/latest/reachability.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information,
see [How to ensure idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: Yes

**Destination**

The ID or ARN of the destination. If the resource is in another account, you must specify an ARN.

Type: String

Required: No

**DestinationIp**

The IP address of the destination.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**DestinationPort**

The destination port.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 65535.

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**FilterAtDestination**

Scopes the analysis to network paths that match specific filters at the destination. If you specify
this parameter, you can't specify the parameter for the destination IP address.

Type: [PathRequestFilter](api-pathrequestfilter.md) object

Required: No

**FilterAtSource**

Scopes the analysis to network paths that match specific filters at the source. If you specify
this parameter, you can't specify the parameters for the source IP address or the destination port.

Type: [PathRequestFilter](api-pathrequestfilter.md) object

Required: No

**Protocol**

The protocol.

Type: String

Valid Values: `tcp | udp`

Required: Yes

**Source**

The ID or ARN of the source. If the resource is in another account, you must specify an ARN.

Type: String

Required: Yes

**SourceIp**

The IP address of the source.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**TagSpecification.N**

The tags to add to the path.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**networkInsightsPath**

Information about the path.

Type: [NetworkInsightsPath](api-networkinsightspath.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateNetworkInsightsPath)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateNetworkInsightsPath)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createnetworkinsightspath.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateNetworkInsightsPath)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createnetworkinsightspath.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateNetworkInsightsAccessScope

CreateNetworkInterface
