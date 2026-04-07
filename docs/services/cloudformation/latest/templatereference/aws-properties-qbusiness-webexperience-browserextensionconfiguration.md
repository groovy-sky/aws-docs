This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::QBusiness::WebExperience BrowserExtensionConfiguration

The container for browser extension configuration for an Amazon Q Business web experience.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EnabledBrowserExtensions" : [ String, ... ]
}

```

### YAML

```yaml

  EnabledBrowserExtensions:
    - String

```

## Properties

`EnabledBrowserExtensions`

Specify the browser extensions allowed for your Amazon Q web experience.

- `CHROME` — Enables the extension for Chromium-based browsers (Google Chrome, Microsoft
Edge, Opera, etc.).

- `FIREFOX` — Enables the extension for Mozilla Firefox.

- `CHROME` and `FIREFOX` — Enable the extension for Chromium-based browsers
and Mozilla Firefox.

_Required_: Yes

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `2`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::QBusiness::WebExperience

CustomizationConfiguration
