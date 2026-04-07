# AssociateFleet

Associates the specified fleet with the specified stack.

## Request Syntax

```nohighlight

{
   "FleetName": "string",
   "StackName": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](https://docs.aws.amazon.com/appstream2/latest/APIReference/CommonParameters.html).

The request accepts the following data in JSON format.

**[FleetName](#API_AssociateFleet_RequestSyntax)**

The name of the fleet.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

**[StackName](#API_AssociateFleet_RequestSyntax)**

The name of the stack.

Type: String

Length Constraints: Minimum length of 1.

Required: Yes

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](https://docs.aws.amazon.com/appstream2/latest/APIReference/CommonErrors.html).

**ConcurrentModificationException**

An API error occurred. Wait a few minutes and try again.

**Message**

The error message in the exception.

HTTP Status Code: 400

**IncompatibleImageException**

The image can't be updated because it's not compatible for updates.

**Message**

The error message in the exception.

HTTP Status Code: 400

**InvalidAccountStatusException**

The resource cannot be created because your AWS account is suspended. For assistance, contact AWS Support.

**Message**

The error message in the exception.

HTTP Status Code: 400

**LimitExceededException**

The requested limit exceeds the permitted limit for an account.

**Message**

The error message in the exception.

HTTP Status Code: 400

**OperationNotPermittedException**

The attempted operation is not permitted.

**Message**

The error message in the exception.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource was not found.

**Message**

The error message in the exception.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/appstream-2016-12-01/AssociateFleet)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appstream-2016-12-01/AssociateFleet)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AssociateApplicationToEntitlement

AssociateSoftwareToImageBuilder
