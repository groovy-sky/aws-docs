This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function Environment

A function's environment variable settings. You can use environment variables to adjust your function's
behavior without updating code. An environment variable is a pair of strings that are stored in a function's
version-specific configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Variables" : {Key: Value, ...}
}

```

### YAML

```yaml

  Variables:
    Key: Value

```

## Properties

`Variables`

Environment variable key-value pairs. For more information, see [Using Lambda environment variables](../../../lambda/latest/dg/configuration-envvars.md).

If the value of the environment variable is a time or a duration, enclose the value in quotes.

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z][a-zA-Z0-9_]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Environment Variables

Add environment variables to a function. Each variable is a key-value pair.
This example specifies values for a `databaseName` and a `databaseUser`.

#### YAML

```yaml

      Environment:
        Variables:
          databaseName: lambdadb
          databaseUser: admin
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DurableConfig

EphemeralStorage

All content copied from https://docs.aws.amazon.com/.
