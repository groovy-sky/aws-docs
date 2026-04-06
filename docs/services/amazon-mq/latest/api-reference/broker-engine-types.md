# Broker Engine Types

Retrieve information about available broker engines. AWS does not support all
instance types in all availability zones and regions. For more information, see the
[ActiveMQ version management](../developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

This API will tell you, for a given region and availability zone, which broker
engine types and engine versions you can create.

## URI

`/v1/broker-engine-types`

## HTTP methods

### GET

**Operation ID:** `DescribeBrokerEngineTypes`

Describe available engine types and versions.

Query parametersNameTypeRequiredDescription`engineType`StringFalse

Filter response by engine type.

`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

ResponsesStatus codeResponse modelDescription`200``

         BrokerEngineTypeOutput`

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Response bodies

```json

{
  "nextToken": "string",
  "maxResults": integer,
  "brokerEngineTypes": [
    {
      "engineVersions": [
        {
          "name": "string"
        }
      ],
      "engineType": enum
    }
  ]
}
```

```json

{
  "errorAttribute": "string",
  "message": "string"
}
```

## Properties

### BrokerEngineType

Types of broker engines.

PropertyTypeRequiredDescription`engineType`

[EngineType](#broker-engine-types-model-enginetype)

False

The broker's engine type.

`engineVersions`

Array of type EngineVersion

False

The list of engine versions.

### BrokerEngineTypeOutput

Returns a list of broker engine type.

PropertyTypeRequiredDescription`brokerEngineTypes`

Array of type BrokerEngineType

False

List of available engine types and versions.

`maxResults`

integer

Minimum: 5

Maximum: 100

True

Required. The maximum number of engine types that can be returned per page (20 by
default). This value must be an integer from 5 to 100.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

### EngineType

The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.

- `ACTIVEMQ`

- `RABBITMQ`

### EngineVersion

Id of the engine version.

PropertyTypeRequiredDescription`name`

string

False

Id for the version.

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

### DescribeBrokerEngineTypes

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/DescribeBrokerEngineTypes)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/DescribeBrokerEngineTypes)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Broker

Broker Instance Options
