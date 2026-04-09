# ListStreams

Returns an array of stream ARNs associated with the current account and endpoint. If the
`TableName` parameter is present, then `ListStreams` will return only the
streams ARNs for that table.

###### Note

You can call `ListStreams` at a maximum rate of 5 times per second.

## Request Syntax

```nohighlight

{
   "ExclusiveStartStreamArn": "string",
   "Limit": number,
   "TableName": "string"
}
```

## Request Parameters

The request accepts the following data in JSON format.

###### Note

In the following list, the required parameters are described first.

**[ExclusiveStartStreamArn](#API_streams_ListStreams_RequestSyntax)**

The ARN (Amazon Resource Name) of the first item that this operation will evaluate. Use the
value that was returned for `LastEvaluatedStreamArn` in the previous operation.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

Required: No

**[Limit](#API_streams_ListStreams_RequestSyntax)**

The maximum number of streams to return. The upper limit is 100.

Type: Integer

Valid Range: Minimum value of 1.

Required: No

**[TableName](#API_streams_ListStreams_RequestSyntax)**

If this parameter is provided, then only the streams associated with this table name are returned.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Pattern: `[a-zA-Z0-9_.-]+`

Required: No

## Response Syntax

```nohighlight

{
   "LastEvaluatedStreamArn": "string",
   "Streams": [
      {
         "StreamArn": "string",
         "StreamLabel": "string",
         "TableName": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[LastEvaluatedStreamArn](#API_streams_ListStreams_ResponseSyntax)**

The stream ARN of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request.

If `LastEvaluatedStreamArn` is empty, then the "last page" of results has been
processed and there is no more data to be retrieved.

If `LastEvaluatedStreamArn` is not empty, it does not necessarily mean that there
is more data in the result set. The only way to know when you have reached the end of the
result set is when `LastEvaluatedStreamArn` is empty.

Type: String

Length Constraints: Minimum length of 37. Maximum length of 1024.

**[Streams](#API_streams_ListStreams_ResponseSyntax)**

A list of stream descriptors associated with the current account and endpoint.

Type: Array of [Stream](api-streams-stream.md) objects

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

### Retrieve All Stream ARNs

The following sample returns all of the stream ARNs.

#### Sample Request

```

POST / HTTP/1.1
x-amzn-RequestId: <RequestID>
x-amzn-crc32: <CRC32>
Content-Type: application/x-amz-json-1.0
Content-Length: <PayloadSizeBytes>
X-Amz-Date: <Date>
X-Amz-Target: DynamoDBStreams_20120810.ListStreams

{}

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
    "Streams": [
        {
            "StreamArn": "arn:aws:dynamodb:us-wesst-2:111122223333:table/Forum/stream/2015-05-20T20:51:10.252",
            "TableName": "Forum",
            "StreamLabel": "2015-05-20T20:51:10.252"
        },
        {
            "StreamArn": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-20T20:50:02.714",
            "TableName": "Forum",
            "StreamLabel": "2015-05-20T20:50:02.714"
        },
        {
            "StreamArn": "arn:aws:dynamodb:us-west-2:111122223333:table/Forum/stream/2015-05-19T23:03:50.641",
            "TableName": "Forum",
            "StreamLabel": "2015-05-19T23:03:50.641"
        },

        ...remaining output omitted...

    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for Python](../../../../services/goto/boto3/streams-dynamodb-2012-08-10/liststreams.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/streams-dynamodb-2012-08-10/liststreams.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetShardIterator

Data Types

All content copied from https://docs.aws.amazon.com/.
