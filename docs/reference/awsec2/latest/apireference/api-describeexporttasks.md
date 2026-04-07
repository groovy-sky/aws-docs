# DescribeExportTasks

Describes the specified export instance tasks or all of your export instance tasks.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ExportTaskId.N**

The export task IDs.

Type: Array of strings

Required: No

**Filter.N**

the filters for the export tasks.

Type: Array of [Filter](api-filter.md) objects

Required: No

## Response Elements

The following elements are returned by the service.

**exportTaskSet**

Information about the export tasks.

Type: Array of [ExportTask](api-exporttask.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example describes a single export task.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeExportTasks
&exportTaskId.1=export-i-1234wxyz
&AUTHPARAMS
```

#### Sample Response

```

<DescribeExportTasksResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
<exportTaskSet>
  <item>
    <exportTaskId>export-i-1234wxyz</exportTaskId>
    <description>Example for docs</description>
    <state>active</state>
    <statusMessage>Running</statusMessage>
    <instanceExport>
      <instanceId>i-12345678</instanceId>
      <targetEnvironment>VMWare</targetEnvironment>
    </instanceExport>
    <exportToS3>
      <diskImageFormat>VMDK</diskImageFormat>
      <containerFormat>OVA</containerFormat>
      <s3Bucket>amzn-s3-demo-for-exported-vm</s3Bucket>
      <s3Key>my-exports/ export-i-1234wxyz.ova</s3Key>
    </exportToS3>
  </item>
</exportTaskSet>
</DescribeExportTasksResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeExportTasks)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeExportTasks)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeexporttasks.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeExportTasks)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeexporttasks.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeExportImageTasks

DescribeFastLaunchImages
