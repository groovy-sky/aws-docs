# Viewing resource-based policies

You can view resource-based policies attached to your clusters to understand the current access controls in place.

###### To view resource-based policies

1. Sign in to the AWS Management Console and open the Aurora DSQL console at [https://console.aws.amazon.com/dsql/](https://console.aws.amazon.com/dsql).

2. Choose your cluster from the cluster list to open the cluster details page.

3. Choose the **Permissions** tab.

4. View the attached policy in the **Resource-based policy** section.

Use the `get-cluster-policy` command to view a cluster's resource-based policy:

```nohighlight

aws dsql get-cluster-policy --identifier your_cluster_id
```

Python

```python

import boto3
import json

client = boto3.client('dsql')

response = client.get_cluster_policy(
    identifier='your_cluster_id'
)

# Parse and pretty-print the policy
policy = json.loads(response['policy'])
print(json.dumps(policy, indent=2))

```

Java

```java

import software.amazon.awssdk.services.dsql.DsqlClient;
import software.amazon.awssdk.services.dsql.model.GetClusterPolicyRequest;
import software.amazon.awssdk.services.dsql.model.GetClusterPolicyResponse;

DsqlClient client = DsqlClient.create();

GetClusterPolicyRequest request = GetClusterPolicyRequest.builder()
    .identifier("your_cluster_id")
    .build();

GetClusterPolicyResponse response = client.getClusterPolicy(request);
System.out.println("Policy: " + response.policy());

```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Add and edit policies

Remove Policy

All content copied from https://docs.aws.amazon.com/.
