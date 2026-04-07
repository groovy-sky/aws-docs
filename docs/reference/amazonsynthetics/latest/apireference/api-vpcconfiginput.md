# VpcConfigInput

If this canary is to test an endpoint in a VPC, this structure contains
information about the subnets and security groups of the VPC endpoint.
For more information, see [Running a Canary in a VPC](../../../../services/amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-vpc.md).

## Contents

**Ipv6AllowedForDualStack**

Set this to `true` to allow outbound IPv6 traffic on VPC canaries that are connected to dual-stack subnets. The default is `false`

Type: Boolean

Required: No

**SecurityGroupIds**

The IDs of the security groups for this canary.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 5 items.

Required: No

**SubnetIds**

The IDs of the subnets where this canary is to run.

Type: Array of strings

Array Members: Minimum number of 0 items. Maximum number of 16 items.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/synthetics-2017-10-11/VpcConfigInput)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/synthetics-2017-10-11/VpcConfigInput)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/synthetics-2017-10-11/VpcConfigInput)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

VisualReferenceOutput

VpcConfigOutput
