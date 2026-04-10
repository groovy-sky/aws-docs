This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Association AWSConfiguration

Configuration for AWS monitor account integration. Specifies the account ID, assumable role ARN,
and resources to be monitored in the primary monitoring account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "AccountType" : String,
  "AssumableRoleArn" : String,
  "Resources" : [ AWSResource, ... ],
  "Tags" : [ KeyValuePair, ... ]
}

```

### YAML

```yaml

  AccountId: String
  AccountType: String
  AssumableRoleArn: String
  Resources:
    - AWSResource
  Tags:
    - KeyValuePair

```

## Properties

`AccountId`

The 12-digit AWS account ID corresponding to the provided resources.

_Required_: Yes

_Type_: String

_Pattern_: `\d{12}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AccountType`

The account type for AWS DevOps Agent monitoring.

_Allowed Values_: `monitor`

_Required_: Yes

_Type_: String

_Allowed values_: `monitor`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AssumableRoleArn`

Role ARN used by AWS DevOps Agent to access resources in the primary account.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Resources`

List of resources to monitor.

_Required_: No

_Type_: Array of [AWSResource](aws-properties-devopsagent-association-awsresource.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

List of tags as key-value pairs, used to identify resources for topology crawl.

_Required_: No

_Type_: Array of [KeyValuePair](aws-properties-devopsagent-association-keyvaluepair.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DevOpsAgent::Association

AWSResource

All content copied from https://docs.aws.amazon.com/.
