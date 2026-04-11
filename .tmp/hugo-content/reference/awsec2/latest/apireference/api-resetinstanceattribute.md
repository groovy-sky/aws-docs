# ResetInstanceAttribute

Resets an attribute of an instance to its default value. To reset the
`kernel` or `ramdisk`, the instance must be in a stopped
state. To reset the `sourceDestCheck`, the instance can be either running or
stopped.

The `sourceDestCheck` attribute controls whether source/destination
checking is enabled. The default value is `true`, which means checking is
enabled. This value must be `false` for a NAT instance to perform NAT. For
more information, see [NAT instances](../../../../services/amazonvpc/latest/userguide/vpc-nat-instance.md) in the
_Amazon VPC User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The attribute to reset.

Type: String

Valid Values: `kernel | ramdisk | sourceDestCheck`

Required: Yes

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

This example resets the `sourceDestCheck` attribute.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ResetInstanceAttribute
&InstanceId=i-1234567890abcdef0
&Attribute=sourceDestCheck
&AUTHPARAMS
```

#### Sample Response

```

<ResetInstanceAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <return>true</return>
</ResetInstanceAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/resetinstanceattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/resetinstanceattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResetImageAttribute

ResetNetworkInterfaceAttribute
