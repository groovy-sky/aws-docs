# TagResource

Adds one or more tags to an Athena resource. A tag is a label that you
assign to a resource. Each tag consists of a key and an optional value, both of which
you define. For example, you can use tags to categorize Athena workgroups,
data catalogs, or capacity reservations by purpose, owner, or environment. Use a
consistent set of tag keys to make it easier to search and filter the resources in your
account. For best practices, see [Tagging\
Best Practices](../../../../services/whitepapers/latest/tagging-best-practices/tagging-best-practices.md). Tag keys can be from 1 to 128 UTF-8 Unicode characters, and
tag values can be from 0 to 256 UTF-8 Unicode characters. Tags can use letters and
numbers representable in UTF-8, and the following characters: + - = . \_ : / @. Tag keys
and values are case-sensitive. Tag keys must be unique per resource. If you specify more
than one tag, separate them by commas.

## Request Syntax

```nohighlight

{
   "ResourceARN": "string",
   "Tags": [
      {
         "Key": "string",
         "Value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ResourceARN](#API_TagResource_RequestSyntax)**

Specifies the ARN of the Athena resource to which tags are to be
added.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Required: Yes

**[Tags](#API_TagResource_RequestSyntax)**

A collection of one or more tags, separated by commas, to be added to an Athena resource.

Type: Array of [Tag](api-tag.md) objects

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServerException**

Indicates a platform issue, which may be due to a transient condition or
outage.

HTTP Status Code: 500

**InvalidRequestException**

Indicates that something is wrong with the input to the request. For example, a
required parameter may be missing or out of range.

**AthenaErrorCode**

The error code returned when the query execution failed to process, or when the
processing request for the named query failed.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource, such as a workgroup, was not found.

**ResourceName**

The name of the Amazon resource.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/athena-2017-05-18/tagresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/athena-2017-05-18/tagresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/tagresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/athena-2017-05-18/tagresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/tagresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/athena-2017-05-18/tagresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/athena-2017-05-18/tagresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/athena-2017-05-18/tagresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/athena-2017-05-18/tagresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/tagresource.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StopQueryExecution

TerminateSession
