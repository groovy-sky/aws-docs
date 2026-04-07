# CoipPool

Describes a customer-owned address pool.

## Contents

**localGatewayRouteTableId**

The ID of the local gateway route table.

Type: String

Required: No

**poolArn**

The ARN of the address pool.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**PoolCidrSet.N**

The address ranges of the address pool.

Type: Array of strings

Required: No

**poolId**

The ID of the address pool.

Type: String

Required: No

**TagSet.N**

The tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CoipPool)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CoipPool)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CoipPool)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CoipCidr

ConnectionLogOptions
