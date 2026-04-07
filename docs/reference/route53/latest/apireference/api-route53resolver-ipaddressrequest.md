# IpAddressRequest

In a
[CreateResolverEndpoint](api-route53resolver-createresolverendpoint.md)
request, the IP address that DNS queries originate from (for outbound endpoints) or that you forward DNS queries to (for inbound endpoints).
`IpAddressRequest` also includes the ID of the subnet that contains the IP address.

## Contents

**SubnetId**

The ID of the subnet that contains the IP address.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 32.

Required: Yes

**Ip**

The IPv4 address that you want to use for DNS queries.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 36.

Required: No

**Ipv6**

The IPv6 address that you want to use for DNS queries.

Type: String

Length Constraints: Minimum length of 7. Maximum length of 39.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53resolver-2018-04-01/IpAddressRequest)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53resolver-2018-04-01/IpAddressRequest)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53resolver-2018-04-01/IpAddressRequest)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FirewallRuleGroupMetadata

IpAddressResponse
