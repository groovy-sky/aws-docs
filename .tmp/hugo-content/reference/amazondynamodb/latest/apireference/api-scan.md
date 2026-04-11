# Scan

The `Scan` operation returns one or more items and item attributes by
accessing every item in a table or a secondary index. To have DynamoDB return fewer
items, you can provide a `FilterExpression` operation.

If the total size of scanned items exceeds the maximum dataset size limit of 1 MB, the
scan completes and results are returned to the user. The `LastEvaluatedKey`
value is also returned and the requestor can use the `LastEvaluatedKey` to
continue the scan in a subsequent operation. Each scan response also includes number of
items that were scanned (ScannedCount) as part of the request. If using a
`FilterExpression`, a scan result can result in no items meeting the
criteria and the `Count` will result in zero. If you did not use a
`FilterExpression` in the scan request, then `Count` is the
same as `ScannedCount`.

###### Note

`Count` and `ScannedCount` only return the count of items
specific to a single scan request and, unless the table is less than 1MB, do not
represent the total number of items in the table.

A single `Scan` operation first reads up to the maximum number of items set
(if using the `Limit` parameter) or a maximum of 1 MB of data and then
applies any filtering to the results if a `FilterExpression` is provided. If
`LastEvaluatedKey` is present in the response, pagination is required to
complete the full table scan. For more information, see [Paginating the\
Results](../../../../services/dynamodb/latest/developerguide/scan.md#Scan.Pagination) in the _Amazon DynamoDB Developer Guide_.

`Scan` operations proceed sequentially; however, for faster performance on
a large table or secondary index, applications can request a parallel `Scan`
operation by providing the `Segment` and `TotalSegments`
parameters. For more information, see [Parallel\
Scan](../../../../services/dynamodb/latest/developerguide/scan.md#Scan.ParallelScan) in the _Amazon DynamoDB Developer Guide_.

By default, a `Scan` uses eventually consistent reads when accessing the
items in a table. Therefore, the results from an eventually consistent `Scan`
may not include the latest item changes at the time the scan iterates through each item
in the table. If you require a strongly consistent read of each item as the scan
iterates through the items in the table, you can set the `ConsistentRead`
parameter to true. Strong consistency only relates to the consistency of the read at the
item level.

###### Note

DynamoDB does not provide snapshot isolation for a scan operation when the
`ConsistentRead` parameter is set to true. Thus, a DynamoDB scan
operation does not guarantee that all reads in a scan see a consistent snapshot of
the table when the scan operation was requested.

## Request Syntax

```nohighlight

{
   "AttributesToGet": [ "string" ],
   "ConditionalOperator": "string",
   "ConsistentRead": boolean,
   "ExclusiveStartKey": {
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
   "ExpressionAttributeNames": {
      "string" : "string"
   },
   "ExpressionAttributeValues": {
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
   "FilterExpression": "string",
   "IndexName": "string",
   "Limit": number,
   "ProjectionExpression": "string",
   "ReturnConsumedCapacity": "string",
   "ScanFilter": {
      "string" : {
         "AttributeValueList": [
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
         "ComparisonOperator": "string"
      }
   },
   "Segment": number,
   "Select": "string",
   "TableName": "string",
   "TotalSegments": number
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[TableName](#API_Scan_RequestSyntax)**

The name of the table containing the requested items or if you provide
`IndexName`, the name of the table to which that index belongs.

You can also provide the Amazon Resource Name (ARN) of the table in this parameter.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: Yes

**[AttributesToGet](#API_Scan_RequestSyntax)**

This is a legacy parameter. Use `ProjectionExpression` instead. For more
information, see [AttributesToGet](../../../../services/dynamodb/latest/developerguide/legacyconditionalparameters-attributestoget.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: Array of strings

Array Members: Minimum number of 1 item.

Length Constraints: Maximum length of 65535.

Required: No

**[ConditionalOperator](#API_Scan_RequestSyntax)**

This is a legacy parameter. Use `FilterExpression` instead. For more
information, see [ConditionalOperator](../../../../services/dynamodb/latest/developerguide/legacyconditionalparameters-conditionaloperator.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String

Valid Values: `AND | OR`

Required: No

**[ConsistentRead](#API_Scan_RequestSyntax)**

A Boolean value that determines the read consistency model during the scan:

- If `ConsistentRead` is `false`, then the data returned
from `Scan` might not contain the results from other recently
completed write operations ( `PutItem`, `UpdateItem`, or
`DeleteItem`).

- If `ConsistentRead` is `true`, then all of the write
operations that completed before the `Scan` began are guaranteed to
be contained in the `Scan` response.

The default setting for `ConsistentRead` is `false`.

The `ConsistentRead` parameter is not supported on global secondary
indexes. If you scan a global secondary index with `ConsistentRead` set to
true, you will receive a `ValidationException`.

Type: Boolean

Required: No

**[ExclusiveStartKey](#API_Scan_RequestSyntax)**

The primary key of the first item that this operation will evaluate. Use the value
that was returned for `LastEvaluatedKey` in the previous operation.

The data type for `ExclusiveStartKey` must be String, Number or Binary. No
set data types are allowed.

In a parallel scan, a `Scan` request that includes
`ExclusiveStartKey` must specify the same segment whose previous
`Scan` returned the corresponding value of
`LastEvaluatedKey`.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

Required: No

**[ExpressionAttributeNames](#API_Scan_RequestSyntax)**

One or more substitution tokens for attribute names in an expression. The following
are some use cases for using `ExpressionAttributeNames`:

- To access an attribute whose name conflicts with a DynamoDB reserved
word.

- To create a placeholder for repeating occurrences of an attribute name in an
expression.

- To prevent special characters in an attribute name from being misinterpreted
in an expression.

Use the **#** character in an expression to dereference
an attribute name. For example, consider the following attribute name:

- `Percentile`

The name of this attribute conflicts with a reserved word, so it cannot be used
directly in an expression. (For the complete list of reserved words, see [Reserved Words](../../../../services/dynamodb/latest/developerguide/reservedwords.md) in the _Amazon DynamoDB Developer_
_Guide_). To work around this, you could specify the following for
`ExpressionAttributeNames`:

- `{"#P":"Percentile"}`

You could then use this substitution in an expression, as in this example:

- `#P = :val`

###### Note

Tokens that begin with the **:** character are
_expression attribute values_, which are placeholders for the
actual value at runtime.

For more information on expression attribute names, see [Specifying Item Attributes](../../../../services/dynamodb/latest/developerguide/expressions-accessingitemattributes.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to string map

Value Length Constraints: Maximum length of 65535.

Required: No

**[ExpressionAttributeValues](#API_Scan_RequestSyntax)**

One or more values that can be substituted in an expression.

Use the **:** (colon) character in an expression to
dereference an attribute value. For example, suppose that you wanted to check whether
the value of the `ProductStatus` attribute was one of the following:

`Available | Backordered | Discontinued`

You would first need to specify `ExpressionAttributeValues` as
follows:

`{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"},
                ":disc":{"S":"Discontinued"} }`

You could then use these values in an expression, such as this:

`ProductStatus IN (:avail, :back, :disc)`

For more information on expression attribute values, see [Condition Expressions](../../../../services/dynamodb/latest/developerguide/expressions-specifyingconditions.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to [AttributeValue](api-attributevalue.md) object map

Required: No

**[FilterExpression](#API_Scan_RequestSyntax)**

A string that contains conditions that DynamoDB applies after the `Scan`
operation, but before the data is returned to you. Items that do not satisfy the
`FilterExpression` criteria are not returned.

###### Note

A `FilterExpression` is applied after the items have already been read;
the process of filtering does not consume any additional read capacity units.

For more information, see [Filter\
Expressions](../../../../services/dynamodb/latest/developerguide/scan.md#Scan.FilterExpression) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String

Required: No

**[IndexName](#API_Scan_RequestSyntax)**

The name of a secondary index to scan. This index can be any local secondary index or
global secondary index. Note that if you use the `IndexName` parameter, you
must also provide `TableName`.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

**[Limit](#API_Scan_RequestSyntax)**

The maximum number of items to evaluate (not necessarily the number of matching
items). If DynamoDB processes the number of items up to the limit while processing the
results, it stops the operation and returns the matching values up to that point, and a
key in `LastEvaluatedKey` to apply in a subsequent operation, so that you can
pick up where you left off. Also, if the processed dataset size exceeds 1 MB before
DynamoDB reaches this limit, it stops the operation and returns the matching values up
to the limit, and a key in `LastEvaluatedKey` to apply in a subsequent
operation to continue the operation. For more information, see [Working with Queries](../../../../services/dynamodb/latest/developerguide/queryandscan.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[ProjectionExpression](#API_Scan_RequestSyntax)**

A string that identifies one or more attributes to retrieve from the specified table
or index. These attributes can include scalars, sets, or elements of a JSON document.
The attributes in the expression must be separated by commas.

If no attribute names are specified, then all attributes will be returned. If any of
the requested attributes are not found, they will not appear in the result.

For more information, see [Specifying Item Attributes](../../../../services/dynamodb/latest/developerguide/expressions-accessingitemattributes.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String

Required: No

**[ReturnConsumedCapacity](#API_Scan_RequestSyntax)**

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

**[ScanFilter](#API_Scan_RequestSyntax)**

This is a legacy parameter. Use `FilterExpression` instead. For more
information, see [ScanFilter](../../../../services/dynamodb/latest/developerguide/legacyconditionalparameters-scanfilter.md) in the _Amazon DynamoDB Developer_
_Guide_.

Type: String to [Condition](api-condition.md) object map

Key Length Constraints: Maximum length of 65535.

Required: No

**[Segment](#API_Scan_RequestSyntax)**

For a parallel `Scan` request, `Segment` identifies an
individual segment to be scanned by an application worker.

Segment IDs are zero-based, so the first segment is always 0. For example, if you want
to use four application threads to scan a table or an index, then the first thread
specifies a `Segment` value of 0, the second thread specifies 1, and so
on.

The value of `LastEvaluatedKey` returned from a parallel `Scan`
request must be used as `ExclusiveStartKey` with the same segment ID in a
subsequent `Scan` operation.

The value for `Segment` must be greater than or equal to 0, and less than
the value provided for `TotalSegments`.

If you provide `Segment`, you must also provide
`TotalSegments`.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 999999.

Required: No

**[Select](#API_Scan_RequestSyntax)**

The attributes to be returned in the result. You can retrieve all item attributes,
specific item attributes, the count of matching items, or in the case of an index, some
or all of the attributes projected into the index.

- `ALL_ATTRIBUTES` \- Returns all of the item attributes from the
specified table or index. If you query a local secondary index, then for each
matching item in the index, DynamoDB fetches the entire item from the parent
table. If the index is configured to project all item attributes, then all of
the data can be obtained from the local secondary index, and no fetching is
required.

- `ALL_PROJECTED_ATTRIBUTES` \- Allowed only when querying an index.
Retrieves all attributes that have been projected into the index. If the index
is configured to project all attributes, this return value is equivalent to
specifying `ALL_ATTRIBUTES`.

- `COUNT` \- Returns the number of matching items, rather than the
matching items themselves. Note that this uses the same quantity of read
capacity units as getting the items, and is subject to the same item size
calculations.

- `SPECIFIC_ATTRIBUTES` \- Returns only the attributes listed in
`ProjectionExpression`. This return value is equivalent to
specifying `ProjectionExpression` without specifying any value for
`Select`.

If you query or scan a local secondary index and request only attributes that
are projected into that index, the operation reads only the index and not the
table. If any of the requested attributes are not projected into the local
secondary index, DynamoDB fetches each of these attributes from the parent
table. This extra fetching incurs additional throughput cost and latency.

If you query or scan a global secondary index, you can only request attributes
that are projected into the index. Global secondary index queries cannot fetch
attributes from the parent table.

If neither `Select` nor `ProjectionExpression` are specified,
DynamoDB defaults to `ALL_ATTRIBUTES` when accessing a table, and
`ALL_PROJECTED_ATTRIBUTES` when accessing an index. You cannot use both
`Select` and `ProjectionExpression` together in a single
request, unless the value for `Select` is `SPECIFIC_ATTRIBUTES`.
(This usage is equivalent to specifying `ProjectionExpression` without any
value for `Select`.)

###### Note

If you use the `ProjectionExpression` parameter, then the value for
`Select` can only be `SPECIFIC_ATTRIBUTES`. Any other
value for `Select` will return an error.

Type: String

Valid Values: `ALL_ATTRIBUTES | ALL_PROJECTED_ATTRIBUTES | SPECIFIC_ATTRIBUTES | COUNT`

Required: No

**[TotalSegments](#API_Scan_RequestSyntax)**

For a parallel `Scan` request, `TotalSegments` represents the
total number of segments into which the `Scan` operation will be divided. The
value of `TotalSegments` corresponds to the number of application workers
that will perform the parallel scan. For example, if you want to use four application
threads to scan a table or an index, specify a `TotalSegments` value of
4.

The value for `TotalSegments` must be greater than or equal to 1, and less
than or equal to 1000000. If you specify a `TotalSegments` value of 1, the
`Scan` operation will be sequential rather than parallel.

If you specify `TotalSegments`, you must also specify
`Segment`.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 1000000.

Required: No

## Response Syntax

```nohighlight

{
   "ConsumedCapacity": {
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
   },
   "Count": number,
   "Items": [
      {
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
   ],
   "LastEvaluatedKey": {
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
   "ScannedCount": number
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ConsumedCapacity](#API_Scan_ResponseSyntax)**

The capacity units consumed by the `Scan` operation. The data returned
includes the total provisioned throughput consumed, along with statistics for the table
and any indexes involved in the operation. `ConsumedCapacity` is only
returned if the `ReturnConsumedCapacity` parameter was specified. For more
information, see [Capacity unit consumption for read operations](../../../../services/dynamodb/latest/developerguide/read-write-operations.md#read-operation-consumption) in the _Amazon_
_DynamoDB Developer Guide_.

Type: [ConsumedCapacity](api-consumedcapacity.md) object

**[Count](#API_Scan_ResponseSyntax)**

The number of items in the response.

If you set `ScanFilter` in the request, then `Count` is the
number of items returned after the filter was applied, and `ScannedCount` is
the number of matching items before the filter was applied.

If you did not use a filter in the request, then `Count` is the same as
`ScannedCount`.

Type: Integer

**[Items](#API_Scan_ResponseSyntax)**

An array of item attributes that match the scan criteria. Each element in this array
consists of an attribute name and the value for that attribute.

Type: Array of string to [AttributeValue](api-attributevalue.md) object maps

Key Length Constraints: Maximum length of 65535.

**[LastEvaluatedKey](#API_Scan_ResponseSyntax)**

The primary key of the item where the operation stopped, inclusive of the previous
result set. Use this value to start a new operation, excluding this value in the new
request.

If `LastEvaluatedKey` is empty, then the "last page" of results has been
processed and there is no more data to be retrieved.

If `LastEvaluatedKey` is not empty, it does not necessarily mean that there
is more data in the result set. The only way to know when you have reached the end of
the result set is when `LastEvaluatedKey` is empty.

Type: String to [AttributeValue](api-attributevalue.md) object map

Key Length Constraints: Maximum length of 65535.

**[ScannedCount](#API_Scan_ResponseSyntax)**

The number of items evaluated, before any `ScanFilter` is applied. A high
`ScannedCount` value with few, or no, `Count` results
indicates an inefficient `Scan` operation. For more information, see [Count and\
ScannedCount](../../../../services/dynamodb/latest/developerguide/queryandscan.md#Count) in the _Amazon DynamoDB Developer_
_Guide_.

If you did not use a filter in the request, then `ScannedCount` is the same
as `Count`.

Type: Integer

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

## Examples

### Return All Items

The following example returns all of the items in a table. No scan filter is
applied.

#### Sample Request

```

POST / HTTP/1.1
Host: dynamodb.<region>.<domain>;
Accept-Encoding: identity
Content-Length: <PayloadSizeBytes>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.0
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=<Headers>, Signature=<Signature>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDB_20120810.Scan

{
    "TableName": "Reply",
    "ReturnConsumedCapacity": "TOTAL"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
x-amz-crc32: <Checksum>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
Date: <Date>
 {
    "ConsumedCapacity": {
        "CapacityUnits": 0.5,
        "TableName": "Reply"
    },
    "Count": 4,
    "Items": [
        {
            "PostedBy": {
                "S": "joe@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115336"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "Have you looked at BatchWriteItem?"
            }
        },
        {
            "PostedBy": {
                "S": "fred@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115342"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "No, I didn't know about that.  Where can I find more information?"
            }
        },
        {
            "PostedBy": {
                "S": "joe@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115347"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "BatchWriteItem is documented in the Amazon DynamoDB API Reference."
            }
        },
        {
            "PostedBy": {
                "S": "fred@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115352"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "OK, I'll take a look at that.  Thanks!"
            }
        }
    ],
    "ScannedCount": 4
}
```

### Use a Filter Expression

The following example returns only those items matching specific
criteria.

#### Sample Request

```

POST / HTTP/1.1
Host: dynamodb.<region>.<domain>;
Accept-Encoding: identity
Content-Length: <PayloadSizeBytes>
User-Agent: <UserAgentString>
Content-Type: application/x-amz-json-1.0
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=<Headers>, Signature=<Signature>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDB_20120810.Scan

{
    "TableName": "Reply",
    "FilterExpression": "PostedBy = :val",
    "ExpressionAttributeValues": {":val": {"S": "joe@example.com"}},
    "ReturnConsumedCapacity": "TOTAL"
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
x-amz-crc32: <Checksum>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
Date: <Date>
 {
    "ConsumedCapacity": {
        "CapacityUnits": 0.5,
        "TableName": "Reply"
    },
    "Count": 2,
    "Items": [
        {
            "PostedBy": {
                "S": "joe@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115336"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "Have you looked at BatchWriteItem?"
            }
        },
        {
            "PostedBy": {
                "S": "joe@example.com"
            },
            "ReplyDateTime": {
                "S": "20130320115347"
            },
            "Id": {
                "S": "Amazon DynamoDB#How do I update multiple items?"
            },
            "Message": {
                "S": "BatchWriteItem is documented in the Amazon DynamoDB API Reference."
            }
        }
    ],
    "ScannedCount": 4
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/dynamodb-2012-08-10/scan.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/dynamodb-2012-08-10/scan.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/dynamodb-2012-08-10/scan.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/dynamodb-2012-08-10/scan.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/dynamodb-2012-08-10/scan.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/dynamodb-2012-08-10/scan.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/dynamodb-2012-08-10/scan.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/dynamodb-2012-08-10/scan.md)

- [AWS SDK for Python](../../../../services/goto/boto3/dynamodb-2012-08-10/scan.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/dynamodb-2012-08-10/scan.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RestoreTableToPointInTime

TagResource

All content copied from https://docs.aws.amazon.com/.
