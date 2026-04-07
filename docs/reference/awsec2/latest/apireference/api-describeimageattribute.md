# DescribeImageAttribute

Describes the specified attribute of the specified AMI. You can specify only one attribute
at a time.

###### Note

The order of the elements in the response, including those within nested structures,
might vary. Applications should not assume the elements appear in a particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The AMI attribute.

**Note**: The `blockDeviceMapping` attribute is
deprecated. Using this attribute returns the `Client.AuthFailure` error. To get
information about the block device mappings for an AMI, describe the image instead.

Type: String

Valid Values: `description | kernel | ramdisk | launchPermission | productCodes | blockDeviceMapping | sriovNetSupport | bootMode | tpmSupport | uefiData | lastLaunchedTime | imdsSupport | deregistrationProtection`

Required: Yes

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

**blockDeviceMapping**

The block device mapping entries.

Type: Array of [BlockDeviceMapping](api-blockdevicemapping.md) objects

**bootMode**

The boot mode.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**deregistrationProtection**

Indicates whether deregistration protection is enabled for the AMI.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**description**

A description for the AMI.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**imageId**

The ID of the AMI.

Type: String

**imdsSupport**

If `v2.0`, it indicates that IMDSv2 is specified in the AMI. Instances launched
from this AMI will have `HttpTokens` automatically set to `required` so
that, by default, the instance requires that IMDSv2 is used when requesting instance metadata.
In addition, `HttpPutResponseHopLimit` is set to `2`. For more
information, see [Configure the AMI](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/configuring-IMDS-new-instances.html#configure-IMDS-new-instances-ami-configuration) in the _Amazon EC2 User Guide_.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**kernel**

The kernel ID.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**lastLaunchedTime**

The date and time, in [ISO 8601 date-time\
format](http://www.iso.org/iso/iso8601), when the AMI was last used to launch an EC2 instance. When the AMI is used
to launch an instance, there is a 24-hour delay before that usage is reported.

###### Note

`lastLaunchedTime` data is available starting April 2017.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**launchPermission**

The launch permissions.

Type: Array of [LaunchPermission](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_LaunchPermission.html) objects

**productCodes**

The product codes.

Type: Array of [ProductCode](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ProductCode.html) objects

**ramdisk**

The RAM disk ID.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**requestId**

The ID of the request.

Type: String

**sriovNetSupport**

Indicates whether enhanced networking with the Intel 82599 Virtual Function interface is
enabled.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**tpmSupport**

If the image is configured for NitroTPM support, the value is `v2.0`.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**uefiData**

Base64 representation of the non-volatile UEFI variable store. To retrieve the UEFI data,
use the [GetInstanceUefiData](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_GetInstanceUefiData) command. You can inspect and modify the UEFI data by using the
[python-uefivars tool](https://github.com/awslabs/python-uefivars) on
GitHub. For more information, see [UEFI Secure Boot for Amazon EC2\
instances](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/uefi-secure-boot.html) in the _Amazon EC2 User Guide_.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example lists the launch permissions for the specified AMI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeImageAttribute
&ImageId=ami-61a54008
&Attribute=launchPermission
&AUTHPARAMS
```

#### Sample Response

```

<DescribeImageAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <imageId>ami-61a54008</imageId>
   <launchPermission>
      <item>
         <group>all</group>
      </item>
      <item>
         <userId>495219933132</userId>
      </item>
   </launchPermission>
</DescribeImageAttributeResponse>
```

### Example 2

This example lists the product codes for the specified AMI.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeImageAttribute
&ImageId=ami-2bb65342
&Attribute=productCodes
&AUTHPARAMS
```

#### Sample Response

```

<DescribeImageAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/>
   <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
   <imageId>ami-2bb65342</imageId>
   <productCodes>
      <item>
        <productCode>a1b2c3d4e5f6g7h8i9j10k11</productCode>
        <type>marketplace</type>
      </item>
   </productCodes>
</DescribeImageAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeImageAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeImageAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeIdFormat

DescribeImageReferences
