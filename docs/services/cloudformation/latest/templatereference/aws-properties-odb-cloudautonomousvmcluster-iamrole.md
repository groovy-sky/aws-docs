This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ODB::CloudAutonomousVmCluster IamRole

Information about an AWS Identity and Access Management (IAM) service role associated with a resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AwsIntegration" : String,
  "IamRoleArn" : String,
  "Status" : String
}

```

### YAML

```yaml

  AwsIntegration: String
  IamRoleArn: String
  Status: String

```

## Properties

`AwsIntegration`

The AWS integration configuration settings for the AWS Identity and Access Management (IAM) service role.

_Required_: No

_Type_: String

_Allowed values_: `KmsTde`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IamRoleArn`

The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) service role.

_Required_: No

_Type_: String

_Pattern_: `arn:(?:aws|aws-cn|aws-us-gov|aws-iso-{0,1}[a-z]{0,1}):iam::[0-9]{12}:role/.+`

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

The current status of the AWS Identity and Access Management (IAM) service role.

_Required_: No

_Type_: String

_Allowed values_: `ASSOCIATING | DISASSOCIATING | FAILED | CONNECTED | DISCONNECTED | PARTIALLY_CONNECTED | UNKNOWN`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::ODB::CloudAutonomousVmCluster

MaintenanceWindow

All content copied from https://docs.aws.amazon.com/.
