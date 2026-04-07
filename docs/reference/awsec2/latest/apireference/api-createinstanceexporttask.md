# CreateInstanceExportTask

Exports a running or stopped instance to an Amazon S3 bucket.

For information about the prerequisites for your Amazon S3 bucket, supported operating systems,
image formats, and known limitations for the types of instances you can export, see [Exporting an instance as a VM Using VM\
Import/Export](https://docs.aws.amazon.com/vm-import/latest/userguide/vmexport.html) in the _VM Import/Export User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**Description**

A description for the conversion task or the resource being exported. The maximum length is 255 characters.

Type: String

Required: No

**ExportToS3**

The format and location for an export instance task.

Type: [ExportToS3TaskSpecification](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ExportToS3TaskSpecification.html) object

Required: Yes

**InstanceId**

The ID of the instance.

Type: String

Required: Yes

**TagSpecification.N**

The tags to apply to the export instance task during creation.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TargetEnvironment**

The target virtualization environment.

Type: String

Valid Values: `citrix | vmware | microsoft`

Required: Yes

## Response Elements

The following elements are returned by the service.

**exportTask**

Information about the export instance task.

Type: [ExportTask](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_ExportTask.html) object

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request creates an Export VM task that makes a Windows instance available as an OVA.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CreateInstanceExportTask
&Description=Example%20for%20docs
&InstanceId=i-1234567890abcdef0
&TargetEnvironment=VMWare
&ExportToS3.DiskImageFormat=VMDK
&ExportToS3.ContainerFormat=OVA
&ExportToS3.S3bucket=amzn-s3-demo-for-exported-vm
&ExportToS3.S3prefix=my-exports/
&AUTHPARAMS
```

#### Sample Response

```

<CreateInstanceExportTaskResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
  <exportTask>
   <exportTaskId>export-i-1234wxyz</exportTaskId>
   <description>Example for docs</description>
   <state>active</state>
   <statusMessage>Running</statusMessage>
    <instanceExport>
     <instanceId>i-1234567890abcdef0</instanceId>
     <targetEnvironment>VMWare</targetEnvironment>
    </instanceExport>
    <exportToS3>
     <diskImageFormat>VMDK</diskImageFormat>
     <containerFormat>OVA</containerFormat>
     <s3Bucket>amzn-s3-demo-for-exported-vm</s3Bucket>
     <s3Key>my-exports/ export-i-1234wxyz.ova</s3Key>
    </exportToS3>
  </exportTask>
</CreateInstanceExportTaskResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CreateInstanceExportTask)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/CreateInstanceExportTask)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CreateInstanceEventWindow

CreateInternetGateway
