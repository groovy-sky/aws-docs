# TimeToLiveSpecification

Represents the settings used to enable or disable Time to Live (TTL) for the specified
table.

## Contents

###### Note

In the following list, the required parameters are described first.

**AttributeName**

The name of the TTL attribute used to store the expiration time for items in the
table.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: Yes

**Enabled**

Indicates whether TTL is to be enabled (true) or disabled (false) on the table.

Type: Boolean

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/timetolivespecification.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/timetolivespecification.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/timetolivespecification.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TimeToLiveDescription

TransactGetItem

All content copied from https://docs.aws.amazon.com/.
