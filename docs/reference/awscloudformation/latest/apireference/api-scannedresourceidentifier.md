---
title: "ScannedResourceIdentifier"
---

# ScannedResourceIdentifier

Identifies a scanned resource. This is used with the
`ListResourceScanRelatedResources` API action.

## Contents

**ResourceIdentifier**
ResourceIdentifier.entry.N.key (key)

ResourceIdentifier.entry.N.value (value)

A list of up to 256 key-value pairs that identifies the scanned resource. The key is the
name of one of the primary identifiers for the resource. (Primary identifiers are specified in
the `primaryIdentifier` list in the resource schema.) The value is the value of that
primary identifier. For example, for a `AWS::DynamoDB::Table` resource, the primary
identifiers is `TableName` so the key-value pair could be `"TableName":
    "MyDDBTable"`. For more information, see [primaryIdentifier](../../../../services/cloudformation-cli/latest/userguide/resource-type-schema.md#schema-properties-primaryidentifier) in the _AWS CloudFormation Command Line Interface (CLI) User Guide_.

Type: String to string map

Required: Yes

**ResourceType**

The type of the resource, such as `AWS::DynamoDB::Table`. For the list of
supported resources, see [Resource type\
support for imports and drift detection](../../../../services/cloudformation/latest/userguide/resource-import-supported-resources.md) In the
_AWS CloudFormation User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/cloudformation-2010-05-15/ScannedResourceIdentifier)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/cloudformation-2010-05-15/ScannedResourceIdentifier)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/cloudformation-2010-05-15/ScannedResourceIdentifier)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ScannedResource

Stack

All content copied from https://docs.aws.amazon.com/.
