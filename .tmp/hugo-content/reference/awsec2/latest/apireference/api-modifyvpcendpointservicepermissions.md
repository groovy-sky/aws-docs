# ModifyVpcEndpointServicePermissions

Modifies the permissions for your VPC endpoint service. You can add or remove permissions
for service consumers (AWS accounts, users, and IAM roles) to connect to
your endpoint service. Principal ARNs with path components aren't supported.

If you grant permissions to all principals, the service is public. Any users who know the name of a
public service can send a request to attach an endpoint. If the service does not require manual approval,
attachments are automatically approved.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AddAllowedPrincipals.N**

The Amazon Resource Names (ARN) of the principals.
Permissions are granted to the principals in this list.
To grant permissions to all principals, specify an asterisk (\*).

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**RemoveAllowedPrincipals.N**

The Amazon Resource Names (ARN) of the principals.
Permissions are revoked for principals in this list.

Type: Array of strings

Required: No

**ServiceId**

The ID of the service.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**addedPrincipalSet**

Information about the added principals.

Type: Array of [AddedPrincipal](api-addedprincipal.md) objects

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example permits all principals in AWS account
`123456789012` to connect
to your endpoint service `vpce-svc-03d5ebb7d9579a2b3`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVpcEndpointServicePermissions
&ServiceId=vpce-svc-03d5ebb7d9579a2b3
&AddAllowedPrincipals.1=arn:aws:iam::123456789012:root
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVpcEndpointServicePermissionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>08d80840-f750-42db-a6f8-2cd32example</requestId>
    <return>true</return>
</ModifyVpcEndpointServicePermissionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpcendpointservicepermissions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpcEndpointServicePayerResponsibility

ModifyVpcPeeringConnectionOptions
