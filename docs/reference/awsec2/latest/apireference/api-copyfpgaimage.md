# CopyFpgaImage

Copies the specified Amazon FPGA Image (AFI) to the current Region.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ClientToken**

Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.
For more information, see [Ensuring idempotency](../../../../services/ec2/latest/devguide/ec2-api-idempotency.md).

Type: String

Required: No

**Description**

The description for the new AFI.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Name**

The name for the new AFI. The default is the name of the source AFI.

Type: String

Required: No

**SourceFpgaImageId**

The ID of the source AFI.

Type: String

Required: Yes

**SourceRegion**

The Region that contains the source AFI.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**fpgaImageId**

The ID of the new AFI.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example copies the specified AFI from the `us-east-1` Region to the current Region ( `eu-west-1`).

#### Sample Request

```

https://ec2.eu-west-1.amazonaws.com/?Action=CopyFpgaImage
&Name=eu-afi
&SourceFpgaImageId=afi-0d123eabcbfc85456
&SourceRegion=us-east-1
&AUTHPARAMS
```

#### Sample Response

```

<CopyFpgaImageResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>2d55d021-9ca9-45a1-8c5c-453example</requestId>
    <fpgaImageId>afi-06b12350a123fbabc</fpgaImageId>
</CopyFpgaImageResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CopyFpgaImage)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CopyFpgaImage)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/copyfpgaimage.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CopyFpgaImage)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/copyfpgaimage.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ConfirmProductInstance

CopyImage
