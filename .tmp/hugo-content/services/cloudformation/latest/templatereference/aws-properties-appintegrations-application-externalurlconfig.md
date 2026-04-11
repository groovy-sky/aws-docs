This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppIntegrations::Application ExternalUrlConfig

The external URL source for the application.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AccessUrl" : String,
  "ApprovedOrigins" : [ String, ... ]
}

```

### YAML

```yaml

  AccessUrl: String
  ApprovedOrigins:
    - String

```

## Properties

`AccessUrl`

The URL to access the application.

_Required_: Yes

_Type_: String

_Pattern_: `^\w+\:\/\/.*$`

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApprovedOrigins`

Additional URLs to allow list if different than the access URL.

_Required_: No

_Type_: Array of String

_Minimum_: `0`

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ContactHandling

IframeConfig

All content copied from https://docs.aws.amazon.com/.
