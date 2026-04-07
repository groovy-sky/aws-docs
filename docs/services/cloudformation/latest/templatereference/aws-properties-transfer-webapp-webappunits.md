This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::WebApp WebAppUnits

Contains an integer value that represents the value for number of concurrent
connections or the user sessions on your web app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Provisioned" : Integer
}

```

### YAML

```yaml

  Provisioned: Integer

```

## Properties

`Provisioned`

An integer that represents the number of units for your desired number of concurrent
connections, or the number of user sessions on your web app at the same time.

Each increment allows an additional 250 concurrent sessions: a value of `1`
sets the number of concurrent sessions to 250; `2` sets a value of 500, and
so on.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WebAppCustomization

AWS::Transfer::Workflow
