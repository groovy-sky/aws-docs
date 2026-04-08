# Configuration Revision

To keep track of the changes you make to your configuration, you can create
configuration revisions. For more information, see [Configuration](../developer-guide/configuration.md) in the Amazon MQ Developer Guide.

###### Important

Making changes to a configuration does not apply the changes to the broker
immediately. To apply your changes, you must wait for the next maintenance window
or reboot the broker.

## URI

`/v1/configurations/configuration-id/revisions/configuration-revision`

## HTTP methods

### GET

**Operation ID:** `DescribeConfigurationRevision`

Returns the specified configuration revision for the specified
configuration.

Path parametersNameTypeRequiredDescription`configuration-revision`StringTrue

The revision of the configuration.

`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200``

         DescribeConfigurationRevisionOutput`

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

Path parametersNameTypeRequiredDescription`configuration-revision`StringTrue

The revision of the configuration.

`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Response bodies

```json

{
  "data": "string",
  "created": "string",
  "description": "string",
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

### DescribeConfigurationRevisionOutput

Returns the specified configuration revision for the specified
configuration.

PropertyTypeRequiredDescription`configurationId`

string

True

Required. The unique ID that Amazon MQ generates for the configuration.

`created`

string

Format: date-time

True

Required. The date and time of the configuration.

`data`

string

Format: byte

True

Amazon MQ for ActiveMQ: the base64-encoded XML configuration.
Amazon MQ for RabbitMQ: base64-encoded Cuttlefish.

`description`

string

False

The description of the configuration.

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

### DescribeConfigurationRevision

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/describeconfigurationrevision.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/describeconfigurationrevision.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configuration

Configuration Revisions

All content copied from https://docs.aws.amazon.com/.
