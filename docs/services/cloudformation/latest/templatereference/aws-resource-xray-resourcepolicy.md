This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::XRay::ResourcePolicy

Use `AWS::XRay::ResourcePolicy` to specify an X-Ray resource-based policy,
which grants one or more AWS services and accounts permissions
to access X-Ray. Each resource-based policy is associated with a
specific AWS account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::XRay::ResourcePolicy",
  "Properties" : {
      "BypassPolicyLockoutCheck" : Boolean,
      "PolicyDocument" : String,
      "PolicyName" : String
    }
}

```

### YAML

```yaml

Type: AWS::XRay::ResourcePolicy
Properties:
  BypassPolicyLockoutCheck: Boolean
  PolicyDocument: String
  PolicyName: String

```

## Properties

`BypassPolicyLockoutCheck`

A flag to indicate whether to bypass the resource-based policy lockout safety check.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyDocument`

The resource-based policy document, which can be up to 5kb in size.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `5120`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PolicyName`

The name of the resource-based policy. Must be unique within a specific AWS account.

_Required_: Yes

_Type_: String

_Pattern_: `[\w+=,.@-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the policy name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## Examples

### Create resource-based policy

This example creates a resource-based policy called MySnsActiveTracingPolicy,
which configures Amazon SNS active tracing.

#### JSON

```json

{
   "AWSTemplateFormatVersion": "2010-09-09",
   "Resources": {
      "MyResourcePolicy": {
         "Type": "AWS::XRay::ResourcePolicy",
         "Properties": {
            "BypassPolicyLockoutCheck": "false",
            "PolicyDocument": "{
                \"Version\": \"2012-10-17\",
                \"Statement\": [
                  {
                    \"Sid\": \"SNSAccess\",
                    \"Effect\": \"Allow\",
                    \"Principal\": {
                      \"Service\": \"sns.amazonaws.com\"
                    },
                    \"Action\": [
                      \"xray:PutTraceSegments\",
                      \"xray:GetSamplingRules\",
                      \"xray:GetSamplingTargets\"
                    ],
                    \"Resource\": \"*\",
                    \"Condition\": {
                      \"StringEquals\": {
                        \"aws:SourceAccount\": \"my-account-id\"
                      },
                      \"StringLike\": {
                        \"aws:SourceArn\": \"arn:$my-partition:sns:$my-region:$my-account-id:my-topic-name\"
                      }
                    }
                  }
                ]
              }",
            "PolicyName": "MySnsActiveTracingPolicy"
            }
         }
      }
   }
}

```

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Resources:
  MyResourcePolicy:
    Type: AWS::XRay::ResourcePolicy
    Properties:
      BypassPolicyLockoutCheck: false
      PolicyDocument: "{
          \"Version\": \"2012-10-17\",
          \"Statement\": [
            {
              \"Sid\": \"SNSAccess\",
              \"Effect\": \"Allow\",
              \"Principal\": {
                \"Service\": \"sns.amazonaws.com\"
              },
              \"Action\": [
                \"xray:PutTraceSegments\",
                \"xray:GetSamplingRules\",
                \"xray:GetSamplingTargets\"
              ],
              \"Resource\": \"*\",
              \"Condition\": {
                \"StringEquals\": {
                  \"aws:SourceAccount\": \"my-account-id\"
                },
                \"StringLike\": {
                  \"aws:SourceArn\": \"arn:$my-partition:sns:$my-region:$my-account-id:my-topic-name\"
                }
              }
            }
          ]
        }"
      PolicyName: "MySNSActiveTracingPolicy"

```

## See also

- [X-Ray resource-based policies](https://docs.aws.amazon.com/xray/latest/devguide/security_iam_service-with-iam.html#security_iam_service-with-iam-resource-based-policies)

- [Identity-based policies and resource-based policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies_identity-vs-resource.html)

- [PutResourcePolicy](https://docs.aws.amazon.com/xray/latest/api/API_PutResourcePolicy.html) action in the X-Ray API Reference

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::XRay::SamplingRule
