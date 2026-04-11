This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Amplify::App CacheConfig

Describes the cache configuration for an Amplify app.

For more information about how Amplify applies an optimal cache
configuration for your app based on the type of content that is being served, see [Managing cache configuration](../../../amplify/latest/userguide/managing-cache-configuration.md) in the _Amplify User_
_guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String
}

```

### YAML

```yaml

  Type: String

```

## Properties

`Type`

The type of cache configuration to use for an Amplify app.

The `AMPLIFY_MANAGED` cache configuration automatically applies an
optimized cache configuration for your app based on its platform, routing rules, and
rewrite rules.

The `AMPLIFY_MANAGED_NO_COOKIES` cache configuration type is the same as
`AMPLIFY_MANAGED`, except that it excludes all cookies from the cache
key. This is the default setting.

_Required_: No

_Type_: String

_Allowed values_: `AMPLIFY_MANAGED | AMPLIFY_MANAGED_NO_COOKIES`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

BasicAuthConfig

CustomRule

All content copied from https://docs.aws.amazon.com/.
