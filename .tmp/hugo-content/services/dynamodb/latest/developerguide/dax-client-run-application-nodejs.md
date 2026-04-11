# Node.js and DAX

Follow these steps to run the Node.js sample application on your Amazon EC2
instance.

###### To run the Node.js sample for DAX

1. Set up Node.js on your Amazon EC2 instance, as follows:
1. Install node version manager ( `nvm`).

      ```nohighlight

      curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.3/install.sh | bash
      ```

2. Use nvm to install Node.js.

      ```nohighlight

      nvm install 12.16.3
      ```

3. Test that Node.js is installed and running correctly.

      ```nohighlight

      node -e "console.log('Running Node.js ' + process.version)"
      ```

      This should display the following message.

      `Running Node.js v12.16.3`
2. Install the DAX Node.js client using the node package manager
    ( `npm`).

```nohighlight

npm install amazon-dax-client
```

3. Download the sample program source code ( `.zip`
    file).

```nohighlight

wget http://docs.aws.amazon.com/amazondynamodb/latest/developerguide/samples/TryDax.zip
```

When the download is complete, extract the source files.

```nohighlight

unzip TryDax.zip
```

4. Run the following Node.js programs. The first program creates an Amazon DynamoDB
    table named `TryDaxTable`. The second program writes data to the
    table.

```nohighlight

node 01-create-table.js
node 02-write-data.js
```

5. Run the following Node.js programs.

```nohighlight

node 03-getitem-test.js
node 04-query-test.js
node 05-scan-test.js
```

    Take note of the timing information—the number of milliseconds required
    for the `GetItem`, `Query`, and `Scan`
    tests.

6. In the previous step, you ran the programs against the DynamoDB endpoint. Run
    the programs again, but this time, the `GetItem`,
    `Query` and `Scan` operations are processed by
    your DAX cluster.

To determine the endpoint for your DAX cluster, choose one of the
    following.

- **Using the DynamoDB console**—Choose your
DAX cluster. The cluster endpoint is shown on the console, as in
the following example.

```nohighlight

dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

- **Using the AWS CLI**—Enter the following
command.

```nohighlight

aws dax describe-clusters --query "Clusters[*].ClusterDiscoveryEndpoint"
```

The cluster endpoint is shown in the output, as in the following
example.

```nohighlight

{
      "Address": "my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com",
      "Port": 8111,
      "URL": "dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com"
}
```

Now run the programs again, but this time, specify the cluster endpoint as
a command line parameter.

```nohighlight

node 03-getitem-test.js dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
node 04-query-test.js dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
node 05-scan-test.js dax://my-cluster.l6fzcv.dax-clusters.us-east-1.amazonaws.com
```

Look at the rest of the output, and take note of the timing information.
The elapsed times for `GetItem`, `Query`, and
`Scan` should be significantly lower with DAX than with
DynamoDB.

7. Run the following Node.js program to delete
    `TryDaxTable`.

```nohighlight

node 06-delete-table
```

For more information about these programs, see the following sections:

- [01-create-table.js](dax-client-run-application-nodejs-01-create-table.md)

- [02-write-data.js](dax-client-run-application-nodejs-02-write-data.md)

- [03-getitem-test.js](dax-client-run-application-nodejs-03-getitem-test.md)

- [04-query-test.js](dax-client-run-application-nodejs-04-query-test.md)

- [05-scan-test.js](dax-client-run-application-nodejs-05-scan-test.md)

- [06-delete-table.js](dax-client-run-application-nodejs-06-delete-table.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS SDK for Node.js 2.x examples

01-create-table.js

All content copied from https://docs.aws.amazon.com/.
