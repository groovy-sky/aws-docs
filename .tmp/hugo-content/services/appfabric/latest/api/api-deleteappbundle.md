# DeleteAppBundle

Deletes an app bundle. You must delete all associated app authorizations before you can
delete an app bundle.

## Request Syntax

```nohighlight

DELETE /appbundles/appBundleIdentifier HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[appBundleIdentifier](#API_DeleteAppBundle_RequestSyntax)**

The ID or Amazon Resource Name (ARN) of the app bundle that needs to be deleted.

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:.+$|^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**AccessDeniedException**

You are not authorized to perform this operation.

HTTP Status Code: 403

**ConflictException**

The request has created a conflict. Check the request parameters and try again.

**resourceId**

The resource ID.

**resourceType**

The resource type.

HTTP Status Code: 409

**InternalServerException**

The request processing has failed because of an unknown error, exception, or failure
with an internal server.

**retryAfterSeconds**

The period of time after which you should retry your request.

HTTP Status Code: 500

**ThrottlingException**

The request rate exceeds the limit.

**quotaCode**

The code for the quota exceeded.

**retryAfterSeconds**

The period of time after which you should retry your request.

**serviceCode**

The code of the service.

HTTP Status Code: 429

**ValidationException**

The request has invalid or missing parameters.

**fieldList**

The field list.

**reason**

The reason for the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for Python](../../../goto/boto3/appfabric-2023-05-19/deleteappbundle.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/appfabric-2023-05-19/deleteappbundle.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteAppAuthorization

DeleteIngestion

All content copied from https://docs.aws.amazon.com/.
