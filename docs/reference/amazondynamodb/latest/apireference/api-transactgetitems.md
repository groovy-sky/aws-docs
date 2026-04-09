# TransactGetItems

`TransactGetItems` is a synchronous operation that atomically retrieves
multiple items from one or more tables (but not from indexes) in a single account and
Region. A `TransactGetItems` call can contain up to 100
`TransactGetItem` objects, each of which contains a `Get`
structure that specifies an item to retrieve from a table in the account and Region. A
call to `TransactGetItems` cannot retrieve items from tables in more than one
AWS account or Region. The aggregate size of the items in the
transaction cannot exceed 4 MB.

DynamoDB rejects the entire `TransactGetItems` request if any of
the following is true:

- A conflicting operation is in the process of updating an item to be
read.

- There is insufficient provisioned capacity for the transaction to be
completed.

- There is a user error, such as an invalid data format.

- The aggregate size of the items in the transaction exceeded 4 MB.

## Request Syntax

```nohighlight

{
   "ReturnConsumedCapacity": "string",
   "TransactItems": [
      {
         "Get": {
            "ExpressionAttributeNames": {
               "string" : "string"
            },
            "Key": {
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
            "ProjectionExpression": "string",
            "TableName": "string"
         }
      }
   ]
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[TransactItems](#API_TransactGetItems_RequestSyntax)**

An ordered array of up to 100 `TransactGetItem` objects, each of which
contains a `Get` structure.

Type: Array of [TransactGetItem](api-transactgetitem.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

Required: Yes

**[ReturnConsumedCapacity](#API_TransactGetItems_RequestSyntax)**

A value of `TOTAL` causes consumed capacity information to be returned, and
a value of `NONE` prevents that information from being returned. No other
value is valid.

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
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ConsumedCapacity](#API_TransactGetItems_ResponseSyntax)**

If the _ReturnConsumedCapacity_ value was `TOTAL`, this
is an array of `ConsumedCapacity` objects, one for each table addressed by
`TransactGetItem` objects in the _TransactItems_
parameter. These `ConsumedCapacity` objects report the read-capacity units
consumed by the `TransactGetItems` call in that table.

Type: Array of [ConsumedCapacity](api-consumedcapacity.md) objects

**[Responses](#API_TransactGetItems_ResponseSyntax)**

An ordered array of up to 100 `ItemResponse` objects, each of which
corresponds to the `TransactGetItem` object in the same position in the
_TransactItems_ array. Each `ItemResponse` object
contains a Map of the name-value pairs that are the projected attributes of the
requested item.

If a requested item could not be retrieved, the corresponding
`ItemResponse` object is Null, or if the requested item has no projected
attributes, the corresponding `ItemResponse` object is an empty Map.

Type: Array of [ItemResponse](api-itemresponse.md) objects

Array Members: Minimum number of 1 item. Maximum number of 100 items.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ProvisionedThroughputExceededException**

The request was denied due to request throttling. For detailed information about
why the request was throttled and the ARN of the impacted resource, find the [ThrottlingReason](api-throttlingreason.md) field in the returned exception. The AWS
SDKs for DynamoDB automatically retry requests that receive this exception.
Your request is eventually successful, unless your retry queue is too large to finish.
Reduce the frequency of requests and use exponential backoff. For more information, go
to [Error Retries and Exponential Backoff](../../../../services/dynamodb/latest/developerguide/programming-errors.md#Programming.Errors.RetryAndBackoff) in the _Amazon DynamoDB Developer Guide_.

**message**

You exceeded your maximum allowed provisioned throughput.

**ThrottlingReasons**

A list of [ThrottlingReason](api-throttlingreason.md) that
provide detailed diagnostic information about why the request was throttled.

HTTP Status Code: 400

**RequestLimitExceeded**

Throughput exceeds the current throughput quota for your account. For detailed
information about why the request was throttled and the ARN of the impacted resource,
find the [ThrottlingReason](api-throttlingreason.md) field in the returned exception. Contact [Support](https://aws.amazon.com/support) to request a quota
increase.

**ThrottlingReasons**

A list of [ThrottlingReason](api-throttlingreason.md) that
provide detailed diagnostic information about why the request was throttled.

HTTP Status Code: 400

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource might not
be specified correctly, or its status might not be `ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

**ThrottlingException**

The request was denied due to request throttling. For detailed information about why
the request was throttled and the ARN of the impacted resource, find the [ThrottlingReason](api-throttlingreason.md) field in the returned exception.

**throttlingReasons**

A list of [ThrottlingReason](api-throttlingreason.md) that
provide detailed diagnostic information about why the request was throttled.

HTTP Status Code: 400

**TransactionCanceledException**

The entire transaction request was canceled.

DynamoDB cancels a `TransactWriteItems` request under the following
circumstances:

- A condition in one of the condition expressions is not met.

- A table in the `TransactWriteItems` request is in a different
account or region.

- More than one action in the `TransactWriteItems` operation
targets the same item.

- There is insufficient provisioned capacity for the transaction to be
completed.

- An item size becomes too large (larger than 400 KB), or a local secondary
index (LSI) becomes too large, or a similar validation error occurs because of
changes made by the transaction.

- There is a user error, such as an invalid data format.

- There is an ongoing `TransactWriteItems` operation that
conflicts with a concurrent `TransactWriteItems` request. In this
case the `TransactWriteItems` operation fails with a
`TransactionCanceledException`.

DynamoDB cancels a `TransactGetItems` request under the
following circumstances:

- There is an ongoing `TransactGetItems` operation that conflicts
with a concurrent `PutItem`, `UpdateItem`,
`DeleteItem` or `TransactWriteItems` request. In this
case the `TransactGetItems` operation fails with a
`TransactionCanceledException`.

- A table in the `TransactGetItems` request is in a different
account or region.

- There is insufficient provisioned capacity for the transaction to be
completed.

- There is a user error, such as an invalid data format.

###### Note

DynamoDB lists the cancellation reasons on the
`CancellationReasons` property. Transaction cancellation reasons are ordered in the order of requested
items, if an item has no error it will have `None` code and
`Null` message.

Cancellation reason codes and possible error messages:

- No Errors:

- Code: `None`

- Message: `null`

- Conditional Check Failed:

- Code: `ConditionalCheckFailed`

- Message: The conditional request failed.

- Item Collection Size Limit Exceeded:

- Code: `ItemCollectionSizeLimitExceeded`

- Message: Collection size exceeded.

- Transaction Conflict:

- Code: `TransactionConflict`

- Message: Transaction is ongoing for the item.

- Provisioned Throughput Exceeded:

- Code: `ProvisionedThroughputExceeded`

- Messages:

- The level of configured provisioned throughput for the
table was exceeded. Consider increasing your provisioning level
with the UpdateTable API.

###### Note

This Message is received when provisioned throughput is
exceeded is on a provisioned DynamoDB
table.

- The level of configured provisioned throughput for one or
more global secondary indexes of the table was exceeded.
Consider increasing your provisioning level for the
under-provisioned global secondary indexes with the UpdateTable
API.

###### Note

This message is returned when provisioned throughput is
exceeded is on a provisioned GSI.

- Throttling Error:

- Code: `ThrottlingError`

- Messages:

- Throughput exceeds the current capacity of your table or
index. DynamoDB is automatically scaling your table or
index so please try again shortly. If exceptions persist, check
if you have a hot key:
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/bp-partition-key-design.html.

###### Note

This message is returned when writes get throttled on an
On-Demand table as DynamoDB is automatically
scaling the table.

- Throughput exceeds the current capacity for one or more
global secondary indexes. DynamoDB is automatically
scaling your index so please try again shortly.

###### Note

This message is returned when writes get throttled on an
On-Demand GSI as DynamoDB is automatically scaling
the GSI.

- Validation Error:

- Code: `ValidationError`

- Messages:

- One or more parameter values were invalid.

- The update expression attempted to update the secondary
index key beyond allowed size limits.

- The update expression attempted to update the secondary
index key to unsupported type.

- An operand in the update expression has an incorrect data
type.

- Item size to update has exceeded the maximum allowed
size.

- Number overflow. Attempting to store a number with
magnitude larger than supported range.

- Type mismatch for attribute to update.

- Nesting Levels have exceeded supported limits.

- The document path provided in the update expression is
invalid for update.

- The provided expression refers to an attribute that does
not exist in the item.

**CancellationReasons**

A list of cancellation reasons.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/transactgetitems.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/transactgetitems.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TagResource

TransactWriteItems

All content copied from https://docs.aws.amazon.com/.
