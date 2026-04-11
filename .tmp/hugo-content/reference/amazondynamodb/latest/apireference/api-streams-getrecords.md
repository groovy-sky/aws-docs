# GetRecords

Retrieves the stream records from a given shard.

Specify a shard iterator using the `ShardIterator` parameter. The shard iterator
specifies the position in the shard from which you want to start reading stream records
sequentially. If there are no stream records available in the portion of the shard that the
iterator points to, `GetRecords` returns an empty list. Note that it might take
multiple calls to get to a portion of the shard that contains stream records.

###### Note

`GetRecords` can retrieve a maximum of 1 MB of data or 1000 stream records,
whichever comes first.

## Request Syntax

```nohighlight

{
   "Limit": number,
   "ShardIterator": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ShardIterator](#API_streams_GetRecords_RequestSyntax)**

A shard iterator that was retrieved from a previous GetShardIterator operation. This iterator can be used to access the stream records in this shard.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Required: Yes

**[Limit](#API_streams_GetRecords_RequestSyntax)**

The maximum number of records to return from the shard. The upper limit is 1000.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

## Response Syntax

```nohighlight

{
   "NextShardIterator": "string",
   "Records": [
      {
         "awsRegion": "string",
         "dynamodb": {
            "ApproximateCreationDateTime": number,
            "Keys": {
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
            "NewImage": {
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
            "OldImage": {
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
            "SequenceNumber": "string",
            "SizeBytes": number,
            "StreamViewType": "string"
         },
         "eventID": "string",
         "eventName": "string",
         "eventSource": "string",
         "eventVersion": "string",
         "userIdentity": {
            "PrincipalId": "string",
            "Type": "string"
         }
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[NextShardIterator](#API_streams_GetRecords_ResponseSyntax)**

The next position in the shard from which to start sequentially reading stream records. If
set to `null`, the shard has been closed and the requested iterator will not return
any more data.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

**[Records](#API_streams_GetRecords_ResponseSyntax)**

The stream records from the shard, which were retrieved using the shard iterator.

Type: Array of [Record](api-streams-record.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ExpiredIteratorException**

The shard iterator has expired and can no longer be used to retrieve stream records. A shard
iterator expires 15 minutes after it is retrieved using the `GetShardIterator`
action.

**message**

The provided iterator exceeds the maximum age allowed.

HTTP Status Code: 400

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**LimitExceededException**

There is no limit to the number of daily on-demand backups that can be taken.

For most purposes, up to 500 simultaneous table operations are allowed per account. These operations
include `CreateTable`, `UpdateTable`,
`DeleteTable`, `UpdateTimeToLive`,
`RestoreTableFromBackup`, and `RestoreTableToPointInTime`.

When you are creating a table with one or more secondary
indexes, you can have up to 250 such requests running at a time. However, if the table or
index specifications are complex, then DynamoDB might temporarily reduce the number
of concurrent operations.

When importing into DynamoDB, up to 50 simultaneous import table operations are allowed per account.

There is a soft account quota of 2,500 tables.

GetRecords was called with a value of more than 1000 for the limit request parameter.

More than 2 processes are reading from the same streams shard at the same time. Exceeding
this limit may result in request throttling.

**message**

Too many operations for a given subscriber.

HTTP Status Code: 400

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource
might not be specified correctly, or its status might not be
`ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

**TrimmedDataAccessException**

The operation attempted to read past the oldest stream record in a shard.

In DynamoDB Streams, there is a 24 hour limit on data retention. Stream records whose age exceeds this limit are subject to removal (trimming) from the stream. You might receive a TrimmedDataAccessException if:

- You request a shard iterator with a sequence number older than the trim point (24 hours).

- You obtain a shard iterator, but before you use the iterator in a `GetRecords`
request, a stream record in the shard exceeds the 24 hour period and is trimmed. This causes
the iterator to access a record that no longer exists.

**message**

"The data you are trying to access has been trimmed.

HTTP Status Code: 400

## Examples

### Retrieve stream records from a shard

The following sample retrieves all the stream records from a shard. To do this,
it uses a `ShardIterator` that was obtained from a previous
`GetShardIterator` call.

#### Sample Request

```

POST / HTTP/1.1
x-amzn-RequestId: <RequestID>
x-amzn-crc32: <CRC32>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDBStreams_20120810.GetRecords

{
    "ShardIterator": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252|1|AAAAAAAAAAEvJp6D+zaQ...  <remaining characters omitted> ..."
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
    "NextShardIterator": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252|1|AAAAAAAAAAGQBYshYDEe ... <remaining characters omitted> ...",
    "Records": [
        {
            "awsRegion": "us-west-2",
            "dynamodb": {
                "ApproximateCreationDateTime": 1.46480431E9,
                "Keys": {
                    "ForumName": {"S": "DynamoDB"},
                    "Subject": {"S": "DynamoDB Thread 3"}
                },
                "SequenceNumber": "300000000000000499659",
                "SizeBytes": 41,
                "StreamViewType": "KEYS_ONLY"
            },
            "eventID": "e2fd9c34eff2d779b297b26f5fef4206",
            "eventName": "INSERT",
            "eventSource": "aws:dynamodb",
            "eventVersion": "1.0"
        },
        {
            "awsRegion": "us-west-2",
            "dynamodb": {
                "ApproximateCreationDateTime": 1.46480527E9,
                "Keys": {
                    "ForumName": {"S": "DynamoDB"},
                    "Subject": {"S": "DynamoDB Thread 1"}
                },
                "SequenceNumber": "400000000000000499660",
                "SizeBytes": 41,
                "StreamViewType": "KEYS_ONLY"
            },
            "eventID": "4b25bd0da9a181a155114127e4837252",
            "eventName": "MODIFY",
            "eventSource": "aws:dynamodb",
            "eventVersion": "1.0"
        },
        {
            "awsRegion": "us-west-2",
            "dynamodb": {
                "ApproximateCreationDateTime": 1.46480646E9,
                "Keys": {
                    "ForumName": {"S": "DynamoDB"},
                    "Subject": {"S": "DynamoDB Thread 2"}
                },
                "SequenceNumber": "500000000000000499661",
                "SizeBytes": 41,
                "StreamViewType": "KEYS_ONLY"
            },
            "eventID": "740280c73a3df7842edab3548a1b08ad",
            "eventName": "REMOVE",
            "eventSource": "aws:dynamodb",
            "eventVersion": "1.0"
        }
    ]
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for Python](../../../../services/goto/boto3/streams-dynamodb-2012-08-10/getrecords.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/getrecords.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeStream

GetShardIterator

All content copied from https://docs.aws.amazon.com/.
