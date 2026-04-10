This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::Service NewRelicApiKeyConfig

The API key configuration for a New Relic service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccountId" : String,
  "AlertPolicyIds" : [ String, ... ],
  "ApiKey" : String,
  "ApplicationIds" : [ String, ... ],
  "EntityGuids" : [ String, ... ],
  "Region" : String
}

```

### YAML

```yaml

  AccountId: String
  AlertPolicyIds:
    - String
  ApiKey: String
  ApplicationIds:
    - String
  EntityGuids:
    - String
  Region: String

```

## Properties

`AccountId`

The New Relic account ID.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`AlertPolicyIds`

List of New Relic alert policy IDs.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApiKey`

The New Relic User API key. Must match the pattern `^NRAK-[A-Z0-9]+$`.

_Required_: Yes

_Type_: String

_Pattern_: `^NRAK-[A-Z0-9]+$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ApplicationIds`

List of New Relic APM application IDs to monitor.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EntityGuids`

List of globally unique IDs for New Relic resources.

_Required_: No

_Type_: Array of String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Region`

The New Relic region.

_Allowed Values_: `US` \| `EU`

_Required_: Yes

_Type_: String

_Allowed values_: `US | EU`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MCPServerSplunkDetails

NewRelicAuthorizationConfig

All content copied from https://docs.aws.amazon.com/.
