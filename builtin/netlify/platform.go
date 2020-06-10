package netlify

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/go-hclog"
	"github.com/netlify/open-api/go/models"
	"github.com/netlify/open-api/go/plumbing/operations"
	netlify "github.com/netlify/open-api/go/porcelain"

	"github.com/hashicorp/waypoint/builtin/files"
	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/datadir"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

// Platform is the Platform implementation for Google Cloud Run.
type Platform struct {
	config Config
}

// Config implements Configurable
func (p *Platform) Config() (interface{}, error) {
	return &p.config, nil
}

// DeployFunc implements component.Platform
func (p *Platform) DeployFunc() interface{} {
	return p.Deploy
}

// Deploy deploys a set of files to netlify
func (p *Platform) Deploy(
	ctx context.Context,
	log hclog.Logger,
	src *component.Source,
	files *files.Files,
	dir *datadir.Component,
	deployConfig *component.DeploymentConfig,
	ui terminal.UI,
) (*Deployment, error) {
	deployment := &Deployment{}
	client := netlify.Default
	clientContext := apiContext("")

	// We'll update the user in realtime
	st := ui.Status()
	defer st.Close()

	// Use configured token, otherwise retrieve one with the user
	token := p.config.AccessToken
	if token == "" {
		st.Update("Logging into your Netlify account")
		auth, err := Authenticate(clientContext, log)
		token = auth

		if err != nil {
			return nil, err
		}
	}

	// Setup a new authenticated context
	clientContext = apiContext(token)
	log.Trace("have token", "token id", token)

	st.Update("Setting up deploy")

	// Default siteID to the app name, unless provided
	siteName := src.App
	if p.config.SiteName != "" {
		siteName = p.config.SiteName
	}

	siteSetup := &models.SiteSetup{
		Site: *&models.Site{
			Name: siteName,
		},
	}

	listParams := operations.NewListSitesParams()
	listParams.Name = &siteName
	sites, err := client.ListSites(clientContext, listParams)
	if err != nil {
		return nil, err
	}

	site := &models.Site{}

	switch len(sites) {
	case 0:
		log.Trace("site does not exist, creating site", "site name", siteName)
		st.Update("Creating site")
		site, err := client.CreateSite(clientContext, siteSetup, false)
		if err != nil {
			return nil, err
		}

		_ = site
	case 1:
		site = sites[0]
		if site.Name != siteName {
			return nil, fmt.Errorf("site returned does not match")
		}
		log.Trace("found site", "site id", site.ID)
	default:
		return nil, fmt.Errorf("incorrect sites returned from netlify")
	}

	deployment.SiteId = site.ID
	deployOptions := netlify.DeployOptions{
		IsDraft: true,
		Dir:     files.Path,
		SiteID:  site.ID,
	}

	log.Trace("deploying site", "site id", site.ID)
	st.Update("Deploying site")
	deploy, err := client.DeploySite(clientContext, deployOptions)
	if err != nil {
		return nil, err
	}

	log.Trace("waiting for deploying to become ready", "site id", site.ID)
	st.Update("Waiting for deploy to be ready")
	deploy, err = client.WaitUntilDeployReady(clientContext, deploy, 10*time.Minute)
	if err != nil {
		return nil, err
	}

	deployment.Url = deploy.DeploySslURL
	log.Trace("url available", "url", deploy.DeploySslURL)

	return deployment, nil
}

// Config is the configuration structure for the Platform.
type Config struct {
	// SiteID is the site to deploy to
	SiteID string `hcl:"site_id,optional"`
	// SiteName is the name of the site we create. Defaults
	// to the application.
	SiteName string `hcl:"site_name,optional"`
	// AccessToken is the access token to use, will
	// prompt oauth exchange if not specified
	AccessToken string `hcl:"access_token,optional"`
}

var (
	_ component.Platform     = (*Platform)(nil)
	_ component.Configurable = (*Platform)(nil)
)
