# IpamExternalResourceVerificationToken

A verification token is an AWS-generated random value that you can use to prove ownership of an external resource. For example, you can use a verification token to validate that you control a public IP address range when you bring an IP address range to AWS (BYOIP).

## Contents

**ipamArn**

ARN of the IPAM that created the token.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**ipamExternalResourceVerificationTokenArn**

Token ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**ipamExternalResourceVerificationTokenId**

The ID of the token.

Type: String

Required: No

**ipamId**

The ID of the IPAM that created the token.

Type: String

Required: No

**ipamRegion**

Region of the IPAM that created the token.

Type: String

Required: No

**notAfter**

Token expiration.

Type: Timestamp

Required: No

**state**

Token state.

Type: String

Valid Values: `create-in-progress | create-complete | create-failed | delete-in-progress | delete-complete | delete-failed`

Required: No

**status**

Token status.

Type: String

Valid Values: `valid | expired`

Required: No

**TagSet.N**

Token tags.

Type: Array of [Tag](api-tag.md) objects

Required: No

**tokenName**

Token name.

Type: String

Required: No

**tokenValue**

Token value.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/ipamexternalresourceverificationtoken.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/ipamexternalresourceverificationtoken.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/ipamexternalresourceverificationtoken.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

IpamDiscoveryFailureReason

IpamOperatingRegion
