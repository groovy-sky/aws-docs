# Rebooting an Amazon Aurora DB cluster or Amazon Aurora DB instance

You might need to reboot your DB cluster or some instances within the cluster, usually for maintenance reasons.
For example, suppose that you modify the parameters within a parameter group or associate a different parameter
group with your cluster. In these cases, you must reboot the cluster for the changes to take effect. Similarly,
you might reboot one or more reader DB instances within the cluster. You can arrange the reboot operations for
individual instances to minimize downtime for the entire cluster.

The time required to reboot each DB instance in your cluster depends on the database activity at the time of
reboot. It also depends on the recovery process of your specific DB engine. If it's practical, reduce
database activity on that particular instance before starting the reboot process. Doing so can reduce the time
needed to restart the database.

You can only reboot each DB instance in your cluster when it's in the available state. A DB instance can be
unavailable for several reasons. These include the cluster being stopped state, a modification being applied to
the instance, and a maintenance-window action such as a version upgrade.

Rebooting a DB instance restarts the database engine process. Rebooting a DB instance results in a momentary
outage, during which the DB instance status is set to _rebooting_.

###### Note

If a DB instance isn't using the latest changes to its associated DB parameter group, the AWS Management Console shows
the DB parameter group with a status of **pending-reboot**. The
**pending-reboot** parameter groups status doesn't result in an automatic reboot during
the next maintenance window. To apply the latest parameter changes to that DB instance, manually reboot the DB
instance. For more information about parameter groups, see
[Parameter groups for Amazon Aurora](user-workingwithparamgroups.md).

###### Topics

- [Rebooting a DB instance within an Aurora cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-reboot-db-instance.html)

- [Rebooting an Aurora cluster with read availability](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-mysql-survivable-replicas.html)

- [Rebooting an Aurora cluster without read availability](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-reboot-cluster.html)

- [Checking uptime for Aurora clusters and instances](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Reboot.Uptime.html)

- [Examples of Aurora reboot operations](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/USER_Reboot.Examples.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS Organizations upgrade rollout

Rebooting a DB instance within an Aurora cluster
