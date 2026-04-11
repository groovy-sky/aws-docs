# DescribeVpcEndpointServicePermissions

Describes the principals (service consumers) that are permitted to discover your VPC
endpoint service. Principal ARNs with path components aren't supported.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `principal` \- The ARN of the principal.

- `principal-type` \- The principal type ( `All` \|
`Service` \| `OrganizationUnit` \| `Account`
\| `User` \| `Role`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining
results of the initial request can be seen by sending another request with the returned
`NextToken` value. This value can be between 5 and 1,000; if
`MaxResults` is given a value larger than 1,000, only 1,000 results are
returned.

Type: Integer

Required: No

**NextToken**

The token to retrieve the next page of results.

Type: String

Required: No

**ServiceId**

The ID of the service.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**allowedPrincipals**

Information about the allowed principals.

Type: Array of [AllowedPrincipal](api-allowedprincipal.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null` when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the permissions for service
`vpce-svc-03d5ebb7d9579a123`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVpcEndpointServicePermissions
&ServiceId=vpce-svc-03d5ebb7d9579a123
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVpcEndpointServicePermissionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>5359c8a3-9151-4964-abed-b4422example</requestId>
    <allowedPrincipals>
        <item>
            <principal>arn:aws:iam::123456789012:root</principal>
            <principalType>Account</principalType>
        </item>
    </allowedPrincipals>
</DescribeVpcEndpointServicePermissionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describevpcendpointservicepermissions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describevpcendpointservicepermissions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeVpcEndpointServiceConfigurations

DescribeVpcEndpointServices
