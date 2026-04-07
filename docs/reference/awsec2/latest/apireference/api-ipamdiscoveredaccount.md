# IpamDiscoveredAccount

An IPAM discovered account. A discovered account is an AWS account that is monitored under a resource discovery. If you have integrated IPAM with AWS Organizations, all accounts in the organization are discovered accounts.

## Contents

**accountId**

The account ID.

Type: String

Required: No

**discoveryRegion**

The AWS Region that the account information is returned from.
An account can be discovered in multiple regions and will have a separate discovered account for each Region.

Type: String

Required: No

**failureReason**

The resource discovery failure reason.

Type: [IpamDiscoveryFailureReason](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_IpamDiscoveryFailureReason.html) object

Required: No

**lastAttemptedDiscoveryTime**

The last attempted resource discovery time.

Type: Timestamp

Required: No

**lastSuccessfulDiscoveryTime**

The last successful resource discovery time.

Type: Timestamp

Required: No

**organizationalUnitId**

The ID of an Organizational Unit in AWS Organizations.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/IpamDiscoveredAccount)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/IpamDiscoveredAccount)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/IpamDiscoveredAccount)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

IpamCidrAuthorizationContext

IpamDiscoveredPublicAddress
