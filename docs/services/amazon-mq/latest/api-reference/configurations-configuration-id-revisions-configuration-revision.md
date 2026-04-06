# Configuration Revision

To keep track of the changes you make to your configuration, you can create
configuration revisions. For more information, see [Configuration](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/configuration.html) in the Amazon MQ Developer Guide.

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

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/DescribeConfigurationRevision)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/DescribeConfigurationRevision)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuration

Configuration Revisions
