This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BedrockAgentCore::Gateway GatewayPolicyEngineConfiguration

The configuration for a policy engine associated with a gateway. A policy engine is a collection of policies that evaluates and authorizes agent tool calls. When associated with a gateway, the policy engine intercepts all agent requests and determines whether to allow or deny each action based on the defined policies.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Arn" : String,
  "Mode" : String
}

```

### YAML

```yaml

  Arn: String
  Mode: String

```

## Properties

`Arn`

The ARN of the policy engine. The policy engine contains Cedar policies that define fine-grained authorization rules specifying who can perform what actions on which resources as agents interact through the gateway.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:[a-z0-9-]{1,20}:bedrock-agentcore:[a-z0-9-]+:[0-9]{12}:policy-engine/[a-zA-Z][a-zA-Z0-9-_]{0,99}-[a-zA-Z0-9_]{10}$`

_Minimum_: `1`

_Maximum_: `170`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Mode`

The enforcement mode for the policy engine. Valid values include:

- `LOG_ONLY` \- The policy engine evaluates each action against your policies and adds traces on whether tool calls would be allowed or denied, but does not enforce the decision. Use this mode to test and validate policies before enabling enforcement.

- `ENFORCE` \- The policy engine evaluates actions against your policies and enforces decisions by allowing or denying agent operations. Test and validate policies in `LOG_ONLY` mode before enabling enforcement to avoid unintended denials or adversely affecting production traffic.

_Required_: Yes

_Type_: String

_Allowed values_: `LOG_ONLY | ENFORCE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

GatewayInterceptorConfiguration

GatewayProtocolConfiguration
