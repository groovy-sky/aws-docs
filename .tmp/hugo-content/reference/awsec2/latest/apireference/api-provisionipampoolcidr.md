# ProvisionIpamPoolCidr

Provision a CIDR to an IPAM pool. You can use this action to provision new CIDRs to a top-level pool or to transfer a CIDR from a top-level pool to a pool within it.

For more information, see [Provision CIDRs to pools](../../../../services/vpc/latest/ipam/prov-cidr-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cidr**

The CIDR you want to assign to the IPAM pool. Either "NetmaskLength" or "Cidr" is required. This value will be null if you specify "NetmaskLength" and will be filled in during the provisioning process.

Type: String

Required: No

**CidrAuthorizationContext**

A signed document that proves that you are authorized to bring a specified IP address range to Amazon using BYOIP. This option only applies to IPv4 and IPv6 pools in the public scope.

Type: [IpamCidrAuthorizationContext](api-ipamcidrauthorizationcontext.md) object

Required: No

**ClientToken**

A unique, case-sensitive identifier that you provide to ensure the idempotency of the request. For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamExternalResourceVerificationTokenId**

Verification token ID. This option only applies to IPv4 and IPv6 pools in the public scope.

Type: String

Required: No

**IpamPoolId**

The ID of the IPAM pool to which you want to assign a CIDR.

Type: String

Required: Yes

**NetmaskLength**

The netmask length of the CIDR you'd like to provision to a pool. Can be used for provisioning Amazon-provided IPv6 CIDRs to top-level pools and for provisioning CIDRs to pools with source pools. Cannot be used to provision BYOIP CIDRs to top-level pools. Either "NetmaskLength" or "Cidr" is required.

Type: Integer

Required: No

**VerificationMethod**

The method for verifying control of a public IP address range. Defaults to `remarks-x509` if not specified. This option only applies to IPv4 and IPv6 pools in the public scope.

Type: String

Valid Values: `remarks-x509 | dns-token`

Required: No

## Response Elements

The following elements are returned by the service.

**ipamPoolCidr**

Information about the provisioned CIDR.

Type: [IpamPoolCidr](api-ipampoolcidr.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/provisionipampoolcidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/provisionipampoolcidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ProvisionIpamByoasn

ProvisionPublicIpv4PoolCidr
