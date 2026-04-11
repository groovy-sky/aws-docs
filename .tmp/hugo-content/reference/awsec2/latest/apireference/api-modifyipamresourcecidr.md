# ModifyIpamResourceCidr

Modify a resource CIDR. You can use this action to transfer resource CIDRs between scopes and ignore resource CIDRs that you do not want to manage. If set to false, the resource will not be tracked for overlap, it cannot be auto-imported into a pool, and it will be removed from any pool it has an allocation in.

For more information, see [Move resource CIDRs between scopes](../../../../services/vpc/latest/ipam/move-resource-ipam.md) and [Change the monitoring state of resource CIDRs](../../../../services/vpc/latest/ipam/change-monitoring-state-ipam.md) in the _Amazon VPC IPAM User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CurrentIpamScopeId**

The ID of the current scope that the resource CIDR is in.

Type: String

Required: Yes

**DestinationIpamScopeId**

The ID of the scope you want to transfer the resource CIDR to.

Type: String

Required: No

**DryRun**

A check for whether you have the required permissions for the action without actually making the request
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Monitored**

Determines if the resource is monitored by IPAM. If a resource is monitored, the resource is discovered by IPAM and you can view details about the resource’s CIDR.

Type: Boolean

Required: Yes

**ResourceCidr**

The CIDR of the resource you want to modify.

Type: String

Required: Yes

**ResourceId**

The ID of the resource you want to modify.

Type: String

Required: Yes

**ResourceRegion**

The AWS Region of the resource you want to modify.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**ipamResourceCidr**

The CIDR of the resource.

Type: [IpamResourceCidr](api-ipamresourcecidr.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyipamresourcecidr.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyipamresourcecidr.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyIpamPrefixListResolverTarget

ModifyIpamResourceDiscovery
