# Broker Reboot

To apply a new configuration to a broker or to initiate a failover in a multi-AZ broker, you can reboot the broker. In addition,
if your broker becomes unresponsive, you can reboot it to recover from a faulty
state.

###### Note

You can reboot only a broker with the RUNNING status.

## URI

`/v1/brokers/broker-id/reboot`

## HTTP methods

### POST

**Operation ID:** `RebootBroker`

Reboots a broker. Note: This API is asynchronous.

Path parametersNameTypeRequiredDescription`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`404``Error`

HTTP Status Code 404: Resource not found due to incorrect input. Correct your
request and then retry it.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

Path parametersNameTypeRequiredDescription`broker-id`StringTrue

The unique ID that Amazon MQ generates for the broker.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Response bodies

```json

{
  "errorAttribute": "string",
  "message": "string"
}
```

## Properties

### Error

Returns information about an error.

PropertyTypeRequiredDescription`errorAttribute`

string

False

The attribute which caused the error.

`message`

string

False

The explanation of the error.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### RebootBroker

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/RebootBroker)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/RebootBroker)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/RebootBroker)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/RebootBroker)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/RebootBroker)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/RebootBroker)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/RebootBroker)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/RebootBroker)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/RebootBroker)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/RebootBroker)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Broker Instance Options

Brokers
