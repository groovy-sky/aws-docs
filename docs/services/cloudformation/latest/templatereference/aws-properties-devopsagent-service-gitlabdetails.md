This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service GitLabDetails

Configuration details for registering a GitLab service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GroupId" : String,
  "TargetUrl" : String,
  "TokenType" : String,
  "TokenValue" : String
}

```

### YAML

```yaml

  GroupId: String
  TargetUrl: String
  TokenType: String
  TokenValue: String

```

## Properties

`GroupId`

The GitLab group ID. Required when `TokenType` is `group`.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetUrl`

The GitLab instance URL. Must be an HTTPS URL.

_Required_: Yes

_Type_: String

_Pattern_: `^https://[a-zA-Z0-9]([a-zA-Z0-9.-]*[a-zA-Z0-9])?(?::[0-9]{1,5})?/?$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TokenType`

The type of GitLab access token.

_Allowed Values_: `personal` \| `group`

_Required_: Yes

_Type_: String

_Allowed values_: `personal | group`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TokenValue`

The GitLab access token value. Must match the pattern `^glpat-[a-zA-Z0-9._-]+$`.

_Required_: Yes

_Type_: String

_Pattern_: `^glpat-[a-zA-Z0-9._-]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DynatraceServiceDetails

MCPServerAuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
