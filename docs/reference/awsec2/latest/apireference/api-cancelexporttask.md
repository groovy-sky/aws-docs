# CancelExportTask

Cancels an active export task. The request removes all artifacts of the export, including any partially-created
Amazon S3 objects. If the export task is complete or is in the process of transferring the final disk image, the
command fails and returns an error.

## Request Parameters

For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**ExportTaskId**

The ID of the export task. This is the ID returned by the
`CreateInstanceExportTask` and `ExportImage` operations.

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**return**

Is `true` if the request succeeds, and an error otherwise.

Type: Boolean

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request cancels the export task with the ID `export-i-1234wxyz`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=CancelExportTask
&exportTaskId=export-i-1234wxyz
&AUTHPARAMS
```

#### Sample Response

```

<CancelExportTask xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
<requestId>59dbff89-35bd-4eac-99ed-be587EXAMPLE</requestId>
<return>true</return>
</CancelExportTask>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/CancelExportTask)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/CancelExportTask)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/cancelexporttask.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/CancelExportTask)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/cancelexporttask.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

CancelDeclarativePoliciesReport

CancelImageLaunchPermission
