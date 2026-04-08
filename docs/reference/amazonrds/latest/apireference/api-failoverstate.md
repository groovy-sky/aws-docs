# FailoverState

Contains the state of scheduled or in-process operations on a
global cluster (Aurora global database). This data type is empty unless a switchover
or failover operation is scheduled or is in progress on the Aurora global database.

## Contents

###### Note

In the following list, the required parameters are described first.

**FromDbClusterArn**

The Amazon Resource Name (ARN) of the Aurora DB cluster that is currently being demoted, and which is associated with this
state.

Type: String

Required: No

**IsDataLossAllowed**

Indicates whether the operation is a global switchover or a global failover. If data loss is allowed, then the operation is a global failover.
Otherwise, it's a switchover.

Type: Boolean

Required: No

**Status**

The current status of the global cluster. Possible values are as follows:

- pending – The service received a request to switch over or fail over the global cluster. The
global cluster's primary DB cluster and the specified secondary DB cluster are being verified before the operation
starts.

- failing-over – Aurora is promoting the chosen secondary Aurora DB cluster to become the new primary DB cluster to fail over the global cluster.

- cancelling – The request to switch over or fail over the global cluster was cancelled and the primary
Aurora DB cluster and the selected secondary Aurora DB cluster are returning to their previous states.

- switching-over – This status covers the range of Aurora internal operations that take place during the switchover process, such
as demoting the primary Aurora DB cluster, promoting the secondary Aurora DB cluster, and synchronizing replicas.

Type: String

Valid Values: `pending | failing-over | cancelling`

Required: No

**ToDbClusterArn**

The Amazon Resource Name (ARN) of the Aurora DB cluster that is currently being promoted, and which is associated
with this state.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/rds-2014-10-31/failoverstate.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/rds-2014-10-31/failoverstate.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/rds-2014-10-31/failoverstate.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ExportTask

Filter

All content copied from https://docs.aws.amazon.com/.
