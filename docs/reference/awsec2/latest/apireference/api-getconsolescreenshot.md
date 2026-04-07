# GetConsoleScreenshot

Retrieve a JPG-format screenshot of a running instance to help with
troubleshooting.

The returned content is Base64-encoded.

For more information, see [Instance console output](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/troubleshoot-unreachable-instance.html#instance-console-console-output) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**WakeUp**

When set to `true`, acts as keystroke input and wakes up an instance that's
in standby or "sleep" mode.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**imageData**

The data that comprises the image.

Type: String

**instanceId**

The ID of the instance.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example returns the image data of a successful request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=GetConsoleScreenshot
&InstanceId=i-0598c7d356eba48d7
&AUTHPARAMS
```

#### Sample Response

```

<GetConsoleScreenshotResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <imageData>997987/8kgj49ikjhewkwwe0008084EXAMPLE</imageData>
  <instanceId>i-765950</instanceId>
</GetConsoleScreenshotResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/GetConsoleScreenshot)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/GetConsoleScreenshot)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GetConsoleOutput

GetDeclarativePoliciesReportSummary
