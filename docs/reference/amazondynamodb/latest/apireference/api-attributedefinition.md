---
title: "AttributeDefinition"
---

# AttributeDefinition

Represents an attribute for describing the schema for the table and indexes.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeName**

A name for the attribute.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**AttributeType**

The data type for the attribute, where:

- `S` \- the attribute is of type String

- `N` \- the attribute is of type Number

- `B` \- the attribute is of type Binary

Type: String

Valid Values: `S | N | B`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/dynamodb-2012-08-10/AttributeDefinition)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/dynamodb-2012-08-10/AttributeDefinition)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/dynamodb-2012-08-10/AttributeDefinition)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ArchivalSummary

AttributeValue

All content copied from https://docs.aws.amazon.com/.
