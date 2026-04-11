# DescribeClientVpnAuthorizationRules

Describes the authorization rules for a specified Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters. Filter names and values are case-sensitive.

- `description` \- The description of the authorization rule.

- `destination-cidr` \- The CIDR of the network to which the authorization rule
applies.

- `group-id` \- The ID of the Active Directory group to which the authorization rule grants access.

Type: Array of [Filter](api-filter.md) objects

Required: No

**MaxResults**

The maximum number of results to return for the request in a single page. The remaining results can be seen by sending another request with the nextToken value.

Type: Integer

Valid Range: Minimum value of 5. Maximum value of 1000.

Required: No

**NextToken**

The token to retrieve the next page of results.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**authorizationRule**

Information about the authorization rules.

Type: Array of [AuthorizationRule](api-authorizationrule.md) objects

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

This example describes the authorization rules for a specific Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeClientVpnAuthorizationRules
&ClientVpnEndpointId.1=cvpn-endpoint-EXAMPLEc8db8d3536
&AUTHPARAMS
```

#### Sample Response

```

<DescribeClientVpnAuthorizationRulesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>9d5c69db-763e-4b63-88ee-EXAMPLE</requestId>
    <authorizationRule>
        <item>
            <accessAll>true</accessAll>
            <description>auth-rule-one</description>
            <destinationCidr>10.0.0.0/16</destinationCidr>
            <clientVpnEndpointId>cvpn-endpoint-EXAMPLEc8db8d3536</clientVpnEndpointId>
            <groupId/>
            <status>
                <code>active</code>
            </status>
        </item>
    </authorizationRule>
</DescribeClientVpnAuthorizationRulesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeclientvpnauthorizationrules.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeClassicLinkInstances

DescribeClientVpnConnections
