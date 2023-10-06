package main

import (
	"fmt"

	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/cloudrun"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/compute"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/organizations"
	"github.com/pulumi/pulumi-gcp/sdk/v6/go/gcp/serviceaccount"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	location := "us-west1"
	rootDir := "/home/yamiro.dev"

	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		machineType, err := cfg.Try("machineType")
		if err != nil {
			machineType = "e2-standard-2"
		}

		osImage, err := cfg.Try("osImage")
		if err != nil {
			osImage = "debian-11"
		}

		instanceTag, err := cfg.Try("instanceTag")
		if err != nil {
			instanceTag = "webserver"
		}

		servicePort, err := cfg.Try("servicePort")
		if err != nil {
			servicePort = "8080"
		}

		network, err := compute.NewNetwork(ctx, "network", &compute.NetworkArgs{
			AutoCreateSubnetworks: pulumi.Bool(false),
		})
		if err != nil {
			return err
		}

		subnet, err := compute.NewSubnetwork(ctx, "subnet", &compute.SubnetworkArgs{
			IpCidrRange: pulumi.String("10.0.1.0/24"),
			Network:     network.ID(),
		})
		if err != nil {
			return err
		}

		firewall, err := compute.NewFirewall(ctx, "firewall", &compute.FirewallArgs{
			Network: network.SelfLink,
			Allows: compute.FirewallAllowArray{
				compute.FirewallAllowArgs{
					Protocol: pulumi.String("tcp"),
					Ports: pulumi.ToStringArray([]string{
						"22",
						servicePort,
					}),
				},
			},
			Direction: pulumi.String("INGRESS"),
			SourceRanges: pulumi.ToStringArray([]string{
				"0.0.0.0/0",
			}),
			TargetTags: pulumi.ToStringArray([]string{
				instanceTag,
			}),
		})
		if err != nil {
			return err
		}

		serviceAccount, err := serviceaccount.NewAccount(ctx, "serviceAccount", &serviceaccount.AccountArgs{
			AccountId:   pulumi.String("service-account"),
			DisplayName: pulumi.String("Service Account"),
		})
		if err != nil {
			return err
		}

		metadataStartupScript := fmt.Sprintf(`#!/bin/bash
			mkdir -p %s/server

			curl -o %s/server/server.jar https://api.papermc.io/v2/projects/paper/versions/1.20.2/builds/217/downloads/paper-1.20.2-217.jar

			echo 'eula=true' > %s/server/eula.txt

			echo '#!/bin/bash
			java -Xmx1024M -Xms1024M -jar server.jar' > %s/server/run.sh

			sudo chmod +x %s/server/run.sh

			echo '[Unit]
			Description=Minecraft Server
			After=network.target

			[Service]
			WorkingDirectory=%s/server
			ExecStart=run.sh
			ExecStop=stop.sh
			Restart=on-failure
			User=minecraft
			Group=minecraft
			TimeoutStartSec=120
			TimeoutStopSec=60

			[Install]
			WantedBy=multi-user.target
			' > /etc/systemd/system/minecraft.service`, rootDir, rootDir, rootDir, rootDir, rootDir, rootDir)

		vm, err := compute.NewInstance(ctx, "minecraft-vm", &compute.InstanceArgs{
			Name:        pulumi.String("minecraft-vm"),
			MachineType: pulumi.String(machineType),
			BootDisk: compute.InstanceBootDiskArgs{
				InitializeParams: compute.InstanceBootDiskInitializeParamsArgs{
					Image: pulumi.String(osImage),
				},
			},
			NetworkInterfaces: compute.InstanceNetworkInterfaceArray{
				compute.InstanceNetworkInterfaceArgs{
					Network:    network.ID(),
					Subnetwork: subnet.ID(),
					AccessConfigs: compute.InstanceNetworkInterfaceAccessConfigArray{
						compute.InstanceNetworkInterfaceAccessConfigArgs{
							// NatIp:       nil,
							// NetworkTier: nil,
						},
					},
				},
			},
			ServiceAccount: compute.InstanceServiceAccountArgs{
				Email: serviceAccount.Email,
				Scopes: pulumi.ToStringArray([]string{
					"https://www.googleapis.com/auth/cloud-platform",
				}),
			},
			AllowStoppingForUpdate: pulumi.Bool(true),
			MetadataStartupScript:  pulumi.String(metadataStartupScript),
			Tags: pulumi.ToStringArray([]string{
				instanceTag,
			}),
		}, pulumi.DependsOn([]pulumi.Resource{firewall}))
		if err != nil {
			return err
		}

		instanceIp := vm.NetworkInterfaces.Index(pulumi.Int(0)).AccessConfigs().Index(pulumi.Int(0)).NatIp()

		cloudrunService, err := cloudrun.NewService(ctx, "service", &cloudrun.ServiceArgs{
			Name:     pulumi.String("webserver"),
			Location: pulumi.String(location),
			Template: &cloudrun.ServiceTemplateArgs{
				Spec: &cloudrun.ServiceTemplateSpecArgs{
					Containers: cloudrun.ServiceTemplateSpecContainerArray{
						&cloudrun.ServiceTemplateSpecContainerArgs{
							Image: pulumi.String("docker.io/vkzrx/mined:main"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}

		noauthIAMPolicy, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
			Bindings: []organizations.GetIAMPolicyBinding{
				{
					Role: "roles/run.invoker",
					Members: []string{
						"allUsers",
					},
				},
			},
		}, nil)
		if err != nil {
			return err
		}

		_, err = cloudrun.NewIamPolicy(ctx, "noauthIamPolicy", &cloudrun.IamPolicyArgs{
			Project:    cloudrunService.Project,
			Location:   cloudrunService.Location,
			Service:    cloudrunService.Name,
			PolicyData: pulumi.String(noauthIAMPolicy.PolicyData),
		})
		if err != nil {
			return err
		}

		ctx.Export("name", vm.Name)
		ctx.Export("ip", instanceIp)
		ctx.Export("url", pulumi.Sprintf("http://%s:%s", instanceIp.Elem(), servicePort))

		return nil
	})
}
