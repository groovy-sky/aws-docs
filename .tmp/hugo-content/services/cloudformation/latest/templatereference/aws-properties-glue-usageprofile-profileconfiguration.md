This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::UsageProfile ProfileConfiguration

Specifies the job and session values that an admin configures in an AWS Glue usage profile.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JobConfiguration" : ConfigurationObject,
  "SessionConfiguration" : ConfigurationObject
}

```

### YAML

```yaml

  JobConfiguration:
    ConfigurationObject
  SessionConfiguration:
    ConfigurationObject

```

## Properties

`JobConfiguration`

A key-value map of configuration parameters for AWS Glue jobs.

_Required_: No

_Type_: [ConfigurationObject](aws-properties-glue-usageprofile-configurationobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SessionConfiguration`

A key-value map of configuration parameters for AWS Glue sessions.

_Required_: No

_Type_: [ConfigurationObject](aws-properties-glue-usageprofile-configurationobject.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ConfigurationObject

Tag

All content copied from https://docs.aws.amazon.com/.
