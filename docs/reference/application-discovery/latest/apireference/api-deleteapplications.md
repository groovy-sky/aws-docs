# DeleteApplications

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Deletes a list of applications and their associations with configuration
items.

## Request Syntax

```nohighlight

{
   "configurationIds": [ "string" ]
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[configurationIds](#API_DeleteApplications_RequestSyntax)**

Configuration ID of an application to be deleted.

Type: Array of strings

Length Constraints: Maximum length of 200.

Pattern: `\S+`

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

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

**ServerInternalErrorException**

The server experienced an internal error. Try again.

HTTP Status Code: 500

## Examples

### Delete applications

The following example deletes two applications as specified by the values passed to
the parameter `configurationIds` in the request.

#### Sample Request

```

{
    "configurationIds": [
	      "d-application-022b024aa42685464",
	      "d-application-0ce38f881a50d2ac1"
	  ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for Python](../../../../services/goto/boto3/discovery-2015-11-01/deleteapplications.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/deleteapplications.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateTags

DeleteTags
