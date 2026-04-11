# CreateVpcEncryptionControl

Creates a VPC Encryption Control configuration for a specified VPC. VPC Encryption Control enables you to enforce encryption for all data in transit within and between VPCs to meet compliance requirements for standards like HIPAA, FedRAMP, and PCI DSS.

For more information, see [Enforce VPC encryption in transit](../../../../services/vpc/latest/userguide/vpc-encryption-controls.md) in the _Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**TagSpecification.N**

The tags to apply to the VPC Encryption Control resource.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**VpcId**

The ID of the VPC for which to create the encryption control configuration.

Type: String

Required: Yes

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/createvpcencryptioncontrol.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/createvpcencryptioncontrol.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateVpcBlockPublicAccessExclusion

CreateVpcEndpoint
