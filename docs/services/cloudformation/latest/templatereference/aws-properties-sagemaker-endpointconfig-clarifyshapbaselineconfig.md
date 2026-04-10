This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ClarifyShapBaselineConfig

The configuration for the [SHAP\
baseline](../../../sagemaker/latest/dg/clarify-feature-attribute-shap-baselines.md) (also called the background or reference dataset) of the Kernal
SHAP algorithm.

###### Note

- The number of records in the baseline data determines the size of the
synthetic dataset, which has an impact on latency of explainability
requests. For more information, see the **Synthetic**
**data** of [Configure and create an endpoint](../../../sagemaker/latest/dg/clarify-online-explainability-create-endpoint.md).

- `ShapBaseline` and `ShapBaselineUri` are mutually
exclusive parameters. One or the either is required to configure a SHAP
baseline.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MimeType" : String,
  "ShapBaseline" : String,
  "ShapBaselineUri" : String
}

```

### YAML

```yaml

  MimeType: String
  ShapBaseline: String
  ShapBaselineUri: String

```

## Properties

`MimeType`

The MIME type of the baseline data. Choose from `'text/csv'` or
`'application/jsonlines'`. Defaults to `'text/csv'`.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9](-*[a-zA-Z0-9])*\/[a-zA-Z0-9](-*[a-zA-Z0-9+.])*`

_Minimum_: `0`

_Maximum_: `255`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShapBaseline`

The inline SHAP baseline data in string format. `ShapBaseline` can have one
or multiple records to be used as the baseline dataset. The format of the SHAP baseline
file should be the same format as the training dataset. For example, if the training
dataset is in CSV format and each record contains four features, and all features are
numerical, then the format of the baseline data should also share these characteristics.
For natural language processing (NLP) of text columns, the baseline value should be the
value used to replace the unit of text specified by the `Granularity` of the
`TextConfig` parameter. The size limit for `ShapBasline` is 4
KB. Use the `ShapBaselineUri` parameter if you want to provide more than 4 KB
of baseline data.

_Required_: No

_Type_: String

_Pattern_: `[\s\S]+`

_Minimum_: `1`

_Maximum_: `4096`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ShapBaselineUri`

The uniform resource identifier (URI) of the S3 bucket where the SHAP baseline file is
stored. The format of the SHAP baseline file should be the same format as the format of
the training dataset. For example, if the training dataset is in CSV format, and each
record in the training dataset has four features, and all features are numerical, then
the baseline file should also have this same format. Each record should contain only the
features. If you are using a virtual private cloud (VPC), the
`ShapBaselineUri` should be accessible to the VPC. For more information
about setting up endpoints with Amazon Virtual Private Cloud, see [Give SageMaker access to\
Resources in your Amazon Virtual Private Cloud](../../../sagemaker/latest/dg/infrastructure-give-access.md).

_Required_: No

_Type_: String

_Pattern_: `(https|s3)://([^/]+)/?(.*)`

_Minimum_: `0`

_Maximum_: `1024`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClarifyInferenceConfig

ClarifyShapConfig

All content copied from https://docs.aws.amazon.com/.
