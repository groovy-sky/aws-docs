This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Events::Endpoint Primary

The primary Region of the endpoint.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "HealthCheck" : String
}

```

### YAML

```yaml

  HealthCheck: String

```

## Properties

`HealthCheck`

The ARN of the health check used by the endpoint to determine whether failover is
triggered.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws([a-z]|\-)*:route53:::healthcheck/[\-a-z0-9]+$`

_Minimum_: `1`

_Maximum_: `1600`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FailoverConfig

ReplicationConfig

All content copied from https://docs.aws.amazon.com/.
