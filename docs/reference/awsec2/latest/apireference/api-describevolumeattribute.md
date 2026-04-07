# DescribeVolumeAttribute

Describes the specified attribute of the specified volume. You can specify only one
attribute at a time.

For more information about EBS volumes, see [Amazon EBS volumes](https://docs.aws.amazon.com/ebs/latest/userguide/ebs-volumes.html) in the _Amazon EBS User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The attribute of the volume. This parameter is required.

Type: String

Valid Values: `autoEnableIO | productCodes`

Required: Yes

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

**autoEnableIO**

The state of `autoEnableIO` attribute.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**productCodes**

A list of product codes.

Type: Array of [ProductCode](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ProductCode.html) objects

**requestId**

The ID of the request.

Type: String

**volumeId**

The ID of the volume.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes the `autoEnableIO` attribute of the volume
vol-1234567890abcdef0.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVolumeAttribute
&Attribute=autoEnableIO
&VolumeId=vol-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVolumeAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>5jkdf074-37ed-4004-8671-a78ee82bf1cbEXAMPLE</requestId>
  <volumeId>vol-1234567890abcdef0</volumeId>
  <autoEnableIO>
    <value>false</value>
  </autoEnableIO>
</DescribeVolumeAttributeResponse>
```

### Example

This example describes the `productCodes` attribute of the volume
vol-1234567890abcdef0.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeVolumeAttribute
&Attribute=productCodes
&VolumeId=vol-1234567890abcdef0
&AUTHPARAMS
```

#### Sample Response

```

<DescribeVolumeAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>5jkdf074-37ed-4004-8671-a78ee82bf1cbEXAMPLE</requestId>
  <volumeId>vol-1234567890abcdef0</volumeId>
  <productCodes>
      <item>
          <productCode>a1b2c3d4e5f6g7h8i9j10k11</productCode>
          <type>marketplace</type>
        </item>
  </productCodes>
</DescribeVolumeAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeVolumeAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeVolumeAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeVerifiedAccessTrustProviders

DescribeVolumes
