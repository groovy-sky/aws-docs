# PathRequestFilter

Describes a set of filters for a path analysis. Use path filters to scope the analysis when
there can be multiple resulting paths.

## Contents

**DestinationAddress**

The destination IPv4 address.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**DestinationPortRange**

The destination port range.

Type: [RequestFilterPortRange](api-requestfilterportrange.md) object

Required: No

**SourceAddress**

The source IPv4 address.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 15.

Pattern: `^([0-9]{1,3}.){3}[0-9]{1,3}$`

Required: No

**SourcePortRange**

The source port range.

Type: [RequestFilterPortRange](api-requestfilterportrange.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/pathrequestfilter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/pathrequestfilter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/pathrequestfilter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

PathFilter

PathStatement
