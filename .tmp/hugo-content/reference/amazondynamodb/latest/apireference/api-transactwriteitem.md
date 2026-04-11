# TransactWriteItem

A list of requests that can perform update, put, delete, or check operations on
multiple items in one or more tables atomically.

## Contents

###### Note

In the following list, the required parameters are described first.

**ConditionCheck**

A request to perform a check item operation.

Type: [ConditionCheck](api-conditioncheck.md) object

Required: No

**Delete**

A request to perform a `DeleteItem` operation.

Type: [Delete](api-delete.md) object

Required: No

**Put**

A request to perform a `PutItem` operation.

Type: [Put](api-put.md) object

Required: No

**Update**

A request to perform an `UpdateItem` operation.

Type: [Update](api-update.md) object

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/transactwriteitem.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/transactwriteitem.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/transactwriteitem.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TransactGetItem

Update

All content copied from https://docs.aws.amazon.com/.
