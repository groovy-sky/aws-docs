# GetStackPolicy

Returns the stack policy for a specified stack. If a stack doesn't have a policy, a null
value is returned.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**StackName**

The name or unique stack ID that's associated with the stack whose policy you want to
get.

Type: String

Required: Yes

## Response Elements

The following element is returned by the service.

**StackPolicyBody**

Structure that contains the stack policy body. For more information, see [Prevent updates to stack resources](../../../../services/cloudformation/latest/userguide/protect-stack-resources.md) in the
_AWS CloudFormation User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16384.

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

## Examples

### GetStackPolicy

This example illustrates one usage of GetStackPolicy.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=GetStackPolicy
 &StackName=MyStack
 &Version=2010-05-15
 &SignatureVersion=2
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<GetStackPolicyResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <GetStackPolicyResult>
      <StackPolicyBody>"{
      "Statement" : [
        {
          "Effect" : "Deny",
          "Action" : "Update:*",
          "Principal" : "*",
          "Resource" : "LogicalResourceId/ProductionDatabase"
        },
        {
          "Effect" : "Allow",
          "Action" : "Update:*",
          "Principal" : "*",
          "Resource" : "*"
        }
      ]
    }</StackPolicyBody>
  </GetStackPolicyResult>
  <ResponseMetadata>
    <RequestId>b9b4b068-3a41-11e5-94eb-example</RequestId>
  </ResponseMetadata>
</GetStackPolicyResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/getstackpolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/getstackpolicy.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

GetHookResult

GetTemplate
