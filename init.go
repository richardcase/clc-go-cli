package cli

import (
	"fmt"

	"github.com/centurylinkcloud/clc-go-cli/base"
	"github.com/centurylinkcloud/clc-go-cli/commands"
	"github.com/centurylinkcloud/clc-go-cli/help"
	"github.com/centurylinkcloud/clc-go-cli/models"
	"github.com/centurylinkcloud/clc-go-cli/models/affinity"
	"github.com/centurylinkcloud/clc-go-cli/models/alert"
	"github.com/centurylinkcloud/clc-go-cli/models/autoscale"
	"github.com/centurylinkcloud/clc-go-cli/models/backup"
	"github.com/centurylinkcloud/clc-go-cli/models/balancer"
	"github.com/centurylinkcloud/clc-go-cli/models/billing"
	"github.com/centurylinkcloud/clc-go-cli/models/crossdc_firewall"
	"github.com/centurylinkcloud/clc-go-cli/models/customfields"
	"github.com/centurylinkcloud/clc-go-cli/models/datacenter"
	"github.com/centurylinkcloud/clc-go-cli/models/db"
	"github.com/centurylinkcloud/clc-go-cli/models/firewall"
	"github.com/centurylinkcloud/clc-go-cli/models/group"
	"github.com/centurylinkcloud/clc-go-cli/models/ips"
	"github.com/centurylinkcloud/clc-go-cli/models/network"
	"github.com/centurylinkcloud/clc-go-cli/models/ospatch"
	"github.com/centurylinkcloud/clc-go-cli/models/server"
	"github.com/centurylinkcloud/clc-go-cli/models/vpn"
)

var AllCommands []base.Command = make([]base.Command, 0)

func init() {
	registerCommandBase(&server.CreateReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}",
		Resource: "server",
		Command:  "create",
		Help: help.Command{
			Brief: []string{
				"Creates a new server.",
				"Use this API operation when you want to create a new server from a standard or custom template, or clone an existing server.",
			},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Required. Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server"},
				},
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the parent group."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the parent group."},
				},
				{
					"--source-server-id",
					[]string{
						"Required unless --template-name or --source-server-name is specified. ID of the server to use as a source.",
						"Actually, it may be the name of a template, or when cloning, an existing server ID.",
					},
				},
				{
					"--source-server-name",
					[]string{
						"Required unless --source-server-id or --template-name is specified. Name of the server to use as a source.",
					},
				},
				{
					"--template-name",
					[]string{
						"Required unless --source-server-id or --source-server-name is specified. A template to create the server from.",
						"If autocomplete is turned on, available template names are shown as options.",
					},
				},
				{
					"--is-managed-os",
					[]string{
						"Whether to create the server as managed or not. Default is false.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--is-managed-backup",
					[]string{
						"Whether to add managed backup to the server. Must be a managed OS server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--network-name",
					[]string{
						"Name of the network to which to deploy the server. An alternative way to identify the network.",
					},
				},
				{
					"--ip-address",
					[]string{
						"IP address to assign to the server. If not provided, one will be assigned automatically.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--root-password",
					[]string{"Password of administrator or root user on server. If not provided, one will be generated automatically."},
				},
				{
					"--source-server-password",
					[]string{
						"Password of the source server, used only when creating a clone from an existing server.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--cpu",
					[]string{"Required. Number of processors to configure the server with (1-16). Ignored for bare metal servers."},
				},
				{
					"--cpu-autoscale-policy-id",
					[]string{
						"ID of the vertical CPU Autoscale policy to associate the server with.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--memory-gb",
					[]string{
						"Required. Number of GB of memory to configure the server with (1-128).",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--type",
					[]string{"Required. Whether to create a standard, hyperscale, or bareMetal server."},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage.",
						"If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
						"Ignored for bare metal servers.",
					},
				},
				{
					"--anti-affinity-policy-id",
					[]string{
						"ID of the Anti-Affinity policy to associate the",
						"server with. Only valid for hyperscale servers.",
					},
				},
				{
					"--anti-affinity-policy-name",
					[]string{
						"Name of the Anti-Affinity policy. An alternative way to identify the policy.",
					},
				},
				{
					"--custom-fields",
					[]string{
						"Collection of custom field ID-value pairs to set for the server.",
						"Each object of a collection has keys 'id' and 'value'.",
					},
				},
				{
					"--additional-disks",
					[]string{"Collection of disk parameters. Ignored for bare metal servers."},
				},
				{
					"--ttl",
					[]string{fmt.Sprintf("Date/time that the server should be deleted. The format is %s. Ignored for bare metal servers.", base.TIME_FORMAT)},
				},
				{
					"--packages",
					[]string{"Collection of packages to run on the server after it has been built. Ignored for bare metal servers."},
				},
				{
					"--configuration-id",
					[]string{
						"Only required for bare metal servers. Specifies the identifier for the specific configuration type of bare metal server to deploy.",
						"Ignored for standard and hyperscale servers.",
					},
				},
				{
					"--os-type",
					[]string{
						"Only required for bare metal servers. Specifies the OS to provision with the bare metal server. Currently, the only supported OS types",
						"are redHat6_64Bit, centOS6_64Bit, windows2012R2Standard_64Bit.",
						"Ignored for standard and hyperscale servers.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.DeleteReq{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Sends the delete operation to a given server and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server to be deleted."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server to be deleted."},
				},
			},
		},
	})
	registerCommandBase(&server.UpdateReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Changes the amount of CPU cores, memory (in GB), server credentials, custom fields, description, disks and server's group."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being updated."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being updated."},
				},
				{
					"--cpu",
					[]string{"The amount of CPU cores to set for the given server."},
				},
				{
					"--memory-gb",
					[]string{"The amount of memory (in GB) to set for the given server."},
				},
				{
					"--root-password",
					[]string{
						"The current and new administrator/root password values.",
						"Has to be an object with 2 fields:",
						"1) current: the current administrator/root password used to login;",
						"2) password: the new administrator/root password to change to.",
					},
				},
				{
					"--custom-fields",
					[]string{
						"A list of id-value pairs for all custom fields including all required values",
						"and other custom field values that you wish to set.",
						"",
						"Note: You must specify the complete list of custom field values to set on the server.",
						"If you want to change only one value, specify all existing field values",
						"along with the new value for the field you wish to change.",
						"To unset the value for an unrequired field, you may leave the field id-value pairing out,",
						"however all required fields must be included.",
					},
				},
				{
					"--description",
					[]string{"The description of the server to set"},
				},
				{
					"--group-id",
					[]string{"The unique identifier of the group to set as the parent."},
				},
				{
					"--group-name",
					[]string{"The name of the group to set as the parent. An alternative way to identify the group."},
				},
				{
					"--disks",
					[]string{
						"A list of information for all disks to be on the server including type (raw or partition), size, and path",
						"",
						"Note: You must specify the complete list of disks to be on the server.",
						"If you want to add or resize a disk, specify all existing disks/sizes",
						"along with a new entry for the disk to add or the new size of an existing disk.",
						"To delete a disk, just specify all the disks that should remain.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.GetReq{}, &server.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}",
		Resource: "server",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details for a individual server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being queried."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being queried."},
				},
			},
		},
	})
	registerCommandBase(&server.GetCredentialsReq{}, &server.GetCredentialsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/credentials",
		Resource: "server",
		Command:  "get-credentials",
		Help: help.Command{
			Brief: []string{"Retrieves the administrator/root password on an existing server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server with the credentials to return."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server with the credentials to return."},
				},
			},
		},
	})
	registerCommandBase(&server.GetImportsReq{}, &server.GetImportsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}/{DataCenter}/available",
		Resource: "server",
		Command:  "get-imports",
		Help: help.Command{
			Brief: []string{"Gets the list of available servers that can be imported."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Data center location identifier."},
				},
			},
		},
	})
	registerCommandBase(&server.GetIPAddressReq{}, &server.GetIPAddressRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "get-public-ip-address",
		Help: help.Command{
			Brief: []string{"Gets the details for the public IP address of a server, including the specific set of protocols and ports allowed and any source IP restrictions."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being queried."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"Required. The specific public IP to return details about."},
				},
			},
		},
	})
	registerCommandBase(&server.AddIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses",
		Resource: "server",
		Command:  "add-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Claims a public IP address and associates it with a server, allowing access to it on a given set of protocols and ports.",
				"It may also be set to restrict access based on a source IP range.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being queried."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being queried."},
				},
				{
					"--internal-ip-address",
					[]string{
						"The internal (private) IP address to map to the new public IP address.",
						"If not provided, one will be assigned for you.",
					},
				},
				{
					"--ports",
					[]string{
						"Required. The set of ports and protocols to allow access to for the new public IP address.",
						"Only these specified ports on the respective protocols will be accessible",
						"when accessing the server using the public IP address claimed here.",
						"Has to be a list of objects with fields port, portTo and protocol.",
					},
				},
				{
					"--source-restrictions",
					[]string{
						"A list of the source IP address ranges allowed to access the new public IP address.",
						"Used to restrict access to only the specified ranges of source IPs.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.RemoveIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "remove-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Releases the given public IP address of a server so that it is no longer associated with the server",
				"and available to be claimed again by another server.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being queried."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"Required. The specific public IP to remove."},
				},
			},
		},
	})
	registerCommandBase(&server.UpdateIPAddressReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/publicIPAddresses/{PublicIp}",
		Resource: "server",
		Command:  "update-public-ip-address",
		Help: help.Command{
			Brief: []string{
				"Updates a public IP address on an existing server, allowing access to it on a given set of protocols and ports",
				"as well as restricting access based on a source IP range.",
			},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server being queried."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server being queried."},
				},
				{
					"--public-ip",
					[]string{"Required. The specific public IP to update."},
				},
				{
					"--ports",
					[]string{
						"Required. The set of ports and protocols to allow access to for the public IP address.",
						"Only these specified ports on the respective protocols will be accessible",
						"when accessing the server using the public IP address claimed here.",
						"Has to be a list of objects with fields port, portTo and protocol.",
					},
				},
				{
					"--source-restrictions",
					[]string{
						"A list of the source IP address ranges allowed to access the public IP address.",
						"Used to restrict access to only the specified ranges of source IPs.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOn",
		Resource: "server",
		Command:  "power-on",
		Help: help.Command{
			Brief: []string{"Sends the power on operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform power on operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/powerOff",
		Resource: "server",
		Command:  "power-off",
		Help: help.Command{
			Brief: []string{"Sends the power off operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform power off operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/pause",
		Resource: "server",
		Command:  "pause",
		Help: help.Command{
			Brief: []string{"Sends the pause operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform pause operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/reboot",
		Resource: "server",
		Command:  "reboot",
		Help: help.Command{
			Brief: []string{"Sends the reboot operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform reboot operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/reset",
		Resource: "server",
		Command:  "reset",
		Help: help.Command{
			Brief: []string{"Sends the reset operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform reset operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/shutDown",
		Resource: "server",
		Command:  "shut-down",
		Help: help.Command{
			Brief: []string{"Sends the shut-down operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform shut-down operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.PowerReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/archive",
		Resource: "server",
		Command:  "archive",
		Help: help.Command{
			Brief: []string{"Sends the archive operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to perform archive operation on."},
				},
			},
		},
	})
	registerCommandBase(&server.RestoreReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/restore",
		Resource: "server",
		Command:  "restore",
		Help: help.Command{
			Brief: []string{"Restores a given archived server to a specified group."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required. ID of the archived server to restore."},
				},
				{
					"--target-group-id",
					[]string{"Required. The unique identifier of the target group to restore the server to."},
				},
			},
		},
	})
	registerCommandBase(&server.CreateSnapshotReq{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/createSnapshot",
		Resource: "server",
		Command:  "create-snapshot",
		Help: help.Command{
			Brief: []string{"Sends the create snapshot operation to a list of servers (along with the number of days to keep the snapshot for) and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server names to perform create snapshot operation on."},
				},
				{
					"--snapshot-expiration-days",
					[]string{"Required. Number of days to keep the snapshot for (must be between 1 and 10)."},
				},
			},
		},
	})
	registerCommandBase(&server.RevertToSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}/restore",
		Resource: "server",
		Command:  "revert-to-snapshot",
		Help: help.Command{
			Brief: []string{"Reverts a server to a snapshot."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server with the snapshot to restore."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server with the snapshot to restore."},
				},
				{
					"--snapshot-id",
					[]string{"Required. ID of the snapshot to restore."},
				},
			},
		},
	})
	registerCommandBase(&server.DeleteSnapshotReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/snapshots/{SnapshotId}",
		Resource: "server",
		Command:  "delete-snapshot",
		Help: help.Command{
			Brief: []string{"Deletes a given server snapshot."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server with the snapshot to delete."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server with the snapshot to delete."},
				},
				{
					"--snapshot-id",
					[]string{"Required. ID of the snapshot to delete."},
				},
			},
		},
	})
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/startMaintenance",
		Resource: "server",
		Command:  "start-maintenance-mode",
		Help: help.Command{
			Brief: []string{"Sends a start maintenance mode operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to start maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.MaintenanceRequest{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/stopMaintenance",
		Resource: "server",
		Command:  "stop-maintenance-mode",
		Help: help.Command{
			Brief: []string{"Sends a stop maintenance mode operation to a list of servers and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. List of server IDs to stop maintenance mode on."},
				},
			},
		},
	})
	registerCommandBase(&server.Import{}, &server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/vmImport/{accountAlias}",
		Resource: "server",
		Command:  "import",
		Help: help.Command{
			Brief: []string{"Imports a new server from an uploaded OVF."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{
						"Required. Name of the server to create. Alphanumeric characters and dashes only.",
						"Must be between 1-8 characters depending on the length of the account alias.",
						"The combination of account alias and server name here must be no more than 10 characters in length.",
						"This name will be appended with a two digit number and prepended with the datacenter code",
						"and account alias to make up the final server name.",
					},
				},
				{
					"--description",
					[]string{"User-defined description of this server."},
				},
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the parent group."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the parent group."},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--network-id",
					[]string{
						"ID of the network to which to deploy the server. If not provided, a network will be chosen automatically.",
						"If your account has not yet been assigned a network, leave this blank and one will be assigned automatically.",
					},
				},
				{
					"--network-name",
					[]string{
						"Name of the network to which to deploy the server. An alternative way to identify the network.",
					},
				},
				{
					"--root-password",
					[]string{
						"Required. Password of administrator or root user on server. This password must match",
						"the one set on the server being imported or the import will fail.",
					},
				},
				{
					"--cpu",
					[]string{
						"Required. Number of processors to configure the server with (1-16). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--memory-gb",
					[]string{
						"Required. Number of GB of memory to configure the server with (1-128). If this value is different from the one specified in the OVF,",
						"the import process will resize the server according to the value specified here.",
					},
				},
				{
					"--type",
					[]string{"Required. Whether to create standard or hyperscale server"},
				},
				{
					"--storage-type",
					[]string{
						"For standard servers, whether to use standard or premium storage. If not provided, will default to premium storage.",
						"For hyperscale servers, storage type must be hyperscale.",
					},
				},
				{
					"--custom-fields",
					[]string{
						"Collection of custom field ID-value pairs to set for the server.",
						"Each object of a collection has keys 'id' and 'value'.",
					},
				},
				{
					"--ovf-id",
					[]string{"Required. The identifier of the OVF that defines the server to import."},
				},
				{
					"--ovf-os-type",
					[]string{
						"Required. The OS type of the server being imported. Currently, the only supported OS types",
						"are redHat6_64Bit, windows2008R2DataCenter_64bit, and windows2012R2DataCenter_64Bit.",
					},
				},
			},
		},
	})
	registerCommandBase(&server.AddNetwork{}, &models.Status{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/networks",
		Resource: "server",
		Command:  "add-secondary-network",
		Help: help.Command{
			Brief: []string{"Adds a secondary network adapter to a given server in a given account."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
				{
					"--ip-address",
					[]string{"Optional IP address for the network ID."},
				},
			},
		},
	})
	registerCommandBase(&server.RemoveNetwork{}, &models.Status{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/networks/{NetworkId}",
		Resource: "server",
		Command:  "remove-secondary-network",
		Help: help.Command{
			Brief: []string{"Removes a secondary network adapter from a given server in a given account."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
			},
		},
	})
	registerCommandBase(&server.ExecutePackage{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/executePackage",
		Resource: "server",
		Command:  "execute-package",
		Help: help.Command{
			Brief: []string{"Executes a single package on one or more servers"},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. A list of server IDs to execute the package on"},
				},
				{
					"--package",
					[]string{
						"Required. The package entity describing which package to run on the specified servers.",
						"It has to contain the following fields:",
						"  package-id: unique identifier of the package to execute",
						"  paramaters: a JSON string containing a set of key-value pairs for setting the package-specific parameters defined",
					},
				},
			},
		},
	})

	registerCommandBase(&group.GetReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details for a individual server group and any sub-groups and servers that it contains."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.CreateReq{}, &group.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}",
		Resource: "group",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a new group."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{"Required. Name of the group to create."},
				},
				{
					"--description",
					[]string{"User-defined description of this group."},
				},
				{
					"--parent-group-id",
					[]string{"Required unless --parent-group-name is specified. ID of the parent group."},
				},
				{
					"--parent-group-name",
					[]string{"Required unless --parent-group-id is specified. Name of the parent group."},
				},
				{
					"--custom-fields",
					[]string{
						"Collection of custom field ID-value pairs to set for the server.",
						"Each object of a collection has keys 'id' and 'value'.",
					},
				},
			},
		},
	})
	registerCommandBase(&group.DeleteReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Sends the delete operation to a given group and adds operation to queue."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being deleted."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being deleted."},
				},
			},
		},
	})
	registerCommandBase(&group.GetBillingReq{}, &group.GetBillingRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/billing",
		Resource: "group",
		Command:  "get-billing-details",
		Help: help.Command{
			Brief: []string{"Gets the current and estimated charges for each server in a designated group hierarchy."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.GetStatsReq{}, &[]group.GetStatsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/statistics?start={Start}&end={End}&sampleInterval={SampleInterval}&type={Type}",
		Resource: "group",
		Command:  "get-monitoring-statistics",
		Help: help.Command{
			Brief: []string{
				"Gets the resource consumption details for whatever window specified in the request.",
				"Data can be retrieved for a variety of time windows and intervals.",
			},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
				{
					"--type",
					[]string{
						"Valid values are latest, hourly, or realtime.",
						"",
						"'latest' will return a single data point that reflects the last monitoring data collected.",
						"No start, end, or sampleInterval values are required for this type.",
						"",
						"'hourly' returns data points for each sampleInterval value between the start and end times provided.",
						"The start and sampleInterval parameters are both required for this type.",
						"",
						"'realtime' will return data from the last 4 hours, available in smaller increments.",
						"To use realtime type, start parameter must be within the last 4 hours.",
						"The start and sampleInterval parameters are both required for this type.",
					},
				},
				{
					"--start",
					[]string{
						fmt.Sprintf("DateTime (UTC) of the query window. The format is `%s`. Note that statistics are only held for 14 days.", base.TIME_FORMAT_REPR),
						"Start date (and optional end date) must be within the past 14 days.",
						"Value is not required if choosing the latest query type.",
					},
				},
				{
					"--end",
					[]string{
						fmt.Sprintf("DateTime (UTC) of the query window. The format is `%s`. Default is the current time in UTC.", base.TIME_FORMAT_REPR),
						"End date (and start date) must be within the past 14 days.",
						"Not a required value if results should be up to the current time.",
					},
				},
				{
					"--sample-interval",
					[]string{
						"Result interval. For the default hourly type, the minimum value is 1 hour (01:00:00)",
						"and maximum is the full window size of 14 days. Note that interval must fit within start/end window,",
						"or you will get an exception that states: 'The 'end' parameter must represent a time that occurs at least one 'sampleInterval' before 'start.'",
						"If realtime type is specified, interval can be as small as 5 minutes (05:00).",
					},
				},
			},
		},
	})
	registerCommandBase(&group.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}",
		Resource: "group",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Changes the custom fields, name, description and parent group of the given group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being updated."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being updated."},
				},
				{
					"--custom-fields",
					[]string{
						"A list of id-value pairs for all custom fields including all required values and other custom field",
						"values that you wish to set.",
						"",
						"Note: You must specify the complete list of custom field values",
						"to set on the group. If you want to change only one value,",
						"specify all existing field values along with the new value for the field you wish to change.",
						"To unset the value for an unrequired field, you may leave the field id-value pairing out,",
						"however all required fields must be included",
					},
				},
				{
					"--name",
					[]string{"The name to set for the group."},
				},
				{
					"--description",
					[]string{"The description to set for the group."},
				},
				{
					"--parent-group-id",
					[]string{"The group identifier for the new parent group."},
				},
				{
					"--parent-group-name",
					[]string{"The group name for the new parent group (an alternative way to identify it)."},
				},
			},
		},
	})
	registerCommandBase(&group.GetReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/archive",
		Resource: "group",
		Command:  "archive",
		Help: help.Command{
			Brief: []string{"Sends the archive operation to a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group to archive."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group to archive."},
				},
			},
		},
	})
	registerCommandBase(&group.RestoreReq{}, &group.RestoreRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/restore",
		Resource: "group",
		Command:  "restore",
		Help: help.Command{
			Brief: []string{"Sends the restore operation to an archived group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group to restore."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group to restore."},
				},
				{
					"--target-group-id",
					[]string{"Required unless --target-group-name is specified. The unique identifier of the target group to restore the group to."},
				},
				{
					"--target-group-name",
					[]string{"Required unless --target-group-id is specified. The name of the target group to restore the group to."},
				},
			},
		},
	})
	registerCommandBase(&group.SetHAPolicy{}, &group.HAPolicy{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/horizontalAutoscalePolicy/",
		Resource: "group",
		Command:  "set-horizontal-autoscale-policy",
		Help: help.Command{
			Brief: []string{"Applies a horizontal autoscale policy to a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
				{
					"--policy-id",
					[]string{"Required. The unique identifier of the horizontal autoscale policy."},
				},
				{
					"--load-balancer",
					[]string{
						"Required. Information about the load balancer.",
						"An object with the following required fields: Id, PublicPort, PrivatePort.",
					},
				},
			},
		},
	})
	registerCommandBase(&group.GetHAPolicy{}, &group.HAPolicy{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/horizontalAutoscalePolicy/",
		Resource: "group",
		Command:  "get-horizontal-autoscale-policy",
		Help: help.Command{
			Brief: []string{"Retrieves the details of a horizontal autoscale policy associated with a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.GetScheduledActivities{}, &[]group.ScheduledActivities{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/ScheduledActivities/",
		Resource: "group",
		Command:  "get-scheduled-activities",
		Help: help.Command{
			Brief: []string{"Gets the scheduled activities associated with a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
			},
		},
	})
	registerCommandBase(&group.SetScheduledActivities{}, &group.ScheduledActivities{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/ScheduledActivities/",
		Resource: "group",
		Command:  "set-scheduled-activities",
		Help: help.Command{
			Brief: []string{"Sets scheduled activities for a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
				{
					"--status",
					[]string{"Required. State of scheduled activity: on or off."},
				},
				{
					"--type",
					[]string{"Required. Type of activity: archive, createsnapshot, delete, deletesnapshot, pause, poweron, reboot, shutdown."},
				},
				{
					"--begin-date-utc",
					[]string{"Required. Time when scheduled activity should start."},
				},
				{
					"--repeat",
					[]string{"Required. How often to repeat: never, daily, weekly, monthly, customWeekly."},
				},
				{
					"--custom-weekly-days",
					[]string{"Required if --repeat is customWeekly. An array of strings for the days of the week: sun, mon, tue, wed, thu, fri, sat."},
				},
				{
					"--expire",
					[]string{"Required. When the scheduled activities are set to expire: never, afterDate, afterCount."},
				},
				{
					"--expire-count",
					[]string{"Required if --expire is afterCount. Number of times scheduled activity should run before expiring."},
				},
				{
					"--expire-date-utc",
					[]string{"Required if --expire is afterDate. When the scheduled activity should expire."},
				},
				{
					"--time-zone-offset",
					[]string{"Required. To display in local time."},
				},
			},
		},
	})
	registerCommandBase(&group.SetDefaults{}, &group.Defaults{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/groups/{accountAlias}/{GroupId}/defaults",
		Resource: "group",
		Command:  "set-defaults",
		Help: help.Command{
			Brief: []string{"Sets the defaults for a group."},
			Arguments: []help.Argument{
				{
					"--group-id",
					[]string{"Required unless --group-name is specified. ID of the group being queried."},
				},
				{
					"--group-name",
					[]string{"Required unless --group-id is specified. Name of the group being queried."},
				},
				{
					"--cpu",
					[]string{"Number of processors to configure the server with (1-16) (ignored for bare metal servers)"},
				},
				{
					"--memory-gb",
					[]string{"Number of GB of memory to configure the server with (1-128) (ignored for bare metal servers)"},
				},
				{
					"--network-id",
					[]string{"ID of the Network."},
				},
				{
					"--primary-dns",
					[]string{"Primary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--secondary-dns",
					[]string{"Secondary DNS to set on the server. If not supplied the default value set on the account will be used."},
				},
				{
					"--template-name",
					[]string{"Name of the template to use as the source. (Ignored for bare metal servers.)"},
				},
			},
		},
	})

	registerCustomCommand(commands.NewDatacenterList(commands.CommandExcInfo{
		Resource: "data-center",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of data centers that a given account has access to."},
			Arguments: []help.Argument{
				{
					"--with-compute-limits",
					[]string{"Get compute limits data"},
				},
				{
					"--with-network-limits",
					[]string{"Get network limits data"},
				},
				{
					"--with-available-ovfs",
					[]string{"Get available servers for import"},
				},
				{
					"--with-load-balancers",
					[]string{"Get load balancers"},
				},
			},
		},
	}))
	registerCustomCommand(commands.NewDatacenterGet(commands.CommandExcInfo{
		Resource: "data-center",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details of a specific data center the account has access to."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--with-compute-limits",
					[]string{"Get compute limits data"},
				},
				{
					"--with-network-limits",
					[]string{"Get network limits data"},
				},
				{
					"--with-available-ovfs",
					[]string{"Get available servers for import"},
				},
				{
					"--with-load-balancers",
					[]string{"Get load balancers"},
				},
			},
		},
	}))
	registerCommandBase(&datacenter.GetDCReq{}, &datacenter.GetDCRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/datacenters/{accountAlias}/{DataCenter}/deploymentCapabilities",
		Resource: "data-center",
		Command:  "get-deployment-capabilities",
		Help: help.Command{
			Brief: []string{
				"Gets the list of capabilities that a specific data center supports for a given account,",
				"including the deployable networks, OS templates, and whether features like",
				"premium storage and shared load balancer configuration are available.",
			},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
			},
		},
	})

	registerCommandBase(&network.ListReq{}, &[]network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}",
		Resource: "network",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of networks available for a given account in a given data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
			},
		},
	})
	registerCommandBase(&network.GetReq{}, &network.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{NetworkId}?ipAddresses={IpAddresses}",
		Resource: "network",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details of a specific network in a given data center for a given account."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
				{
					"--ip-addresses",
					[]string{
						"Optional component of the query to request details of IP Addresses in a certain state.",
						"Should be one of the following:",
						"none: returns details of the network only,",
						"claimed: returns details of the network as well as information about claimed IP addresses,",
						"free: returns details of the network as well as information about free IP addresses or",
						"all: returns details of the network as well as information about all IP addresses.",
					},
				},
			},
		},
	})
	registerCommandBase(&network.ListIpAddresses{}, &[]network.IpAddress{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{NetworkId}/ipAddresses?type={Type}",
		Resource: "network",
		Command:  "list-ip-addresses",
		Help: help.Command{
			Brief: []string{"Gets the list of IP addresses for a network in a given data center for a given account."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
				{
					"--type",
					[]string{
						"Optional component of the query to request details of IP Addresses in a certain state.",
						"Should be one of the following:",
						"claimed: returns details of the network as well as information about claimed IP addresses,",
						"free: returns details of the network as well as information about free IP addresses or",
						"all: returns details of the network as well as information about all IP addresses",
					},
				},
			},
		},
	})
	registerCommandBase(&network.CreateReq{}, &models.LinkEntity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/claim",
		Resource: "network",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Claims a network for a given account in a given data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
			},
		},
	})
	registerCommandBase(&network.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{NetworkId}",
		Resource: "network",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates the attributes of a given Network."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
				{
					"--name",
					[]string{"Required. User-defined name of the network; the default is the VLAN number combined with the network address."},
				},
				{
					"--description",
					[]string{"Required. Description of VLAN, a free text field that defaults to the VLAN number combined with the network address."},
				},
			},
		},
	})
	registerCommandBase(&network.ReleaseReq{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/networks/{accountAlias}/{DataCenter}/{NetworkId}/release",
		Resource: "network",
		Command:  "release",
		Help: help.Command{
			Brief: []string{"Releases a network from a given account in a given data center to a pool for another user to claim."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--network-id",
					[]string{"Required unless --network-name is specified. ID of the network."},
				},
				{
					"--network-name",
					[]string{"Required unless --network-id is specified. Name of the network."},
				},
			},
		},
	})

	registerCommandBase(&alert.CreateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates an alert policy in a given account."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{"Required. Name of the alert policy."},
				},
				{
					"--actions",
					[]string{
						"Required. The actions to perform when the alert is triggered.",
						"",
						"Has to be a list of objects with 2 fields in each: action and settings.",
						"The only action currently supported by alerts is 'email'.",
						"The only settings value supported currently is an object with the 'recipients' field,",
						"which is an array of email addresses to be notified when the alert is triggered.",
					},
				},
				{
					"--triggers",
					[]string{
						"Required. The definition of the triggers that fire the alert.",
						"",
						"Has to be a list of objects with 3 fields each: metric, duration and threshold.",
						"metric: the metric on which to measure the condition that will trigger the alert: cpu, memory, or disk.",
						"duration: the length of time in minutes that the condition must exceed the threshold: 00:05:00, 00:10:00, 00:15:00.",
						"threshold: the threshold that will trigger the alert when the metric equals or exceeds it.",
						"This number represents a percentage and must be a value between 5.0 - 95.0 that is a multiple of 5.0.",
					},
				},
			},
		},
	})
	registerCommandBase(nil, &alert.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}",
		Resource: "alert-policy",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets a list of alert policies within a given account."},
		},
	})
	registerCommandBase(&alert.GetReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets a given alert policy by ID."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the alert policy being queried."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the alert policy being queried."},
				},
			},
		},
	})
	registerCommandBase(&alert.UpdateReq{}, &alert.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates the name of an alert policy in a given account."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the alert policy being updated."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the alert policy being updated."},
				},
				{
					"--name",
					[]string{"Required. Name of the alert policy."},
				},
				{
					"--actions",
					[]string{
						"Required. The actions to perform when the alert is triggered.",
						"",
						"Has to be an object with 2 fields: action and settings.",
						"The only action currently supported by alerts is 'email'.",
						"The only setting currently supported is the 'recipients' list, which is an array of",
						"email addresses to be notified when the alert is triggered.",
					},
				},
				{
					"--triggers",
					[]string{
						"Required. The definition of the triggers that fire the alert.",
						"",
						"Has to be an object with 3 fields: metric, duration and threshold.",
						"metric: the metric on which to measure the condition that will trigger the alert: cpu, memory, or disk.",
						"duration: the length of time in minutes that the condition must exceed the threshold: 00:05:00, 00:10:00, 00:15:00.",
						"threshold: the threshold that will trigger the alert when the metric equals or exceeds it.",
						"This number represents a percentage and must be a value between 5.0 - 95.0 that is a multiple of 5.0.",
					},
				},
			},
		},
	})
	registerCommandBase(&alert.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/alertPolicies/{accountAlias}/{PolicyId}",
		Resource: "alert-policy",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a given alert policy by ID."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the alert policy being deleted."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the alert policy being deleted."},
				},
			},
		},
	})

	registerCommandBase(&affinity.CreateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates an anti-affinity policy in a given account."},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{"Required. Name of the anti-affinity policy."},
				},
				{
					"--data-center",
					[]string{"Required. Data center location of the anti-affinity policy."},
				},
			},
		},
	})
	registerCommandBase(nil, &affinity.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}",
		Resource: "anti-affinity-policy",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets a list of anti-affinity policies within a given account."},
		},
	})
	registerCommandBase(&affinity.GetReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets a given anti-affinity policy by ID."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the anti-affinity policy being queried."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the anti-affinity policy being queried."},
				},
			},
		},
	})
	registerCommandBase(&affinity.UpdateReq{}, &affinity.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates the name of an anti-affinity policy in a given account."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the anti-affinity policy being updated."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the anti-affinity policy being updated."},
				},
				{
					"--name",
					[]string{"Required. Name of the anti-affinity policy."},
				},
			},
		},
	})
	registerCommandBase(&affinity.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/antiAffinityPolicies/{accountAlias}/{PolicyId}",
		Resource: "anti-affinity-policy",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a given anti-affinity policy by ID."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the anti-affinity policy being deleted."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the anti-affinity policy being deleted."},
				},
			},
		},
	})

	registerCommandBase(&firewall.CreateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{DataCenter}",
		Resource: "firewall-policy",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a firewall policy for a given account in a given data center ('intra data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the target data center for the new policy."},
				},
				{
					"--destination-account",
					[]string{"Required. Short code for a particular account."},
				},
				{
					"--sources",
					[]string{"Required. Source addresses for traffic on the originating firewall, specified using CIDR notation, on the originating firewall."},
				},
				{
					"--destinations",
					[]string{"Required. Destination addresses for traffic on the terminating firewall, specified using CIDR notation."},
				},
				{
					"--ports",
					[]string{
						"Required. Type of ports associated with the policy. Supported ports include: any, icmp, TCP and UDP",
						"with single ports (tcp/123, udp/123) and port ranges (tcp/123-456, udp/123-456).",
						"Some common ports include: tcp/21 (for FTP), tcp/990 (FTPS), tcp/80 (HTTP 80), tcp/8080 (HTTP 8080), tcp/443 (HTTPS 443),",
						"icmp (PING), tcp/3389 (RDP), and tcp/22 (SSH/SFTP).",
					},
				},
			},
		},
	})
	registerCommandBase(&firewall.ListReq{}, &[]firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{DataCenter}?destinationAccount={DestinationAccountAlias}",
		Resource: "firewall-policy",
		Command:  "list",
		Help: help.Command{
			Brief: []string{
				"Gets the list of firewall policies associated with a given account in a given data center ('intra data center firewall policies').",
				"Users can optionally filter results by requesting policies associated with a second 'destination' account.",
			},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--destination-account-alias",
					[]string{"Short code for a particular account."},
				},
			},
		},
	})
	registerCommandBase(&firewall.GetReq{}, &firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details of a specific firewall policy associated with a given account in a given data center (an 'intra data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
			},
		},
	})
	registerCommandBase(&firewall.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates a given firewall policy associated with a given account in a given data center (an 'intra data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center associated with the policy of interest."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
				{
					"--enabled",
					[]string{"Indicates if the policy is enabled (true) or disabled (false)."},
				},
				{
					"--sources",
					[]string{"Required. Source addresses for traffic on the originating firewall, specified using CIDR notation."},
				},
				{
					"--destinations",
					[]string{"Required. Destination addresses for traffic on the terminating firewall, specified using CIDR notation."},
				},
				{
					"--ports",
					[]string{
						"Required. Type of ports associated with the policy. Supported ports include: any, icmp, TCP and UDP",
						"with single ports (tcp/123, udp/123) and port ranges (tcp/123-456, udp/123-456).",
						"Some common ports include: tcp/21 (for FTP), tcp/990 (FTPS), tcp/80 (HTTP 80), tcp/8080 (HTTP 8080), tcp/443 (HTTPS 443),",
						"icmp (PING), tcp/3389 (RDP), and tcp/22 (SSH/SFTP).",
					},
				},
			},
		},
	})
	registerCommandBase(&firewall.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2-experimental/firewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "firewall-policy",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a firewall policy for a given account in a given data center ('intra data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
			},
		},
	})

	registerCommandBase(&crossdc_firewall.CreateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2-experimental/crossDcFirewallPolicies/{accountAlias}/{DataCenter}",
		Resource: "crossdc-firewall-policy",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a firewall policy for a given account, between networks in different data centers ('cross data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the target data center for the new policy."},
				},
				{
					"--destination-account-id",
					[]string{"Required. Short code for a particular account."},
				},
				{
					"--destination-location-id",
					[]string{"Required. Short code for a particular location."},
				},
				{
					"--destination-cidr",
					[]string{"Required. Destination address for traffic on the terminating firewall, specified using CIDR notation."},
				},
				{
					"--source-cidr",
					[]string{"Required. Source address for traffic on the originating firewall, specified using CIDR notation, on the originating firewall."},
				},
				{
					"--enabled",
					[]string{"Indicates if the policy is enabled (true) or disabled (false)."},
				},
			},
		},
	})
	registerCommandBase(&crossdc_firewall.ListReq{}, &[]crossdc_firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/crossDcFirewallPolicies/{accountAlias}/{DataCenter}?destinationAccount={DestinationAccountAlias}",
		Resource: "crossdc-firewall-policy",
		Command:  "list",
		Help: help.Command{
			Brief: []string{
				"Gets the list of firewall policies associated with a given account, between networks in different data centers ('cross data center firewall policies').",
				"Users can optionally filter results by requesting policies associated with a second 'destination' account",
			},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--destination-account-alias",
					[]string{"Short code for a particular account."},
				},
			},
		},
	})
	registerCommandBase(&crossdc_firewall.GetReq{}, &crossdc_firewall.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2-experimental/crossDcFirewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "crossdc-firewall-policy",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the details of a specific firewall policy associated with a given account in a given data center (an 'cross data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
			},
		},
	})
	registerCommandBase(&crossdc_firewall.UpdateReq{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2-experimental/crossDcFirewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}?enabled={Enabled}",
		Resource: "crossdc-firewall-policy",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates a given firewall policy associated with a given account in a given data center (an 'cross data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center associated with the policy of interest."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
				{
					"--enabled",
					[]string{"Indicates if the policy is enabled (true) or disabled (false)."},
				},
			},
		},
	})
	registerCommandBase(&crossdc_firewall.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2-experimental/crossDcFirewallPolicies/{accountAlias}/{DataCenter}/{FirewallPolicy}",
		Resource: "crossdc-firewall-policy",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a firewall policy for a given account in a given data center ('cross data center firewall policy')."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center you are querying."},
				},
				{
					"--firewall-policy",
					[]string{"Required. ID of the firewall policy."},
				},
			},
		},
	})

	registerCommandBase(&balancer.CreatePool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a new shared load balancer configuration for a given account and data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
				{
					"--port",
					[]string{"Required. Port to configure on the public-facing side of the load balancer pool. Must be either 80 (HTTP) or 443 (HTTPS)."},
				},
				{
					"--method",
					[]string{"The balancing method for this load balancer, either leastConnection or roundRobin. Default is roundRobin."},
				},
				{
					"--persistence",
					[]string{"The persistence method for this load balancer, either standard or sticky. Default is standard."},
				},
			},
		},
	})
	registerCommandBase(&balancer.Create{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a new shared load balancer configuration for a given account and data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--name",
					[]string{"Required. Friendly name for new the load balancer."},
				},
				{
					"--description",
					[]string{"Required. Description for new the load balancer."},
				},
				{
					"--status",
					[]string{"Status to create the load balancer with: enabled or disabled."},
				},
			},
		},
	})
	registerCommandBase(&balancer.ListPools{}, &[]balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools",
		Resource: "load-balancer-pool",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of pools configured for a given shared load balancer."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
			},
		},
	})
	registerCommandBase(&balancer.List{}, &[]balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}",
		Resource: "load-balancer",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of shared load balancers that exist for a given account and data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
			},
		},
	})
	registerCommandBase(&balancer.GetPool{}, &balancer.Pool{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets a specified pool configured for the given shared load balancer."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
				{
					"--pool-id",
					[]string{"Required. ID of the pool."},
				},
			},
		},
	})
	registerCommandBase(&balancer.Get{}, &balancer.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets the specified shared load balancer for a given account and data center."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
			},
		},
	})
	registerCommandBase(&balancer.UpdatePool{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates a given shared load balancer pool."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
				{
					"--pool-id",
					[]string{"Required. ID of the pool to update."},
				},
				{
					"--method",
					[]string{"The balancing method for this load balancer, either leastConnection or roundRobin."},
				},
				{
					"--persistence",
					[]string{"The persistence method for this load balancer, either standard or sticky."},
				},
			},
		},
	})
	registerCommandBase(&balancer.Update{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates a given shared load balancer."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer to update."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer to update."},
				},
				{
					"--name",
					[]string{"Required. Friendly name for new the load balancer."},
				},
				{
					"--description",
					[]string{"Required. Description for new the load balancer."},
				},
				{
					"--status",
					[]string{"Status to create the load balancer with: enabled or disabled."},
				},
			},
		},
	})
	registerCommandBase(&balancer.GetNodes{}, &[]balancer.Node{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "get-nodes",
		Help: help.Command{
			Brief: []string{"Gets the list of nodes configured behind a given shared load balancer pool."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
				{
					"--pool-id",
					[]string{"Required. ID of the pool containing the nodes."},
				},
			},
		},
	})
	registerCommandBase(&balancer.UpdateNodes{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}/nodes",
		Resource: "load-balancer",
		Command:  "update-nodes",
		Help: help.Command{
			Brief: []string{"Updates the nodes behind a given shared load balancer pool."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer."},
				},
				{
					"--pool-id",
					[]string{"Required. ID of the pool to update."},
				},
				{
					"--nodes",
					[]string{
						"A list of objects each representing a node.",
						"Each object must have the following fields:",
						"'ip-address': the internal (private) IP address of the node server;",
						"'private-port': the internal (private) port of the node server. Must be a value between 1 and 65535.",
						"The object may optionally have a 'status' field:",
						"status of the node: enabled, disabled or deleted.",
					},
				},
			},
		},
	})
	registerCommandBase(&balancer.DeletePool{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}/pools/{PoolId}",
		Resource: "load-balancer-pool",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a given shared load balancer by ID."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer with the pool to delete."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer with the pool to delete."},
				},
				{
					"--pool-id",
					[]string{"Required. ID of the pool to delete."},
				},
			},
		},
	})
	registerCommandBase(&balancer.Delete{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/sharedLoadBalancers/{accountAlias}/{DataCenter}/{LoadBalancerId}",
		Resource: "load-balancer",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a given shared load balancer by ID."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required. Short string representing the data center where the load balancer is."},
				},
				{
					"--load-balancer-id",
					[]string{"Required unless --load-balancer-name is specified. ID of the load balancer to delete."},
				},
				{
					"--load-balancer-name",
					[]string{"Required unless --load-balancer-id is specified. Name of the load balancer to delete."},
				},
			},
		},
	})

	registerCommandBase(nil, &[]customfields.GetRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/accounts/{accountAlias}/customFields",
		Resource: "custom-fields",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Retrieves the custom fields defined for a given account."},
		},
	})

	registerCommandBase(&billing.GetInvoiceData{}, &billing.InvoiceData{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/invoice/{accountAlias}/{Year}/{Month}?pricingAccount={PricingAccountAlias}",
		Resource: "billing",
		Command:  "get-invoice-data",
		Help: help.Command{
			Brief: []string{
				"Gets a list of invoicing data for a given account alias for a given month.",
				"API NOTE: The data returned in this request are usage estimates only, and does not represent an actual bill.",
			},
			Arguments: []help.Argument{
				{
					"--year",
					[]string{"Required. Year of usage, in YYYY format."},
				},
				{
					"--month",
					[]string{"Monthly period of usage, a number between 1 and 12."},
				},
				{
					"--pricing-account-alias",
					[]string{"Short code of the account that sends the invoice for the accountAlias"},
				},
			},
		},
	})

	registerCustomCommand(commands.NewGroupList(commands.CommandExcInfo{
		Resource: "group",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of groups of the given data-center or of all data centers."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required unless the --all option is set. A short code of the data center to query."},
				},
				{
					"--all",
					[]string{"Forces the command to query all of the data centers."},
				},
			},
		},
	}))
	registerCustomCommand(commands.NewServerList(commands.CommandExcInfo{
		Resource: "server",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Gets the list of servers of the given data-center or of all data centers."},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Required unless the --all option is set. A short code of the data center to query."},
				},
				{
					"--all",
					[]string{"Forces the command to query all of the data centers."},
				},
			},
		},
	}))
	registerCustomCommand(commands.NewWait(commands.CommandExcInfo{
		Resource: "wait",
		Help: help.Command{
			Brief:           []string{"Waits for the previous command to complete."},
			AccountAgnostic: true,
		},
	}))
	registerCustomCommand(commands.NewLogin(commands.CommandExcInfo{
		Resource: "login",
		Help: help.Command{
			Brief: []string{
				"Logs the user in by saving his credentials to the config.",
				"Specify the credentials using the --user and --password options.",
			},
			NoEnvVars:       true,
			AccountAgnostic: true,
		},
	}))

	registerCustomCommand(commands.NewSetDefaultDC(commands.CommandExcInfo{
		Resource: "data-center",
		Command:  "set-default",
		Help: help.Command{
			Brief: []string{
				"Sets a default data center to work with.",
			},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Short code for the data center being set."},
				},
			},
			AccountAgnostic: true,
		},
	}))
	registerCustomCommand(commands.NewUnsetDefaultDC(commands.CommandExcInfo{
		Resource: "data-center",
		Command:  "unset-default",
		Help: help.Command{
			Brief: []string{
				"Unsets the default data center.",
			},
			AccountAgnostic: true,
		},
	}))
	registerCustomCommand(commands.NewShowDefaultDC(commands.CommandExcInfo{
		Resource: "data-center",
		Command:  "show-default",
		Help: help.Command{
			Brief: []string{
				"Show the default data center set, if any.",
			},
			AccountAgnostic: true,
		},
	}))

	registerCommandBase(&autoscale.ListReq{}, &[]autoscale.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/autoscalePolicies/{accountAlias}",
		Resource: "autoscale-policy",
		Command:  "list",
		Help: help.Command{
			Brief: []string{
				"Gets a list of vertical autoscale policies for a given account.",
			},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(&autoscale.GetReq{}, &autoscale.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/autoscalePolicies/{accountAlias}/{PolicyId}",
		Resource: "autoscale-policy",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Gets a given vertical autoscale policy."},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the autoscale policy being queried."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the autoscale policy being queried."},
				},
			},
		},
	})
	registerCommandBase(&autoscale.RemoveOnServerReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/cpuAutoscalePolicy",
		Resource: "autoscale-policy",
		Command:  "remove-on-server",
		Help: help.Command{
			Brief: []string{"Removes the autoscale policy from a given server, if the policy has first been applied to the server. "},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server."},
				},
			},
		},
	})
	registerCommandBase(&autoscale.SetOnServerReq{}, &autoscale.SetOnServerRes{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/cpuAutoscalePolicy",
		Resource: "autoscale-policy",
		Command:  "set-on-server",
		Help: help.Command{
			Brief: []string{"Sets the autoscale policy for a specified server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server."},
				},
				{
					"--policy-id",
					[]string{"Required unless --policy-name is specified. ID of the autoscale policy being queried."},
				},
				{
					"--policy-name",
					[]string{"Required unless --policy-id is specified. Name of the autoscale policy being queried."},
				},
			},
		},
	})
	registerCommandBase(&autoscale.ViewOnServerReq{}, &autoscale.ViewOnServerRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/servers/{accountAlias}/{ServerId}/cpuAutoscalePolicy",
		Resource: "autoscale-policy",
		Command:  "view-on-server",
		Help: help.Command{
			Brief: []string{"Gets the autoscale policy of a given server, if a policy has been applied on the server."},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. ID of the server."},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. Name of the server."},
				},
			},
		},
	})
	registerCommandBase(&db.Create{}, &db.CreateRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription",
		Resource: "db",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a database subscription"},
			Arguments: []help.Argument{
				{
					"--instance-type",
					[]string{"Required. 'MySQL' or 'MySQL_REPLICATION'."},
				},
				{
					"--external-id",
					[]string{"Requried. A service ID."},
				},
				{
					"--machine-config",
					[]string{"Required. An object with the keys: cpu, memory, and storage. All keys are integers."},
				},
				{
					"--backup-retention-days",
					[]string{"Required. A number of days defining the backup living time."},
				},
				{
					"--users",
					[]string{"Required. An array of objects with the name and password fields."},
				},
				{
					"--data-center",
					[]string{"A data-center location identifier. The server default (is applied if no client default is set) is VA1."},
				},
				{
					"--destinations",
					[]string{
						"An array of the objects to set up the notifications.",
						"Each object has the following keys:",
						"  destination-type: 'EMAIL' or 'SMS'",
						"  location: an email address or a phone number",
						"  notifications: an array of the notification objects",
						"Each notification object has the following fields:",
						"  notification-type: 'CPU_UTILIZATION' or 'MEMORY_UTILIZATION' or 'STORAGE_UTILIZATION'",
					},
				},
				{
					"--instances",
					[]string{"An array of the instance objects with the name property."},
				},
				{
					"--backup-time",
					[]string{"An object with the hour and minute fields. Both fields are integers. Defines the time to make backups at."},
				},
			},
		},
	})
	registerCommandBase(&db.List{}, &db.ListRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/list?status={Status}&dataCenter={DataCenter}",
		Resource: "db",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Shows the database subscriptions"},
			Arguments: []help.Argument{
				{
					"--data-center",
					[]string{"Queries by the data center."},
				},
				{
					"--status",
					[]string{
						"Queries by the subscription status.",
						"Either 'PENDING' or 'READY' or 'ACTIVE' or 'DELETED' or 'FAILED' or 'UNKNOWN' or 'SUCCESS' or 'CONFIGURING' or 'TERMINATED'.",
					},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, &db.CreateRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}",
		Resource: "db",
		Command:  "get",
		Help: help.Command{
			Brief: []string{"Shows the database subscription info"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, new(base.BinaryResponse), commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/certificate",
		Resource: "db",
		Command:  "get-certificate",
		Help: help.Command{
			Brief: []string{"Shows the database certificate info"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}",
		Resource: "db",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes the database subscription"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to delete"},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/failover",
		Resource: "db",
		Command:  "failover",
		Help: help.Command{
			Brief: []string{"Performs an update to the connection to cause a failover to the secondary server"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, &[]db.BackupResponse{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/backup",
		Resource: "db",
		Command:  "list-backups",
		Help: help.Command{
			Brief: []string{"Shows the backups list"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
			},
		},
	})
	registerCommandBase(&db.DeleteBackup{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/backup/{BackupId}",
		Resource: "db",
		Command:  "delete-backup",
		Help: help.Command{
			Brief: []string{"Deletes the backup"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
				{
					"--backup-id",
					[]string{"Required. ID of a backup to delete"},
				},
			},
		},
	})
	registerCommandBase(nil, &[]db.ActionLogEntry{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/customer/history",
		Resource: "db",
		Command:  "history",
		Help: help.Command{
			Brief:     []string{"Gets up to 10 most recent actions performed by the account"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(nil, &[]db.Promotion{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/customer/promotion",
		Resource: "db",
		Command:  "promotions",
		Help: help.Command{
			Brief:     []string{"Gets promotions applied to the account"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(nil, &[]db.DataCenterInfo{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/datacenter/list",
		Resource: "db",
		Command:  "list-datacenters",
		Help: help.Command{
			Brief:     []string{"Displays available datacenters"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(&db.CreateNotification{}, &db.Destination{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/notification",
		Resource: "db",
		Command:  "create-notification",
		Help: help.Command{
			Brief: []string{"Adds a destination to a subscription to receive notifications"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
				{
					"--destination-type",
					[]string{"Required. 'EMAIL' or 'SMS'"},
				},
				{
					"--location",
					[]string{"Required. An email address or a phone number"},
				},
				{
					"--notifications",
					[]string{
						"Required. An array of the notification objects.",
						"Each notification object has the following fields:",
						"  notification-type: 'CPU_UTILIZATION' or 'MEMORY_UTILIZATION' or 'STORAGE_UTILIZATION'",
					},
				},
			},
		},
	})
	registerCommandBase(&db.Get{}, &[]db.Destination{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/notification",
		Resource: "db",
		Command:  "get-destinations",
		Help: help.Command{
			Brief: []string{"Retrieves destinations receiving notifications for a subscription"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
			},
		},
	})
	registerCommandBase(&db.UpdateNotification{}, &db.Destination{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/notification/{DestinationId}",
		Resource: "db",
		Command:  "update-notification",
		Help: help.Command{
			Brief: []string{"Updates destinations receiving notifications for a subscription"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
				{
					"--destination-id",
					[]string{"Required. ID of a destination to update"},
				},
				{
					"--destination-type",
					[]string{"Required. 'EMAIL' or 'SMS'"},
				},
				{
					"--location",
					[]string{"Required. An email address or a phone number"},
				},
				{
					"--notifications",
					[]string{
						"Required. An array of the notification objects.",
						"Each notification object has the following fields:",
						"  notification-type: 'CPU_UTILIZATION' or 'MEMORY_UTILIZATION' or 'STORAGE_UTILIZATION'",
					},
				},
			},
		},
	})
	registerCommandBase(&db.DeleteNotification{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/notification/{DestinationId}",
		Resource: "db",
		Command:  "delete-notification",
		Help: help.Command{
			Brief: []string{"Removes a destination from a subscription"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
				{
					"--destination-id",
					[]string{"Required. ID of a destination to delete"},
				},
			},
		},
	})
	registerCommandBase(&db.VerifyDestination{}, &db.Destination{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}/notification/{DestinationId}/verify/{Token}",
		Resource: "db",
		Command:  "verify-destination",
		Help: help.Command{
			Brief: []string{"Verifies a destination for notifications"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to query"},
				},
				{
					"--destination-id",
					[]string{"Required. ID of a destination to verify"},
				},
				{
					"--token",
					[]string{"Required. A token to use for verification"},
				},
			},
		},
	})
	registerCommandBase(&db.Update{}, new(string), commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api.rdbs.ctl.io/{accountAlias}/subscription/{SubscriptionId}",
		Resource: "db",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates the database subscription"},
			Arguments: []help.Argument{
				{
					"--subscription-id",
					[]string{"Required. ID of a subscription to update"},
				},
				{
					"--machine-config",
					[]string{"An object with the keys: cpu, memory, and storage. All keys are integers."},
				},
				{
					"--backup-retention-days",
					[]string{"A number of days defining the backup living time."},
				},
				{
					"--backup-time",
					[]string{"An object with the hour and minute fields. Both fields are integers. Defines the time to make backups at."},
				},
			},
		},
	})

	registerCommandBase(&ips.Install{}, new(string), commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.client-security.ctl.io/ips/api/app",
		Resource: "ips",
		Command:  "install",
		Help: help.Command{
			Brief: []string{"Installs an IPS agent on the designated host"},
			Arguments: []help.Argument{
				{
					"--server-name",
					[]string{"Required. The name of the server that the destination should be set for"},
				},
			},
		},
	})
	registerCommandBase(&ips.Install{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.client-security.ctl.io/ips/api/app",
		Resource: "ips",
		Command:  "uninstall",
		Help: help.Command{
			Brief: []string{"Uninstalls an IPS agent from a designated host"},
			Arguments: []help.Argument{
				{
					"--server-name",
					[]string{"Required. The name of the server that the agent should be uninstalled from"},
				},
			},
		},
	})
	registerCommandBase(&ips.SetNotifications{}, new(string), commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.client-security.ctl.io/ips/api/notifications/{accountAlias}/{ServerName}",
		Resource: "ips",
		Command:  "set-notifications",
		Help: help.Command{
			Brief: []string{"Sets a notification destination"},
			Arguments: []help.Argument{
				{
					"--server-name",
					[]string{"Required. The name of the server that the destination should be set for"},
				},
				{
					"--notification-destinations",
					[]string{
						"Required. An array of the notification destination objects.",
						"Each object has the following keys:",
						"  url: the URL endpoint for WEBHOOK or SLACK notification",
						"  type-code: the type of destination. 'SYSLOG' or 'EMAIL' or 'WEBHOOK'",
						"  email-address: a string with the options for an EMAIL notification",
						"  sys-log-settings: an object with the options for SYSLOG; contains the following fields:",
						"    ip-address: an IP address of a syslog server",
						"    udp-port: the port the server is listening on",
						"    facility: a syslog facility code; 16-23",
					},
				},
			},
		},
	})
	registerCommandBase(&ips.GetNotifications{}, &[]ips.Notification{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.client-security.ctl.io/ips/api/notifications/{accountAlias}/{ServerName}",
		Resource: "ips",
		Command:  "get-notifications",
		Help: help.Command{
			Brief: []string{"Gets a notification destination"},
			Arguments: []help.Argument{
				{
					"--server-name",
					[]string{"Required. The name of the server to query"},
				},
			},
		},
	})
	registerCommandBase(&ips.GetNotifications{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.client-security.ctl.io/ips/api/notifications/{accountAlias}/{ServerName}",
		Resource: "ips",
		Command:  "delete-notifications",
		Help: help.Command{
			Brief: []string{"Deletes the server's notification destination"},
			Arguments: []help.Argument{
				{
					"--server-name",
					[]string{"Required. The name of the server to query"},
				},
			},
		},
	})

	registerCommandBase(&ospatch.Patch{}, &[]server.ServerRes{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/operations/{accountAlias}/servers/executePackage",
		Resource: "os-patch",
		Command:  "apply",
		Help: help.Command{
			Brief: []string{"Patches the given servers with the latest available patches provided by the OS vendor"},
			Arguments: []help.Argument{
				{
					"--server-ids",
					[]string{"Required. A list of server IDs to execute the package on"},
				},
				{
					"--os-type",
					[]string{
						"Required. 'Windows2012' or 'RedHat'",
						"'Windows2012' patches Windows 2012 and Windows 2012R2",
						"'RedHat' patches RedHat Enterprise Linux 5,6,7, and CentOS 5,6",
					},
				},
			},
		},
	})
	registerCommandBase(&ospatch.List{}, &[]ospatch.PatchInfo{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://patching.useast.appfog.ctl.io/rest/servers/{accountAlias}/server/{ServerId}/patch",
		Resource: "os-patch",
		Command:  "list",
		Help: help.Command{
			Brief: []string{"Shows a history of all the patches deployed to a given server"},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. The ID of a server to query"},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. The name of a server to query"},
				},
			},
		},
	})
	registerCommandBase(&ospatch.ListDetails{}, &ospatch.PatchDetails{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://patching.useast.appfog.ctl.io/rest/servers/{accountAlias}/server/{ServerId}/patch/{ExecutionId}",
		Resource: "os-patch",
		Command:  "list-details",
		Help: help.Command{
			Brief: []string{"Returns details on all attempted patches for a single execution against a server"},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. The ID of a server to query"},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. The name of a server to query"},
				},
				{
					"--execution-id",
					[]string{
						"Required. Correlation ID for all the patches included with a single update execution, obtained from the os-patch list response",
						"or emails regarding a patch request. The execution ID format will be aa#-######",
					},
				},
			},
		},
	})

	registerCommandBase(&backup.AccountPolicyReq{}, &backup.AccountPolicy{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies",
		Resource: "backup",
		Command:  "create-account-policy",
		Help: help.Command{
			Brief: []string{"Creates a backup policy for the account"},
			Arguments: []help.Argument{
				{
					"--name",
					[]string{"Required. A name for the policy"},
				},
				{
					"--os-type",
					[]string{"Required. Either 'Linux' or 'Windows'"},
				},
				{
					"--paths",
					[]string{"Required. A list of paths to backup"},
				},
				{
					"--excluded-directory-paths",
					[]string{"A list of paths to exclude from the backup"},
				},
				{
					"--backup-interval-hours",
					[]string{"Required. A number of hours defining the backup interval"},
				},
				{
					"--retention-days",
					[]string{"Required. A number of days to keep backups for"},
				},
			},
		},
	})
	registerCommandBase(&backup.AccountPoliciesReq{}, &backup.AccountPoliciesRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies?limit={Limit}&offset={Offset}&withStatus={WithStatus}&sortBy={SortBy}&ascendingSort={AscendingSort}",
		Resource: "backup",
		Command:  "get-account-policies",
		Help: help.Command{
			Brief: []string{"Shows the list of the account backup policies"},
			Arguments: []help.Argument{
				{
					"--limit",
					[]string{"A number of results to show"},
				},
				{
					"--offset",
					[]string{"A number of a result to start from"},
				},
				{
					"--with-status",
					[]string{"A status to filter the results by"},
				},
				{
					"--sort-by",
					[]string{
						"A field to sort the results by.",
						"'status' or 'osType' or 'name' or 'policyId' or 'backupIntervalHours' or 'retentionDays'",
					},
				},
				{
					"--ascending-sort",
					[]string{"Defines in which order to sort the results. 'true' or 'false'"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetAccountPolicy{}, &backup.AccountPolicy{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{PolicyId}",
		Resource: "backup",
		Command:  "get-account-policy",
		Help: help.Command{
			Brief: []string{"Shows the backup account policy description"},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required. An ID of the account policy to query"},
				},
			},
		},
	})
	registerCommandBase(&backup.AllowedAccountPoliciesReq{}, &backup.AccountPoliciesRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/servers/{ServerId}?limit={Limit}&offset={Offset}&withStatus={WithStatus}&sortBy={SortBy}&ascendingSort={AscendingSort}",
		Resource: "backup",
		Command:  "get-allowed-account-policies",
		Help: help.Command{
			Brief: []string{"Shows the list of the account backup policies eligible for the given server"},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. An ID of a server to query"},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. A name of a server to query"},
				},
				{
					"--limit",
					[]string{"A number of results to show"},
				},
				{
					"--offset",
					[]string{"A number of a result to start from"},
				},
				{
					"--with-status",
					[]string{"A status to filter the results by"},
				},
				{
					"--sort-by",
					[]string{
						"A field to sort the results by.",
						"'status' or 'osType' or 'name' or 'policyId' or 'backupIntervalHours' or 'retentionDays'",
					},
				},
				{
					"--ascending-sort",
					[]string{"Defines in which order to sort the results. 'true' or 'false'"},
				},
			},
		},
	})
	registerCommandBase(&backup.UpdateAccountPolicy{}, &backup.AccountPolicy{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{PolicyId}",
		Resource: "backup",
		Command:  "update-account-policy",
		Help: help.Command{
			Brief: []string{"Updates the given backup account policy"},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required. An ID of a policy to update"},
				},
				{
					"--name",
					[]string{"Required. A name for the policy"},
				},
				{
					"--clc-account-alias",
					[]string{"An alias of a Centurylink Cloud account to bind the policy to"},
				},
				{
					"--os-type",
					[]string{"Required. Either 'Linux' or 'Windows'"},
				},
				{
					"--paths",
					[]string{"Required. A list of paths to backup"},
				},
				{
					"--excluded-directory-paths",
					[]string{"A list of paths to exclude from the backup"},
				},
				{
					"--backup-interval-hours",
					[]string{"Required. A number of hours defining the backup interval"},
				},
				{
					"--retention-days",
					[]string{"Required. A number of days to keep backups for"},
				},
				{
					"--status",
					[]string{"Required. The policy status"},
				},
			},
		},
	})
	registerCommandBase(nil, &backup.DataCenters{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/datacenters",
		Resource: "backup",
		Command:  "get-data-centers",
		Help: help.Command{
			Brief:     []string{"Show the list of the supported backup data centers"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(nil, &backup.OSTypes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/osTypes",
		Resource: "backup",
		Command:  "get-os-types",
		Help: help.Command{
			Brief:     []string{"Show the list of the supported OS types for backup"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(nil, &[]backup.Region{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/regions",
		Resource: "backup",
		Command:  "get-regions",
		Help: help.Command{
			Brief:     []string{"Show the list of the supported backup regions"},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(&backup.GetServerPolicies{}, &[]backup.ServerPolicyDetails{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/serverPolicyDetails?serverId={ServerId}&withStatus={WithStatus}",
		Resource: "backup",
		Command:  "get-applied-server-policies",
		Help: help.Command{
			Brief: []string{"Shows the backup policies applied to the given server"},
			Arguments: []help.Argument{
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. An ID of a server to query"},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. A name of a server to query"},
				},
				{
					"--with-status",
					[]string{"A policy status to filter by"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetAccountServerPolicies{}, &backup.AccountServerPoliciesRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{PolicyId}/serverPolicies?limit={Limit}&offset={Offset}&withStatus={WithStatus}&sortBy={SortBy}&ascendingSort={AscendingSort}",
		Resource: "backup",
		Command:  "get-applied-account-policies",
		Help: help.Command{
			Brief: []string{"Shows the list of applied backup policies in the given account"},
			Arguments: []help.Argument{
				{
					"--policy-id",
					[]string{"Required. An ID of an account policy to query"},
				},
				{
					"--limit",
					[]string{"A number of results to show"},
				},
				{
					"--offset",
					[]string{"The number of a result to start from"},
				},
				{
					"--with-status",
					[]string{"A policy status to filter by"},
				},
				{
					"--sort-by",
					[]string{
						"A field to sort the results by.",
						"'status' or 'storageRegion' or 'serverId' or 'serverPolicyId'",
					},
				},
				{
					"--ascending-sort",
					[]string{"Defines in which order to sort the results. 'true' or 'false'"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetAccountServerPolicy{}, &backup.ServerPolicy{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies/{ServerPolicyId}",
		Resource: "backup",
		Command:  "get-applied-policy",
		Help: help.Command{
			Brief: []string{"Shows the applied backup policy to the server in the given account"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of an account policy to query"},
				},
				{
					"--server-policy-id",
					[]string{"Required. The ID of a server policy to query"},
				},
			},
		},
	})
	registerCommandBase(&backup.CreateServerPolicy{}, &backup.ServerPolicy{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies",
		Resource: "backup",
		Command:  "apply-policy",
		Help: help.Command{
			Brief: []string{"Applies a backup policy to the given server"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of an account policy to bind the server policy to"},
				},
				{
					"--server-id",
					[]string{"Required unless --server-name is specified. The ID of a server to backup"},
				},
				{
					"--server-name",
					[]string{"Required unless --server-id is specified. The name of a server to backup"},
				},
				{
					"--storage-region",
					[]string{"Required. A region to store backups in"},
				},
				{
					"--storage-account-id",
					[]string{"The ID of a storage account"},
				},
			},
		},
	})
	registerCommandBase(&backup.DeleteServerPolicy{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies/{ServerPolicyId}",
		Resource: "backup",
		Command:  "unapply-policy",
		Help: help.Command{
			Brief: []string{"Unapplies the given backup policy from the given server"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of the account policy to unbind"},
				},
				{
					"--server-policy-id",
					[]string{"Required. The ID of the server policy to remove"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetStoredData{}, &backup.StoredData{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies/{ServerPolicyId}/storedData?searchDate={SearchDate}",
		Resource: "backup",
		Command:  "get-stored-data",
		Help: help.Command{
			Brief: []string{"Shows the size of the backup data"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of the account policy to query"},
				},
				{
					"--server-policy-id",
					[]string{"Required. The ID of the server policy to query"},
				},
				{
					"--search-date",
					[]string{"Required. A date to show the backup state at"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetRestoreDetails{}, &backup.GetRestoreDetailsRes{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies/{ServerPolicyId}/restorePointDetails?backupFinishedStartDate={BackupFinishedStartDate}&backupFinishedEndDate={BackupFinishedEndDate}&limit={Limit}&offset={Offset}&inRetentionOnly={InRetentionOnly}&sortBy={SortBy}&ascendingSort={AscendingSort}",
		Resource: "backup",
		Command:  "get-restore-point-details",
		Help: help.Command{
			Brief: []string{"Shows some details about the restore points"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of the account policy to query"},
				},
				{
					"--server-policy-id",
					[]string{"Required. The ID of the server policy to query"},
				},
				{
					"--backup-finished-start-date",
					[]string{"Required. A start date to query the finished backups from."},
				},
				{
					"--backup-finished-end-date",
					[]string{"Required. An end date to query the finished backups at."},
				},
				{
					"--limit",
					[]string{"A number of results to show"},
				},
				{
					"--offset",
					[]string{"The number of a result to show results from"},
				},
				{
					"--in-retention-only",
					[]string{"'true' or 'false'. A retention filter"},
				},
				{
					"--sort-by",
					[]string{
						"A field to sort the results by.",
						"'policyId' or 'retentionDay' or 'backupStartedDate' or 'backupFinishedDate' or 'retentionExpiredDate' or",
						"'backupStatus' or 'filesTransferredToStorage' or 'bytesTransferredToStorage' or 'filesFailedTransferToStorage' or",
						"'bytesFailedToTransfer' or 'unchangedFilesNotTransferred' or 'unchangedBytesNotTransferred' or",
						"'filesRemovedFromDisk' or 'bytesRemovedFromDisk'",
					},
				},
				{
					"--ascending-sort",
					[]string{"'true' or 'false'. Defines a sort order"},
				},
			},
		},
	})
	registerCommandBase(&backup.UpdateServerPolicy{}, &backup.ServerPolicy{}, commands.CommandExcInfo{
		Verb:     "PATCH",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/accountPolicies/{AccountPolicyId}/serverPolicies/{ServerPolicyId}",
		Resource: "backup",
		Command:  "update-server-policy",
		Help: help.Command{
			Brief: []string{"Updates the given backup policy of the given server"},
			Arguments: []help.Argument{
				{
					"--account-policy-id",
					[]string{"Required. The ID of the account policy the server policy is bound to"},
				},
				{
					"--server-policy-id",
					[]string{"Required. The ID of the server policy to update"},
				},
				{
					"--status",
					[]string{"Required. A new status value. 'ACTIVE' or 'INACTIVE'"},
				},
			},
		},
	})
	registerCommandBase(&backup.GetServers{}, &backup.DataCenterServers{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api-va1.backup.ctl.io/clc-backup-api/api/datacenters/{DataCenterName}/servers",
		Resource: "backup",
		Command:  "get-servers",
		Help: help.Command{
			Brief: []string{"Shows the servers in the given data center"},
			Arguments: []help.Argument{
				{
					"--data-center-name",
					[]string{"Required. The name of a data center to query. For example, `UC1 - US West (Santa Clara)`"},
				},
			},
		},
	})

	registerCommandBase(&vpn.CreateReq{}, &vpn.Entity{}, commands.CommandExcInfo{
		Verb:     "POST",
		Url:      "https://api.ctl.io/v2/siteToSiteVpn?account={accountAlias}",
		Resource: "site-to-site-vpn",
		Command:  "create",
		Help: help.Command{
			Brief: []string{"Creates a Site to Site VPN for a given account."},
			Arguments: []help.Argument{
				{
					"--local",
					[]string{
						"Required. Local site properties",
						"",
						"alias – short code for a particular location (required)",
						"subnets - local address for Site to Site VPN, specified using CIDR notation (required)",
					},
				},
				{
					"--remote",
					[]string{
						"Required. Remote site properties",
						"",
						"siteName   - friendly name of the site (required)",
						"deviceType - friendly name of the device type (required)",
						"address    - remote address for Site to Site VPN, specified using CIDR notation (required)",
						"subnets    - remote network address for Site to Site VPN, specified using CIDR notation (required)",
					},
				},
				{
					"--ipsec",
					[]string{
						"Required. IPSec properties",
						"",
						"encryption - encryption algorithm.",
						"             Valid values are: 'aes128', 'aes192', 'aes256', 'tripleDES' (required)",
						"hashing    - hashing algorithm.",
						"             Valid values are: 'sha1_96', 'sha1_256', 'md5' (required)",
						"protocol   - IPSec protocol.",
						"             Valid values are: 'esp', 'ah' (required)",
						"pfs        - PFS enabled or disabled",
						"             (we suggest enabled, using Group 2, though Group 5 is ",
						"             recommended with SHA2 hashing or AES-192 or AES-256).",
						"             Valid values are: 'disabled', 'group1', 'group2', 'group5' (required)",
						"lifetime   - Lifetime is set to 1 hour (and unlimited KB).",
						"             This setting is not required to match, as the negotiation ",
						"             process will choose the shortest value supplied by either peer.",
						"             Valid values are: 3600, 28800, 86400 (required)",
					},
				},
				{
					"--ike",
					[]string{
						"Required. IKE properties",
						"",
						"encryption        - encryption algorithm.",
						"                    Valid values are: 'aes128', 'aes192', 'aes256', 'tripleDES' (required)",
						"hashing           - hashing algorithm.",
						"                    Valid values are: 'sha1_96', 'sha1_256', 'md5' (required)",
						"diffieHelmanGroup - Group 1 (legacy), Group 2 or Group 5.",
						"                    If using AES with a cipher strength greater than 128-bit,",
						"                    or SHA2 for hashing, we recommend Group 5, otherwise Group 2 is sufficient",
						"                    Valid values are: 'group1', 'group2', 'group5' (required)",
						"preSharedKey      - The pre-shared key is a shared secret that secures the VPN tunnel.",
						"                    This value must be identical on both ends of the connection (required)",
						"lifetime          - Lifetime is set to 1 hour (and unlimited KB).",
						"                    This setting is not required to match, as the negotiation ",
						"                    process will choose the shortest value supplied by either peer.",
						"                    Valid values are: 3600, 28800, 86400 (required)",
						"mode              - protocol mode.",
						"                    Valid values are: 'main', 'aggressive' (required)",
						"deadPeerDetection - specify if you wish this enabled or disabled. ",
						"                    Check your device defaults; for example, Cisco ASA defaults to 'on',",
						"                    while Netscreen/Juniper SSG or Juniper SRX default to 'off'. Our default is 'off'.",
						"                    Valid values are: true, false",
						"natTraversal      - NAT-Traversal: Allows connections to VPN end-points behind a NAT device.",
						"                    Defaults to 'off'. If you require NAT-T, you also need to provide the",
						"                    private IP address that your VPN endpoint will use to identify itself.",
						"                    Valid values are: true, false",
						"remoteIdentity    - The private IP address that your VPN endpoint will use to identify itself.",
						"                    Required only when NAT-T state is on",
					},
				},
			},
		},
	})
	registerCommandBase(&vpn.UpdateReq{}, &vpn.Entity{}, commands.CommandExcInfo{
		Verb:     "PUT",
		Url:      "https://api.ctl.io/v2/siteToSiteVpn/{VpnId}?account={accountAlias}",
		Resource: "site-to-site-vpn",
		Command:  "update",
		Help: help.Command{
			Brief: []string{"Updates a Site to Site VPN for a given account."},
			Arguments: []help.Argument{
				{
					"--vpn-id",
					[]string{"Required. The ID of the VPN to update"},
				},
				{
					"--local",
					[]string{
						"Local site properties",
						"",
						"subnets - local address for Site to Site VPN, specified using CIDR notation",
					},
				},
				{
					"--remote",
					[]string{
						"Remote site properties",
						"",
						"siteName   - friendly name of the site (required)",
						"deviceType - friendly name of the device type (required)",
						"address    - remote address for Site to Site VPN, specified using CIDR notation",
						"subnets    - remote network address for Site to Site VPN, specified using CIDR notation",
					},
				},
				{
					"--ipsec",
					[]string{
						"IPSec properties",
						"",
						"encryption - encryption algorithm.",
						"             Valid values are: 'aes128', 'aes192', 'aes256', 'tripleDES'",
						"hashing    - hashing algorithm.",
						"             Valid values are: 'sha1_96', 'sha1_256', 'md5'",
						"protocol   - IPSec protocol.",
						"             Valid values are: 'esp', 'ah'",
						"pfs        - PFS enabled or disabled",
						"             (we suggest enabled, using Group 2, though Group 5 is ",
						"             recommended with SHA2 hashing or AES-192 or AES-256).",
						"             Valid values are: 'disabled', 'group1', 'group2', 'group5'",
						"lifetime   - Lifetime is set to 1 hour (and unlimited KB).",
						"             This setting is not required to match, as the negotiation ",
						"             process will choose the shortest value supplied by either peer.",
						"             Valid values are: 3600, 28800, 86400",
					},
				},
				{
					"--ike",
					[]string{
						"Required. IKE properties",
						"",
						"encryption        - encryption algorithm.",
						"                    Valid values are: 'aes128', 'aes192', 'aes256', 'tripleDES'",
						"hashing           - hashing algorithm.",
						"                    Valid values are: 'sha1_96', 'sha1_256', 'md5'",
						"diffieHelmanGroup - Group 1 (legacy), Group 2 or Group 5.",
						"                    If using AES with a cipher strength greater than 128-bit,",
						"                    or SHA2 for hashing, we recommend Group 5, otherwise Group 2 is sufficient",
						"                    Valid values are: 'group1', 'group2', 'group5'",
						"preSharedKey      - The pre-shared key is a shared secret that secures the VPN tunnel.",
						"                    This value must be identical on both ends of the connection",
						"lifetime          - Lifetime is set to 1 hour (and unlimited KB).",
						"                    This setting is not required to match, as the negotiation ",
						"                    process will choose the shortest value supplied by either peer.",
						"                    Valid values are: 3600, 28800, 86400",
						"mode              - protocol mode.",
						"                    Valid values are: 'main', 'aggressive'",
						"deadPeerDetection - specify if you wish this enabled or disabled. ",
						"                    Check your device defaults; for example, Cisco ASA defaults to 'on',",
						"                    while Netscreen/Juniper SSG or Juniper SRX default to 'off'. Our default is 'off'.",
						"                    Valid values are: true, false",
						"natTraversal      - NAT-Traversal: Allows connections to VPN end-points behind a NAT device.",
						"                    Defaults to 'off'. If you require NAT-T, you also need to provide the",
						"                    private IP address that your VPN endpoint will use to identify itself.",
						"                    Valid values are: true, false",
						"remoteIdentity    - The private IP address that your VPN endpoint will use to identify itself.",
						"                    Required only when NAT-T state is on",
					},
				},
			},
		},
	})
	registerCommandBase(&vpn.ListReq{}, &[]vpn.Entity{}, commands.CommandExcInfo{
		Verb:     "GET",
		Url:      "https://api.ctl.io/v2/siteToSiteVpn?account={accountAlias}",
		Resource: "site-to-site-vpn",
		Command:  "list",
		Help: help.Command{
			Brief:     []string{"Gets all Site to Site VPNs for a given account."},
			Arguments: []help.Argument{},
		},
	})
	registerCommandBase(&vpn.DeleteReq{}, new(string), commands.CommandExcInfo{
		Verb:     "DELETE",
		Url:      "https://api.ctl.io/v2/siteToSiteVpn/{VpnId}?account={accountAlias}",
		Resource: "site-to-site-vpn",
		Command:  "delete",
		Help: help.Command{
			Brief: []string{"Deletes a Site to Site VPN for a given account."},
			Arguments: []help.Argument{
				{
					"--vpn-id",
					[]string{"Required. ID of the VPN."},
				},
			},
		},
	})
}

func registerCommandBase(inputModel interface{}, outputModel interface{}, info commands.CommandExcInfo) {
	cmd := &commands.CommandBase{
		Input:   inputModel,
		Output:  outputModel,
		ExcInfo: info,
	}
	AllCommands = append(AllCommands, cmd)
}

func registerCustomCommand(cmd base.Command) {
	AllCommands = append(AllCommands, cmd)
}
