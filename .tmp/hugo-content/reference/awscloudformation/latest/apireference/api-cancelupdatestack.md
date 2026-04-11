# CancelUpdateStack

Cancels an update on the specified stack. If the call completes successfully, the stack
rolls back the update and reverts to the previous stack configuration.

###### Note

You can cancel only stacks that are in the `UPDATE_IN_PROGRESS` state.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**ClientRequestToken**

A unique identifier for this `CancelUpdateStack` request. Specify this token if
you plan to retry requests so that CloudFormation knows that you're not attempting to cancel an
update on a stack with the same name. You might retry `CancelUpdateStack` requests
to ensure that CloudFormation successfully received them.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `[a-zA-Z0-9][-a-zA-Z0-9]*`

Required: No

**StackName**

###### Note

If you don't pass a parameter to `StackName`, the API returns a response that
describes all resources in the account.

The IAM policy below can be added to IAM policies when you want to limit
resource-level permissions and avoid returning a response when no parameter is sent in the
request:

`{ "Version": "2012-10-17",		 	 	  "Statement": [{ "Effect": "Deny",
          "Action": "cloudformation:DescribeStacks", "NotResource":
          "arn:aws:cloudformation:*:*:stack/*/*" }] }`

The name or the unique stack ID that's associated with the stack.

Type: String

Required: Yes

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**TokenAlreadyExists**

A client request token already exists.

HTTP Status Code: 400

## Examples

### CancelUpdateStack

This example illustrates one usage of CancelUpdateStack.

#### Sample Request

```

https://cloudformation.us-east-1.amazonaws.com/
 ?Action=CancelUpdateStack
 &StackName=MyStack
 &Version=2010-05-15
 &SignatureVersion=2
 &Timestamp=2010-07-27T22%3A26%3A28.000Z
 &AWSAccessKeyId=[AWS Access KeyID]
 &Signature=[Signature]
```

#### Sample Response

```

<CancelUpdateStackResponse xmlns="http://cloudformation.amazonaws.com/doc/2010-05-15/">
  <ResponseMetadata>
    <RequestId>5ccc7dcd-744c-11e5-be70-1b08c228efb3</RequestId>
  </ResponseMetadata>
</CancelUpdateStackResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/cancelupdatestack.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/cancelupdatestack.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BatchDescribeTypeConfigurations

ContinueUpdateRollback

All content copied from https://docs.aws.amazon.com/.
