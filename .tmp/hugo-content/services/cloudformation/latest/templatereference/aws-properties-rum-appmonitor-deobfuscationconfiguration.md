This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::RUM::AppMonitor DeobfuscationConfiguration

A structure that contains the configuration for how an app monitor can deobfuscate stack traces.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "JavaScriptSourceMaps" : JavaScriptSourceMaps
}

```

### YAML

```yaml

  JavaScriptSourceMaps:
    JavaScriptSourceMaps

```

## Properties

`JavaScriptSourceMaps`

A structure that contains the configuration for how an app monitor can unminify JavaScript error stack traces using source maps.

_Required_: No

_Type_: [JavaScriptSourceMaps](aws-properties-rum-appmonitor-javascriptsourcemaps.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CustomEvents

JavaScriptSourceMaps

All content copied from https://docs.aws.amazon.com/.
