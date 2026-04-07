# Managing an Amazon Aurora DB cluster

This section shows how to manage and maintain your Aurora DB cluster.
Aurora involves clusters of database servers that are connected in a replication topology. Thus,
managing Aurora often involves deploying changes to multiple servers and making sure that all
Aurora Replicas are keeping up with the source server. Because Aurora transparently scales
the underlying storage as your data grows, managing Aurora requires relatively little management
of disk storage. Likewise, because Aurora automatically performs continuous backups, an Aurora
cluster does not require extensive planning or downtime for performing backups.

###### Topics

- [Stopping and starting an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-cluster-stop-start.html)

- [Automatically connecting an EC2 instance and an Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/ec2-rds-connect.html)

- [Automatically connecting a Lambda function and an Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/lambda-rds-connect.html)

- [Modifying an Amazon Aurora DB cluster](aurora-modifying.md)

- [Adding Aurora Replicas to a DB cluster](aurora-replicas-adding.md)

- [Managing performance and scaling for Aurora DB clusters](aurora-managing-performance.md)

- [Cloning a volume for an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Managing.Clone.html)

- [Integrating Aurora with other AWS services](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Integrating.html)

- [Maintaining an Amazon Aurora DB cluster](user-upgradedbinstance-maintenance.md)

- [Rebooting an Amazon Aurora DB cluster or Amazon Aurora DB instance](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_RebootCluster.html)

- [Failing over an Amazon Aurora DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-failover.html)

- [Deleting Aurora DB clusters and DB instances](user-deletecluster.md)

- [Tagging Amazon Aurora andAmazon RDS resources](user-tagging.md)

- [Amazon Resource Names (ARNs) in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Tagging.ARN.html)

- [Amazon Aurora updates](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Aurora.Updates.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tutorial: Creating a MySQL DB cluster with a custom parameter
group

Stopping and starting a cluster
