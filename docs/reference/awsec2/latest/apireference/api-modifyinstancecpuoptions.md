# ModifyInstanceCpuOptions

By default, all vCPUs for the instance type are active when you launch an instance. When you
configure the number of active vCPUs for the instance, it can help you save on licensing costs and
optimize performance. The base cost of the instance remains unchanged.

The number of active vCPUs equals the number of threads per CPU core multiplied by the number
of cores. The instance must be in a `Stopped` state before you make changes.

###### Note

Some instance type options do not support this capability. For more information, see
[Supported CPU \
options](../../../../services/ec2/latest/userguide/cpu-options-supported-instances-values.md) in the _Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**CoreCount**

The number of CPU cores to activate for the specified instance.

Type: Integer

Required: No

**DryRun**

Checks whether you have the required permissions for the operation, without actually making the
request, and provides an error response. If you have the required permissions, the error response is
`DryRunOperation`. Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**InstanceId**

The ID of the instance to update.

Type: String

Required: Yes

**NestedVirtualization**

Indicates whether to enable or disable nested virtualization for the instance.
When nested virtualization is enabled, Virtual Secure Mode (VSM) is automatically disabled for the instance.

Type: String

Valid Values: `enabled | disabled`

Required: No

**ThreadsPerCore**

The number of threads to run for each CPU core.

Type: Integer

Required: No

## Response Elements

The following elements are returned by the service.

**coreCount**

The number of CPU cores that are running for the specified instance after the
update.

Type: Integer

**instanceId**

The ID of the instance that was updated.

Type: String

**nestedVirtualization**

Indicates whether nested virtualization has been enabled or disabled.

Type: String

Valid Values: `enabled | disabled`

**requestId**

The ID of the request.

Type: String

**threadsPerCore**

The number of threads that are running per CPU core for the specified
instance after the update.

Type: Integer

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This request example modifies the number of active vCPUs for the specified
instance in the specified Region to `4` CPU cores running `2`
threads each. This example assumes that the instance was stopped prior to running the
request.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ModifyInstanceCpuOptions
&Region=us-east-1
&InstanceId=i-1234567890abcdef0
&CoreCount=4
&ThreadsPerCore=2
&AUTHPARAMS
```

#### Sample Response

```

<ModifyInstanceCpuOptionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
	<requestId>11111111-2222-3333-4444-5555EXAMPLE</requestId>
	<instanceId>i-1234567890abcdef0</instanceId>
	<coreCount/>4</coreCount/>
	<threadsPerCore/>2</coreCount/>
</ModifyInstanceCpuOptionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ModifyInstanceCpuOptions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ModifyInstanceCpuOptions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/modifyinstancecpuoptions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ModifyInstanceCpuOptions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/modifyinstancecpuoptions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ModifyInstanceConnectEndpoint

ModifyInstanceCreditSpecification
