This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function TenancyConfig

Specifies the tenant isolation mode configuration for a Lambda function.
This allows you to configure specific tenant isolation strategies for your function invocations.
Tenant isolation configuration cannot be modified after function creation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TenantIsolationMode" : String
}

```

### YAML

```yaml

  TenantIsolationMode: String

```

## Properties

`TenantIsolationMode`

Tenant isolation mode allows for invocation to be sent to a
corresponding execution environment dedicated to a specific tenant ID.

_Required_: Yes

_Type_: String

_Allowed values_: `PER_TENANT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TracingConfig

All content copied from https://docs.aws.amazon.com/.
