# ModifyVolumeAttribute

Modifies a volume attribute.

By default, all I/O operations for the volume are suspended when the data on the volume is
determined to be potentially inconsistent, to prevent undetectable, latent data corruption.
The I/O access to the volume can be resumed by first enabling I/O access and then checking the
data consistency on your volume.

You can change the default behavior to resume I/O operations. We recommend that you change
this only for boot volumes or for volumes that are stateless or disposable.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AutoEnableIO**

Indicates whether the volume should be auto-enabled for I/O operations.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**VolumeId**

The ID of the volume.

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

This example modifies the attribute of the volume vol-1234567890abcdef0.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyVolumeAttribute
&VolumeId=vol-1234567890abcdef0
&AutoEnableIO.Value=true
&AUTHPARAMS
```

#### Sample Response

```

<ModifyVolumeAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>5jkdf074-37ed-4004-8671-a78ee82bf1cbEXAMPLE</requestId>
   <return>true</return>
</ModifyVolumeAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/modifyvolumeattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyvolumeattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyVolume

ModifyVpcAttribute
