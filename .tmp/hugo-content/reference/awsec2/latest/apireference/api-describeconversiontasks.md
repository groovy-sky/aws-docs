# DescribeConversionTasks

Describes the specified conversion tasks or all your conversion tasks. For more information, see the
[VM Import/Export User Guide](../../../../services/vm-import/latest/userguide.md).

For information about the import manifest referenced by this API action, see [VM Import Manifest](manifest.md).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ConversionTaskId.N**

The conversion task IDs.

Type: Array of strings

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**conversionTasks**

Information about the conversion tasks.

Type: Array of [ConversionTask](api-conversiontask.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes all your conversion tasks.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeConversionTasks
&AUTHPARAMS
```

#### Sample Response

```

<DescribeConversionTasksResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <conversionTasks>
     <item>
        <conversionTask>
           <conversionTaskId>import-i-fh95npoc</conversionTaskId>
           <expirationTime>2010-12-22T12:01Z</expirationTime>
           <importVolume>
              <bytesConverted>1000</bytesConverted>
              <availabilityZone>us-east-1a</availabilityZone>
              <description/>
              <image>
                 <format>VDMK</format>
                 <size>128696320</size>
                 <importManifestUrl>
                  https://s3.amazonaws.com/amzn-s3-demo-bucket/​a3a5e1b6-590d-43cc-97c1-15c7325d3f41/​Win_2008_Server_Data_Center_SP2_32-bit.​vmdkmanifest.xml?AWSAccessKeyId=​AWS_ACCESS_KEY_ID_REDACTED&​Expires=1294855591&​Signature=5snej01TlTtL0uR7KExtEXAMPLE%3D
                 </importManifestUrl>
             </image>
             <volume>
                <size>8</size>
                <id>vol-1234567890abcdef0</id>
             </volume>
           </importVolume>
           <state>active</state>
           <statusMessage/>
        </conversionTask>
     </item>
  </conversionTasks>
</DescribeConversionTasksResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeconversiontasks.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeconversiontasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeCoipPools

DescribeCustomerGateways
