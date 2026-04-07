# ResetFpgaImageAttribute

Resets the specified attribute of the specified Amazon FPGA Image (AFI) to its default value.
You can only reset the load permission attribute.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The attribute.

Type: String

Valid Values: `loadPermission`

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**FpgaImageId**

The ID of the AFI.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example resets the load permissions for the specified AFI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ResetFpgaImageAttribute
&FpgaImageId=afi-0d123e21abcc85abc
&Attribute=loadPermission
&AUTHPARAMS
```

#### Sample Response

```

<ResetFpgaImageAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>ccb58a32-30ee-4f9b-831c-639example</requestId>
    <return>true</return>
</ResetFpgaImageAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ResetFpgaImageAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/ResetFpgaImageAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ResetEbsDefaultKmsKeyId

ResetImageAttribute
