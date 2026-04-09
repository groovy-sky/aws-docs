# Python and DAX

Follow this procedure to run the Python sample application on your Amazon EC2
instance.

###### To run the Python sample for DAX

1. Install the DAX Python client using the `pip` utility.

```nohighlight

pip install amazon-dax-client
```

2. Download the sample program source code ( `.zip`
    file).

```nohighlight

wget http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/samples/TryDax.zip
```

When the download is complete, extract the source files.

```nohighlight

unzip TryDax.zip
```

3. Run the following Python programs. The first program creates an Amazon DynamoDB
    table named `TryDaxTable`. The second program writes data
    to the table.

```nohighlight

python 01-create-table.py
python 02-write-data.py
```

4. Run the following Python programs.

```nohighlight

python 03-getitem-test.py
python 04-query-test.py
python 05-scan-test.py
```

    Take note of the timing information—the number of milliseconds required
    for the `GetItem`, `Query`, and `Scan`
    tests.

5. In the previous step, you ran the programs against the DynamoDB endpoint. Now
    run the programs again, but this time, the `GetItem`,
    `Query`, and `Scan` operations are processed by
    your DAX cluster.

To determine the endpoint for your DAX cluster, choose one of the
    following:

- **Using the DynamoDB console** — Choose your
DAX cluster. The cluster endpoint is shown on the console, as in
the following example.

```nohighlight

dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

- **Using the AWS CLI** — Enter the following
command.

```nohighlight

aws dax describe-clusters --query "Clusters[*].ClusterDiscoveryEndpoint"
```

The cluster endpoint is shown in the output, as in this
example.

```nohighlight

{
      "Address": "my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com",
      "Port": 8111,
      "URL": "dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com"
}
```

Run the programs again, but this time, specify the cluster endpoint as a
command line parameter.

```nohighlight

python 03-getitem-test.py dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
python 04-query-test.py dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
python 05-scan-test.py dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

Look at the rest of the output, and take note of the timing information.
The elapsed times for `GetItem`, `Query`, and
`Scan` should be significantly lower with DAX than with
DynamoDB.

6. Run the following Python program to delete
    `TryDaxTable`.

```nohighlight

python 06-delete-table.py
```

For more information about these programs, see the following sections:

- [01-create-table.py](dax-client-run-application-python-01-create-table.md)

- [02-write-data.py](dax-client-run-application-python-02-write-data.md)

- [03-getitem-test.py](dax-client-run-application-python-03-getitem-test.md)

- [04-query-test.py](dax-client-run-application-python-04-query-test.md)

- [05-scan-test.py](dax-client-run-application-python-05-scan-test.md)

- [06-delete-table.py](dax-client-run-application-python-06-delete-table.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

06-DeleteTable.cs

01-create-table.py

All content copied from https://docs.aws.amazon.com/.
