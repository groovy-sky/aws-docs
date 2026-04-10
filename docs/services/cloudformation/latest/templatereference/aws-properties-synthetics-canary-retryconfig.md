This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary RetryConfig

The canary's retry configuration information.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxRetries" : Integer
}

```

### YAML

```yaml

  MaxRetries: Integer

```

## Properties

`MaxRetries`

The maximum number of retries. The value must be less than or equal to two.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Dependency

RunConfig

All content copied from https://docs.aws.amazon.com/.
