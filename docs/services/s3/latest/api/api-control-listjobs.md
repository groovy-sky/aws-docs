# ListJobs

Lists current S3 Batch Operations jobs as well as the jobs that have ended within the last 90
days for the AWS account making the request. For more information, see [S3 Batch Operations](../userguide/batch-ops.md) in the _Amazon S3 User Guide_.

Permissions

To use the
`ListJobs` operation, you must have permission to
perform the `s3:ListJobs` action.

Related actions include:

- [CreateJob](api-control-createjob.md)

- [DescribeJob](api-control-describejob.md)

- [UpdateJobPriority](api-control-updatejobpriority.md)

- [UpdateJobStatus](api-control-updatejobstatus.md)

## Request Syntax

```nohighlight

GET /v20180820/jobs?jobStatuses=JobStatuses&maxResults=MaxResults&nextToken=NextToken HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[jobStatuses](#API_control_ListJobs_RequestSyntax)**

The `List Jobs` request returns jobs that match the statuses listed in this
element.

Valid Values: `Active | Cancelled | Cancelling | Complete | Completing | Failed | Failing | New | Paused | Pausing | Preparing | Ready | Suspended`

**[maxResults](#API_control_ListJobs_RequestSyntax)**

The maximum number of jobs that Amazon S3 will include in the `List Jobs`
response. If there are more jobs than this number, the response will include a pagination
token in the `NextToken` field to enable you to retrieve the next page of
results.

Valid Range: Minimum value of 0. Maximum value of 1000.

**[nextToken](#API_control_ListJobs_RequestSyntax)**

A pagination token to request the next page of results. Use the token that Amazon S3 returned
in the `NextToken` element of the `ListJobsResult` from the previous
`List Jobs` request.

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^[A-Za-z0-9\+\:\/\=\?\#-_]+$`

**[x-amz-account-id](#API_control_ListJobs_RequestSyntax)**

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
<ListJobsResult>
   <NextToken>string</NextToken>
   <Jobs>
      <JobListDescriptor>
         <CreationTime>timestamp</CreationTime>
         <Description>string</Description>
         <JobId>string</JobId>
         <Operation>string</Operation>
         <Priority>integer</Priority>
         <ProgressSummary>
            <NumberOfTasksFailed>long</NumberOfTasksFailed>
            <NumberOfTasksSucceeded>long</NumberOfTasksSucceeded>
            <Timers>
               <ElapsedTimeInActiveSeconds>long</ElapsedTimeInActiveSeconds>
            </Timers>
            <TotalNumberOfTasks>long</TotalNumberOfTasks>
         </ProgressSummary>
         <Status>string</Status>
         <TerminationDate>timestamp</TerminationDate>
      </JobListDescriptor>
   </Jobs>
</ListJobsResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[ListJobsResult](#API_control_ListJobs_ResponseSyntax)**

Root level tag for the ListJobsResult parameters.

Required: Yes

**[Jobs](#API_control_ListJobs_ResponseSyntax)**

The list of current jobs and jobs that have ended within the last 30 days.

Type: Array of [JobListDescriptor](api-control-joblistdescriptor.md) data types

**[NextToken](#API_control_ListJobs_ResponseSyntax)**

If the `List Jobs` request produced more than the maximum number of results,
you can pass this value into a subsequent `List Jobs` request in order to
retrieve the next page of results.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Pattern: `^[A-Za-z0-9\+\:\/\=\?\#-_]+$`

## Errors

**InternalServiceException**

HTTP Status Code: 500

**InvalidNextTokenException**

HTTP Status Code: 400

**InvalidRequestException**

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/listjobs.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/listjobs.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/listjobs.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/listjobs.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/listjobs.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/listjobs.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/listjobs.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/listjobs.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/listjobs.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/listjobs.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ListCallerAccessGrants

ListMultiRegionAccessPoints

All content copied from https://docs.aws.amazon.com/.
