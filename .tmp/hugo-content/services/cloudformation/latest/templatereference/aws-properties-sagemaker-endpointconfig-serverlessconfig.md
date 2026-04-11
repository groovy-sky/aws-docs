This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ServerlessConfig

Specifies the serverless configuration for an endpoint variant.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxConcurrency" : Integer,
  "MemorySizeInMB" : Integer,
  "ProvisionedConcurrency" : Integer
}

```

### YAML

```yaml

  MaxConcurrency: Integer
  MemorySizeInMB: Integer
  ProvisionedConcurrency: Integer

```

## Properties

`MaxConcurrency`

The maximum number of concurrent invocations your serverless endpoint can process.

_Required_: Yes

_Type_: Integer

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MemorySizeInMB`

The memory size of your serverless endpoint. Valid values are in 1 GB increments: 1024 MB, 2048 MB, 3072 MB, 4096 MB, 5120 MB, or 6144 MB.

_Required_: Yes

_Type_: Integer

_Minimum_: `1024`

_Maximum_: `6144`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProvisionedConcurrency`

The amount of provisioned concurrency to allocate for the serverless endpoint.
Should be less than or equal to `MaxConcurrency`.

###### Note

This field is not supported for serverless endpoint recommendations for Inference Recommender jobs.
For more information about creating an Inference Recommender job, see
[CreateInferenceRecommendationsJobs](../../../../reference/sagemaker/latest/apireference/api-createinferencerecommendationsjob.md).

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProductionVariant

Tag

All content copied from https://docs.aws.amazon.com/.
