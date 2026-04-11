# Configurations

This is a collection of configurations. A configuration contains all of the
settings for your broker. For more information, see [Configuration](../developer-guide/configuration.md) and [Amazon MQ Broker Configuration Parameters](../developer-guide/amazon-mq-broker-configuration-parameters.md) in the Amazon MQ Developer
Guide.

You can create a configuration before creating any brokers. You can then apply the
configuration to one or more brokers.

###### Important

Making changes to a configuration does not apply the changes to the broker
immediately. To apply your changes, you must wait for the next maintenance window
or reboot the broker.

## URI

`/v1/configurations`

## HTTP methods

### GET

**Operation ID:** `ListConfigurations`

Returns a list of all configurations.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

ResponsesStatus codeResponse modelDescription`200``

         ListConfigurationsOutput`

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

### POST

**Operation ID:** `CreateConfiguration`

Creates a new configuration for the specified configuration name. Amazon MQ uses
the default configuration (the engine type and version).

ResponsesStatus codeResponse modelDescription`200``

         CreateConfigurationOutput`

HTTP Status Code 200: OK.

`400``Error`

HTTP Status Code 400: Bad request due to incorrect input. Correct your request and
then retry it.

`403``Error`

HTTP Status Code 403: Access forbidden. Correct your credentials and then retry
your request.

`409``Error`

HTTP Status Code 409: Configuration ID is already in use.
Remove the configuration from all brokers and retry the request.

`500``Error`

HTTP Status Code 500: Unexpected internal server error. Retrying your request
might resolve the issue.

### OPTIONS

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Request bodies

```json

{
  "engineVersion": "string",
  "authenticationStrategy": enum,
  "name": "string",
  "engineType": enum,
  "tags": {
  }
}
```

### Response bodies

```json

{
  "nextToken": "string",
  "maxResults": integer,
  "configurations": [
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
  ]
}
```

```json

{
  "created": "string",
  "authenticationStrategy": enum,
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

[AuthenticationStrategy](#configurations-model-authenticationstrategy)

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

[EngineType](#configurations-model-enginetype)

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

[ConfigurationRevision](#configurations-model-configurationrevision)

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

### CreateConfigurationInput

Creates a new configuration for the specified configuration name. Amazon MQ uses
the default configuration (the engine type and version).

PropertyTypeRequiredDescription`authenticationStrategy`

[AuthenticationStrategy](#configurations-model-authenticationstrategy)

False

Optional. The authentication strategy associated with the configuration. The
default is `SIMPLE`.

`engineType`

[EngineType](#configurations-model-enginetype)

True

Required. The type of broker engine. Currently, Amazon MQ supports `ACTIVEMQ` and `RABBITMQ`.

`engineVersion`

string

False

The broker engine version. Defaults to the latest available version for the specified broker engine type. For more information, see the
[ActiveMQ version management](../developer-guide/activemq-version-management.md)
and the [RabbitMQ version management](../developer-guide/rabbitmq-version-management.md)
sections in the Amazon MQ Developer Guide.

`name`

string

True

Required. The name of the configuration. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
1-150 characters long.

`tags`

object

False

Create tags when creating the configuration.

### CreateConfigurationOutput

Returns information about the created configuration.

PropertyTypeRequiredDescription`arn`

string

True

Required. The Amazon Resource Name (ARN) of the configuration.

`authenticationStrategy`

[AuthenticationStrategy](#configurations-model-authenticationstrategy)

True

Optional. The authentication strategy associated with the configuration. The
default is `SIMPLE`.

`created`

string

Format: date-time

True

Required. The date and time of the configuration.

`id`

string

True

Required. The unique ID that Amazon MQ generates for the configuration.

`latestRevision`

[ConfigurationRevision](#configurations-model-configurationrevision)

False

The latest revision of the configuration.

`name`

string

True

Required. The name of the configuration. This value can contain only alphanumeric
characters, dashes, periods, underscores, and tildes (- . \_ ~). This value must be
1-150 characters long.

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

### ListConfigurationsOutput

Returns a list of all configurations.

PropertyTypeRequiredDescription`configurations`

Array of type Configuration

False

The list of all revisions for the specified configuration.

`maxResults`

integer

False

The maximum number of configurations that Amazon MQ can return per page (20 by
default). This value must be an integer from 5 to 100.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### ListConfigurations

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/listconfigurations.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/listconfigurations.md)

### CreateConfiguration

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/createconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/createconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuration Revisions

Tag

All content copied from https://docs.aws.amazon.com/.
