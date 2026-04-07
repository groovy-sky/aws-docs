# RevokeClientVpnIngress

Removes an ingress authorization rule from a Client VPN endpoint.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AccessGroupId**

The ID of the Active Directory group for which to revoke access.

Type: String

Required: No

**ClientVpnEndpointId**

The ID of the Client VPN endpoint with which the authorization rule is associated.

Type: String

Required: Yes

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RevokeAllGroups**

Indicates whether access should be revoked for all groups for a single `TargetNetworkCidr` that earlier authorized ingress for all groups using `AuthorizeAllGroups`.
This does not impact other authorization rules that allowed ingress to the same `TargetNetworkCidr` with a specific `AccessGroupId`.

Type: Boolean

Required: No

**TargetNetworkCidr**

The IPv4 address range, in CIDR notation, of the network for which access is being removed.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**status**

The current state of the authorization rule.

Type: [ClientVpnAuthorizationRuleStatus](api-clientvpnauthorizationrulestatus.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example removes an authorization rule from a Client VPN endpoint.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=RevokeClientVpnIngress
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&TargetNetworkCidr=10.0.0.0/16
&RevokeAllGroups=true
&AUTHPARAMS
```

#### Sample Response

```

<RevokeClientVpnIngressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>691de4ea-32ef-447b-b4f8-d8463XAMPLE</requestId>
    <status>
        <code>revoking</code>
    </status>
</RevokeClientVpnIngressResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/RevokeClientVpnIngress)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/RevokeClientVpnIngress)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/revokeclientvpningress.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/RevokeClientVpnIngress)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/revokeclientvpningress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RestoreVolumeFromRecycleBin

RevokeSecurityGroupEgress
