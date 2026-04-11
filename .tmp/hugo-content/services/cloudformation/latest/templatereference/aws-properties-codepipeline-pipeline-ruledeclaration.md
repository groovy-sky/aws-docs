This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::Pipeline RuleDeclaration

Represents information about the rule to be created for an associated condition. An
example would be creating a new rule for an entry condition, such as a rule that checks
for a test result before allowing the run to enter the deployment stage. For more
information about conditions, see [Stage conditions](../../../codepipeline/latest/userguide/stage-conditions.md)
and [How do\
stage conditions work?](../../../codepipeline/latest/userguide/concepts-how-it-works-conditions.md). For more information about rules, see the [AWS CodePipeline rule reference](../../../codepipeline/latest/userguide/rule-reference.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Commands" : [ String, ... ],
  "Configuration" : Json,
  "InputArtifacts" : [ InputArtifact, ... ],
  "Name" : String,
  "Region" : String,
  "RoleArn" : String,
  "RuleTypeId" : RuleTypeId
}

```

### YAML

```yaml

  Commands:
    - String
  Configuration: Json
  InputArtifacts:
    - InputArtifact
  Name: String
  Region: String
  RoleArn: String
  RuleTypeId:
    RuleTypeId

```

## Properties

`Commands`

The shell commands to run with your commands rule in CodePipeline. All commands
are supported except multi-line formats. While CodeBuild logs and permissions
are used, you do not need to create any resources in CodeBuild.

###### Note

Using compute time for this action will incur separate charges in AWS CodeBuild.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Configuration`

The action configuration fields for the rule.

_Required_: No

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputArtifacts`

The input artifacts fields for the rule, such as specifying an input file for the
rule.

_Required_: No

_Type_: Array of [InputArtifact](aws-properties-codepipeline-pipeline-inputartifact.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the rule that is created for the condition, such as
`VariableCheck`.

_Required_: No

_Type_: String

_Pattern_: `[A-Za-z0-9.@\-_]+`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The Region for the condition associated with the rule.

_Required_: No

_Type_: String

_Minimum_: `4`

_Maximum_: `30`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RoleArn`

The pipeline role ARN associated with the rule.

_Required_: No

_Type_: String

_Pattern_: `arn:aws(-[\w]+)*:iam::[0-9]{12}:role/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RuleTypeId`

The ID for the rule type, which is made up of the combined values for category, owner,
provider, and version.

_Required_: No

_Type_: [RuleTypeId](aws-properties-codepipeline-pipeline-ruletypeid.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetryConfiguration

RuleTypeId

All content copied from https://docs.aws.amazon.com/.
