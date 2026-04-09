# BatchExecuteStatement

This operation allows you to perform batch reads or writes on data stored in DynamoDB,
using PartiQL. Each read statement in a `BatchExecuteStatement` must specify
an equality condition on all key attributes. This enforces that each `SELECT`
statement in a batch returns at most a single item. For more information, see [Running batch operations with PartiQL for DynamoDB](../../../../services/dynamodb/latest/developerguide/ql-reference-multiplestatements-batching.md).

###### Note

The entire batch must consist of either read statements or write statements, you
cannot mix both in one batch.

###### Important

A HTTP 200 response does not mean that all statements in the BatchExecuteStatement
succeeded. Error details for individual statements can be found under the [Error](api-batchstatementresponse.md#DDB-Type-BatchStatementResponse-Error) field of the `BatchStatementResponse` for each
statement.

## Request Syntax

```nohighlight

{
   "ReturnConsumedCapacity": "string",
   "Statements": [
      {
         "ConsistentRead": boolean,
         "Parameters": [
            {
               "B": blob,
               "BOOL": boolean,
               "BS": [ blob ],
               "L": [
                  "AttributeValue"
               ],
               "M": {
                  "string" : "AttributeValue"
               },
               "N": "string",
               "NS": [ "string" ],
               "NULL": boolean,
               "S": "string",
               "SS": [ "string" ]
            }
         ],
         "ReturnValuesOnConditionCheckFailure": "string",
         "Statement": "string"
      }
   ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[Statements](#API_BatchExecuteStatement_RequestSyntax)**

The list of PartiQL statements representing the batch to run.

Type: Array of [BatchStatementRequest](api-batchstatementrequest.md) objects

Array Members: Minimum number of 1 item. Maximum number of 25 items.

Required: Yes

**[ReturnConsumedCapacity](#API_BatchExecuteStatement_RequestSyntax)**

Determines the level of detail about either provisioned or on-demand throughput
consumption that is returned in the response:

- `INDEXES` \- The response includes the aggregate
`ConsumedCapacity` for the operation, together with
`ConsumedCapacity` for each table and secondary index that was
accessed.

Note that some operations, such as `GetItem` and
`BatchGetItem`, do not access any indexes at all. In these cases,
specifying `INDEXES` will only return `ConsumedCapacity`
information for table(s).

- `TOTAL` \- The response includes only the aggregate
`ConsumedCapacity` for the operation.

- `NONE` \- No `ConsumedCapacity` details are included in the
response.

Type: String

Valid Values: `INDEXES | TOTAL | NONE`

Required: No

## Response Syntax

```nohighlight

{
   "ConsumedCapacity": [
      {
         "CapacityUnits": number,
         "GlobalSecondaryIndexes": {
            "string" : {
               "CapacityUnits": number,
               "ReadCapacityUnits": number,
               "WriteCapacityUnits": number
            }
         },
         "LocalSecondaryIndexes": {
            "string" : {
               "CapacityUnits": number,
               "ReadCapacityUnits": number,
               "WriteCapacityUnits": number
            }
         },
         "ReadCapacityUnits": number,
         "Table": {
            "CapacityUnits": number,
            "ReadCapacityUnits": number,
            "WriteCapacityUnits": number
         },
         "TableName": "string",
         "WriteCapacityUnits": number
      }
   ],
   "Responses": [
      {
         "Error": {
            "Code": "string",
            "Item": {
               "string" : {
                  "B": blob,
                  "BOOL": boolean,
                  "BS": [ blob ],
                  "L": [
                     "AttributeValue"
                  ],
                  "M": {
                     "string" : "AttributeValue"
                  },
                  "N": "string",
                  "NS": [ "string" ],
                  "NULL": boolean,
                  "S": "string",
                  "SS": [ "string" ]
               }
            },
            "Message": "string"
         },
         "Item": {
            "string" : {
               "B": blob,
               "BOOL": boolean,
               "BS": [ blob ],
               "L": [
                  "AttributeValue"
               ],
               "M": {
                  "string" : "AttributeValue"
               },
               "N": "string",
               "NS": [ "string" ],
               "NULL": boolean,
               "S": "string",
               "SS": [ "string" ]
            }
         },
         "TableName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ConsumedCapacity](#API_BatchExecuteStatement_ResponseSyntax)**

The capacity units consumed by the entire operation. The values of the list are
ordered according to the ordering of the statements.

Type: Array of [ConsumedCapacity](api-consumedcapacity.md) objects

**[Responses](#API_BatchExecuteStatement_ResponseSyntax)**

The response to each PartiQL statement in the batch. The values of the list are
ordered according to the ordering of the request statements.

Type: Array of [BatchStatementResponse](api-batchstatementresponse.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**RequestLimitExceeded**

Throughput exceeds the current throughput quota for your account. For detailed
information about why the request was throttled and the ARN of the impacted resource,
find the [ThrottlingReason](api-throttlingreason.md) field in the returned exception. Contact [Support](https://aws.amazon.com/support) to request a quota
increase.

**ThrottlingReasons**

A list of [ThrottlingReason](api-throttlingreason.md) that
provide detailed diagnostic information about why the request was throttled.

HTTP Status Code: 400

**ThrottlingException**

The request was denied due to request throttling. For detailed information about why
the request was throttled and the ARN of the impacted resource, find the [ThrottlingReason](api-throttlingreason.md) field in the returned exception.

**throttlingReasons**

A list of [ThrottlingReason](api-throttlingreason.md) that
provide detailed diagnostic information about why the request was throttled.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/batchexecutestatement.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/batchexecutestatement.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon DynamoDB

BatchGetItem

All content copied from https://docs.aws.amazon.com/.
