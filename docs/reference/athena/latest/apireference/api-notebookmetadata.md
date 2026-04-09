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

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/notebookmetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/notebookmetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/notebookmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NamedQuery

NotebookSessionSummary

All content copied from https://docs.aws.amazon.com/.
