# Getting an existing ARN for Amazon RDS

You can get the ARN of an RDS resource by using the AWS Management Console, AWS Command Line Interface (AWS CLI), or RDS API.

To get an ARN from the AWS Management Console, navigate to the resource you want an ARN for, and view the details for that resource.

For example, you can get the ARN for a DB cluster from the **Configuration** tab
of the DB cluster details.

![DB cluster ARN.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/DB-cluster-arn.png)

To get an ARN from the AWS CLI for a particular RDS resource, you use the `describe` command for that resource.
The following table shows each AWS CLI command, and the ARN property used with the command to get an ARN.

AWS CLI commandARN property[describe-event-subscriptions](../../../cli/latest/reference/rds/describe-event-subscriptions.md)EventSubscriptionArn[describe-certificates](../../../cli/latest/reference/rds/describe-certificates.md)CertificateArn[describe-db-parameter-groups](../../../cli/latest/reference/rds/describe-db-parameter-groups.md)DBParameterGroupArn[describe-db-cluster-parameter-groups](../../../cli/latest/reference/rds/describe-db-cluster-parameter-groups.md)DBClusterParameterGroupArn[describe-db-instances](../../../cli/latest/reference/rds/describe-db-instances.md)DBInstanceArn[describe-db-security-groups](../../../cli/latest/reference/rds/describe-db-security-groups.md)DBSecurityGroupArn[describe-db-snapshots](../../../cli/latest/reference/rds/describe-db-snapshots.md)DBSnapshotArn[describe-events](../../../cli/latest/reference/rds/describe-events.md)SourceArn[describe-reserved-db-instances](../../../cli/latest/reference/rds/describe-reserved-db-instances.md)ReservedDBInstanceArn[describe-db-subnet-groups](../../../cli/latest/reference/rds/describe-db-subnet-groups.md)DBSubnetGroupArn[describe-db-clusters](../../../cli/latest/reference/rds/describe-db-clusters.md)DBClusterArn[describe-db-cluster-snapshots](../../../cli/latest/reference/rds/describe-db-cluster-snapshots.md)DBClusterSnapshotArn

For example, the following AWS CLI command gets the ARN for a DB instance.

###### Example

For Linux, macOS, or Unix:

```nohighlight

aws rds describe-db-instances \
--db-instance-identifier DBInstanceIdentifier \
--region us-west-2 \
--query "*[].{DBInstanceIdentifier:DBInstanceIdentifier,DBInstanceArn:DBInstanceArn}"

```

For Windows:

```nohighlight

aws rds describe-db-instances ^
--db-instance-identifier DBInstanceIdentifier ^
--region us-west-2 ^
--query "*[].{DBInstanceIdentifier:DBInstanceIdentifier,DBInstanceArn:DBInstanceArn}"

```

The output of that command is like the following:

```json

[
    {
        "DBInstanceArn": "arn:aws:rds:us-west-2:account_id:db:instance_id",
        "DBInstanceIdentifier": "instance_id"
    }
]

```

To get an ARN for a particular RDS resource, you can call the following RDS API operations and use the ARN properties shown following.

RDS API operationARN property[DescribeEventSubscriptions](../../../../reference/amazonrds/latest/apireference/api-describeeventsubscriptions.md)EventSubscriptionArn[DescribeCertificates](../../../../reference/amazonrds/latest/apireference/api-describecertificates.md)CertificateArn[DescribeDBParameterGroups](../../../../reference/amazonrds/latest/apireference/api-describedbparametergroups.md)DBParameterGroupArn[DescribeDBClusterParameterGroups](../../../../reference/amazonrds/latest/apireference/api-describedbclusterparametergroups.md)DBClusterParameterGroupArn[DescribeDBInstances](../../../../reference/amazonrds/latest/apireference/api-describedbinstances.md)DBInstanceArn[DescribeDBSecurityGroups](../../../../reference/amazonrds/latest/apireference/api-describedbsecuritygroups.md)DBSecurityGroupArn[DescribeDBSnapshots](../../../../reference/amazonrds/latest/apireference/api-describedbsnapshots.md)DBSnapshotArn[DescribeEvents](../../../../reference/amazonrds/latest/apireference/api-describeevents.md)SourceArn[DescribeReservedDBInstances](../../../../reference/amazonrds/latest/apireference/api-describereserveddbinstances.md)ReservedDBInstanceArn[DescribeDBSubnetGroups](../../../../reference/amazonrds/latest/apireference/api-describedbsubnetgroups.md)DBSubnetGroupArn[DescribeDBClusters](../../../../reference/amazonrds/latest/apireference/api-describedbclusters.md)DBClusterArn[DescribeDBClusterSnapshots](../../../../reference/amazonrds/latest/apireference/api-describedbclustersnapshots.md)DBClusterSnapshotArn

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Constructing an ARN

Aurora updates

All content copied from https://docs.aws.amazon.com/.
