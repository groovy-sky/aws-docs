# Tag

A tag is a key-value pair associated with a resource. You can use these metadata
tags to identify the purpose of a broker or configuration. For more information see
[Tagging resources](../developer-guide/amazon-mq-tagging.md) in the Amazon MQ Developer Guide.

## URI

`/v1/tags/resource-arn`

## HTTP methods

### GET

**Operation ID:** `ListTags`

Lists tags for a resource.

Path parametersNameTypeRequiredDescription`resource-arn`StringTrue

The Amazon Resource Name (ARN) of the resource tag.

ResponsesStatus codeResponse modelDescription`200``Tags`

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

### POST

**Operation ID:** `CreateTags`

Add a tag to a resource.

Path parametersNameTypeRequiredDescription`resource-arn`StringTrue

The Amazon Resource Name (ARN) of the resource tag.

ResponsesStatus codeResponse modelDescription`204`None

HTTP Status Code 204: Successful response.

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

### DELETE

**Operation ID:** `DeleteTags`

Removes a tag from a resource.

Path parametersNameTypeRequiredDescription`resource-arn`StringTrue

The Amazon Resource Name (ARN) of the resource tag.

Query parametersNameTypeRequiredDescription`tagKeys`StringTrue

An array of tag keys to delete

ResponsesStatus codeResponse modelDescription`204`None

HTTP Status Code 204: Successful response.

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

Path parametersNameTypeRequiredDescription`resource-arn`StringTrue

The Amazon Resource Name (ARN) of the resource tag.

ResponsesStatus codeResponse modelDescription`200`None

Default response for CORS method

## Schemas

### Request bodies

```json

{
  "tags": {
  }
}
```

### Response bodies

```json

{
  "tags": {
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

### Tags

A map of the key-value pairs for the resource tag.

PropertyTypeRequiredDescription`tags`

object

False

The key-value pair for the resource tag.

## See also

For more information about using this API in one of the language-specific AWS SDKs and references, see the following:

### ListTags

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/listtags.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/listtags.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/listtags.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/listtags.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/listtags.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/listtags.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/listtags.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/listtags.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/listtags.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/listtags.md)

### CreateTags

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/createtags.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/createtags.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/createtags.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/createtags.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/createtags.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/createtags.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/createtags.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/createtags.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/createtags.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/createtags.md)

### DeleteTags

- [AWS Command Line Interface V2](../../../goto/cli2/mq-2017-11-27/deletetags.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/mq-2017-11-27/deletetags.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/mq-2017-11-27/deletetags.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/mq-2017-11-27/deletetags.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/mq-2017-11-27/deletetags.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/mq-2017-11-27/deletetags.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/mq-2017-11-27/deletetags.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/mq-2017-11-27/deletetags.md)

- [AWS SDK for Python](../../../goto/boto3/mq-2017-11-27/deletetags.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/mq-2017-11-27/deletetags.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Configurations

User

All content copied from https://docs.aws.amazon.com/.
