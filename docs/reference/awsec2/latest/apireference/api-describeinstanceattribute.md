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

Type: Array of [InstanceBlockDeviceMapping](api-instanceblockdevicemapping.md) objects

**disableApiStop**

Indicates whether stop protection is enabled for the instance.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

**disableApiTermination**

Indicates whether termination protection is enabled. If the value is `true`,
you can't terminate the instance using the Amazon EC2 console, command line tools, or API.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

**ebsOptimized**

Indicates whether the instance is optimized for Amazon EBS I/O.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

**enaSupport**

Indicates whether enhanced networking with ENA is enabled.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

**enclaveOptions**

Indicates whether the instance is enabled for AWS Nitro Enclaves.

Type: [EnclaveOptions](api-enclaveoptions.md) object

**groupSet**

The security groups associated with the instance.

Type: Array of [GroupIdentifier](api-groupidentifier.md) objects

**instanceId**

The ID of the instance.

Type: String

**instanceInitiatedShutdownBehavior**

Indicates whether an instance stops or terminates when you initiate shutdown from the
instance (using the operating system command for system shutdown).

Type: [AttributeValue](api-attributevalue.md) object

**instanceType**

The instance type.

Type: [AttributeValue](api-attributevalue.md) object

**kernel**

The kernel ID.

Type: [AttributeValue](api-attributevalue.md) object

**productCodes**

The product codes.

Type: Array of [ProductCode](api-productcode.md) objects

**ramdisk**

The RAM disk ID.

Type: [AttributeValue](api-attributevalue.md) object

**requestId**

The ID of the request.

Type: String

**rootDeviceName**

The device name of the root device volume (for example,
`/dev/sda1`).

Type: [AttributeValue](api-attributevalue.md) object

**sourceDestCheck**

Indicates whether source/destination checks are enabled.

Type: [AttributeBooleanValue](api-attributebooleanvalue.md) object

**sriovNetSupport**

Indicates whether enhanced networking with the Intel 82599 Virtual Function interface
is enabled.

Type: [AttributeValue](api-attributevalue.md) object

**userData**

The user data.

Type: [AttributeValue](api-attributevalue.md) object

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

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeinstanceattribute.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeInstanceAttribute)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeinstanceattribute.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeImportSnapshotTasks

DescribeInstanceConnectEndpoints
