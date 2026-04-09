# StartDeployment

###### Important

AWS App Runner will no longer be open to new
customers starting March 31, 2026. If you would like to use App Runner, sign up prior to that date. Existing customers can
continue to use the service as normal. For more information, see
[AWS App Runner availability change](../dg/apprunner-availability-change.md).

Initiate a manual deployment of the latest commit in a source code repository or the latest image in a source image repository to an AWS App Runner
service.

For a source code repository, App Runner retrieves the commit and builds a Docker image. For a source image repository, App Runner retrieves the latest Docker
image. In both cases, App Runner then deploys the new image to your service and starts a new container instance.

This is an asynchronous operation. On a successful call, you can use the returned `OperationId` and the [ListOperations](api-listoperations.md)
call to track the operation's progress.

## Request Syntax

```nohighlight

{
   "ServiceArn": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[ServiceArn](#API_StartDeployment_RequestSyntax)**

The Amazon Resource Name (ARN) of the App Runner service that you want to manually deploy to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: Yes

## Response Syntax

```nohighlight

{
   "OperationId": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[OperationId](#API_StartDeployment_ResponseSyntax)**

The unique ID of the asynchronous operation that this request started. You can use it combined with the [ListOperations](api-listoperations.md) call to track
the operation's progress.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InternalServiceErrorException**

An unexpected service exception occurred.

HTTP Status Code: 500

**InvalidRequestException**

One or more input parameters aren't valid. Refer to the API action's document page, correct the input parameters, and try the action again.

HTTP Status Code: 400

**ResourceNotFoundException**

A resource doesn't exist for the specified Amazon Resource Name (ARN) in your AWS account.

HTTP Status Code: 400

## Examples

### Initiate a manual deployment

This example illustrates how to perform a manual deployment to an App Runner service.

#### Sample Request

```json

$ aws apprunner start-deployment --cli-input-json "`cat`"
{
  "ServiceArn": "arn:aws:apprunner:us-east-1:123456789012:service/python-app/8fe1e10304f84fd2b0df550fe98a71fa"
}
```

#### Sample Response

```json

{
  "OperationId": "853a7d5b-fc9f-4730-831b-fd8037ab832a"
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for Python](../../../goto/boto3/apprunner-2020-05-15/startdeployment.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/startdeployment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResumeService

TagResource

All content copied from https://docs.aws.amazon.com/.
