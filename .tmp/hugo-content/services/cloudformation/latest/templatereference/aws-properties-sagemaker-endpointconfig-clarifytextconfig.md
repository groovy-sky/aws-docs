This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::EndpointConfig ClarifyTextConfig

A parameter used to configure the SageMaker Clarify explainer to treat text features as text so
that explanations are provided for individual units of text. Required only for natural
language processing (NLP) explainability.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Granularity" : String,
  "Language" : String
}

```

### YAML

```yaml

  Granularity: String
  Language: String

```

## Properties

`Granularity`

The unit of granularity for the analysis of text features. For example, if the unit is
`'token'`, then each token (like a word in English) of the text is
treated as a feature. SHAP values are computed for each unit/feature.

_Required_: Yes

_Type_: String

_Allowed values_: `token | sentence | paragraph`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Language`

Specifies the language of the text features in [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) or
[ISO 639-3](https://en.wikipedia.org/wiki/ISO_639-3) code of a
supported language.

###### Note

For a mix of multiple languages, use code `'xx'`.

_Required_: Yes

_Type_: String

_Allowed values_: `af | sq | ar | hy | eu | bn | bg | ca | zh | hr | cs | da | nl | en | et | fi | fr | de | el | gu | he | hi | hu | is | id | ga | it | kn | ky | lv | lt | lb | mk | ml | mr | ne | nb | fa | pl | pt | ro | ru | sa | sr | tn | si | sk | sl | es | sv | tl | ta | tt | te | tr | uk | ur | yo | lij | xx`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ClarifyShapConfig

DataCaptureConfig

All content copied from https://docs.aws.amazon.com/.
