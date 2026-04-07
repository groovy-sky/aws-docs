# DeleteTags

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Deletes the association between configuration items and one or more tags. This API
accepts a list of multiple configuration items.

## Request Syntax

```nohighlight

{
   "configurationIds": [ "string" ],
   "tags": [
      {
         "key": "string",
         "value": "string"
      }
   ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/application-discovery/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[configurationIds](#API_DeleteTags_RequestSyntax)**

A list of configuration items with tags that you want to delete.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S*`

Required: Yes

**[tags](#API_DeleteTags_RequestSyntax)**

Tags that you want to delete from one or more configuration items. Specify the tags
that you want to delete in a _key_- _value_ format. For
example:

`{"key": "serverType", "value": "webServer"}`

Type: Array of [Tag](https://docs.aws.amazon.com/application-discovery/latest/APIReference/API_Tag.html) objects

Required: No

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/application-discovery/latest/APIReference/CommonErrors.html).

**AuthorizationErrorException**

The user does not have permission to perform the action. Check the IAM policy
associated with this user.

HTTP Status Code: 400

**HomeRegionNotSetException**

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

The home Region is not set. Set the home Region to continue.

HTTP Status Code: 400

**InvalidParameterException**

One or more parameters are not valid. Verify the parameters and try again.

HTTP Status Code: 400

**InvalidParameterValueException**

The value of one or more parameters are either invalid or out of range. Verify the
parameter values and try again.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified configuration ID was not located. Verify the configuration ID and try
again.

HTTP Status Code: 400

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

### Delete tags

The following example deletes a tag, identified by its key/value pair in the
`tags` parameter, from the server specified by the value passed to the
`configurationIds` parameter in the request.

#### Sample Request

```

{
"configurationIds": [
	"d-server-08a4bce106f63340e"
   ],
   "tags": [
      {
         "key": "ServerHostType",
         "value": "OracleProdDB"
      }
   ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/discovery-2015-11-01/DeleteTags)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/discovery-2015-11-01/DeleteTags)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/discovery-2015-11-01/DeleteTags)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/discovery-2015-11-01/DeleteTags)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/discovery-2015-11-01/DeleteTags)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/discovery-2015-11-01/DeleteTags)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/discovery-2015-11-01/DeleteTags)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/discovery-2015-11-01/DeleteTags)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/discovery-2015-11-01/DeleteTags)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/discovery-2015-11-01/DeleteTags)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DeleteApplications

DescribeAgents
