This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace AWSResources

The Amazon Web Services resources associated with an agent space, including VPCs, log groups, S3 buckets, secrets, Lambda functions, and IAM roles.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IamRoles" : [ String, ... ],
  "LambdaFunctionArns" : [ String, ... ],
  "LogGroups" : [ String, ... ],
  "S3Buckets" : [ String, ... ],
  "SecretArns" : [ String, ... ],
  "Vpcs" : [ VpcConfig, ... ]
}

```

### YAML

```yaml

  IamRoles:
    - String
  LambdaFunctionArns:
    - String
  LogGroups:
    - String
  S3Buckets:
    - String
  SecretArns:
    - String
  Vpcs:
    - VpcConfig

```

## Properties

`IamRoles`

The IAM roles associated with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LambdaFunctionArns`

The Amazon Resource Names (ARNs) of the Lambda functions associated with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`LogGroups`

The Amazon Resource Names (ARNs) of the CloudWatch log groups associated with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Buckets`

The Amazon Resource Names (ARNs) of the S3 buckets associated with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretArns`

The Amazon Resource Names (ARNs) of the Secrets Manager secrets associated with the agent space.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Vpcs`

The VPC configurations associated with the agent space.

_Required_: No

_Type_: Array of [VpcConfig](aws-properties-securityagent-agentspace-vpcconfig.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityAgent::AgentSpace

CodeReviewSettings

All content copied from https://docs.aws.amazon.com/.
