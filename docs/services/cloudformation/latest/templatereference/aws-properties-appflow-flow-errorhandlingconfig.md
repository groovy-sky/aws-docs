This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::AppFlow::Flow ErrorHandlingConfig

The settings that determine how Amazon AppFlow handles an error when placing data in
the destination. For example, this setting would determine if the flow should fail after one
insertion error, or continue and attempt to insert every record regardless of the initial
failure. `ErrorHandlingConfig` is a part of the destination connector details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BucketName" : String,
  "BucketPrefix" : String,
  "FailOnFirstError" : Boolean
}

```

### YAML

```yaml

  BucketName: String
  BucketPrefix: String
  FailOnFirstError: Boolean

```

## Properties

`BucketName`

Specifies the name of the Amazon S3 bucket.

_Required_: No

_Type_: String

_Pattern_: `\S+`

_Minimum_: `3`

_Maximum_: `63`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`BucketPrefix`

Specifies the Amazon S3 bucket prefix.

_Required_: No

_Type_: String

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FailOnFirstError`

Specifies if the flow should fail after the first instance of a failure when attempting
to place data in the destination.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [ErrorHandlingConfig](../../../../reference/appflow/1-0/apireference/api-errorhandlingconfig.md) in the _Amazon AppFlow API_
_Reference_.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

DynatraceSourceProperties

EventBridgeDestinationProperties
