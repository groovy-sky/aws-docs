# Configuring and managing a Multi-AZ deployment for Amazon RDS

Multi-AZ deployments can have one standby or two standby DB instances. When the deployment
has one standby DB instance, it's called a _Multi-AZ DB instance_
_deployment_. A Multi-AZ DB instance deployment has one standby DB instance
that provides failover support, but doesn't serve read traffic. When the deployment has
two standby DB instances, it's called a _Multi-AZ DB cluster deployment_.
A Multi-AZ DB cluster deployment has standby DB instances that provide failover support and
can also serve read traffic.

You can use the AWS Management Console to determine whether a Multi-AZ deployment is a Multi-AZ DB instance deployment or a Multi-AZ DB cluster deployment.
In the navigation pane, choose **Databases**, and then choose a **DB identifier**.

- A Multi-AZ DB instance deployment has the following characteristics:

- There is only one row for the DB instance.

- The value of **Role** is **Instance** or **Primary**.

- The value of **Multi-AZ** is **Yes**.

- A Multi-AZ DB cluster deployment has the following characteristics:

- There is a cluster-level row with three DB instance rows under it.

- For the cluster-level row, the value of **Role** is **Multi-AZ DB cluster**.

- For each instance-level row, the value of **Role** is **Writer instance**
or **Reader instance**.

- For each instance-level row, the value of **Multi-AZ** is **3 Zones**.

###### Topics

- [Multi-AZ DB instance deployments for Amazon RDS](concepts-multiazsinglestandby.md)

- [Multi-AZ DB cluster deployments for Amazon RDS](multi-az-db-clusters-concepts.md)

In addition, the following topics apply to both DB instances and Multi-AZ DB clusters.

- [Tagging Amazon RDS resources](user-tagging.md)

- [Amazon Resource Names (ARNs) in Amazon RDS](user-tagging-arn.md)

- [Working with storage for Amazon RDS DB instances](user-piops-storagetypes.md)

- [Maintaining a DB instance](user-upgradedbinstance-maintenance.md)

- [Upgrading a DB instance engine version](user-upgradedbinstance-upgrading.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tutorial: Managing a MySQL DB instance

Multi-AZ DB instance deployments

All content copied from https://docs.aws.amazon.com/.
