# Configuring private image replication in Amazon ECR

Configure replication per Region for your private registry. You can configure
cross-Region replication or cross-account replication.

For examples of how replication is commonly used, see [Private image replication examples for Amazon ECR](registry-settings-examples.md).

01. Open the Amazon ECR console at
     [https://console.aws.amazon.com/ecr/repositories](https://console.aws.amazon.com/ecr/repositories).

02. From the navigation bar, choose the Region to configure your registry
     replication settings for.

03. In the navigation pane, choose **Private**
    **registry**.

04. On the **Private registry** page, choose **Settings** and then choose **Edit** under **Replication**
    **configuration**.

05. On the **Replication** page, choose **Add**
    **replication rule**.

06. On the **Destination types** page, choose whether to
     enable cross-Region replication, cross-account replication, or both and
     then choose **Next**.

07. If cross-Region replication is enabled, then for **Configure**
    **destination regions**, choose one or more **Destination**
    **regions** and then choose **Next**.

08. If cross-account replication is enabled, then for **Cross-account**
    **replication**, choose the cross-account replication setting for
     the registry. For **Destination account**, enter the
     account ID for the destination account and one or more **Destination**
    **regions** to replicate to. Choose **Destination**
    **account +** to configure additional accounts as replication
     destinations.

    ###### Important

    For cross-account replication to occur, the destination account
    must configure a registry permissions policy to allow replication to
    occur. For more information, see [Private registry permissions in Amazon ECR](registry-permissions.md).

09. (Optional) On the **Add filters** page, specify one
     or more filters for the replication rule and then choose **Add**.
     Repeat this step for each filter you want to associate with the
     replication action. A filter must be specified as a repository name
     prefix. If no filters are added, the contents of all repositories are
     replicated. Choose **Next** once all filters have been
     added.

10. On the **Review and submit** page, review the
     replication rule configuration and then choose **Submit**
    **rule**.

1. Create a JSON file containing the replication rules to define for your
    registry. A replication configuration may contain up to 10 rules, with
    up to 25 unique destinations across all rules and 100 filters per each
    rule. To configure cross-Region replication within your own account, you
    specify your own account ID. For more examples, see [Private image replication examples for Amazon ECR](registry-settings-examples.md).

```nohighlight

{
   	"rules": [{
   		"destinations": [{
   			"region": "destination_region",
   			"registryId": "destination_accountId"
   		}],
   		"repositoryFilters": [{
   			"filter": "repository_prefix_name",
   			"filterType": "PREFIX_MATCH"
   		}]
   	}]
}
```

2. Create a replication configuration for your registry.

```nohighlight

aws ecr put-replication-configuration \
        --replication-configuration file://replication-settings.json \
        --region us-west-2
```

3. Confirm your registry settings.

```nohighlight

aws ecr describe-registry \
        --region us-west-2
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Replication examples

Removing replication

All content copied from https://docs.aws.amazon.com/.
