This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppRunner::ObservabilityConfiguration TraceConfiguration

Describes the configuration of the tracing feature within an AWS App Runner observability configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Vendor" : String
}

```

### YAML

```yaml

  Vendor: String

```

## Properties

`Vendor`

The implementation provider chosen for tracing App Runner services.

_Required_: Yes

_Type_: String

_Allowed values_: `AWSXRAY`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

AWS::AppRunner::Service
