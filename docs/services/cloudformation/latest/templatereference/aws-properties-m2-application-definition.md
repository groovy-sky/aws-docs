This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::M2::Application Definition

The application definition for a particular application. You can specify either inline
JSON or an Amazon S3 bucket location.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Content" : String,
  "S3Location" : String
}

```

### YAML

```yaml

  Content: String
  S3Location: String

```

## Properties

`Content`

The content of the application definition. This is a JSON object that contains the
resource configuration/definitions that identify an application.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `65000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Location`

The S3 bucket that contains the application definition.

_Required_: No

_Type_: String

_Pattern_: `^\S{1,2000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::M2::Application

AWS::M2::Deployment
