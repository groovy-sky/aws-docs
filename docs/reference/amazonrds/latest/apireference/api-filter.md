# Filter

A filter name and value pair that is used to return a more specific list of results
from a describe operation. Filters can be used to match a set of resources by specific
criteria, such as IDs. The filters supported by a describe operation are documented
with the describe operation.

###### Note

Currently, wildcards are not supported in filters.

The following actions can be filtered:

- `DescribeDBClusterBacktracks`

- `DescribeDBClusterEndpoints`

- `DescribeDBClusters`

- `DescribeDBInstances`

- `DescribeDBRecommendations`

- `DescribeDBShardGroups`

- `DescribePendingMaintenanceActions`

## Contents

###### Note

In the following list, the required parameters are described first.

**Name**

The name of the filter. Filter names are case-sensitive.

Type: String

Required: Yes

**Values.Value.N**

One or more filter values. Filter values are case-sensitive.

Type: Array of strings

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/filter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/filter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/filter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FailoverState

GlobalCluster
