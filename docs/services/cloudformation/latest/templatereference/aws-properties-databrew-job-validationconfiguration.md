This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job ValidationConfiguration

Configuration for data quality validation. Used to select the Rulesets and Validation Mode
to be used in the profile job. When ValidationConfiguration is null, the profile
job will run without data quality validation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RulesetArn" : String,
  "ValidationMode" : String
}

```

### YAML

```yaml

  RulesetArn: String
  ValidationMode: String

```

## Properties

`RulesetArn`

The Amazon Resource Name (ARN) for the ruleset to be validated in the profile job.
The TargetArn of the selected ruleset should be the same as the Amazon Resource Name (ARN) of
the dataset that is associated with the profile job.

_Required_: Yes

_Type_: String

_Minimum_: `20`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ValidationMode`

Mode of data quality validation. Default mode is “CHECK\_ALL” which verifies all rules
defined in the selected ruleset.

_Required_: No

_Type_: String

_Allowed values_: `CHECK_ALL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::DataBrew::Project
