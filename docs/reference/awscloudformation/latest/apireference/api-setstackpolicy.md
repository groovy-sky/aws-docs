# SetStackPolicy

Sets a stack policy for a specified stack.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**StackName**

The name or unique stack ID that you want to associate a policy with.

Type: String

Required: Yes

**StackPolicyBody**

Structure that contains the stack policy body. For more information, see [Prevent updates to stack resources](../../../../services/cloudformation/latest/userguide/protect-stack-resources.md) in the _AWS CloudFormation User Guide_.
You can specify either the `StackPolicyBody` or the `StackPolicyURL`
parameter, but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16384.

Required: No

**StackPolicyURL**

Location of a file that contains the stack policy. The URL must point to a policy (maximum
size: 16 KB) located in an Amazon S3 bucket in the same AWS Region as the stack. The location for
an Amazon S3 bucket must start with `https://`. URLs from S3 static websites are not
supported.

You can specify either the `StackPolicyBody` or the `StackPolicyURL`
parameter, but not both.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 5120.

Required: No

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### SetStackPolicy

This example illustrates one usage of SetStackPolicy.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=SetStackPolicy
 &StackName=MyStack
 &StackPolicyBody=[Stack Policy Document]
 &Version=2010-05-15
 &SignatureVersion=2
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<SetStackPolicyResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <ResponseMetadata>
    <RequestId>e7d8c346-744b-11e5-b40b-example</RequestId>
  </ResponseMetadata>
</SetStackPolicyResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/setstackpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/setstackpolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RollbackStack

SetTypeConfiguration

All content copied from https://docs.aws.amazon.com/.
