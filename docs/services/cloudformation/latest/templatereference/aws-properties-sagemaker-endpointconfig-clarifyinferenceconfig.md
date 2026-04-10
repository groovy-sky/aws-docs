This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ClarifyInferenceConfig

The inference configuration parameter for the model container.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ContentTemplate" : String,
  "FeatureHeaders" : [ String, ... ],
  "FeaturesAttribute" : String,
  "FeatureTypes" : [ String, ... ],
  "LabelAttribute" : String,
  "LabelHeaders" : [ String, ... ],
  "LabelIndex" : Integer,
  "MaxPayloadInMB" : Integer,
  "MaxRecordCount" : Integer,
  "ProbabilityAttribute" : String,
  "ProbabilityIndex" : Integer
}

```

### YAML

```yaml

  ContentTemplate: String
  FeatureHeaders:
    - String
  FeaturesAttribute: String
  FeatureTypes:
    - String
  LabelAttribute: String
  LabelHeaders:
    - String
  LabelIndex: Integer
  MaxPayloadInMB: Integer
  MaxRecordCount: Integer
  ProbabilityAttribute: String
  ProbabilityIndex: Integer

```

## Properties

`ContentTemplate`

A template string used to format a JSON record into an acceptable model container
input. For example, a `ContentTemplate` string
`'{"myfeatures":$features}'` will format a list of features
`[1,2,3]` into the record string `'{"myfeatures":[1,2,3]}'`.
Required only when the model container input is in JSON Lines format.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureHeaders`

The names of the features. If provided, these are included in the endpoint response
payload to help readability of the `InvokeEndpoint` output. See the [Response](../../../sagemaker/latest/dg/clarify-online-explainability-invoke-endpoint.md#clarify-online-explainability-response) section under **Invoke the endpoint**
in the Developer Guide for more information.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeaturesAttribute`

Provides the JMESPath expression to extract the features from a model container input
in JSON Lines format. For example, if `FeaturesAttribute` is the JMESPath
expression `'myfeatures'`, it extracts a list of features
`[1,2,3]` from request data `'{"myfeatures":[1,2,3]}'`.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FeatureTypes`

A list of data types of the features (optional). Applicable only to NLP
explainability. If provided, `FeatureTypes` must have at least one
`'text'` string (for example, `['text']`). If
`FeatureTypes` is not provided, the explainer infers the feature types
based on the baseline data. The feature types are included in the endpoint response
payload. For additional information see the [response](../../../sagemaker/latest/dg/clarify-online-explainability-invoke-endpoint.md#clarify-online-explainability-response) section under **Invoke the endpoint**
in the Developer Guide for more information.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `256`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LabelAttribute`

A JMESPath expression used to locate the list of label headers in the model container
output.

**Example**: If the model container output of a batch
request is `'{"labels":["cat","dog","fish"],"probability":[0.6,0.3,0.1]}'`,
then set `LabelAttribute` to `'labels'` to extract the list of
label headers `["cat","dog","fish"]`

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LabelHeaders`

For multiclass classification problems, the label headers are the names of the
classes. Otherwise, the label header is the name of the predicted label. These are used
to help readability for the output of the `InvokeEndpoint` API. See the
[response](../../../sagemaker/latest/dg/clarify-online-explainability-invoke-endpoint.md#clarify-online-explainability-response) section under **Invoke the endpoint**
in the Developer Guide for more information. If there are no label headers in the model
container output, provide them manually using this parameter.

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `16`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`LabelIndex`

A zero-based index used to extract a label header or list of label headers from model
container output in CSV format.

**Example for a multiclass model:** If the model
container output consists of label headers followed by probabilities:
`'"[\'cat\',\'dog\',\'fish\']","[0.1,0.6,0.3]"'`, set
`LabelIndex` to `0` to select the label headers
`['cat','dog','fish']`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxPayloadInMB`

The maximum payload size (MB) allowed of a request from the explainer to the model
container. Defaults to `6` MB.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Maximum_: `25`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`MaxRecordCount`

The maximum number of records in a request that the model container can process when
querying the model container for the predictions of a [synthetic dataset](../../../sagemaker/latest/dg/clarify-online-explainability-create-endpoint.md#clarify-online-explainability-create-endpoint-synthetic). A record is a unit of input data that inference can be
made on, for example, a single line in CSV data. If `MaxRecordCount` is
`1`, the model container expects one record per request. A value of 2 or
greater means that the model expects batch requests, which can reduce overhead and speed
up the inferencing process. If this parameter is not provided, the explainer will tune
the record count per request according to the model container's capacity at
runtime.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProbabilityAttribute`

A JMESPath expression used to extract the probability (or score) from the model
container output if the model container is in JSON Lines format.

**Example**: If the model container output of a single
request is `'{"predicted_label":1,"probability":0.6}'`, then set
`ProbabilityAttribute` to `'probability'`.

_Required_: No

_Type_: String

_Pattern_: `.*`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ProbabilityIndex`

A zero-based index used to extract a probability value (score) or list from model
container output in CSV format. If this value is not provided, the entire model
container output will be treated as a probability value (score) or list.

**Example for a single class model:** If the model
container output consists of a string-formatted prediction label followed by its
probability: `'1,0.6'`, set `ProbabilityIndex` to `1`
to select the probability value `0.6`.

**Example for a multiclass model:** If the model
container output consists of a string-formatted prediction label followed by its
probability: `'"[\'cat\',\'dog\',\'fish\']","[0.1,0.6,0.3]"'`, set
`ProbabilityIndex` to `1` to select the probability values
`[0.1,0.6,0.3]`.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClarifyExplainerConfig

ClarifyShapBaselineConfig

All content copied from https://docs.aws.amazon.com/.
