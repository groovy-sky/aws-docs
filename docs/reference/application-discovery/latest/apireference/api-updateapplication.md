# UpdateApplication

###### Important

AWS Application Discovery Service is no longer open to new customers. Existing customers can continue to use the service as normal. For more information, see
[AWS Application Discovery Service availability change](../../../../services/application-discovery/latest/userguide/application-discovery-service-availability-change.md).

Updates metadata about an application.

## Request Syntax

```nohighlight

{
   "configurationId": "string",
   "description": "string",
   "name": "string",
   "wave": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[configurationId](#API_UpdateApplication_RequestSyntax)**

Configuration ID of the application to be updated.

Type: String

Length Constraints: Maximum length of 200.

Pattern: `\S+`

Required: Yes

**[description](#API_UpdateApplication_RequestSyntax)**

New description of the application to be updated.

Type: String

Length Constraints: Maximum length of 1000.

Pattern: `(^$|[\s\S]*\S[\s\S]*)`

Required: No

**[name](#API_UpdateApplication_RequestSyntax)**

New name of the application to be updated.

Type: String

Length Constraints: Maximum length of 127.

Pattern: `[\s\S]*\S[\s\S]*`

Required: No

**[wave](#API_UpdateApplication_RequestSyntax)**

The new migration wave of the application that you want to update.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 256.

Pattern: `^($|[^\s\x00]( *[^\s\x00])*$)`

Required: No

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

### Update application name and description

In the following example, both the name and the description are changed (updated)
of an existing application. The application is identified by passing the application id to
the required parameter `configurationId`.

#### Sample Request

```

{
   "configurationId": "d-application-03767f7bddd6c0531",
   "description": "PSoft financials db migration",
   "name": "payroll_db_migration"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for Python](../../../../services/goto/boto3/discovery-2015-11-01/updateapplication.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/discovery-2015-11-01/updateapplication.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

StopDataCollectionByAgentIds

Data Types
