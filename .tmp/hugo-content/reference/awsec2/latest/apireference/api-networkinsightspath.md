# NetworkInsightsPath

Describes a path.

## Contents

**createdDate**

The time stamp when the path was created.

Type: Timestamp

Required: No

**destination**

The ID of the destination.

Type: String

Required: No

**destinationArn**

The Amazon Resource Name (ARN) of the destination.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**destinationIp**

The IP address of the destination.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**destinationPort**

The destination port.

Type: Integer

Required: No

**filterAtDestination**

Scopes the analysis to network paths that match specific filters at the destination.

Type: [PathFilter](api-pathfilter.md) object

Required: No

**filterAtSource**

Scopes the analysis to network paths that match specific filters at the source.

Type: [PathFilter](api-pathfilter.md) object

Required: No

**networkInsightsPathArn**

The Amazon Resource Name (ARN) of the path.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**networkInsightsPathId**

The ID of the path.

Type: String

Required: No

**protocol**

The protocol.

Type: String

Valid Values: `tcp | udp`

Required: No

**source**

The ID of the source.

Type: String

Required: No

**sourceArn**

The Amazon Resource Name (ARN) of the source.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1283.

Required: No

**sourceIp**

The IP address of the source.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**TagSet.N**

The tags associated with the path.

Type: Array of [Tag](api-tag.md) objects

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/networkinsightspath.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/networkinsightspath.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/networkinsightspath.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

NetworkInsightsAnalysis

NetworkInterface
