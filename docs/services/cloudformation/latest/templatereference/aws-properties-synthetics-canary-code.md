This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary Code

Use this structure to input your script code for the canary. This structure contains the
Lambda handler with the location where the canary should start running the script. If the
script is stored in an S3 bucket, the bucket name, key, and version are also included. If
the script is passed into the canary directly, the script code is contained in the value
of `Script`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlueprintTypes" : [ String, ... ],
  "Dependencies" : [ Dependency, ... ],
  "Handler" : String,
  "S3Bucket" : String,
  "S3Key" : String,
  "S3ObjectVersion" : String,
  "Script" : String,
  "SourceLocationArn" : String
}

```

### YAML

```yaml

  BlueprintTypes:
    - String
  Dependencies:
    - Dependency
  Handler: String
  S3Bucket: String
  S3Key: String
  S3ObjectVersion: String
  Script: String
  SourceLocationArn: String

```

## Properties

`BlueprintTypes`

`BlueprintTypes` are a list of templates that enable simplified canary creation. You can create canaries for common monitoring scenarios by providing only a JSON configuration file instead of writing custom scripts. `multi-checks` is the only supported value.

When you specify `BlueprintTypes`, the `Handler` field cannot be specified since the blueprint provides a pre-defined entry point.

_Required_: No

_Type_: Array of String

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Dependencies`

Property description not available.

_Required_: No

_Type_: Array of [Dependency](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-synthetics-canary-dependency.html)

_Maximum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Handler`

The entry point to use for the source code when running the canary. For canaries that use the
`syn-python-selenium-1.0` runtime
or a `syn-nodejs.puppeteer` runtime earlier than `syn-nodejs.puppeteer-3.4`,
the handler must be specified as `fileName.handler`. For
`syn-python-selenium-1.1`, `syn-nodejs.puppeteer-3.4`, and later runtimes, the handler can be specified as
`fileName.functionName`, or
you can specify a folder where canary scripts reside as
`folder/fileName.functionName`.

This field is required when you don't specify `BlueprintTypes` and is not allowed when you specify `BlueprintTypes`.

_Required_: Conditional

_Type_: String

_Pattern_: `^(([0-9a-zA-Z_-]+(\/|\.))*[0-9A-Za-z_\\-]+(\.|::)[A-Za-z_][A-Za-z0-9_]*)?$`

_Minimum_: `0`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Bucket`

If your canary script is located in S3, specify the bucket name here. The bucket
must already exist.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3Key`

The Amazon S3 key of your script. For more information, see [Working with Amazon S3 Objects](https://docs.aws.amazon.com/AmazonS3/latest/dev/UsingObjects.html).

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3ObjectVersion`

The Amazon S3 version ID of your script.

_Required_: Conditional

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Script`

If you input your canary script directly into the canary instead of referring to an S3
location, the value of this parameter is the script in plain text. It can be
up to 5 MB.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SourceLocationArn`

The ARN of the Lambda layer where Synthetics stores the canary script code.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

BrowserConfig

Dependency
