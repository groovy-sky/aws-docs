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

Type: [AttributeValue](api-attributevalue.md) object

**deregistrationProtection**

Indicates whether deregistration protection is enabled for the AMI.

Type: [AttributeValue](api-attributevalue.md) object

**description**

A description for the AMI.

Type: [AttributeValue](api-attributevalue.md) object

**imageId**

The ID of the AMI.

Type: String

**imdsSupport**

If `v2.0`, it indicates that IMDSv2 is specified in the AMI. Instances launched
from this AMI will have `HttpTokens` automatically set to `required` so
that, by default, the instance requires that IMDSv2 is used when requesting instance metadata.
In addition, `HttpPutResponseHopLimit` is set to `2`. For more
information, see [Configure the AMI](../../../../services/ec2/latest/userguide/configuring-imds-new-instances.md#configure-IMDS-new-instances-ami-configuration) in the _Amazon EC2 User Guide_.

Type: [AttributeValue](api-attributevalue.md) object

**kernel**

The kernel ID.

Type: [AttributeValue](api-attributevalue.md) object

**lastLaunchedTime**

The date and time, in [ISO 8601 date-time\
format](http://www.iso.org/iso/iso8601), when the AMI was last used to launch an EC2 instance. When the AMI is used
to launch an instance, there is a 24-hour delay before that usage is reported.

###### Note

`lastLaunchedTime` data is available starting April 2017.

Type: [AttributeValue](api-attributevalue.md) object

**launchPermission**

The launch permissions.

Type: Array of [LaunchPermission](api-launchpermission.md) objects

**productCodes**

The product codes.

Type: Array of [ProductCode](api-productcode.md) objects

**ramdisk**

The RAM disk ID.

Type: [AttributeValue](api-attributevalue.md) object

**requestId**

The ID of the request.

Type: String

**sriovNetSupport**

Indicates whether enhanced networking with the Intel 82599 Virtual Function interface is
enabled.

Type: [AttributeValue](api-attributevalue.md) object

**tpmSupport**

If the image is configured for NitroTPM support, the value is `v2.0`.

Type: [AttributeValue](api-attributevalue.md) object

**uefiData**

Base64 representation of the non-volatile UEFI variable store. To retrieve the UEFI data,
use the [GetInstanceUefiData](api-getinstanceuefidata.md) command. You can inspect and modify the UEFI data by using the
[python-uefivars tool](https://github.com/awslabs/python-uefivars) on
GitHub. For more information, see [UEFI Secure Boot for Amazon EC2\
instances](../../../../services/ec2/latest/userguide/uefi-secure-boot.md) in the _Amazon EC2 User Guide_.

Type: [AttributeValue](api-attributevalue.md) object

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

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeimageattribute.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeimageattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeIdFormat

DescribeImageReferences
