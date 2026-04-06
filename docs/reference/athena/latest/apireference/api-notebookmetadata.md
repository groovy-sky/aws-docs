# NotebookMetadata

Contains metadata for notebook, including the notebook name, ID, workgroup, and time
created.

## Contents

**CreationTime**

The time when the notebook was created.

Type: Timestamp

Required: No

**LastModifiedTime**

The time when the notebook was last modified.

Type: Timestamp

Required: No

**Name**

The name of the notebook.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Pattern: `(?!.*[/:\\])[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]+`

Required: No

**NotebookId**

The notebook ID.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 36.

Pattern: `[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}`

Required: No

**Type**

The type of notebook. Currently, the only valid type is `IPYNB`.

Type: String

Valid Values: `IPYNB`

Required: No

**WorkGroup**

The name of the Spark enabled workgroup to which the notebook belongs.

Type: String

Pattern: `[a-zA-Z0-9._-]{1,128}`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/athena-2017-05-18/NotebookMetadata)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/athena-2017-05-18/NotebookMetadata)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/athena-2017-05-18/NotebookMetadata)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NamedQuery

NotebookSessionSummary
