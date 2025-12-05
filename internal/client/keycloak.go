package client

import (
	"context"
	"log"

	"backend/config"

	"github.com/Nerzal/gocloak/v13"
)

func NewKeycloak(ctx context.Context, srvCfg *config.Server) *gocloak.GoCloak {
	client := gocloak.NewClient(srvCfg.KCBasePath)
	token, err := client.LoginClient(ctx, srvCfg.KCClientId, srvCfg.KCClientSecret, srvCfg.KCRealm)
	if err != nil || token == nil {
		log.Printf("error when connecting to keycloak:%s", err)
		return nil
	}
	return client
}
