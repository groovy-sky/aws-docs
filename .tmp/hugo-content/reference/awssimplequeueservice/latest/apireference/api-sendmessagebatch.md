# SendMessageBatch

You can use `SendMessageBatch` to send up to 10 messages to the specified
queue by assigning either identical or different values to each message (or by not
assigning values at all). This is a batch version of `
         SendMessage.` For a FIFO queue, multiple messages within a single batch are enqueued
in the order they are sent.

The result of sending each message is reported individually in the response.
Because the batch request can result in a combination of successful and unsuccessful actions, you should check for batch errors even when the call returns an HTTP status code of `200`.

The maximum allowed individual message size and the maximum total payload size (the
sum of the individual lengths of all of the batched messages) are both 1 MiB
1,048,576 bytes.

###### Important

A message can include only XML, JSON, and unformatted text. The following Unicode characters are allowed. For more information, see the [W3C specification for characters](http://www.w3.org/TR/REC-xml).

`#x9` \| `#xA` \| `#xD` \| `#x20` to `#xD7FF` \| `#xE000` to `#xFFFD` \| `#x10000` to `#x10FFFF`

If a message contains characters outside the allowed set, Amazon SQS rejects the message and returns an InvalidMessageContents error. Ensure that your message body includes only valid characters to avoid this exception.

If you don't specify the `DelaySeconds` parameter for an entry, Amazon SQS uses
the default value for the queue.

## Request Syntax

```nohighlight

{
   "Entries": [
      {
         "DelaySeconds": number,
         "Id": "string",
         "MessageAttributes": {
            "string" : {
               "BinaryListValues": [ blob ],
               "BinaryValue": blob,
               "DataType": "string",
               "StringListValues": [ "string" ],
               "StringValue": "string"
            }
         },
         "MessageBody": "string",
         "MessageDeduplicationId": "string",
         "MessageGroupId": "string",
         "MessageSystemAttributes": {
            "string" : {
               "BinaryListValues": [ blob ],
               "BinaryValue": blob,
               "DataType": "string",
               "StringListValues": [ "string" ],
               "StringValue": "string"
            }
         }
      }
   ],
   "QueueUrl": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[Entries](#API_SendMessageBatch_RequestSyntax)**

A list of `
                     SendMessageBatchRequestEntry
                  ` items.

Type: Array of [SendMessageBatchRequestEntry](api-sendmessagebatchrequestentry.md) objects

Required: Yes

**[QueueUrl](#API_SendMessageBatch_RequestSyntax)**

The URL of the Amazon SQS queue to which batched messages are sent.

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
         "Id": "string",
         "MD5OfMessageAttributes": "string",
         "MD5OfMessageBody": "string",
         "MD5OfMessageSystemAttributes": "string",
         "MessageId": "string",
         "SequenceNumber": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[Failed](#API_SendMessageBatch_ResponseSyntax)**

A list of `
                     BatchResultErrorEntry
                  ` items with error
details about each message that can't be enqueued.

Type: Array of [BatchResultErrorEntry](api-batchresulterrorentry.md) objects

**[Successful](#API_SendMessageBatch_ResponseSyntax)**

A list of `
                     SendMessageBatchResultEntry
                  ` items.

Type: Array of [SendMessageBatchResultEntry](api-sendmessagebatchresultentry.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Errors](commonerrors.md).

**BatchEntryIdsNotDistinct**

Two or more batch entries in the request have the same `Id`.

HTTP Status Code: 400

**BatchRequestTooLong**

The length of all the messages put together is more than the limit.

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

**KmsAccessDenied**

The caller doesn't have the required KMS access.

HTTP Status Code: 400

**KmsDisabled**

The request was denied due to request throttling.

HTTP Status Code: 400

**KmsInvalidKeyUsage**

The request was rejected for one of the following reasons:

- The KeyUsage value of the KMS key is incompatible with the API
operation.

- The encryption algorithm or signing algorithm specified for the operation is
incompatible with the type of key material in the KMS key (KeySpec).

HTTP Status Code: 400

**KmsInvalidState**

The request was rejected because the state of the specified resource is not valid for
this request.

HTTP Status Code: 400

**KmsNotFound**

The request was rejected because the specified entity or resource could not be found.

HTTP Status Code: 400

**KmsOptInRequired**

The request was rejected because the specified key policy isn't syntactically or
semantically correct.

HTTP Status Code: 400

**KmsThrottled**

AWS KMS throttles requests for the following conditions.

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

The following example `SendMessageBatch` request sends two messages to
the queue. You must URL-encode the entire URL. However, in this example only the message body is URL-encoded to make the example easier to read. The structure of `AUTHPARAMS` depends on the signature of the API request.
For more information, see [Examples of Signed Signature Version 4 Requests](../../../../general/general/latest/gr/sigv4-signed-request-examples.md) in the _AWS General Reference_.

### Example

**Using AWS JSON protocol**
**(Default)**

#### Sample Request

```json

POST / HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
X-Amz-Target: AmazonSQS.SendMessageBatch
X-Amz-Date: <Date>
Content-Type: application/x-amz-json-1.0
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
{
    "QueueUrl": "https://sqs.us-east-1.amazonaws.com/177715257436/MyQueue/",
    "Entries": [
        {
            "Id": "test_msg_001",
            "MessageBody": "test message body 1"
        },
        {
            "Id": "test_msg_002",
            "MessageBody": "test message body 2",
            "DelaySeconds": 60,
            "MessageAttributes": {
                "my_attribute_name_1": {
                    "DataType": "String",
                    "StringValue": "my_attribute_value_1"
                }
            }
        }
    ]
}
```

#### Sample Response

```json

HTTP/1.1 200 OK
x-amzn-RequestId: <requestId>
Content-Length: <PayloadSizeBytes>
Date: <Date>
Content-Type: application/x-amz-json-1.0
{
   "Failed": [],
    "Successful": [
        {
            "Id": "test_msg_001",
            "MD5OfMessageBody": "0e024d309850c78cba5eabbeff7cae71",
            "MessageId": "f4eb349f-cd33-4bc4-bdc2-e557c900d41d"
        },
        {
            "Id": "test_msg_002",
            "MD5OfMessageAttributes": "8ef4d60dbc8efda9f260e1dfd09d29f3",
            "MD5OfMessageBody": "27118326006d3829667a400ad23d5d98",
            "MessageId": "1dcfcd50-5a67-45ae-ae4c-1c152b5effb9"
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
X-Amz-Date: <Date>
Content-Type: application/x-www-form-urlencoded
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
Action=SendMessageBatch
&SendMessageBatchRequestEntry.1.Id=test_msg_001
&SendMessageBatchRequestEntry.1.MessageBody=test%20message%20body%201
&SendMessageBatchRequestEntry.2.Id=test_msg_002
&SendMessageBatchRequestEntry.2.MessageBody=test%20message%20body%202
&SendMessageBatchRequestEntry.2.DelaySeconds=60
&SendMessageBatchRequestEntry.2.MessageAttribute.1.Name=test_attribute_name_1
&SendMessageBatchRequestEntry.2.MessageAttribute.1.Value.StringValue=test_attribute_value_1
&SendMessageBatchRequestEntry.2.MessageAttribute.1.Value.DataType=String
```

#### Sample Response

```xml

HTTP/1.1 200 OK
<?xml version="1.0"?>
<SendMessageBatchResponse xmlns="http://queue.amazonaws.com/doc/2012-11-05/">
    <SendMessageBatchResult>
        <SendMessageBatchResultEntry>
            <Id>test_msg_001</Id>
            <MessageId>60e827c3-c8a5-410a-af0e-fb43746e70b1</MessageId>
            <MD5OfMessageBody>0e024d309850c78cba5eabbeff7cae71</MD5OfMessageBody>
        </SendMessageBatchResultEntry>
        <SendMessageBatchResultEntry>
            <Id>test_msg_00</Id>
            <MessageId>c6e7fc6a-b802-4724-be06-59833004578b</MessageId>
            <MD5OfMessageBody>7fb8146a82f95e0af155278f406862c2</MD5OfMessageBody>
            <MD5OfMessageAttributes>ba056227cfd9533dba1f72ad9816d233</MD5OfMessageAttributes>
        </SendMessageBatchResultEntry>
    </SendMessageBatchResult>
    <ResponseMetadata>
        <RequestId>5150a701-14f7-5609-b136-fb71a0ca744a</RequestId>
    </ResponseMetadata>
</SendMessageBatchResponse>
```

### Example

The following example sends multiple messages with _message_
_timers_—applying a visibility delay of variable length to
the messages in the batch—by calling the `SendMessageBatch`
action _without_ a value for `DelaySeconds` for
the first message and with the values of 45 seconds and 2 minutes for the second
and third messages.

###### Note

If you don't set a value for the `DelaySeconds` parameter, the
message might still be subject to a delay if you add the message to a
_delay queue_. For more information about using delay
queues, see [Amazon SQS Delay Queues](../../../../services/awssimplequeueservice/latest/sqsdeveloperguide/sqs-delay-queues.md) in the _Amazon SQS Developer_
_Guide_.

**Using AWS JSON protocol**
**(Default)**

#### Sample Request

```xml

POST / HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
X-Amz-Target: AmazonSQS.SendMessageBatch
X-Amz-Date: <Date>
Content-Type: application/x-amz-json-1.0
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
{
    "QueueUrl": "https://sqs.us-east-1.amazonaws.com/177715257436/MyQueue/",
    "Entries": [
        {
            "Id": "test_msg_no_message_timer",
            "MessageBody": "test message body 1"
        },
        {
            "Id": "test_msg_delay_45_seconds",
            "MessageBody": "test message body 2",
            "DelaySeconds": 45
        },
        {
            "Id": "test_msg_delay_2_minutes",
            "MessageBody": "test message body 3",
            "DelaySeconds": 120
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
X-Amz-Date: <Date>
Content-Type: application/x-www-form-urlencoded
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
Action=SendMessageBatch
&SendMessageBatchRequestEntry.1.Id=test_msg_no_message_timer
&SendMessageBatchRequestEntry.1.MessageBody=test%20message%20body%201
&SendMessageBatchRequestEntry.2.Id=test_msg_delay_45_seconds
&SendMessageBatchRequestEntry.2.MessageBody=test%20message%20body%202
&SendMessageBatchRequestEntry.2.DelaySeconds=45
&SendMessageBatchRequestEntry.3.Id=test_msg_delay_2_minutes
&SendMessageBatchRequestEntry.3.MessageBody=test%20message%20body%203
&SendMessageBatchRequestEntry.3.DelaySeconds=120
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for Python](../../../../services/goto/boto3/sqs-2012-11-05/sendmessagebatch.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/sqs-2012-11-05/sendmessagebatch.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SendMessage

SetQueueAttributes
