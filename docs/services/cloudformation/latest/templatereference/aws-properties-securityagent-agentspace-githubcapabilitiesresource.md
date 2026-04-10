This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityAgent::AgentSpace GitHubCapabilitiesResource

The capabilities enabled for a GitHub resource integration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LeaveComments" : Boolean,
  "RemediateCode" : Boolean
}

```

### YAML

```yaml

  LeaveComments: Boolean
  RemediateCode: Boolean

```

## Properties

`LeaveComments`

Indicates whether the integration can leave comments on pull requests.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RemediateCode`

Indicates whether the integration can create code remediation pull requests.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CodeReviewSettings

GitHubRepositoryResource

All content copied from https://docs.aws.amazon.com/.
