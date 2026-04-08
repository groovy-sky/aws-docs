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

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/rebootbroker.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/rebootbroker.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Broker Instance Options

Brokers

All content copied from https://docs.aws.amazon.com/.
