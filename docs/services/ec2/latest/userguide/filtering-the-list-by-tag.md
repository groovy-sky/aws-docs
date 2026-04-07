# Filter Amazon EC2 resources by tag

After you add tags, you can filter your Amazon EC2 resources based tag keys and tag
values.

Console

###### To filter resources by tag

1. Open the Amazon EC2 console at
    [https://console.aws.amazon.com/ec2/](https://console.aws.amazon.com/ec2).

2. In the navigation pane, select a resource type (for example, **Instances**).

3. Choose the search field.

4. In the list, under **Tags**, choose the tag key.

5. Choose the corresponding tag value from the list.

6. When you are finished, remove the filter.

For more information about using filters in the Amazon EC2 console, see [Find your Amazon EC2 resources](using-filtering.md).

AWS CLI

###### To describe resources of a single type with the specified tag key

Add the following filter to a `describe` command to describe the resources
of that type with a Stack tag, regardless of the value of the tag.

```nohighlight

--filters Name=tag-key,Values=Stack
```

###### To describe resources of a single type with the specified tag

Add the following filter to a `describe` command to describe the resources
of that type with the tag Stack=production.

```nohighlight

--filters Name=tag:Stack,Values=production
```

###### To describe resources of a single type with the specified tag value

Add the following filter to a `describe` command to describe the resources
of that type with a tag with the value production, regardless of the tag key.

```nohighlight

--filters Name=tag-value,Values=production
```

###### To describe all EC2 resources with the specified tag

Add the following filter to the [describe-tags](https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-tags.html)
command to describe all EC2 resources with the tag Stack=test.

```nohighlight

--filters Name=key,Values=Stack Name=value,Values=test
```

PowerShell

###### To filter resources of a single type by tag key

Add the following filter to a `Get` cmdlet to describe the resources
of that type with a Stack tag, regardless of the value of the tag.

```powershell

-Filter @{Name="tag-key"; Values="Stack"}
```

###### To filter resources of a single type by tag

Add the following filter to a `Get` cmdlet to describe the resources
of that type with the tag Stack=production.

```powershell

-Filter @{Name="tag:Stack"; Values="production"}
```

###### To filter resources of a single type by tag value

Add the following filter to a `Get` cmdlet to describe the resources
of that type with a tag with the value production, regardless of the value of
the tag key.

```powershell

-Filter @{Name="tag-value"; Values="production"}
```

###### To filter all EC2 resources by tag

Add the following filter to the [Get-EC2Tag](https://docs.aws.amazon.com/powershell/latest/reference/items/Get-EC2Tag.html)
cmdlet to describe all EC2 resources with the tag Stack=test.

```powershell

-Filter @{Name="tag:Stack"; Values="test"}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Add and remove tags

View tags using instance metadata
