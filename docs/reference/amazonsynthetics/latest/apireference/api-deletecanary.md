# DeleteCanary

Permanently deletes the specified canary.

If the canary's `ProvisionedResourceCleanup` field is set to `AUTOMATIC`
or you specify `DeleteLambda` in this operation as `true`, CloudWatch Synthetics also deletes
the Lambda functions and layers that are used by the canary.

Other resources used and created by the canary are not automatically deleted.
After you delete a canary, you
should also delete the following:

- The CloudWatch alarms created for this canary. These alarms have a name of
`Synthetics-Alarm-first-198-characters-of-canary-name-canaryId-alarm number
              `

- Amazon S3 objects and buckets, such as the canary's artifact location.

- IAM roles created for the canary. If they were created in the console, these roles
have the name `
                       role/service-role/CloudWatchSyntheticsRole-First-21-Characters-of-CanaryName
              `

- CloudWatch Logs log groups created for the canary. These logs groups have the name
`/aws/lambda/cwsyn-First-21-Characters-of-CanaryName
              `

Before you delete a canary, you might want to use `GetCanary` to display
the information about this canary. Make
note of the information returned by this operation so that you can delete these resources
after you delete the canary.

## Request Syntax

```nohighlight

DELETE /canary/name?deleteLambda=DeleteLambda HTTP/1.1

```

## URI Request Parameters

The request uses the following URI parameters.

**[DeleteLambda](#API_DeleteCanary_RequestSyntax)**

Specifies whether to also delete the Lambda functions and layers used by this canary. The default
is `false`.

Your setting for this parameter is used only if the canary doesn't have `AUTOMATIC` for its
`ProvisionedResourceCleanup` field. If that field is set to `AUTOMATIC`, then the
Lambda functions and layers will be deleted when this canary is deleted.

Type: Boolean

**[name](#API_DeleteCanary_RequestSyntax)**

The name of the canary that you want to delete. To find the names of your canaries, use [DescribeCanaries](api-describecanaries.md).

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `^[0-9a-z_\-]+$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**ConflictException**

A conflicting operation is already in progress.

HTTP Status Code: 409

**InternalServerException**

An unknown internal error occurred.

HTTP Status Code: 500

**ResourceNotFoundException**

One of the specified resources was not found.

HTTP Status Code: 404

**ValidationException**

A parameter could not be validated.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for Python](../../../../services/goto/boto3/synthetics-2017-10-11/deletecanary.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/synthetics-2017-10-11/deletecanary.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CreateGroup

DeleteGroup
