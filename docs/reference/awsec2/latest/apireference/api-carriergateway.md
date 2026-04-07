# CarrierGateway

Describes a carrier gateway.

## Contents

**carrierGatewayId**

The ID of the carrier gateway.

Type: String

Required: No

**ownerId**

The AWS account ID of the owner of the carrier gateway.

Type: String

Required: No

**state**

The state of the carrier gateway.

Type: String

Valid Values: `pending | available | deleting | deleted`

Required: No

**TagSet.N**

The tags assigned to the carrier gateway.

Type: Array of [Tag](api-tag.md) objects

Required: No

**vpcId**

The ID of the VPC associated with the carrier gateway.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CarrierGateway)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CarrierGateway)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CarrierGateway)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CapacityReservationTopology

CertificateAuthentication
