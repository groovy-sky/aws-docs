# Delete a Capacity Reservation group

You can use the following examples to delete a Capacity Reservation group.

AWS CLI

###### To delete a group

Use the [delete-group](../../../cli/latest/reference/resource-groups/delete-group.md) command.

```nohighlight

aws resource-groups delete-group --group MyCRGroup
```

PowerShell

###### To delete a group

Use the [Remove-RGGroup](../../../powershell/latest/reference/items/remove-rggroup.md)
cmdlet.

```powershell

Remove-RGGroup -GroupName MyCRGroup
```

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

Remove Capacity Reservation from group

Using Capacity Reservation in cluster placement groups
with a Capacity Reservation group
