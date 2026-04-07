This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::IntegrationResourceProperty SourceProcessingProperties

The structure used to define the resource properties associated with the integration source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RoleArn" : String
}

```

### YAML

```yaml

  RoleArn: String

```

## Properties

`RoleArn`

The IAM role to access the AWS Glue connection.

_Required_: Yes

_Type_: String

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::Glue::IntegrationResourceProperty

Tag
