# DescribeInstanceAttribute

Describes the specified attribute of the specified instance. You can specify only one
attribute at a time. Available attributes include SQL license exemption configuration
for instances registered with the SQL LE service.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Attribute**

The instance attribute.

Note that the `enaSupport` attribute is not supported.

Type: String

Valid Values: `instanceType | kernel | ramdisk | userData | disableApiTermination | instanceInitiatedShutdownBehavior | rootDeviceName | blockDeviceMapping | productCodes | sourceDestCheck | groupSet | ebsOptimized | sriovNetSupport | enaSupport | enclaveOptions | disableApiStop`

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

**blockDeviceMapping**

The block device mapping of the instance.

Type: Array of [InstanceBlockDeviceMapping](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceBlockDeviceMapping.html) objects

**disableApiStop**

Indicates whether stop protection is enabled for the instance.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**disableApiTermination**

Indicates whether termination protection is enabled. If the value is `true`,
you can't terminate the instance using the Amazon EC2 console, command line tools, or API.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**ebsOptimized**

Indicates whether the instance is optimized for Amazon EBS I/O.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**enaSupport**

Indicates whether enhanced networking with ENA is enabled.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**enclaveOptions**

Indicates whether the instance is enabled for AWS Nitro Enclaves.

Type: [EnclaveOptions](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_EnclaveOptions.html) object

**groupSet**

The security groups associated with the instance.

Type: Array of [GroupIdentifier](api-groupidentifier.md) objects

**instanceId**

The ID of the instance.

Type: String

**instanceInitiatedShutdownBehavior**

Indicates whether an instance stops or terminates when you initiate shutdown from the
instance (using the operating system command for system shutdown).

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**instanceType**

The instance type.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**kernel**

The kernel ID.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**productCodes**

The product codes.

Type: Array of [ProductCode](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ProductCode.html) objects

**ramdisk**

The RAM disk ID.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**requestId**

The ID of the request.

Type: String

**rootDeviceName**

The device name of the root device volume (for example,
`/dev/sda1`).

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**sourceDestCheck**

Indicates whether source/destination checks are enabled.

Type: [AttributeBooleanValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeBooleanValue.html) object

**sriovNetSupport**

Indicates whether enhanced networking with the Intel 82599 Virtual Function interface
is enabled.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

**userData**

The user data.

Type: [AttributeValue](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_AttributeValue.html) object

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example lists the instance type of the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeInstanceAttribute
&InstanceId=i-1234567890abcdef0
&Attribute=instanceType
&AUTHPARAMS
```

#### Sample Response

```

<DescribeInstanceAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instanceId>i-1234567890abcdef0</instanceId>
  <instanceType>
    <value>t1.micro</value>
  </instanceType>
</DescribeInstanceAttributeResponse>
```

### Example 2

This example lists the current value of the
`InstanceInitiatedShutdownBehavior` attribute for the specified
instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeInstanceAttribute
&InstanceId=i-1234567890abcdef0
&Attribute=instanceInitiatedShutdownBehavior
&AUTHPARAMS
```

#### Sample Response

```

<DescribeInstanceAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instanceId>i-1234567890abcdef0</instanceId>
  <instanceInitiatedShutdownBehavior>
    <value>stop</value>
  </instanceInitiatedShutdownBehavior>
</DescribeInstanceAttributeResponse>
```

### Example 3

This example lists the current value of the `DisableApiTermination`
attribute for the specified instance.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeInstanceAttribute
&InstanceId=i-1234567890abcdef0
&Attribute=disableApiTermination
&AUTHPARAMS
```

#### Sample Response

```

<DescribeInstanceAttributeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <instanceId>i-1234567890abcdef0</instanceId>
  <disableApiTermination>
    <value>false</value>
  </disableApiTermination>
</DescribeInstanceAttributeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/DescribeInstanceAttribute)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DescribeImportSnapshotTasks

DescribeInstanceConnectEndpoints
