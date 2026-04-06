# SAPODataParallelismConfig

Sets the number of _concurrent processes_ that transfer OData records
from your SAP instance. A concurrent process is query that retrieves a batch of records as
part of a flow run. Amazon AppFlow can run multiple concurrent processes in parallel to
transfer data faster.

## Contents

**maxParallelism**

The maximum number of processes that Amazon AppFlow runs at the same time when it
retrieves your data from your SAP application.

Type: Integer

Valid Range: Minimum value of 1. Maximum value of 10.

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/appflow-2020-08-23/SAPODataParallelismConfig)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/appflow-2020-08-23/SAPODataParallelismConfig)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/appflow-2020-08-23/SAPODataParallelismConfig)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

SAPODataPaginationConfig

SAPODataSourceProperties
