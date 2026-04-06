# StartDeclarativePoliciesReport

Generates an account status report. The report is generated asynchronously, and can
take several hours to complete.

The report provides the current status of all attributes supported by declarative
policies for the accounts within the specified scope. The scope is determined by the
specified `TargetId`, which can represent an individual account, or all the
accounts that fall under the specified organizational unit (OU) or root (the entire
AWS Organization).

The report is saved to your specified S3 bucket, using the following path structure
(with the capitalized placeholders representing your specific values):

`s3://AMZN-S3-DEMO-BUCKET/YOUR-OPTIONAL-S3-PREFIX/ec2_TARGETID_REPORTID_YYYYMMDDTHHMMZ.csv`

###### Prerequisites for generating a report

- The `StartDeclarativePoliciesReport` API can only be called by the
management account or delegated administrators for the organization.

- An S3 bucket must be available before generating the report (you can create a
new one or use an existing one), it must be in the same Region where the report
generation request is made, and it must have an appropriate bucket policy. For a
sample S3 policy, see _Sample Amazon S3 policy_ under [Examples](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_StartDeclarativePoliciesReport.html#API_StartDeclarativePoliciesReport_Examples).

- Trusted access must be enabled for the service for which the declarative
policy will enforce a baseline configuration. If you use the AWS Organizations
console, this is done automatically when you enable declarative policies. The
API uses the following service principal to identify the EC2 service:
`ec2.amazonaws.com`. For more information on how to enable
trusted access with the AWS CLI and AWS SDKs, see [Using\
Organizations with other AWS services](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html) in the
_AWS Organizations User Guide_.

- Only one report per organization can be generated at a time. Attempting to
generate a report while another is in progress will result in an error.

For more information, including the required IAM permissions to run this API, see
[Generating the account status report for declarative policies](https://docs.aws.amazon.com/organizations/latest/userguide/orgs_manage_policies_declarative_status-report.html) in the
_AWS Organizations User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**S3Bucket**

The name of the S3 bucket where the report will be saved. The bucket must be in the
same Region where the report generation request is made.

Type: String

Required: Yes

**S3Prefix**

The prefix for your S3 object.

Type: String

Required: No

**TagSpecification.N**

The tags to apply.

Type: Array of [TagSpecification](api-tagspecification.md) objects

Required: No

**TargetId**

The root ID, organizational unit ID, or account ID.

Format:

- For root: `r-ab12`

- For OU: `ou-ab12-cdef1234`

- For account: `123456789012`

Type: String

Required: Yes

## Response Elements

The following elements are returned by the service.

**reportId**

The ID of the report.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Sample Amazon S3 policy

Before generating a report, you must grant the EC2 Declarative Policies
principal access to the Amazon S3 bucket where the report will be stored. To do
this, attach the following policy to the bucket. Remember to replace
_amzn-s3-demo-bucket_ in the policy with your actual S3
bucket name, and _identity\_ARN_ with the [IAM\
identity](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns) used to call the
`StartDeclarativePoliciesReport` API. The statement in the
`Condition` element ensures that the service can call the
operation on your behalf using the credentials used to call the
`StartDeclarativePoliciesReport` API.

```json

{
    "Version": "2012-10-17",
    "Statement": [{
        "Sid": "DeclarativePoliciesReportDelivery",
        "Effect": "Allow",
        "Principal": {
            "AWS": "identity_ARN"
        },
        "Action": [
            "s3:PutObject"
        ],
        "Resource": "arn:aws:s3:::amzn-s3-demo-bucket/*",
        "Condition": {
            "StringEquals": {
                "aws:CalledViaLast": "report.declarative-policies-ec2.amazonaws.com"
            }
        }
    }]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/StartDeclarativePoliciesReport)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/StartDeclarativePoliciesReport)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SendDiagnosticInterrupt

StartInstances
