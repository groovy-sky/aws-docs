This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodePipeline::CustomActionType Settings

`Settings` is a property of the `AWS::CodePipeline::CustomActionType`
resource that provides URLs that users can access to view information about the CodePipeline
custom action.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EntityUrlTemplate" : String,
  "ExecutionUrlTemplate" : String,
  "RevisionUrlTemplate" : String,
  "ThirdPartyConfigurationUrl" : String
}

```

### YAML

```yaml

  EntityUrlTemplate: String
  ExecutionUrlTemplate: String
  RevisionUrlTemplate: String
  ThirdPartyConfigurationUrl: String

```

## Properties

`EntityUrlTemplate`

The URL returned to the CodePipeline console that provides a deep link to the
resources of the external system, such as the configuration page for a CodeDeploy
deployment group. This link is provided as part of the action display in the
pipeline.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ExecutionUrlTemplate`

The URL returned to the CodePipeline console that contains a link to the
top-level landing page for the external system, such as the console page for CodeDeploy.
This link is shown on the pipeline view page in the CodePipeline console and
provides a link to the execution entity of the external action.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RevisionUrlTemplate`

The URL returned to the CodePipeline console that contains a link to the page
where customers can update or change the configuration of the external action.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ThirdPartyConfigurationUrl`

The URL of a sign-up page where users can sign up for an external service and
perform initial configuration of the action provided by that service.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationProperties

Tag
