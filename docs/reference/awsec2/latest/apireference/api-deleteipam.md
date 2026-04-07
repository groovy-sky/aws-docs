# DeleteIpam

Delete an IPAM. Deleting an IPAM removes all monitored data associated with the IPAM including the historical data for CIDRs.

For more information, see [Delete an IPAM](../../../../services/vpc/latest/ipam/delete-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Cascade**

Enables you to quickly delete an IPAM, private scopes, pools in private scopes, and
any allocations in the pools in private scopes. You cannot delete the IPAM with this option if there is a pool in your public scope. If you use this option, IPAM does the following:

- Deallocates any CIDRs allocated to VPC resources (such as VPCs) in pools in private scopes.

###### Note

No VPC resources are deleted as a result of enabling this option. The CIDR associated with the resource will no longer be allocated from an IPAM pool, but the CIDR itself will remain unchanged.

- Deprovisions all IPv4 CIDRs provisioned to IPAM pools in private scopes.

- Deletes all IPAM pools in private scopes.

- Deletes all non-default private scopes in the IPAM.

- Deletes the default public and private scopes and the IPAM.

Type: Boolean

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**IpamId**

The ID of the IPAM to delete.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipam**

Information about the results of the deletion.

Type: [Ipam](api-ipam.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DeleteIpam)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DeleteIpam)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deleteipam.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DeleteIpam)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deleteipam.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteInternetGateway

DeleteIpamExternalResourceVerificationToken
