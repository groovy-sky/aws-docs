# ModifyVpcEncryptionControl

Modifies the encryption control configuration for a VPC. You can update the encryption mode and exclusion settings for various gateway types and peering connections.

For more information, see [Enforce VPC encryption in transit](../../../../services/vpc/latest/userguide/vpc-encryption-controls.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**EgressOnlyInternetGatewayExclusion**

Specifies whether to exclude egress-only internet gateway traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**ElasticFileSystemExclusion**

Specifies whether to exclude Elastic File System traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**InternetGatewayExclusion**

Specifies whether to exclude internet gateway traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**LambdaExclusion**

Specifies whether to exclude Lambda function traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**Mode**

The encryption mode for the VPC Encryption Control configuration.

Type: String

Valid Values: `monitor | enforce`

Required: No

**NatGatewayExclusion**

Specifies whether to exclude NAT gateway traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**VirtualPrivateGatewayExclusion**

Specifies whether to exclude virtual private gateway traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**VpcEncryptionControlId**

The ID of the VPC Encryption Control resource to modify.

Type: String

Required: Yes

**VpcLatticeExclusion**

Specifies whether to exclude VPC Lattice traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

**VpcPeeringExclusion**

Specifies whether to exclude VPC peering connection traffic from encryption enforcement.

Type: String

Valid Values: `enable | disable`

Required: No

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**vpcEncryptionControl**

Information about the VPC Encryption Control configuration.

Type: [VpcEncryptionControl](api-vpcencryptioncontrol.md) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyVpcEncryptionControl)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyVpcEncryptionControl)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyVpcEncryptionControl)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvpcencryptioncontrol.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVpcBlockPublicAccessOptions

ModifyVpcEndpoint
