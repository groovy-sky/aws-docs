# UpdateJobPriority

Updates an existing S3 Batch Operations job's priority. For more information, see [S3 Batch Operations](../userguide/batch-ops.md) in the _Amazon S3 User Guide_.

Permissions

To use the
`UpdateJobPriority` operation, you must have permission to
perform the `s3:UpdateJobPriority` action.

Related actions include:

- [CreateJob](api-control-createjob.md)

- [ListJobs](api-control-listjobs.md)

- [DescribeJob](api-control-describejob.md)

- [UpdateJobStatus](api-control-updatejobstatus.md)

## Request Syntax

```nohighlight

POST /v20180820/jobs/id/priority?priority=Priority HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[id](#API_control_UpdateJobPriority_RequestSyntax)**

The ID for the job whose priority you want to update.

Length Constraints: Minimum length of 5. Maximum length of 36.

Pattern: `[a-zA-Z0-9\-\_]+`

Required: Yes

**[priority](#API_control_UpdateJobPriority_RequestSyntax)**

The priority you want to assign to this job.

Valid Range: Minimum value of 0. Maximum value of 2147483647.

Required: Yes

**[x-amz-account-id](#API_control_UpdateJobPriority_RequestSyntax)**

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
<UpdateJobPriorityResult>
   <JobId>string</JobId>
   <Priority>integer</Priority>
</UpdateJobPriorityResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[UpdateJobPriorityResult](#API_control_UpdateJobPriority_ResponseSyntax)**

Root level tag for the UpdateJobPriorityResult parameters.

Required: Yes

**[JobId](#API_control_UpdateJobPriority_ResponseSyntax)**

The ID for the job whose priority Amazon S3 updated.

Type: String

Length Constraints: Minimum length of 5. Maximum length of 36.

Pattern: `[a-zA-Z0-9\-\_]+`

**[Priority](#API_control_UpdateJobPriority_ResponseSyntax)**

The new priority assigned to the specified job.

Type: Integer

Valid Range: Minimum value of 0. Maximum value of 2147483647.

## Errors

**BadRequestException**

HTTP Status Code: 400

**InternalServiceException**

HTTP Status Code: 500

**NotFoundException**

HTTP Status Code: 400

**TooManyRequestsException**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/updatejobpriority.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/updatejobpriority.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

UpdateAccessGrantsLocation

UpdateJobStatus

All content copied from https://docs.aws.amazon.com/.
