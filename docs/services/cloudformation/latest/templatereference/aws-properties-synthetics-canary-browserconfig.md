This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary BrowserConfig

A structure that specifies the browser type to use for a canary run.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BrowserType" : String
}

```

### YAML

```yaml

  BrowserType: String

```

## Properties

`BrowserType`

The browser type associated with this browser configuration.

_Required_: Yes

_Type_: String

_Allowed values_: `CHROME | FIREFOX`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BaseScreenshot

Code

All content copied from https://docs.aws.amazon.com/.
