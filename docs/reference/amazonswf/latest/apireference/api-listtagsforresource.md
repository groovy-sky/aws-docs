# ListTagsForResource

List tags for a given domain.

## Request Syntax

```nohighlight

{
   "resourceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[resourceArn](#API_ListTagsForResource_RequestSyntax)**

The Amazon Resource Name (ARN) for the Amazon SWF domain.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: Yes

## Response Syntax

```nohighlight

{
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[tags](#API_ListTagsForResource_ResponseSyntax)**

An array of tags associated with the domain.

Type: Array of [ResourceTag](api-resourcetag.md) objects

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**LimitExceededFault**

Returned by any operation if a system imposed limitation has been reached. To address this fault you should either clean up unused resources or increase the limit by contacting AWS.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**OperationNotPermittedFault**

Returned when the caller doesn't have sufficient permissions to invoke the action.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

**UnknownResourceFault**

Returned when the named resource cannot be found with in the scope of this operation (region or domain). This could happen if the named resource was never created or is no longer available for this operation.

**message**

A description that may help with diagnosing the cause of the fault.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for Python](../../../../services/goto/boto3/swf-2012-01-25/listtagsforresource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/listtagsforresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListOpenWorkflowExecutions

ListWorkflowTypes

All content copied from https://docs.aws.amazon.com/.
