# PutAccessPointPolicyForObjectLambda

###### Note

This operation is not supported by directory buckets.

Creates or replaces resource policy for an Object Lambda Access Point. For an example policy, see [Creating Object Lambda Access Points](../userguide/olap-create.md#olap-create-cli) in the _Amazon S3 User Guide_.

The following actions are related to
`PutAccessPointPolicyForObjectLambda`:

- [DeleteAccessPointPolicyForObjectLambda](api-control-deleteaccesspointpolicyforobjectlambda.md)

- [GetAccessPointPolicyForObjectLambda](api-control-getaccesspointpolicyforobjectlambda.md)

## Request Syntax

```nohighlight

PUT /v20180820/accesspointforobjectlambda/name/policy HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<PutAccessPointPolicyForObjectLambdaRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Policy>string</Policy>
</PutAccessPointPolicyForObjectLambdaRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_PutAccessPointPolicyForObjectLambda_RequestSyntax)**

The name of the Object Lambda Access Point.

Length Constraints: Minimum length of 3. Maximum length of 45.

Pattern: `^[a-z0-9]([a-z0-9\-]*[a-z0-9])?$`

Required: Yes

**[x-amz-account-id](#API_control_PutAccessPointPolicyForObjectLambda_RequestSyntax)**

The account ID for the account that owns the specified Object Lambda Access Point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[PutAccessPointPolicyForObjectLambdaRequest](#API_control_PutAccessPointPolicyForObjectLambda_RequestSyntax)**

Root level tag for the PutAccessPointPolicyForObjectLambdaRequest parameters.

Required: Yes

**[Policy](#API_control_PutAccessPointPolicyForObjectLambda_RequestSyntax)**

Object Lambda Access Point resource policy document.

Type: String

Required: Yes

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Sample resource policy

The following illustrates a sample resource policy.

```

{
    "Version" : "2008-10-17",
    "Statement":[{
        "Sid": "Grant account 123456789012 GetObject access",
        "Effect":"Allow",
        "Principal" : {
            "AWS": "arn:aws:iam::123456789012:root"
        },
        "Action":["s3-object-lambda:GetObject"],
        "Resource":["arn:aws:s3-object-lambda:us-east-1:123456789012:accesspoint/my-object-lambda-ap"]
        },
        {
        "Sid": "Grant account 444455556666 GetObject access",
        "Effect":"Allow",
        "Principal" : {
            "AWS": "arn:aws:iam::444455556666:root"
        },
        "Action":["s3-object-lambda:GetObject"],
        "Resource":["arn:aws:s3-object-lambda:us-east-1:123456789012:accesspoint/my-object-lambda-ap"]
        }
    ]
}
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3control-2018-08-20/PutAccessPointPolicyForObjectLambda)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

PutAccessPointPolicy

PutAccessPointScope
