# CidrRoutingConfig

The object that is specified in resource record set object when you are linking a
resource record set to a CIDR location.

A `LocationName` with an asterisk “\*” can be used to create a default CIDR
record. `CollectionId` is still required for default record.

## Contents

**CollectionId**

The CIDR collection ID.

Type: String

Pattern: `[0-9a-f]{8}-(?:[0-9a-f]{4}-){3}[0-9a-f]{12}`

Required: Yes

**LocationName**

The CIDR collection location name.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16.

Pattern: `[0-9A-Za-z_\-\*]+`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/route53-2013-04-01/CidrRoutingConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/route53-2013-04-01/CidrRoutingConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/route53-2013-04-01/CidrRoutingConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CidrCollectionChange

CloudWatchAlarmConfiguration
