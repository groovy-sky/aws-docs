This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Scraper RoleConfiguration

The role configuration in an Amazon Managed Service for Prometheus scraper.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceRoleArn" : String,
  "TargetRoleArn" : String
}

```

### YAML

```yaml

  SourceRoleArn: String
  TargetRoleArn: String

```

## Properties

`SourceRoleArn`

The ARN of the source role.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetRoleArn`

The ARN of the target role.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EksConfiguration

ScrapeConfiguration
