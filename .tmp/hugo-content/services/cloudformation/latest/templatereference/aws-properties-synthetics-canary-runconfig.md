This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary RunConfig

A structure that contains input information for a canary run. This structure
is required.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ActiveTracing" : Boolean,
  "EnvironmentVariables" : {Key: Value, ...},
  "EphemeralStorage" : Integer,
  "MemoryInMB" : Integer,
  "TimeoutInSeconds" : Integer
}

```

### YAML

```yaml

  ActiveTracing: Boolean
  EnvironmentVariables:
    Key: Value
  EphemeralStorage: Integer
  MemoryInMB: Integer
  TimeoutInSeconds: Integer

```

## Properties

`ActiveTracing`

Specifies whether this canary is to use active AWS X-Ray tracing when it runs. Active tracing
enables this canary run to be displayed in the ServiceLens and X-Ray service maps even if the
canary does not hit an endpoint that has X-Ray tracing enabled. Using X-Ray tracing
incurs charges. For more information, see
[Canaries and X-Ray tracing](../../../amazoncloudwatch/latest/monitoring/cloudwatch-synthetics-canaries-tracing.md).

You can enable active tracing only for canaries that use version `syn-nodejs-2.0` or later
for their canary runtime.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnvironmentVariables`

Specifies the keys and values to use for any environment variables
used in the canary script. Use the following format:

{ "key1" : "value1", "key2" : "value2", ...}

Keys must start with a letter and be at least two characters. The total size
of your environment variables cannot exceed 4 KB. You can't specify any Lambda
reserved environment variables as the keys for your environment variables. For
more information about reserved keys, see [Runtime environment variables](../../../lambda/latest/dg/configuration-envvars.md#configuration-envvars-runtime).

_Required_: No

_Type_: Object of String

_Pattern_: `[a-zA-Z][a-zA-Z0-9_]+`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EphemeralStorage`

Specifies the amount of ephemeral storage (in MB) to allocate for the canary run during execution. This temporary storage is used for storing canary run artifacts (which are uploaded to an Amazon S3
bucket at the end of the run), and any canary browser operations. This temporary storage is cleared after the run is completed. Default storage value is 1024 MB.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MemoryInMB`

The maximum amount of memory that the
canary can use while running. This value must be a multiple of 64. The range is 960 to 3008.

_Required_: No

_Type_: Integer

_Minimum_: `960`

_Maximum_: `3008`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TimeoutInSeconds`

How long the canary is allowed to run before it must stop. You can't set this time to be longer
than the frequency of the runs of this canary.

If you omit this field, the
frequency of the canary is used as this value, up to a maximum of 900 seconds.

_Required_: No

_Type_: Integer

_Minimum_: `3`

_Maximum_: `840`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

RetryConfig

S3Encryption

All content copied from https://docs.aws.amazon.com/.
