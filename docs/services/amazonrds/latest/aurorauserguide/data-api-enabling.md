# Enabling the Amazon RDS Data API

To use the Amazon RDS Data API (Data API), enable it for your Aurora DB cluster. You can
enable Data API when you create or modify the DB cluster.

###### Note

Whether Data API is available for your cluster depends on your Aurora version, database engine, and AWS Region.
For older Aurora versions, Data API only works with Aurora Serverless v1 clusters.
For newer Aurora versions, Data API works with clusters that use both provisioned and Aurora Serverless v2 instances.
To check whether your cluster can use Data API, see
[Supported Regions and Aurora DB engines for RDS Data API](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/Concepts.Aurora_Fea_Regions_DB-eng.Feature.Data_API.html).

###### Topics

- [Enabling RDS Data API when you create a database](#data-api.enabling.creating)

- [Enabling RDS Data API on an existing database](#data-api.enabling.modifying)

## Enabling RDS Data API when you create a database

While you are creating a database that supports RDS Data API (Data API), you
can enable this feature. The following procedures describe how to do so when you use
the AWS Management Console, the AWS CLI, or the RDS API.

To enable Data API when you create a DB cluster, select the
**Enable the RDS Data API** checkbox in the
**Connectivity** section of the **Create**
**database** page, as in the following screenshot.

![The Connectivity section on the Create database page, with the Enable the RDS Data API checkbox selected.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/data-api-enable-on-create.png)

For instructions on how to create an Aurora DB cluster that can use the RDS Data API, see the
following:

- For Aurora Serverless v2 and provisioned clusters – [Creating an Amazon Aurora DB cluster](aurora-createinstance.md)

- For Aurora Serverless v1 – [Creating an Aurora Serverless v1 DB cluster](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-serverless.create.html)

To enable Data API while you're creating an Aurora DB cluster, run the
[create-db-cluster](https://docs.aws.amazon.com/cli/latest/reference/rds/create-db-cluster.html) AWS CLI command with the
`--enable-http-endpoint` option.

The following example creates an Aurora PostgreSQL DB cluster with Data
API enabled.

For Linux, macOS, or Unix:

```nohighlight

aws rds create-db-cluster \
    --db-cluster-identifier my_pg_cluster \
    --engine aurora-postgresql \
    --enable-http-endpoint
```

For Windows:

```nohighlight

aws rds create-db-cluster ^
    --db-cluster-identifier my_pg_cluster ^
    --engine aurora-postgresql ^
    --enable-http-endpoint
```

To enable Data API while you're creating an Aurora DB cluster, use the
[CreateDBCluster](../../../../reference/amazonrds/latest/apireference/api-createdbcluster.md) operation with the
value of the `EnableHttpEndpoint` parameter set to `true`.

## Enabling RDS Data API on an existing database

You can modify a DB cluster that supports RDS Data API (Data API) to enable or
disable this feature.

###### Topics

- [Enabling or disabling Data API (Aurora Serverless v2 and provisioned)](#data-api.enabling.modifying.all)

- [Enabling or disabling Data API (Aurora Serverless v1 only)](#data-api.enabling.modifying.sv1)

### Enabling or disabling Data API (Aurora Serverless v2 and provisioned)

Use the following procedures to enable or disable Data API on Aurora Serverless v2 and provisioned databases.
To enable or disable Data API on Aurora Serverless v1 databases, use the procedures in [Enabling or disabling Data API (Aurora Serverless v1 only)](#data-api.enabling.modifying.sv1).

You can enable or disable Data API by using the RDS console for a
DB cluster that supports this feature. To do so, open the cluster
details page of the database on which you want to enable or disable Data
API, and on the **Connectivity & security** tab, go
to the **RDS Data API** section. This section displays
the status of Data API, and allows you to enable or disable it.

The following screenshot shows that the **RDS Data**
**API** isn't enabled.

![The RDS Data API section on the Connectivity and security tab of the details page for a DB cluster. The status of Data API displays as disabled, and the Enable the RDS Data API button is present.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/data-api-enable-from-details.png)

To enable or disable Data API on an existing database, run the
[enable-http-endpoint](https://docs.aws.amazon.com/cli/latest/reference/rds/enable-http-endpoint.html) or
[disable-http-endpoint](https://docs.aws.amazon.com/cli/latest/reference/rds/disable-http-endpoint.html) AWS CLI command,
and specify the ARN of your DB cluster.

The following example enables Data API.

For Linux, macOS, or Unix:

```nohighlight

aws rds enable-http-endpoint \
    --resource-arn cluster_arn
```

For Windows:

```nohighlight

aws rds enable-http-endpoint ^
    --resource-arn cluster_arn
```

To enable or disable Data API on an existing database, use the
[EnableHttpEndpoint](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_EnableHttpEndpoint.html) and
[DisableHttpEndpoint](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DisableHttpEndpoint.html) operations.

### Enabling or disabling Data API (Aurora Serverless v1 only)

Use the following procedures to enable or disable Data API on existing
Aurora Serverless v1 databases. To enable or disable Data API on Aurora Serverless v2 and provisioned databases,
use the procedures in [Enabling or disabling Data API (Aurora Serverless v2 and provisioned)](#data-api.enabling.modifying.all).

When you modify an Aurora Serverless v1 DB cluster, you enable Data
API in the RDS console's **Connectivity**
section.

The following screenshot shows the enabled **Data**
**API** when modifying an Aurora DB cluster.

![The Connectivity section on the Modify DB Cluster page, the Data API checkbox is selected.](https://docs.aws.amazon.com/images/AmazonRDS/latest/AuroraUserGuide/images/data-api-modify-serverlessv1.png)

For instructions on how to modify an Aurora Serverless
v1 DB cluster, see [Modifying an Aurora Serverless v1 DB cluster](aurora-serverless-modifying.md).

To enable or disable Data API, run the
[modify-db-cluster](https://docs.aws.amazon.com/cli/latest/reference/rds/modify-db-cluster.html) AWS CLI command, with the
`--enable-http-endpoint` or
`--no-enable-http-endpoint`, as applicable.

The following example enables Data API on
`sample-cluster`.

For Linux, macOS, or Unix:

```nohighlight

aws rds modify-db-cluster \
    --db-cluster-identifier sample-cluster \
    --enable-http-endpoint
```

For Windows:

```nohighlight

aws rds modify-db-cluster ^
    --db-cluster-identifier sample-cluster ^
    --enable-http-endpoint
```

To enable Data API, use the [ModifyDBCluster](../../../../reference/amazonrds/latest/apireference/api-modifydbcluster.md)
operation, and set the value of `EnableHttpEndpoint` to
`true` or `false`, as applicable.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Authentication and authorization

Creating a Amazon VPC endpoint (PrivateLink)
