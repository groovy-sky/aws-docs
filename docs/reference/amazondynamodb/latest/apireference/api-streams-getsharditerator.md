# GetShardIterator

Returns a shard iterator. A shard iterator provides information
about how to retrieve the stream records from within a shard. Use
the shard iterator in a subsequent
`GetRecords` request to read the stream records
from the shard.

###### Note

A shard iterator expires 15 minutes after it is returned to the requester.

## Request Syntax

```nohighlight

{
   "SequenceNumber": "string",
   "ShardId": "string",
   "ShardIteratorType": "string",
   "StreamArn": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ShardId](#API_streams_GetShardIterator_RequestSyntax)**

The identifier of the shard. The iterator will be returned for this shard ID.

Type: String

Length Constraints: Minimum length of 28. Maximum length of 65.

Required: Yes

**[ShardIteratorType](#API_streams_GetShardIterator_RequestSyntax)**

Determines how the shard iterator is used to start reading stream records from the shard:

- `AT_SEQUENCE_NUMBER` \- Start reading exactly from the position denoted by a
specific sequence number.

- `AFTER_SEQUENCE_NUMBER` \- Start reading right after the position denoted by a
specific sequence number.

- `TRIM_HORIZON` \- Start reading at the last (untrimmed) stream record, which is
the oldest record in the shard. In DynamoDB Streams, there is a 24 hour limit on data retention.
Stream records whose age exceeds this limit are subject to removal (trimming) from the
stream.

- `LATEST` \- Start reading just after the most recent stream record in the
shard, so that you always read the most recent data in the shard.

Type: String

Valid Values: `TRIM_HORIZON | LATEST | AT_SEQUENCE_NUMBER | AFTER_SEQUENCE_NUMBER`

Required: Yes

**[StreamArn](#API_streams_GetShardIterator_RequestSyntax)**

The Amazon Resource Name (ARN) for the stream.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: Yes

**[SequenceNumber](#API_streams_GetShardIterator_RequestSyntax)**

The sequence number of a stream record in the shard from which to start reading.

Type: String

Length Constraints: Minimum length of 21. Maximum length of 40.

Required: No

## Response Syntax

```nohighlight

{
   "ShardIterator": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[ShardIterator](#API_streams_GetShardIterator_ResponseSyntax)**

The position in the shard from which to start reading stream records sequentially. A shard iterator specifies this position using the sequence number of a stream record in a shard.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

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

### Retrieve a Shard Iterator For a Stream

The following sample returns a shard iterator for the provided stream ARN and
shard ID. The shard iterator will allow access to stream records beginning with the given
sequence number.

#### Sample Request

```

POST / HTTP/1.1
x-amzn-RequestId: <RequestID>
x-amzn-crc32: <CRC32>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDBStreams_20120810.GetShardIterator

{
    "StreamArn": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252",
    "ShardId": "00000001414576573621-f55eea83",
    "ShardIteratorType": "TRIM_HORIZON"
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
    "ShardIterator": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252|1|AAAAAAAAAAEvJp6D+zaQ...  <remaining characters omitted> ..."
}

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for Python](../../../../services/goto/boto3/streams-dynamodb-2012-08-10/getsharditerator.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/getsharditerator.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetRecords

ListStreams

All content copied from https://docs.aws.amazon.com/.
