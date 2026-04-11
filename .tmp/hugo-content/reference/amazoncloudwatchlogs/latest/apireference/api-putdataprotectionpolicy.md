# PutDataProtectionPolicy

Creates a data protection policy for the specified log group. A data protection policy can
help safeguard sensitive data that's ingested by the log group by auditing and masking the
sensitive log data.

###### Important

Sensitive data is detected and masked when it is ingested into the log group. When you
set a data protection policy, log events ingested into the log group before that time are
not masked.

By default, when a user views a log event that includes masked data, the sensitive data is
replaced by asterisks. A user who has the `logs:Unmask` permission can use a [GetLogEvents](api-getlogevents.md) or [FilterLogEvents](api-filterlogevents.md) operation with the `unmask` parameter set to
`true` to view the unmasked log events. Users with the `logs:Unmask`
can also view unmasked data in the CloudWatch Logs console by running a CloudWatch Logs
Insights query with the `unmask` query command.

For more information, including a list of types of data that can be audited and masked,
see [Protect sensitive log data\
with masking](../../../../services/amazoncloudwatch/latest/logs/mask-sensitive-log-data.md).

The `PutDataProtectionPolicy` operation applies to only the specified log
group. You can also use [PutAccountPolicy](api-putaccountpolicy.md) to create an account-level data protection policy that applies to
all log groups in the account, including both existing log groups and log groups that are
created level. If a log group has its own data protection policy and the account also has an
account-level data protection policy, then the two policies are cumulative. Any sensitive term
specified in either policy is masked.

## Request Syntax

```nohighlight

{
   "logGroupIdentifier": "string",
   "policyDocument": "string"
}
```

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

The request accepts the following data in JSON format.

**[logGroupIdentifier](#API_PutDataProtectionPolicy_RequestSyntax)**

Specify either the log group name or log group ARN.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: Yes

**[policyDocument](#API_PutDataProtectionPolicy_RequestSyntax)**

Specify the data protection policy, in JSON.

This policy must include two JSON blocks:

- The first block must include both a `DataIdentifer` array and an
`Operation` property with an `Audit` action. The
`DataIdentifer` array lists the types of sensitive data that you want to
mask. For more information about the available options, see [Types of data that\
you can mask](../../../../services/amazoncloudwatch/latest/logs/mask-sensitive-log-data-types.md).

The `Operation` property with an `Audit` action is required to
find the sensitive data terms. This `Audit` action must contain a
`FindingsDestination` object. You can optionally use that
`FindingsDestination` object to list one or more destinations to send audit
findings to. If you specify destinations such as log groups, Firehose streams,
and S3 buckets, they must already exist.

- The second block must include both a `DataIdentifer` array and an
`Operation` property with an `Deidentify` action. The
`DataIdentifer` array must exactly match the `DataIdentifer` array
in the first block of the policy.

The `Operation` property with the `Deidentify` action is what
actually masks the data, and it must contain the ` "MaskConfig": {}` object.
The ` "MaskConfig": {}` object must be empty.

For an example data protection policy, see the **Examples**
section on this page.

###### Important

The contents of the two `DataIdentifer` arrays must match exactly.

In addition to the two JSON blocks, the `policyDocument` can also include
`Name`, `Description`, and `Version` fields. The
`Name` is used as a dimension when CloudWatch Logs reports audit findings
metrics to CloudWatch.

The JSON specified in `policyDocument` can be up to 30,720 characters.

Type: String

Required: Yes

## Response Syntax

```nohighlight

{
   "lastUpdatedTime": number,
   "logGroupIdentifier": "string",
   "policyDocument": "string"
}
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in JSON format by the service.

**[lastUpdatedTime](#API_PutDataProtectionPolicy_ResponseSyntax)**

The date and time that this policy was most recently updated.

Type: Long

Valid Range: Minimum value of 0.

**[logGroupIdentifier](#API_PutDataProtectionPolicy_ResponseSyntax)**

The log group name or ARN that you specified in your request.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

**[policyDocument](#API_PutDataProtectionPolicy_ResponseSyntax)**

The data protection policy used for this log group.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**InvalidParameterException**

A parameter is specified incorrectly.

HTTP Status Code: 400

**LimitExceededException**

You have reached the maximum number of resources that can be created.

HTTP Status Code: 400

**OperationAbortedException**

Multiple concurrent requests to update the same resource were in conflict.

HTTP Status Code: 400

**ResourceNotFoundException**

The specified resource does not exist.

HTTP Status Code: 400

**ServiceUnavailableException**

The service cannot complete the request.

HTTP Status Code: 500

## Examples

### To create a data protection policy

The following example creates a data protection policy in the log group.

#### Sample Request

```nohighlight

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.PutDataProtectionPolicy
{
    "logGroupIdentifier": "my-log-group",
    "policyDocument": {
        "Name": "data-protection-policy",
        "Description": "test description",
        "Version": "2021-06-01",
        "Statement": [
            {
                "Sid": "audit-policy test",
                "DataIdentifier": [
                    "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
                ],
                "Operation": {
                    "Audit": {
                        "FindingsDestination": {
                            "CloudWatchLogs": {
                                "LogGroup": "EXISTING_LOG_GROUP_IN_YOUR_ACCOUNT"
                            },
                            "Firehose": {
                                "DeliveryStream": "EXISTING_STREAM_IN_YOUR_ACCOUNT"
                            },
                            "S3": {
                                "Bucket": "EXISTING_BUCKET"
                            }
                        }
                    }
                }
            },
            {
                "Sid": "redact-policy",
                "DataIdentifier": [
                    "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
                    "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US"
                ],
                "Operation": {
                    "Deidentify": {
                        "MaskConfig": {}
                    }
                }
            }
        ]
    }
}
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
```

### To create a log transformer policy

The following example creates a log transformer policy in the account that applies
to all log groups with names that start with `test-`.

#### Sample Request

```

POST / HTTP/1.1
Host: logs.<region>.<domain>
X-Amz-Date: <DATE>
Authorization: AWS4-HMAC-SHA256 Credential=<Credential>, SignedHeaders=content-type;date;host;user-agent;x-amz-date;x-amz-target;x-amzn-requestid, Signature=<Signature>
User-Agent: <UserAgentString>
Accept: application/json
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Connection: Keep-Alive
X-Amz-Target: Logs_20140328.PutDataProtectionPolicy
{
    "policyName": "ExampleTransformerPolicy",
    "policyType": "TRANSFORMER_POLICY",
    "selectionCriteria": 'LogGroupNamePrefix = "test-"'
    "policyDocument": [
        {
            "parseJSON": {}
        },
        {
            "addKeys": {
                "entries": [
                    {
                        "key": "metadata.transformed_in",
                        "value": "CloudWatchLogs"
                    }
                ]
            }
        },
        {
            "trimString": {
                "withKeys": [
                    "status"
                ]
            }
        },
        {
            "lowerCaseString": {
                "withKeys": [
                    "status"
                ]
            }
        }
    ]
        ]
    }
```

#### Sample Response

```

HTTP/1.1 200 OK
x-amzn-RequestId: <RequestId>
Content-Type: application/x-amz-json-1.1
Content-Length: <PayloadSizeBytes>
Date: <Date>
```

### To create a field index policy

The following example creates an account-level field index policy that is scoped to
log groups that have names that begin with `test`. The policy indexed two
fields in these log groups, `RequestId` and
`TransactionId`

#### Sample Request

```

{
    "policyName": "my_indexing_account_policy",
    "policyType": "FIELD_INDEX_POLICY",
    "policyDocument": {
        "Fields": ["RequestId", "TransactionId"]
    },
    "selectionCriteria": 'LogGroupNamePrefix = "test"'
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/logs-2014-03-28/putdataprotectionpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/putdataprotectionpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBearerTokenAuthentication

PutDeliveryDestination

All content copied from https://docs.aws.amazon.com/.
