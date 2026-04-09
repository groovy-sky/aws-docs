# QueryExecutionContext

The database and data catalog context in which the query execution occurs.

## Contents

**Catalog**

The name of the data catalog used in the query execution.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDC00-\uDBFF\uDFFF\t]*`

Required: No

**Database**

The name of the database used in the query execution. The database must exist in the
catalog.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 255.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/athena-2017-05-18/queryexecutioncontext.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/athena-2017-05-18/queryexecutioncontext.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/athena-2017-05-18/queryexecutioncontext.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

QueryExecution

QueryExecutionStatistics

All content copied from https://docs.aws.amazon.com/.
