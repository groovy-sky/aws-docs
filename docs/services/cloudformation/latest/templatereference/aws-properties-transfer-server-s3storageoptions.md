This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Transfer::Server S3StorageOptions

The Amazon S3 storage options that are configured for your server.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DirectoryListingOptimization" : String
}

```

### YAML

```yaml

  DirectoryListingOptimization: String

```

## Properties

`DirectoryListingOptimization`

Specifies whether or not performance for your Amazon S3 directories is optimized.

- If using the console, this is enabled by default.

- If using the API or CLI, this is disabled by default.

By default, home directory mappings have a `TYPE` of `DIRECTORY`. If you enable this option, you would then need to explicitly set the `HomeDirectoryMapEntry` `Type` to `FILE` if you want a mapping to have a file target.

_Required_: No

_Type_: String

_Allowed values_: `ENABLED | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProtocolDetails

Tag
