# Broker Instance Options

You can retrieve information about broker instances. For more information about the different components of an Amazon MQ broker, see
[How Amazon MQ works](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/amazon-mq-how-it-works.html) in the _Amazon MQ Developer Guide_.

## URI

`/v1/broker-instance-options`

## HTTP methods

### GET

**Operation ID:** `DescribeBrokerInstanceOptions`

Describe available broker instance options.

Query parametersNameTypeRequiredDescription`hostInstanceType`StringFalse

Filter response by host instance type.

`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`storageType`StringFalse

Filter response by storage type.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

`engineType`StringFalse

Filter response by engine type.

ResponsesStatus codeResponse modelDescription`200``

         BrokerInstanceOptionsOutput`

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
  "brokerInstanceOptions": [
    {
      "supportedDeploymentModes": [
        enum
      ],
      "supportedEngineVersions": [
        "string"
      ],
      "storageType": enum,
      "engineType": enum,
      "availabilityZones": [
        {
          "name": "string"
        }
      ],
      "hostInstanceType": "string"
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

### AvailabilityZone

Name of the availability zone.

PropertyTypeRequiredDescription`name`

string

False

Id for the availability zone.

### BrokerInstanceOption

Option for host instance type.

PropertyTypeRequiredDescription`availabilityZones`

Array of type AvailabilityZone

False

The list of available az.

`engineType`

[EngineType](#broker-instance-options-model-enginetype)

False

The broker's engine type.

`hostInstanceType`

string

False

The broker's instance type.

`storageType`

[BrokerStorageType](#broker-instance-options-model-brokerstoragetype)

False

The broker's storage type.

`supportedDeploymentModes`

Array of type DeploymentMode

False

The list of supported deployment modes.

`supportedEngineVersions`

Array of type string

False

The list of supported engine versions.

### BrokerInstanceOptionsOutput

Returns a list of broker instance options.

PropertyTypeRequiredDescription`brokerInstanceOptions`

Array of type BrokerInstanceOption

False

List of available broker instance options.

`maxResults`

integer

Minimum: 5

Maximum: 100

True

Required. The maximum number of instance options that can be returned per page (20
by default). This value must be an integer from 5 to 100.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

### BrokerStorageType

The broker's storage type.

###### Important

`EFS` is not supported for RabbitMQ engine type.

- `EBS`

- `EFS`

### DeploymentMode

The broker's deployment mode.

- `SINGLE_INSTANCE`

- `ACTIVE_STANDBY_MULTI_AZ`

- `CLUSTER_MULTI_AZ`

### EngineType

The type of broker engine. Amazon MQ supports ActiveMQ and RabbitMQ.

- `ACTIVEMQ`

- `RABBITMQ`

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

### DescribeBrokerInstanceOptions

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/DescribeBrokerInstanceOptions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/DescribeBrokerInstanceOptions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Broker Engine Types

Broker Reboot
