# Configuring engine parameters using ElastiCache parameter groups

Amazon ElastiCache uses parameters to control the runtime properties of your
nodes and clusters. Generally, newer engine versions include additional parameters to
support the newer functionality. For tables of Memcached parameters, see [Memcached specific parameters](parametergroups-engine.md#ParameterGroups.Memcached).
For tables of Valkey and Redis OSS parameters, see [Valkey and Redis OSS parameters](parametergroups-engine.md#ParameterGroups.Redis).

As you would expect, some parameter values, such as
`maxmemory`, are determined by the engine and node type. For a table of these Memcached
parameter values by node type, see [Memcached node-type specific parameters](parametergroups-engine.md#ParameterGroups.Memcached.NodeSpecific). For a table of these Valkey and Redis OSS parameter values by node type, see [Redis OSS node-type specific parameters](parametergroups-engine.md#ParameterGroups.Redis.NodeSpecific).

###### Note

For a list of Memcached-specific parameters, see [Memcached\
Specific Parameters](parametergroups-engine.md#ParameterGroups.Memcached).

###### Topics

- [Parameter management in ElastiCache](parametergroups-management.md)

- [Cache parameter group tiers in ElastiCache](parametergroups-tiers.md)

- [Creating an ElastiCache parameter group](parametergroups-creating.md)

- [Listing ElastiCache parameter groups by name](parametergroups-listinggroups.md)

- [Listing an ElastiCache parameter group's values](parametergroups-listingvalues.md)

- [Modifying an ElastiCache parameter group](parametergroups-modifying.md)

- [Deleting an ElastiCache parameter group](parametergroups-deleting.md)

- [Engine specific parameters](parametergroups-engine.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Managing ElastiCache cluster maintenance

Parameter management in ElastiCache

All content copied from https://docs.aws.amazon.com/.
