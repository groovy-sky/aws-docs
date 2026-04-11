# IpamDiscoveryFailureReason

The discovery failure reason.

## Contents

**code**

The discovery failure code.

- `assume-role-failure` \- IPAM could not assume the AWS IAM service-linked role. This could be because of any of the following:

- SLR has not been created yet and IPAM is still creating it.

- You have opted-out of the IPAM home Region.

- Account you are using as your IPAM account has been suspended.

- `throttling-failure` \- IPAM account is already using the allotted transactions per second and IPAM is receiving a throttling error when assuming the AWS IAM SLR.

- `unauthorized-failure` \- AWS account making the request is not authorized. For more information, see [AuthFailure](errors-overview.md) in the _Amazon Elastic Compute Cloud API Reference_.

Type: String

Valid Values: `assume-role-failure | throttling-failure | unauthorized-failure`

Required: No

**message**

The discovery failure message.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipamdiscoveryfailurereason.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipamdiscoveryfailurereason.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipamdiscoveryfailurereason.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamDiscoveredResourceCidr

IpamExternalResourceVerificationToken
