This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudFront::ContinuousDeploymentPolicy SessionStickinessConfig

Session stickiness provides the ability to define multiple requests from a single
viewer as a single session. This prevents the potentially inconsistent experience of
sending some of a given user's requests to your staging distribution, while others are
sent to your primary distribution. Define the session duration using TTL values.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdleTTL" : Integer,
  "MaximumTTL" : Integer
}

```

### YAML

```yaml

  IdleTTL: Integer
  MaximumTTL: Integer

```

## Properties

`IdleTTL`

The amount of time after which you want sessions to cease if no requests are
received. Allowed values are 300–3600 seconds (5–60 minutes).

_Required_: Yes

_Type_: Integer

_Minimum_: `300`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MaximumTTL`

The maximum amount of time to consider requests from the viewer as being part of the same
session. Allowed values are 300–3600 seconds (5–60 minutes).

_Required_: Yes

_Type_: Integer

_Minimum_: `300`

_Maximum_: `3600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContinuousDeploymentPolicyConfig

SingleHeaderConfig

All content copied from https://docs.aws.amazon.com/.
