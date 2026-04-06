# Configuration Revisions

This is a collection of configuration revisions. To keep track of the changes you
make to your configuration, you can create configuration revisions. For more
information, see [Configuration](https://docs.aws.amazon.com/amazon-mq/latest/developer-guide/configuration.html) in the Amazon MQ Developer Guide.

###### Important

Making changes to a configuration does not apply the changes to the broker
immediately. To apply your changes, you must wait for the next maintenance window
or reboot the broker.

## URI

`/v1/configurations/configuration-id/revisions`

## HTTP methods

### GET

**Operation ID:** `ListConfigurationRevisions`

Returns a list of all revisions for the specified configuration.

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

Query parametersNameTypeRequiredDescription`nextToken`StringFalse

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`maxResults`StringFalse

The maximum number of brokers that Amazon MQ can return per page (20 by default).
This value must be an integer from 5 to 100.

ResponsesStatus codeResponse modelDescription`200``

         ListConfigurationRevisionsOutput`

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

Path parametersNameTypeRequiredDescription`configuration-id`StringTrue

The unique ID that Amazon MQ generates for the configuration.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Response bodies

```json

{
  "nextToken": "string",
  "maxResults": integer,
  "revisions": [
    {
      "created": "string",
      "description": "string",
      "revision": integer
    }
  ],
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

### ListConfigurationRevisionsOutput

Returns a list of all revisions for the specified configuration.

PropertyTypeRequiredDescription`configurationId`

string

False

The unique ID that Amazon MQ generates for the configuration.

`maxResults`

integer

False

The maximum number of configuration revisions that can be returned per page (20 by
default). This value must be an integer from 5 to 100.

`nextToken`

string

False

The token that specifies the next page of results Amazon MQ should return. To
request the first page, leave nextToken empty.

`revisions`

Array of type ConfigurationRevision

False

The list of all revisions for the specified configuration.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### ListConfigurationRevisions

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/mq-2017-11-27/ListConfigurationRevisions)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/mq-2017-11-27/ListConfigurationRevisions)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Configuration Revision

Configurations
