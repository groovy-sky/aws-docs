# BatchStatementResponse

A PartiQL batch statement response..

## Contents

###### Note

In the following list, the required parameters are described first.

**Error**

The error associated with a failed PartiQL batch statement.

Type: [BatchStatementError](api-batchstatementerror.md) object

Required: No

**Item**

A DynamoDB item associated with a BatchStatementResponse

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: No

**TableName**

The table name associated with a failed PartiQL batch statement.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/batchstatementresponse.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/batchstatementresponse.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/batchstatementresponse.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchStatementRequest

BillingModeSummary

All content copied from https://docs.aws.amazon.com/.
