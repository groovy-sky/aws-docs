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

- [AWS SDK for C++](../../../goto/sdkforcpp/appflow-2020-08-23/sapodataparallelismconfig.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appflow-2020-08-23/sapodataparallelismconfig.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appflow-2020-08-23/sapodataparallelismconfig.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SAPODataPaginationConfig

SAPODataSourceProperties
