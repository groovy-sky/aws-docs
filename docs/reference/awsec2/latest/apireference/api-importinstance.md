# ImportInstance

###### Note

We recommend that you use the [`ImportImage`](api-importimage.md)
API instead. For more information, see [Importing a VM as an image using VM\
Import/Export](../../../../services/vm-import/latest/userguide/vmimport-image-import.md) in the _VM Import/Export User Guide_.

Creates an import instance task using metadata from the specified disk image.

This API action supports only single-volume VMs. To import multi-volume VMs, use [ImportImage](api-importimage.md)
instead.

For information about the import manifest referenced by this API action, see [VM Import Manifest](manifest.md).

This API action is not supported by the AWS Command Line Interface (AWS CLI).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Description**

A description for the instance being imported.

Type: String

Required: No

**DiskImage.N**

The disk image.

Type: Array of [DiskImage](api-diskimage.md) objects

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**LaunchSpecification**

The launch specification.

Type: [ImportInstanceLaunchSpecification](api-importinstancelaunchspecification.md) object

Required: No

**Platform**

The instance operating system.

Type: String

Valid Values: `Windows`

Required: Yes

## Response Elements

The following elements are returned by the service.

**conversionTask**

Information about the conversion task.

Type: [ConversionTask](api-conversiontask.md) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example creates an import instance task that migrates a Windows Server 2008 SP2 (32-bit) VM into the
`us-east-1` Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ImportInstance
&LaunchSpecification.Architecture=x86_64
&LaunchSpecification.InstanceType=m1.xlarge
&DiskImage.1.Image.Format=VMDK
&DiskImage.1.Image.Bytes=1179593728
&DiskImage.1.Image.ImportManifestUrl=https://s3.amazonaws.com/amzn-s3-demo-bucket/​a3a5e1b6-590d-43cc-97c1-15c7325d3f41/​Win_2008_Server_Data_Center_SP2_32-bit.​vmdkmanifest.xml?AWSAccessKeyId=​AWS_ACCESS_KEY_ID_REDACTED&​Expires=1294855591&​Signature=5snej01TlTtL0uR7KExtEXAMPLE%3D
&DiskImage.1.Volume.Size=12
&Platform=Windows
&AUTHPARAMS
```

#### Sample Response

```

<ImportInstanceResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
  <conversionTask>
    <conversionTaskId>import-i-ffvko9js</conversionTaskId>
    <expirationTime>2010-12-22T12:01Z</expirationTime>
    <importInstance>
       <volumes>
          <item>
             <bytesConverted>0</bytesConverted>
             <availabilityZone>us-east-1a</availabilityZone>
             <image>
                <format>VMDK</format>
                <size>1179593728</size>
                <importManifestUrl>
                 https://s3.amazonaws.com/amzn-s3-demo-bucket/​a3a5e1b6-590d-43cc-97c1-15c7325d3f41/​Win_2008_Server_Data_Center_SP2_32-bit.​vmdkmanifest.xml?AWSAccessKeyId=​AWS_ACCESS_KEY_ID_REDACTED&​Expires=1294855591&​Signature=5snej01TlTtL0uR7KExtEXAMPLE%3D
                </importManifestUrl>
             </image>
             <description/>
             <volume>
                <size>12</size>
                <id>vol-1234567890abcdef0</id>
             </volume>
             <status>active</status>
             <statusMessage/>
          </item>
       </volumes>
       <instanceId>i-1234567890abcdef0</instanceId>
       <description/>
    </importInstance>
  </conversionTask>
</ImportInstanceResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/ImportInstance)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/ImportInstance)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/importinstance.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/importinstance.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/importinstance.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/importinstance.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/importinstance.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/importinstance.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/ImportInstance)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/importinstance.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportImage

ImportKeyPair
