# DescribeAccountLimits

Retrieves your account's CloudFormation limits, such as the maximum number of stacks that you
can create in your account. For more information about account limits, see [Understand CloudFormation quotas](../../../../services/cloudformation/latest/userguide/cloudformation-limits.md) in the _AWS CloudFormation User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**NextToken**

The token for the next set of items to return. (You received this token from a previous
call.)

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

## Response Elements

The following elements are returned by the service.

**AccountLimits.member.N**

An account limit structure that contain a list of CloudFormation account limits and their
values.

Type: Array of [AccountLimit](api-accountlimit.md) objects

**NextToken**

If the output exceeds 1 MB in size, a string that identifies the next page of limits. If
no additional page exists, this value is null.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### DescribeAccountLimits

This example illustrates one usage of DescribeAccountLimits.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=DescribeAccountLimits
 &NextToken=[NextToken]
 &Version=2010-05-15
 &SignatureVersion=2
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<DescribeAccountLimitsResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <DescribeAccountLimitsResult>
    <AccountLimits>
      <member>
        <Name>StackLimit</Name>
        <Value>20</Value>
      </member>
    </AccountLimits>
  </DescribeAccountLimitsResult>
  <ResponseMetadata>
    <RequestId>b9b4b068-3a41-11e5-94eb-example</RequestId>
  </ResponseMetadata>
</DescribeAccountLimitsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describeaccountlimits.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describeaccountlimits.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeregisterType

DescribeChangeSet
