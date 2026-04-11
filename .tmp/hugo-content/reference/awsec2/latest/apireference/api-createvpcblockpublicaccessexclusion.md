# CreateVpcBlockPublicAccessExclusion

Create a VPC Block Public Access (BPA) exclusion. A VPC BPA exclusion is a mode that can be applied to a single VPC or subnet that exempts it from the account’s BPA mode and will allow bidirectional or egress-only access. You can create BPA exclusions for VPCs and subnets even when BPA is not enabled on the account to ensure that there is no traffic disruption to the exclusions when VPC BPA is turned on. To learn more about VPC BPA, see [Block public access to VPCs and subnets](../../../../services/vpc/latest/userguide/security-vpc-bpa.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InternetGatewayExclusionMode**

The exclusion mode for internet gateway traffic.

- `allow-bidirectional`: Allow all internet traffic to and from the excluded VPCs and subnets.

- `allow-egress`: Allow outbound internet traffic from the excluded VPCs and subnets. Block inbound internet traffic to the excluded VPCs and subnets. Only applies when VPC Block Public Access is set to Bidirectional.

Type: String

Valid Values: `allow-bidirectional | allow-egress`

Required: Yes

**SubnetId**

A subnet ID.

Type: String

Required: No

**TagSpecification.N**

`tag` \- The key/value combination of a tag assigned to the resource. Use the tag key in the filter name and the tag value as the filter value.
For example, to find all resources that have a tag with the key `Owner` and the value `TeamA`, specify `tag:Owner` for the filter name and `TeamA` for the filter value.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VpcId**

A VPC ID.

Type: String

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpcBlockPublicAccessExclusion**

Details about an exclusion.

Type: [VpcBlockPublicAccessExclusion](api-vpcblockpublicaccessexclusion.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createvpcblockpublicaccessexclusion.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVpc

CreateVpcEncryptionControl
