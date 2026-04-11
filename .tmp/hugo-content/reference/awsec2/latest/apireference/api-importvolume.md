# ImportVolume

###### Note

This API action supports only single-volume VMs. To import multi-volume VMs, use
[ImportImage](api-importimage.md) instead. To import a disk to a snapshot, use
[ImportSnapshot](api-importsnapshot.md) instead.

Creates an import volume task using metadata from the specified disk image.

For information about the import manifest referenced by this API action, see [VM Import Manifest](manifest.md).

This API action is not supported by the AWS Command Line Interface (AWS CLI).

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AvailabilityZone**

The Availability Zone for the resulting EBS volume.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified,
but not both.

Type: String

Required: No

**AvailabilityZoneId**

The ID of the Availability Zone for the resulting EBS volume.

Either `AvailabilityZone` or `AvailabilityZoneId` must be specified,
but not both.

Type: String

Required: No

**Description**

A description of the volume.

Type: String

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Image**

The disk image.

Type: [DiskImageDetail](api-diskimagedetail.md) object

Required: Yes

**Volume**

The volume size.

Type: [VolumeDetail](api-volumedetail.md) object

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

This example creates an import volume task that migrates a Windows Server 2008 SP2 (32-bit) volume into the
`us-east-1` Region.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=ImportVolume
&AvailabilityZone=us-east-1c
&Image.Format=VMDK
&Image.Bytes=128696320
&Image.ImportManifestUrl=https://s3.amazonaws.com/amzn-s3-demo-bucket/​a3a5e1b6-590d-43cc-97c1-15c7325d3f41/​Win_2008_Server_Data_Center_SP2_32-bit.​vmdkmanifest.xml?AWSAccessKeyId=​AWS_ACCESS_KEY_ID_REDACTED&​Expires=1294855591&​Signature=5snej01TlTtL0uR7KExtEXAMPLE%3D
&Volume.Size=8
&AUTHPARAMS>
```

#### Sample Response

```

<ImportVolumeResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
   <conversionTask>
     <conversionTaskId>import-i-fh95npoc</conversionTaskId>
     <expirationTime>2010-12-22T12:01Z</expirationTime>
     <importVolume>
        <bytesConverted>0</bytesConverted>
        <availabilityZone>us-east-1c</availabilityZone>
        <description/>
        <image>
            <format>VDMK</format>
            <size>128696320</size>
            <importManifestUrl>
               https://s3.amazonaws.com/amzn-s3-demo-bucket/​a3a5e1b6-590d-43cc-97c1-15c7325d3f41/​Win_2008_Server_Data_Center_SP2_32-bit.​vmdkmanifest.xml?AWSAccessKeyId=​AWS_ACCESS_KEY_ID_REDACTED&​Expires=1294855591&​Signature=5snej01TlTtL0uR7KExtEXAMPLE%3D
            </importManifestUrl>
            <checksum>ccb1b0536a4a70e86016b85229b5c6b10b14a4eb</checksum>
        </image>
        <volume>
           <size>8</size>
           <id>vol-1234567890abcdef0</id>
        </volume>
     </importVolume>
     <state>active</state>
     <statusMessage/>
  </conversionTask>
</ImportVolumeResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/importvolume.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/importvolume.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/importvolume.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/importvolume.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/importvolume.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/importvolume.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/importvolume.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/importvolume.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/importvolume.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/importvolume.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ImportSnapshot

ListImagesInRecycleBin
