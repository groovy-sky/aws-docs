---
title: "LogFieldType"
---

# LogFieldType

Defines the data type structure for a log field, including the type, element information,
and nested fields for complex types.

## Contents

**element**

For array or collection types, specifies the element type information.

Type: [LogFieldType](api-logfieldtype.md) object

Required: No

**fields**

For complex types, contains the nested field definitions.

Type: Array of [LogFieldsListItem](api-logfieldslistitem.md) objects

Required: No

**type**

The data type of the log field.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/LogFieldType)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/LogFieldType)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/LogFieldType)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LogFieldsListItem

LogGroup

All content copied from https://docs.aws.amazon.com/.
