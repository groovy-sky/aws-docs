# BatchStatementError

An error associated with a statement in a PartiQL batch that was run.

## Contents

###### Note

In the following list, the required parameters are described first.

**Code**

The error code associated with the failed PartiQL batch statement.

Type: String

Valid Values: `ConditionalCheckFailed | ItemCollectionSizeLimitExceeded | RequestLimitExceeded | ValidationError | ProvisionedThroughputExceeded | TransactionConflict | ThrottlingError | InternalServerError | ResourceNotFound | AccessDenied | DuplicateItem`

Required: No

**Item**

The item which caused the condition check to fail. This will be set if
ReturnValuesOnConditionCheckFailure is specified as `ALL_OLD`.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: No

**Message**

The error message associated with the PartiQL batch response.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/batchstatementerror.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/batchstatementerror.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/batchstatementerror.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BackupSummary

BatchStatementRequest

All content copied from https://docs.aws.amazon.com/.
