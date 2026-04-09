# PutRequest

Represents a request to perform a `PutItem` operation on an item.

## Contents

###### Note

In the following list, the required parameters are described first.

**Item**

A map of attribute name to attribute values, representing the primary key of an item
to be processed by `PutItem`. All of the table's primary key attributes must
be specified, and their data types must match those of the table's key schema. If any
attributes are present in the item that are part of an index key schema for the table,
their types must match the index key schema.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/putrequest.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/putrequest.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/putrequest.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Put

Replica

All content copied from https://docs.aws.amazon.com/.
