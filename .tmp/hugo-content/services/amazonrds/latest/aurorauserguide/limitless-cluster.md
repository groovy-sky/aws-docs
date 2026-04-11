# Creating a DB cluster that uses Aurora PostgreSQL Limitless Database

You create a new Aurora DB cluster using the Limitless Database version of Aurora PostgreSQL, and add a DB shard group to the cluster. When adding a DB shard
group, you specify the maximum compute capacity for the entire DB shard group (sum of the capacities for all routers and shards) in Aurora capacity units
(ACUs). Each ACU is a combination of approximately 2 gibibytes (GiB) of memory, corresponding CPU, and networking. Scaling increases or decreases
capacity for your DB shard group, depending on your application workload, similar to how [Aurora Serverless v2](aurora-serverless-v2-how-it-works.md) works.

###### Topics

- [Correlating DB shard group maximum capacity with the number of routers and shards created](#limitless-capacity-mapping)

- [Creating your DB cluster](limitless-create-cluster.md)

## Correlating DB shard group maximum capacity with the number of routers and shards created

The initial number of routers and shards in a DB shard group is determined by the maximum capacity that you set when you create the DB shard
group. The higher the maximum capacity, the greater the number of routers and shards that are created in the DB shard group.

Each node (shard or router) has its own current capacity value, also measured in ACUs.

- Limitless Database scales a node up to a higher capacity when its current capacity is too low to handle the load. However, nodes won't scale up
when the total capacity is at the maximum.

- Limitless Database scales the node down to a lower capacity when its current capacity is higher than needed. However, nodes won't scale down
when the total capacity is at the minimum.

The following table shows the correlation between the DB shard group maximum capacity in Aurora capacity units (ACUs) and the number of nodes
(routers and shards) created.

###### Note

These values are subject to change.

If you set the compute redundancy to a nonzero value, the total number of shards will be doubled or tripled. This will incur extra costs.

The nodes in compute standbys are scaled up and down to the same capacity as the writer. You don't set the capacity range separately for the standbys.

Total nodesRoutersShardsDefault minimum capacity (ACUs)Maximum capacity range (ACUs)4221616–40052320401–50062424501–60073428601–70083532701–80093636801–900104640901–10001147441001–11001248481101–12001358521201–13001459561301–140015510601401–150016610641501–160017611681601–170018612721701–180019712761801–190020713801901–200021714842001–210022814882101–220023815922201–230024816962301–6144

Maximum capacity–based dynamic configuration for the DB shard group is available only during creation. The number of routers and shards
remains the same when the maximum capacity is modified. For more information, see [Changing the capacity of a DB shard group](limitless-capacity.md).

You can use SQL commands to add shards and routers to a DB shard group. For more information, see the following:

- [Splitting a shard in a DB shard group](limitless-shard-split.md)

- [Adding a router to a DB shard group](limitless-add-router.md)

###### Note

You can't delete shards or routers.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Prerequisites for using Limitless Database

Creating your DB cluster

All content copied from https://docs.aws.amazon.com/.
