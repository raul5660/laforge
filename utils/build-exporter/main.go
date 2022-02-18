package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/gen0cide/laforge/ent"
	"github.com/sirupsen/logrus"
)

type buildConf struct {
	EnvironmentName string     `json:"environment_name"`
	RevisionNumber  int        `json:"revision_number"`
	Teams           []teamConf `json:"teams"`
}
type teamConf struct {
	TeamNumber int           `json:"team_number"`
	Networks   []networkConf `json:"networks"`
}
type networkConf struct {
	NetworkName string     `json:"network_name"`
	CIDR        string     `json:"cidr"`
	VDIVisible  bool       `json:"vdi_visible"`
	Hosts       []hostConf `json:"hosts"`
}
type hostConf struct {
	HostName         string   `json:"host_name"`
	SubnetIP         string   `json:"subnet_ip"`
	OS               string   `json:"os"`
	Instance_Size    string   `json:"instance_size"`
	AllowMacChanges  bool     `json:"allow_mac_changes"`
	ExposedTCPPorts  []string `json:"exposed_tcp_ports"`
	ExposedUDPPorts  []string `json:"exposed_udp_ports"`
	OverridePassword string   `json:"override_password"`
	DiskSize         int      `json:"disk_size"`
	AgentURL         string   `json:"agent_url"`
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	pgHost, ok := os.LookupEnv("PG_URI")
	client := &ent.Client{}

	if !ok {
		client = ent.PGOpen("postgresql://laforger:laforge@127.0.0.1/laforge")
	} else {
		client = ent.PGOpen(pgHost)
	}

	ctx := context.Background()
	defer ctx.Done()
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	entBuild, err := client.Build.Query().First(ctx)
	if err != nil {
		log.Fatalf("failed to get a Build: %v", err)
	}
	buildJson, err := GenerateBuildConf(ctx, client, entBuild)
	if err != nil {
		log.Fatalf("failed to generate build conf: %v", err)
	}
	fmt.Printf("%+v", buildJson)
}

func GenerateBuildConf(ctx context.Context, client *ent.Client, entBuild *ent.Build) (string, error) {

	entEnvrioment, err := entBuild.QueryBuildToEnvironment().Only(ctx)
	if err != nil {
		return "", err
	}
	laforgeServerUrl, exists := entEnvrioment.Config["laforge_server_url"]
	if !exists {
		return "", errors.New("laforge_server_url doesn't exist in the environment configuration")
	}
	currentBuildConf := buildConf{
		EnvironmentName: entEnvrioment.Name,
		RevisionNumber:  entBuild.Revision,
		Teams:           []teamConf{},
	}
	entTeams, err := entBuild.QueryBuildToTeam().All(ctx)
	if err != nil {
		return "", err
	}
	for _, entTeam := range entTeams {
		currentTeamConf := teamConf{
			TeamNumber: entTeam.TeamNumber,
			Networks:   []networkConf{},
		}
		entProNetworks, err := entTeam.QueryTeamToProvisionedNetwork().All(ctx)
		if err != nil {
			return "", err
		}
		for _, entProNetwork := range entProNetworks {
			entNetwork, err := entProNetwork.QueryProvisionedNetworkToNetwork().Only(ctx)
			if err != nil {
				return "", err
			}
			currentNetworkConf := networkConf{
				NetworkName: entNetwork.Name,
				CIDR:        entProNetwork.Cidr,
				VDIVisible:  entNetwork.VdiVisible,
				Hosts:       []hostConf{},
			}
			entProHosts, err := entProNetwork.QueryProvisionedNetworkToProvisionedHost().All(ctx)
			if err != nil {
				return "", err
			}
			for _, entProHost := range entProHosts {
				entHost, err := entProHost.QueryProvisionedHostToHost().Only(ctx)
				if err != nil {
					return "", err
				}
				entDisk, err := entHost.QueryHostToDisk().Only(ctx)
				if err != nil {
					return "", err
				}
				entAgent, err := entProHost.QueryProvisionedHostToGinFileMiddleware().First(ctx)
				if err != nil {
					return "", err
				}

				agentUrl := fmt.Sprintf("%s/api/download/%s", laforgeServerUrl, entAgent.URLID)
				currentHostConf := hostConf{
					HostName:         entHost.Hostname,
					SubnetIP:         entProHost.SubnetIP,
					OS:               entHost.OS,
					Instance_Size:    entHost.InstanceSize,
					AllowMacChanges:  entHost.AllowMACChanges,
					ExposedTCPPorts:  entHost.ExposedTCPPorts,
					ExposedUDPPorts:  entHost.ExposedUDPPorts,
					OverridePassword: entHost.OverridePassword,
					DiskSize:         entDisk.Size,
					AgentURL:         agentUrl,
				}
				currentNetworkConf.Hosts = append(currentNetworkConf.Hosts, currentHostConf)
			}
			currentTeamConf.Networks = append(currentTeamConf.Networks, currentNetworkConf)

		}
		currentBuildConf.Teams = append(currentBuildConf.Teams, currentTeamConf)
	}
	jsonByteArray, err := json.MarshalIndent(currentBuildConf, "", "  ")
	if err != nil {
		return "", err
	}
	jsonString := string(jsonByteArray)
	return jsonString, nil

}
