This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::Workteam OidcMemberDefinition

A list of user groups that exist in your OIDC Identity Provider (IdP).
One to ten groups can be used to create a single private work team.
When you add a user group to the list of `Groups`, you can add that user group to one or more
private work teams. If you add a user group to a private work team, all workers in that user group
are added to the work team.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OidcGroups" : [ String, ... ]
}

```

### YAML

```yaml

  OidcGroups:
    - String

```

## Properties

`OidcGroups`

A list of OpenID Connect (OIDC) groups for the work team member definition.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NotificationConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
