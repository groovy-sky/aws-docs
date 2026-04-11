# DocumentAttributeValue

The value of a document attribute. You can only provide one value for a document
attribute.

## Contents

###### Important

This data type is a UNION, so only one of the following members can be specified when used or returned.

**dateValue**

A date expressed as an ISO 8601 string.

It's important for the time zone to be included in the _ISO 8601_
_date-time_ format. For example, 2012-03-25T12:30:10+01:00 is the ISO 8601
date-time format for March 25th 2012 at 12:30PM (plus 10 seconds) in Central European Time.

Type: Timestamp

Required: No

**longValue**

A long integer value.

Type: Long

Required: No

**stringListValue**

A list of strings.

Type: Array of strings

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: No

**stringValue**

A string.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/qapps-2023-11-27/documentattributevalue.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/qapps-2023-11-27/documentattributevalue.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/qapps-2023-11-27/documentattributevalue.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DocumentAttribute

FileUploadCard

All content copied from https://docs.aws.amazon.com/.
