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

- [AWS SDK for C++](../../../goto/sdkforcpp/route53-2013-04-01/cidrroutingconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/route53-2013-04-01/cidrroutingconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/route53-2013-04-01/cidrroutingconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CidrCollectionChange

CloudWatchAlarmConfiguration
