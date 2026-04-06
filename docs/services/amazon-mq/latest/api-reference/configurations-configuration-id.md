# Configuration

A configuration contains all of the settings for your broker. For more information, see [Amazon MQ for RabbitMQ broker configurations](../developer-guide/rabbitmq-broker-configuration-parameters.md) and
[Amazon MQ for ActiveMQ broker configurations](../developer-guide/amazon-mq-broker-configuration-parameters.md) in the Amazon MQ Developer
Guide.

You can create a configuration before creating any brokers. You can then apply the
configuration to one or more brokers.

###### Important

Making changes to a configuration does not apply the changes to the broker
immediately. To apply your changes, you must wait for the next maintenance window
or reboot the broker.

## URI

`/v1/configurations/configuration-id`

## HTTP methods

### GET

**Operation ID:** `DescribeConfiguration`

Returns information about the specified configuration.

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200``

         Configuration`

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

### PUT

**Operation ID:** `UpdateConfiguration`

Updates the specified configuration.

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200``

         UpdateConfigurationOutput`

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

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### DELETE

**Operation ID:** `DeleteConfiguration`

Deletes the specified configuration.

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200``

         DeleteConfigurationOutput`

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

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Request bodies

```json

{
  "data": "string",
  "description": "string"
}
```

### Response bodies

```json

{
  "engineVersion": "string",
  "created": "string",
  "authenticationStrategy": enum,
  "name": "string",
  "description": "string",
  "engineType": enum,
  "id": "string",
  "arn": "string",
  "latestRevision": {
    "created": "string",
    "description": "string",
    "revision": integer
  },
  "tags": {
  }
}
```

```json

{
  "created": "string",
  "warnings": [
    {
      "reason": enum,
      "attributeName": "string",
      "elementName": "string"
    }
  ],
  "name": "string",
  "id": "string",
  "arn": "string",
  "latestRevision": {
    "created": "string",
    "description": "string",
    "revision": integer
  }
}
```

```json

{
  "configurationId": "string"
}
```

```json

{
  "errorAttribute": "string",
  "message": "string"
}
```

## Properties

### AuthenticationStrategy

Optional. The authentication strategy used to secure the broker. The
default is `SIMPLE`.

- `SIMPLE`

- `LDAP`

- `CONFIG_MANAGED`

### Configuration

Returns information about all configurations.

PropertyTypeRequiredDescription`arn`

string

True

Required. The ARN of the configuration.

`authenticationStrategy`

[AuthenticationStrategy](#configurations-configuration-id-model-authenticationstrategy)

True

Optional. The authentication strategy associated with the configuration. The
default is `SIMPLE`.

`created`

string

Format: date-time

True

Required. The date and time of the configuration revision.

`description`

string

True

Required. The description of the configuration.

`engineType`

[EngineType](#configurations-configuration-id-model-enginetype)

True

Required. The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ`.

`engineVersion`

string

True

The broker engine version. Defaults to the latest available version for the specified broker engine type. For a list of supported engine versions,
see the
[ActiveMQ version management](../developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

`id`

string

True

Required. The unique ID that Amazon MQ generates for the configuration.

`latestRevision`

[ConfigurationRevision](#configurations-configuration-id-model-configurationrevision)

True

Required. The latest revision of the configuration.

`name`

string

True

Required. The name of the configuration. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
1-150 characters long.

`tags`

object

False

The list of all tags associated with this configuration.

### ConfigurationRevision

Returns information about the specified configuration revision.

PropertyTypeRequiredDescription`created`

string

Format: date-time

True

Required. The date and time of the configuration revision.

`description`

string

False

The description of the configuration revision.

`revision`

integer

True

Required. The revision number of the configuration.

### DeleteConfigurationOutput

Returns information about the deleted configuration.

PropertyTypeRequiredDescription`configurationId`

string

False

The unique ID that Amazon MQ generates for the configuration.

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

### SanitizationWarning

Returns information about the configuration element or attribute that was sanitized in the
configuration.

PropertyTypeRequiredDescription`attributeName`

string

False

The name of the configuration attribute that has been sanitized.

`elementName`

string

False

The name of the configuration element that has been sanitized.

`reason`

[SanitizationWarningReason](#configurations-configuration-id-model-sanitizationwarningreason)

True

The reason for which the configuration elements or attributes were
sanitized.

### SanitizationWarningReason

The reason for which the configuration elements or attributes were sanitized.

- `DISALLOWED_ELEMENT_REMOVED`

- `DISALLOWED_ATTRIBUTE_REMOVED`

- `INVALID_ATTRIBUTE_VALUE_REMOVED`

### UpdateConfigurationInput

Updates the specified configuration.

PropertyTypeRequiredDescription`data`

string

Format: byte

True

Amazon MQ for Active MQ: The base64-encoded XML configuration.
Amazon MQ for RabbitMQ: the base64-encoded Cuttlefish configuration.

`description`

string

False

The description of the configuration.

### UpdateConfigurationOutput

Returns information about the updated configuration.

PropertyTypeRequiredDescription`arn`

string

True

The Amazon Resource Name (ARN) of the configuration.

`created`

string

Format: date-time

True

Required. The date and time of the configuration.

`id`

string

True

The unique ID that Amazon MQ generates for the configuration.

`latestRevision`

[ConfigurationRevision](#configurations-configuration-id-model-configurationrevision)

False

The latest revision of the configuration.

`name`

string

True

The name of the configuration. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
1-150 characters long.

`warnings`

Array of type SanitizationWarning

False

The list of the first 20 warnings about the configuration elements or
attributes that were sanitized.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### DescribeConfiguration

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/DescribeConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/DescribeConfiguration)

### UpdateConfiguration

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/UpdateConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/UpdateConfiguration)

### DeleteConfiguration

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/DeleteConfiguration)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/DeleteConfiguration)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Brokers broker-id Promote

Configuration Revision
