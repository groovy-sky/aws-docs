# AuthorizeClientVpnIngress

Adds an ingress authorization rule to a Client VPN endpoint. Ingress authorization rules act as
firewall rules that grant access to networks. You must configure ingress authorization rules to
enable clients to access resources in AWS or on-premises networks.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AccessGroupId**

The ID of the group to grant access to, for example, the Active Directory group or identity provider (IdP) group. Required if `AuthorizeAllGroups` is `false` or not specified.

Type: String

Required: No

**AuthorizeAllGroups**

Indicates whether to grant access to all clients. Specify `true` to grant all
clients who successfully establish a VPN connection access to the network. Must be set
to `true` if `AccessGroupId` is not specified.

Type: Boolean

Required: No

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**ClientVpnEndpointId**

The ID of the Client VPN endpoint.

Type: String

Required: Yes

**Description**

A brief description of the authorization rule.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request, and provides an error response. If you have the required permissions, the error response is `DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TargetNetworkCidr**

The IPv4 address range, in CIDR notation, of the network for which access is being authorized.

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

This example adds an authorization rule to a subnet and grants access to all users.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=AuthorizeClientVpnIngress
&ClientVpnEndpointId=cvpn-endpoint-00c5d11fc4EXAMPLE
&TargetNetworkCidr=10.0.0.0/16
&AuthorizeAllGroups=true
&AUTHPARAMS
```

#### Sample Response

```

<AuthorizeClientVpnIngressResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>afafad8c-274c-4584-bbd1-75a21EXAMPLE</requestId>
    <status>
        <code>authorizing</code>
    </status>
</AuthorizeClientVpnIngressResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/AuthorizeClientVpnIngress)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/AuthorizeClientVpnIngress)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/authorizeclientvpningress.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/AuthorizeClientVpnIngress)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/authorizeclientvpningress.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

AttachVpnGateway

AuthorizeSecurityGroupEgress
