This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FIS::TargetAccountConfiguration

Creates a target account configuration for the experiment template. A target account configuration
is required when `accountTargeting` of `experimentOptions` is set to `multi-account`.
For more information, see [experiment options](https://docs.aws.amazon.com/fis/latest/userguide/experiment-options.html)
in the _AWS Fault Injection Service User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FIS::TargetAccountConfiguration",
  "Properties" : {
      "AccountId" : String,
      "Description" : String,
      "ExperimentTemplateId" : String,
      "RoleArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::FIS::TargetAccountConfiguration
Properties:
  AccountId: String
  Description: String
  ExperimentTemplateId: String
  RoleArn: String

```

## Properties

`AccountId`

The AWS account ID of the target account.

_Required_: Yes

_Type_: String

_Maximum_: `512`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The description of the target account.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExperimentTemplateId`

The ID of the experiment template.

_Required_: Yes

_Type_: String

_Pattern_: `[\S]+`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The Amazon Resource Name (ARN) of an IAM role for the target account.

_Required_: Yes

_Type_: String

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Configuration

Next
