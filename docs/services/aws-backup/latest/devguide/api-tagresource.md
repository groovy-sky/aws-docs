# TagResource

Assigns a set of key-value pairs to a resource.

## Request Syntax

```nohighlight

POST /tags/resourceArn HTTP/1.1
Content-type: application/json

{
   "Tags": {
      "string" : "string"
   }
}
```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_TagResource_RequestSyntax)**

The ARN that uniquely identifies the resource.

Required: Yes

## Request Body

The request accepts the following data in JSON format.

**[Tags](#API_TagResource_RequestSyntax)**

Key-value pairs that are used to help organize your resources. You can assign your own
metadata to the resources you create. For clarity, this is the structure to assign tags:
`[{"Key":"string","Value":"string"}]`.

Type: String to string map

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterValueException**

Indicates that something is wrong with a parameter's value. For example, the value is
out of range.

**Context**

**Type**

HTTP Status Code: 400

**LimitExceededException**

A limit in the request has been exceeded; for example, a maximum number of items allowed
in a request.

**Context**

**Type**

HTTP Status Code: 400

**MissingParameterValueException**

Indicates that a required parameter is missing.

**Context**

**Type**

HTTP Status Code: 400

**ResourceNotFoundException**

A resource that is required for the action doesn't exist.

**Context**

**Type**

HTTP Status Code: 400

**ServiceUnavailableException**

The request failed due to a temporary failure of the server.

**Context**

**Type**

HTTP Status Code: 500

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/backup-2018-11-15/tagresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/backup-2018-11-15/tagresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/backup-2018-11-15/tagresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/backup-2018-11-15/tagresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/backup-2018-11-15/tagresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/backup-2018-11-15/tagresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/backup-2018-11-15/tagresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/backup-2018-11-15/tagresource.md)

- [AWS SDK for Python](../../../goto/boto3/backup-2018-11-15/tagresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/backup-2018-11-15/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

StopBackupJob

UntagResource

All content copied from https://docs.aws.amazon.com/.
