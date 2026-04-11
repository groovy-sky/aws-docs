# Automatically identify nodes in your cluster (Memcached)

For clusters running the Memcached engine,
ElastiCache supports _Auto Discovery_—the ability for client
programs to automatically identify all of the nodes in a cluster, and to initiate and
maintain connections to all of these nodes.

###### Note

Auto Discovery is added for clusters running on Amazon ElastiCache Memcached. Auto Discovery is not available for Valkey or Redis OSS engines.

With Auto Discovery, your application does not need to manually connect to individual cache nodes;
instead, your application connects to one Memcached node and retrieves the list of nodes.
From that list your application is aware of the rest of the nodes in the cluster
and can connect to any of them.
You do not need to hard code the individual cache node endpoints in your application.

If you are using dual stack network type on your cluster, Auto Discovery will return only IPv4 or IPv6 addresses, depending on which one you select.
For more information, see [Choosing a network type in ElastiCache](network-type.md)
.

All of the cache nodes in the cluster maintain a list of metadata about all of the
other nodes. This metadata is updated whenever nodes are added or removed from the
cluster.

###### Topics

- [Benefits of Auto Discovery with Memcached](autodiscovery-benefits.md)

- [How Auto Discovery Works](autodiscovery-howautodiscoveryworks.md)

- [Using Auto Discovery](autodiscovery-using.md)

- [Connecting to Memcached Cache Nodes Manually](autodiscovery-manual.md)

- [Adding Auto Discovery to your Memcached client library](autodiscovery-addingtoyourclientlibrary.md)

- [ElastiCache clients with auto discovery](clients.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Choosing a network type in ElastiCache

Benefits of Auto Discovery with Memcached

All content copied from https://docs.aws.amazon.com/.
