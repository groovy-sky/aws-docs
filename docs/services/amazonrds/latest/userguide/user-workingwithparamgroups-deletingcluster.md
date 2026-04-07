# Deleting a DB cluster parameter group

You can delete a DB cluster parameter group using the AWS Management Console, AWS CLI, or RDS API. A DB cluster parameter group parameter group
is eligible for deletion only if it isn't associated with a DB cluster.

###### To delete parameter groups

1. Sign in to the AWS Management Console and open the Amazon RDS console at
    [https://console.aws.amazon.com/rds/](https://console.aws.amazon.com/rds).

2. In the navigation pane, choose **Parameter**
**groups**.

The parameter groups appear in a list.

3. Choose the name of the DB cluster parameter groups to be deleted.

4. Choose **Actions** and then
    **Delete**.

5. Review the parameter group names and then choose
    **Delete**.

To delete a DB cluster parameter group, use the AWS CLI [`delete-db-cluster-parameter-group`](https://docs.aws.amazon.com/cli/latest/reference/rds/delete-db-cluster-parameter-group.html) command with the
following required parameter.

- `--db-parameter-group-name`

###### Example

The following example deletes a DB cluster parameter group named
_mydbparametergroup._

```nohighlight

aws rds delete-db-cluster-parameter-group --db-parameter-group-name mydbparametergroup
```

To delete a DB cluster parameter group, use the RDS API [`DeleteDBClusterParameterGroup`](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DeleteDBClusterParameterGroup.html) command with the
following required parameter.

- `DBParameterGroupName`

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Viewing parameter values for a DB cluster parameter group

Comparing DB parameter groups
