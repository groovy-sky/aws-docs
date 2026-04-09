# DescribeStream

Returns information about a stream, including the current status of the stream, its Amazon Resource Name (ARN), the composition of its shards, and its corresponding DynamoDB table.

###### Note

You can call `DescribeStream` at a maximum rate of 10 times per second.

Each shard in the stream has a `SequenceNumberRange` associated with it. If the
`SequenceNumberRange` has a `StartingSequenceNumber` but no
`EndingSequenceNumber`, then the shard is still open (able to receive more stream
records). If both `StartingSequenceNumber` and `EndingSequenceNumber`
are present, then that shard is closed and can no longer receive more data.

## Request Syntax

```nohighlight

{
   "ExclusiveStartShardId": "string",
   "Limit": number,
   "ShardFilter": {
      "ShardId": "string",
      "Type": "string"
   },
   "StreamArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[StreamArn](#API_streams_DescribeStream_RequestSyntax)**

The Amazon Resource Name (ARN) for the stream.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: Yes

**[ExclusiveStartShardId](#API_streams_DescribeStream_RequestSyntax)**

The shard ID of the first item that this operation will evaluate. Use the value that was
returned for `LastEvaluatedShardId` in the previous operation.

Type: String

Length Constraints: Minimum length of 28. Maximum length of 65.

Required: No

**[Limit](#API_streams_DescribeStream_RequestSyntax)**

The maximum number of shard objects to return. The upper limit is 100.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[ShardFilter](#API_streams_DescribeStream_RequestSyntax)**

This optional field contains the filter definition for the
`DescribeStream` API.

Type: [ShardFilter](api-streams-shardfilter.md) object

Required: No

## Response Syntax

```nohighlight

{
   "StreamDescription": {
      "CreationRequestDateTime": number,
      "KeySchema": [
         {
            "AttributeName": "string",
            "KeyType": "string"
         }
      ],
      "LastEvaluatedShardId": "string",
      "Shards": [
         {
            "ParentShardId": "string",
            "SequenceNumberRange": {
               "EndingSequenceNumber": "string",
               "StartingSequenceNumber": "string"
            },
            "ShardId": "string"
         }
      ],
      "StreamArn": "string",
      "StreamLabel": "string",
      "StreamStatus": "string",
      "StreamViewType": "string",
      "TableName": "string"
   }
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[StreamDescription](#API_streams_DescribeStream_ResponseSyntax)**

A complete description of the stream, including its creation date and time, the DynamoDB table associated with the stream, the shard IDs within the stream, and the beginning and ending sequence numbers of stream records within the shards.

Type: [StreamDescription](api-streams-streamdescription.md) object

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerError**

An error occurred on the server side.

**message**

The server encountered an internal error trying to fulfill the request.

HTTP Status Code: 500

**ResourceNotFoundException**

The operation tried to access a nonexistent table or index. The resource
might not be specified correctly, or its status might not be
`ACTIVE`.

**message**

The resource which is being requested does not exist.

HTTP Status Code: 400

## Examples

### Describe A Stream

The following sample returns a description of a stream with a given stream ARN.
All of the shards in the stream are listed in the response, along with the beginning and
ending sequence numbers of stream records within the shards. Note that one of the shards
is still open, because it does not have an `EndingSequenceNumber`.

#### Sample Request

```

POST / HTTP/1.1
x-amzn-RequestId: <RequestID>
x-amzn-crc32: <CRC32>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDBStreams_20120810.DescribeStream

{
    "StreamArn": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252"
    "ShardFilter":"{
    "Type": "CHILD_SHARDS",
    "ShardId": "shardId-00000001741631711871-1f6a72cf"
    }"
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
    "StreamDescription": {
        "StreamArn": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252",
        "StreamLabel": "2015-05-20T20:51:10.252",
        "StreamStatus": "ENABLED",
        "StreamViewType": "NEW_AND_OLD_IMAGES",
        "CreationRequestDateTime": "Wed May 20 13:51:10 PDT 2015",
        "TableName": "Forum",
        "KeySchema": [
            {"AttributeName": "ForumName","KeyType": "HASH"},
            {"AttributeName": "Subject","KeyType": "RANGE"}
        ],
        "Shards": [
            {
                "SequenceNumberRange": {
                    "EndingSequenceNumber": "20500000000000000910398",
                    "StartingSequenceNumber": "20500000000000000910398"
                },
                "ShardId": "shardId-00000001414562045508-2bac9cd2"
            },
            {
                "ParentShardId": "shardId-00000001414562045508-2bac9cd2",
                "SequenceNumberRange": {
                    "EndingSequenceNumber": "820400000000000001192334",
                    "StartingSequenceNumber": "820400000000000001192334"
                },
                "ShardId": "shardId-00000001414576573621-f55eea83"
            },
            {
                "ParentShardId": "shardId-00000001414576573621-f55eea83",
                "SequenceNumberRange": {
                    "EndingSequenceNumber": "1683700000000000001135967",
                    "StartingSequenceNumber": "1683700000000000001135967"
                },
                "ShardId": "shardId-00000001414592258131-674fd923"
            },
            {
                "ParentShardId": "shardId-00000001414592258131-674fd923",
                "SequenceNumberRange": {"StartingSequenceNumber": "2574600000000000000935255"},
                "ShardId": "shardId-00000001414608446368-3a1afbaf"
            }
        ],
    }
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for Python](../../../../services/goto/boto3/streams-dynamodb-2012-08-10/describestream.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/describestream.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Amazon DynamoDB Streams

GetRecords

All content copied from https://docs.aws.amazon.com/.
