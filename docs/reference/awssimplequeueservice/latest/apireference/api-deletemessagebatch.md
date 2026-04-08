# DeleteMessageBatch

Deletes up to ten messages from the specified queue. This is a batch version of
`
         DeleteMessage.` The result of the action on each
message is reported individually in the response.

###### Important

Because the batch request can result in a combination of successful and unsuccessful actions, you should check for batch errors even when the call returns an HTTP status code of `200`.

## Request Syntax

```nohighlight

{
   "Entries": [
      {
         "Id": "string",
         "ReceiptHandle": "string"
      }
   ],
   "QueueUrl": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Entries](#API_DeleteMessageBatch_RequestSyntax)**

Lists the receipt handles for the messages to be deleted.

Type: Array of [DeleteMessageBatchRequestEntry](api-deletemessagebatchrequestentry.md) objects

Required: Yes

**[QueueUrl](#API_DeleteMessageBatch_RequestSyntax)**

The URL of the Amazon SQS queue from which messages are deleted.

Queue URLs and names are case-sensitive.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "Failed": [
      {
         "Code": "string",
         "Id": "string",
         "Message": "string",
         "SenderFault": boolean
      }
   ],
   "Successful": [
      {
         "Id": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Failed](#API_DeleteMessageBatch_ResponseSyntax)**

A list of `
                     BatchResultErrorEntry
                  ` items.

Type: Array of [BatchResultErrorEntry](api-batchresulterrorentry.md) objects

**[Successful](#API_DeleteMessageBatch_ResponseSyntax)**

A list of `
                     DeleteMessageBatchResultEntry
                  ` items.

Type: Array of [DeleteMessageBatchResultEntry](api-deletemessagebatchresultentry.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Errors](commonerrors.md).

**BatchEntryIdsNotDistinct**

Two or more batch entries in the request have the same `Id`.

HTTP Status Code: 400

**EmptyBatchRequest**

The batch request doesn't contain any entries.

HTTP Status Code: 400

**InvalidAddress**

The specified ID is invalid.

HTTP Status Code: 400

**InvalidBatchEntryId**

The `Id` of a batch entry in a batch request doesn't abide by the
specification.

HTTP Status Code: 400

**InvalidSecurity**

The request was not made over HTTPS or did not use SigV4 for signing.

HTTP Status Code: 400

**QueueDoesNotExist**

Ensure that the `QueueUrl` is correct and that the queue has not been
deleted.

HTTP Status Code: 400

**RequestThrottled**

The request was denied due to request throttling.

- Exceeds the permitted request rate for the queue or for the recipient of the
request.

- Ensure that the request rate is within the Amazon SQS limits for
sending messages. For more information, see [Amazon SQS quotas](../../../../services/awssimplequeueservice/latest/sqsdeveloperguide/sqs-quotas.md#quotas-requests) in the _Amazon SQS_
_Developer Guide_.

HTTP Status Code: 400

**TooManyEntriesInBatchRequest**

The batch request contains more entries than permissible. For Amazon SQS, the
maximum number of entries you can include in a single [SendMessageBatch](api-sendmessagebatch.md), [DeleteMessageBatch](api-deletemessagebatch.md), or [ChangeMessageVisibilityBatch](api-changemessagevisibilitybatch.md) request is 10.

HTTP Status Code: 400

**UnsupportedOperation**

Error code 400. Unsupported operation.

HTTP Status Code: 400

## Examples

The following examples show how a `DeleteMessageBatch` request deletes
two messages. You must URL-encode the entire URL. However, in this example only the message body is URL-encoded to make the example easier to read. The structure of `AUTHPARAMS` depends on the signature of the API request.
For more information, see [Examples of Signed Signature Version 4 Requests](../../../../general/general/latest/gr/sigv4-signed-request-examples.md) in the _AWS General Reference_.

### Example

**Using AWS JSON protocol**
**(Default)**

#### Sample Request

```json

POST / HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
X-Amz-Target: AmazonSQS.DeleteMessageBatch
X-Amz-Date: <Date>
Content-Type: application/x-amz-json-1.0
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
{
    "QueueUrl": "https://sqs.us-east-1.amazonaws.com/177715257436/MyQueue/",
    "Entries": [
        {
            "Id": "msg1",
            "ReceiptHandle": "AQEBaZ+j5qUoOAoxlmrCQPkBm9njMWXqemmIG6shMHCO6fV20JrQYg/AiZ8JELwLwOu5U61W+aIX5Qzu7GGofxJuvzymr4Ph53RiR0mudj4InLSgpSspYeTRDteBye5tV/txbZDdNZxsi+qqZA9xPnmMscKQqF6pGhnGIKrnkYGl45Nl6GPIZv62LrIRb6mSqOn1fn0yqrvmWuuY3w2UzQbaYunJWGxpzZze21EOBtywknU3Je/g7G9is+c6K9hGniddzhLkK1tHzZKjejOU4jokaiB4nmi0dF3JqLzDsQuPF0Gi8qffhEvw56nl8QCbluSJScFhJYvoagGnDbwOnd9z50L239qtFIgETdpKyirlWwl/NGjWJ45dqWpiW3d2Ws7q"
        },
        {
            "Id": "msg2",
            "ReceiptHandle": "AQEB3LQoW7GQWgodQCEJXHjMvO/QkeCHiRldRfLC/E6RUggm+BjpthqxfoUOUn6Vs271qmrBaufFqEmnMKgk2n1EuUBne1pe+hZcrDE8IveUUPmqkUT54FGhAAjPX3oEIryz/XeQ/muKAuLclcZvt2Q+ZDPW8DvZqMa1RoHxOqSq+6kQ4PwgQxB+VqDYvIc/LpHOoL4PTROBXgLPjWrzz/knK6HTzKpqC4ESvFdJ/dkk2nvS0iqYOly5VQknK/lv/rTUOgEYevjJSrNLIPDgZGyvgcLwbm6+yo1cW/c9cPDiVm96gIhVkuiCZ1gtskoOtyroZVPcY71clDG2EPZJeY8akMd3u+sXEMWhiOPFs1cgWQs2ugsL+vdwMCbsZRkXbJv7"
        }
    ]
}
```

#### Sample Response

```json

HTTP/1.1 200 OK
x-amzn-RequestId: <requestId>
Content-Length: <>
Date: <Date>
Content-Type: application/x-amz-json-1.0
{
    "Failed": [],
    "Successful": [
        {
            "Id": "msg2"
        },
        {
            "Id": "msg1"
        }
    ]
}
```

### Example

**Using AWS query**
**protocol**

#### Sample Request

```xml

POST /177715257436/MyQueue/ HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
Content-Type: application/x-www-form-urlencoded
X-Amz-Date: <Date>
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
Action=DeleteMessageBatch
&DeleteMessageBatchRequestEntry.1.Id=msg1
&DeleteMessageBatchRequestEntry.1.ReceiptHandle=gfk0T0R0waama4fVFffkjPQrrvzMrOg0fTFk2LxT33EuB8wR0ZCFgKWyXGWFoqqpCIiprQUEhir%2F5LeGPpYTLzjqLQxyQYaQALeSNHb0us3uE84uujxpBhsDkZUQkjFFkNqBXn48xlMcVhTcI3YLH%2Bd%2BIqetIOHgBCZAPx6r%2B09dWaBXei6nbK5Ygih21DCDdAwFV68Jo8DXhb3ErEfoDqx7vyvC5nCpdwqv%2BJhU%2FTNGjNN8t51v5c%2FAXvQsAzyZVNapxUrHIt4NxRhKJ72uICcxruyE8eRXlxIVNgeNP8ZEDcw7zZU1Zw%3D%3D
&DeleteMessageBatchRequestEntry.2.Id=msg2
&DeleteMessageBatchRequestEntry.2.ReceiptHandle=gfk0T0R0waama4fVFffkjKzmhMCymjQvfTFk2LxT33G4ms5subrE0deLKWSscPU1oD3J9zgeS4PQQ3U30qOumIE6AdAv3w%2F%2Fa1IXW6AqaWhGsEPaLm3Vf6IiWqdM8u5imB%2BNTwj3tQRzOWdTOePjOjPcTpRxBtXix%2BEvwJOZUma9wabv%2BSw6ZHjwmNcVDx8dZXJhVp16Bksiox%2FGrUvrVTCJRTWTLc59oHLLF8sEkKzRmGNzTDGTiV%2BYjHfQj60FD3rVaXmzTsoNxRhKJ72uIHVMGVQiAGgB%2BqAbSqfKHDQtVOmJJgkHug%3D%3D
```

#### Sample Response

```xml

HTTP/1.1 200 OK
<?xml version="1.0"?>
<DeleteMessageBatchResponse xmlns="http://queue.amazonaws.com/doc/2012-11-05/">
    <DeleteMessageBatchResult>
        <DeleteMessageBatchResultEntry>
            <Id>msg1</Id>
        </DeleteMessageBatchResultEntry>
        <DeleteMessageBatchResultEntry>
            <Id>msg2</Id>
        </DeleteMessageBatchResultEntry>
    </DeleteMessageBatchResult>
    <ResponseMetadata>
        <RequestId>d6f86b7a-74d1-4439-b43f-196a1e29cd85</RequestId>
    </ResponseMetadata>
</DeleteMessageBatchResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for Python](../../../../services/goto/boto3/sqs-2012-11-05/deletemessagebatch.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/sqs-2012-11-05/deletemessagebatch.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteMessage

DeleteQueue
