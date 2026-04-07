# DisableImageDeprecation

Cancels the deprecation of the specified AMI.

For more information, see [Deprecate an Amazon EC2 AMI](../../../../services/ec2/latest/userguide/ami-deprecate.md) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**ImageId**

The ID of the AMI.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Returns `true` if the request succeeds; otherwise, it returns an error.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example cancels the planned deprecation of the specified AMI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DisableImageDeprecation
&ImageId=ami-0123456789EXAMPLE
&AUTHPARAMS
```

#### Sample Response

```

<DisableImageDeprecationResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>11aabb229-4eac-35bd-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</DisableImageDeprecationResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DisableImageDeprecation)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DisableImageDeprecation)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/disableimagedeprecation.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DisableImageDeprecation)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/disableimagedeprecation.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DisableImageBlockPublicAccess

DisableImageDeregistrationProtection
