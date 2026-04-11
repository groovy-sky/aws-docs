# Deleting an RDS Proxy

You can delete a proxy when you no longer need it. Or, you might delete a proxy if you take the DB instance or cluster
associated with it out of service.

###### To delete a proxy

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Proxies**.

3. Choose the proxy to delete from the list.

4. Choose **Delete Proxy**.

To delete a DB proxy, use the AWS CLI command
[delete-db-proxy](../../../cli/latest/reference/rds/delete-db-proxy.md). To remove related
associations, also use the
[deregister-db-proxy-targets](../../../cli/latest/reference/rds/deregister-db-proxy-targets.md)
command.

```nohighlight

aws rds delete-db-proxy --name proxy_name
```

```nohighlight

aws rds deregister-db-proxy-targets
    --db-proxy-name proxy_name
    [--target-group-name target_group_name]
    [--target-ids comma_separated_list]       # or
    [--db-instance-identifiers instance_id]       # or
    [--db-cluster-identifiers cluster_id]
```

To delete a DB proxy, call the Amazon RDS API function
[DeleteDBProxy](../../../../reference/amazonrds/latest/apireference/api-deletedbproxy.md). To delete related items and
associations, you also call the functions
[DeleteDBProxyTargetGroup](../../../../reference/amazonrds/latest/apireference/api-deletedbproxytargetgroup.md) and
[DeregisterDBProxyTargets](../../../../reference/amazonrds/latest/apireference/api-deregisterdbproxytargets.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Avoid pinning RDS Proxy

Working with RDS Proxy endpoints

All content copied from https://docs.aws.amazon.com/.
