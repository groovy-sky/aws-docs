# ChangeMessageVisibility

Changes the visibility timeout of a specified message in a queue to a new value. The
default visibility timeout for a message is 30 seconds. The minimum is 0 seconds. The
maximum is 12 hours. For more information, see [Visibility Timeout](../../../../services/awssimplequeueservice/latest/sqsdeveloperguide/sqs-visibility-timeout.md) in the _Amazon SQS Developer_
_Guide_.

For example, if the default timeout for a queue is 60 seconds, 15 seconds have elapsed
since you received the message, and you send a ChangeMessageVisibility call with
`VisibilityTimeout` set to 10 seconds, the 10 seconds begin to count from
the time that you make the `ChangeMessageVisibility` call. Thus, any attempt
to change the visibility timeout or to delete that message 10 seconds after you
initially change the visibility timeout (a total of 25 seconds) might result in an
error.

An Amazon SQS message has three basic states:

1. Sent to a queue by a producer.

2. Received from the queue by a consumer.

3. Deleted from the queue.

A message is considered to be _stored_ after it is sent to a queue by a producer, but not yet received from the queue by a consumer (that is, between states 1 and 2). There is no limit to the number of stored messages.
A message is considered to be _in flight_ after it is received from a queue by a consumer, but not yet deleted from the queue (that is, between states 2 and 3). There is a limit to the number of in flight messages.

Limits that apply to in flight messages are unrelated to the _unlimited_ number of stored messages.

For most standard queues (depending on queue traffic and message backlog), there can be a maximum of approximately 120,000 in flight messages (received from a queue by a consumer, but not yet deleted from the queue).
If you reach this limit, Amazon SQS returns the `OverLimit` error message.
To avoid reaching the limit, you should delete messages from the queue after they're processed. You can also increase the number of queues you use to process your messages.
To request a limit increase, [file a support request](https://console.aws.amazon.com/support/home).

For FIFO queues, there can be a maximum of 120,000 in flight messages (received from a queue by a consumer, but not yet deleted from the queue). If you reach this limit, Amazon SQS returns no error messages.

###### Important

If you attempt to set the `VisibilityTimeout` to a value greater than
the maximum time left, Amazon SQS returns an error. Amazon SQS doesn't automatically
recalculate and increase the timeout to the maximum remaining time.

Unlike with a queue, when you change the visibility timeout for a specific message
the timeout value is applied immediately but isn't saved in memory for that message.
If you don't delete a message after it is received, the visibility timeout for the
message reverts to the original timeout value (not to the value you set using the
`ChangeMessageVisibility` action) the next time the message is
received.

## Request Syntax

```nohighlight

{
   "QueueUrl": "string",
   "ReceiptHandle": "string",
   "VisibilityTimeout": number
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[QueueUrl](#API_ChangeMessageVisibility_RequestSyntax)**

The URL of the Amazon SQS queue whose message's visibility is changed.

Queue URLs and names are case-sensitive.

Type: String

Required: Yes

**[ReceiptHandle](#API_ChangeMessageVisibility_RequestSyntax)**

The receipt handle associated with the message, whose visibility timeout is changed.
This parameter is returned by the `
                     ReceiveMessage
                  `
action.

Type: String

Required: Yes

**[VisibilityTimeout](#API_ChangeMessageVisibility_RequestSyntax)**

The new value for the message's visibility timeout (in seconds). Values range:
`0` to `43200`. Maximum: 12 hours.

Type: Integer

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Errors](commonerrors.md).

**InvalidAddress**

The specified ID is invalid.

HTTP Status Code: 400

**InvalidSecurity**

The request was not made over HTTPS or did not use SigV4 for signing.

HTTP Status Code: 400

**MessageNotInflight**

The specified message isn't in flight.

HTTP Status Code: 400

**QueueDoesNotExist**

Ensure that the `QueueUrl` is correct and that the queue has not been
deleted.

HTTP Status Code: 400

**ReceiptHandleIsInvalid**

The specified receipt handle isn't valid.

HTTP Status Code: 400

**RequestThrottled**

The request was denied due to request throttling.

- Exceeds the permitted request rate for the queue or for the recipient of the
request.

- Ensure that the request rate is within the Amazon SQS limits for
sending messages. For more information, see [Amazon SQS quotas](../../../../services/awssimplequeueservice/latest/sqsdeveloperguide/sqs-quotas.md#quotas-requests) in the _Amazon SQS_
_Developer Guide_.

HTTP Status Code: 400

**UnsupportedOperation**

Error code 400. Unsupported operation.

HTTP Status Code: 400

## Examples

The following example queries request changes the visibility timeout for a message
to 60 seconds. The structure of `AUTHPARAMS` depends on the signature of the API request.
For more information, see [Examples of Signed Signature Version 4 Requests](../../../../general/general/latest/gr/sigv4-signed-request-examples.md) in the _AWS General Reference_.

### Example

**Using AWS JSON protocol**
**(Default)**

#### Sample Request

```json

POST / HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
X-Amz-Target: AmazonSQS.ChangeMessageVisibility
X-Amz-Date: <Date>
Content-Type: application/x-amz-json-1.0
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
{
    "QueueUrl": "https://sqs.us-east-1.amazonaws.com/177715257436/MyQueue/",
    "ReceiptHandle": "AQEBaZ+j5qUoOAoxlmrCQPkBm9njMWXqemmIG6shMHCO6fV20JrQYg/AiZ8JELwLwOu5U61W+aIX5Qzu7GGofxJuvzymr4Ph53RiR0mudj4InLSgpSspYeTRDteBye5tV/txbZDdNZxsi+qqZA9xPnmMscKQqF6pGhnGIKrnkYGl45Nl6GPIZv62LrIRb6mSqOn1fn0yqrvmWuuY3w2UzQbaYunJWGxpzZze21EOBtywknU3Je/g7G9is+c6K9hGniddzhLkK1tHzZKjejOU4jokaiB4nmi0dF3JqLzDsQuPF0Gi8qffhEvw56nl8QCbluSJScFhJYvoagGnDbwOnd9z50L239qtFIgETdpKyirlWwl/NGjWJ45dqWpiW3d2Ws7q",
    "VisibilityTimeout": 60
}
```

#### Sample Response

```json

HTTP/1.1 200 OK
x-amzn-RequestId: <requestId>
Content-Length: 0
Date: <Date>
Content-Type: application/x-amz-json-1.0
```

### Example

**Using AWS query**
**protocol**

#### Sample Request

```xml

POST /177715257436/MyQueue HTTP/1.1
Host: sqs.us-east-1.amazonaws.com
X-Amz-Date: <Date>
Content-Type: application/x-www-form-urlencoded
Authorization: <AuthParams>
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
Action=ChangeMessageVisibility
&VisibilityTimeout=60
&ReceiptHandle=AQEBwPTK2fT2gy97H1iyU5in9umgT+Y4IOxyKGOzpZa8iemEqoR5/aPn0xAodmiVTzyrW7S4e8XwcWbB04XK92jIQzUpiGwRFA4Dl7r3GOw84Qzq/0OBQe/JaKxJw6iilafYA5fo1SJQo5Wg8xXbJHTVlJqgvTXd/UtlByLMhWMi0JMra1UUjYiPsGtYUpLVnOaRkYSPvzRnFFYUbcqCW9lm2BijQKK6KNOZyCCfIh8TooE5i4P2L9N3o9yUHwMdv6p0nb5lKaGurQ2sJwwsyhXf38ZHnVN6pWwsqQnWKYuEXpxPofxd2lcLdgUurMpydS22DzCrkAaf6gmrdxbmCAoeQxE0sFf8alwX9yQmcOjny9aLGe7ro4Vl5o5KMr5hHM4vHEyhwi4wHeKM6MGX0vATA==
```

#### Sample Response

```xml

HTTP/1.1 200 OK
<?xml version="1.0"?>
<ChangeMessageVisibilityResponse xmlns="http://queue.amazonaws.com/doc/2012-11-05/">
    <ResponseMetadata>
        <RequestId>6a7a282a-d013-4a59-aba9-335b0fa48bed</RequestId>
    </ResponseMetadata>
</ChangeMessageVisibilityResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for Python](../../../../services/goto/boto3/sqs-2012-11-05/changemessagevisibility.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/sqs-2012-11-05/changemessagevisibility.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelMessageMoveTask

ChangeMessageVisibilityBatch
