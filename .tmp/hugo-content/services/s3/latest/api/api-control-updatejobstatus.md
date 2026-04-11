# UpdateJobStatus

Updates the status for the specified job. Use this operation to confirm that you want to
run a job or to cancel an existing job. For more information, see [S3 Batch Operations](../userguide/batch-ops.md) in the _Amazon S3 User Guide_.

Permissions

To use the
`UpdateJobStatus` operation, you must have permission to
perform the `s3:UpdateJobStatus` action.

Related actions include:

- [CreateJob](api-control-createjob.md)

- [ListJobs](api-control-listjobs.md)

- [DescribeJob](api-control-describejob.md)

- [UpdateJobStatus](api-control-updatejobstatus.md)

## Request Syntax

```nohighlight

POST /v20180820/jobs/id/status?requestedJobStatus=RequestedJobStatus&statusUpdateReason=StatusUpdateReason HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[id](#API_control_UpdateJobStatus_RequestSyntax)**

The ID of the job whose status you want to update.

Length Constraints: Minimum length of 5. Maximum length of 36.

Pattern: `[a-zA-Z0-9\-\_]+`

Required: Yes

**[requestedJobStatus](#API_control_UpdateJobStatus_RequestSyntax)**

The status that you want to move the specified job to.

Valid Values: `Cancelled | Ready`

Required: Yes

**[statusUpdateReason](#API_control_UpdateJobStatus_RequestSyntax)**

A description of the reason why you want to change the specified job's status. This
field can be any string up to the maximum length.

Length Constraints: Minimum length of 1. Maximum length of 256.

**[x-amz-account-id](#API_control_UpdateJobStatus_RequestSyntax)**

The AWS account ID associated with the S3 Batch Operations job.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<UpdateJobStatusResult>
   <JobId>string</JobId>
   <Status>string</Status>
   <StatusUpdateReason>string</StatusUpdateReason>
</UpdateJobStatusResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[UpdateJobStatusResult](#API_control_UpdateJobStatus_ResponseSyntax)**

Root level tag for the UpdateJobStatusResult parameters.

Required: Yes

**[JobId](#API_control_UpdateJobStatus_ResponseSyntax)**

The ID for the job whose status was updated.

Type: String

Length Constraints: Minimum length of 5. Maximum length of 36.

Pattern: `[a-zA-Z0-9\-\_]+`

**[Status](#API_control_UpdateJobStatus_ResponseSyntax)**

The current status for the specified job.

Type: String

Valid Values: `Active | Cancelled | Cancelling | Complete | Completing | Failed | Failing | New | Paused | Pausing | Preparing | Ready | Suspended`

**[StatusUpdateReason](#API_control_UpdateJobStatus_ResponseSyntax)**

The reason that the specified job's status was updated.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

## Errors

**BadRequestException**

HTTP Status Code: 400

**InternalServiceException**

HTTP Status Code: 500

**JobStatusException**

HTTP Status Code: 400

**NotFoundException**

HTTP Status Code: 400

**TooManyRequestsException**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/updatejobstatus.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/updatejobstatus.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateJobPriority

UpdateStorageLensGroup

All content copied from https://docs.aws.amazon.com/.
