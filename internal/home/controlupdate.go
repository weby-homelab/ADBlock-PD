package home

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/weby-homelab/adblock-pd/internal/aghalg"
	"github.com/weby-homelab/adblock-pd/internal/aghhttp"
	"github.com/weby-homelab/adblock-pd/internal/aghnet"
)



// handleVersionJSON is the handler for the POST /control/version.json HTTP API.
//
// TODO(a.garipov): Find out if this API used with a GET method by anyone.
func (web *webAPI) handleVersionJSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := web.logger

	resp := &versionResponse{}
	if web.conf.disableUpdate {
		resp.Disabled = true
		aghhttp.WriteJSONResponseOK(ctx, l, w, r, resp)

		return
	}

	req := &struct {
		Recheck bool `json:"recheck_now"`
	}{}

	var err error
	if r.ContentLength != 0 {
		err = json.NewDecoder(r.Body).Decode(req)
		if err != nil {
			aghhttp.ErrorAndLog(ctx, l, r, w, http.StatusBadRequest, "parsing request: %s", err)

			return
		}
	}

	err = web.requestVersionInfo(ctx, resp, req.Recheck)
	if err != nil {
		// Don't wrap the error, because it's informative enough as is.
		aghhttp.ErrorAndLog(ctx, l, r, w, http.StatusBadGateway, "%s", err)

		return
	}

	err = resp.setAllowedToAutoUpdate(ctx, l, web.tlsManager)
	if err != nil {
		// Don't wrap the error, because it's informative enough as is.
		aghhttp.ErrorAndLog(ctx, l, r, w, http.StatusInternalServerError, "%s", err)

		return
	}

	aghhttp.WriteJSONResponseOK(ctx, l, w, r, resp)
}

// requestVersionInfo sets the VersionInfo field of resp if it can reach the
// update server.
func (web *webAPI) requestVersionInfo(
	_ context.Context,
	resp *versionResponse,
	_ bool,
) (err error) {
	resp.VersionInfo = VersionInfo{
		CanAutoUpdate: aghalg.NBFalse,
	}
	return nil
}

// handleUpdate performs an update to the latest available version procedure.
func (web *webAPI) handleUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	l := web.logger
	aghhttp.ErrorAndLog(
		ctx,
		l,
		r,
		w,
		http.StatusForbidden,
		"updates are disabled in this version",
	)
}

// versionResponse is the response for /control/version.json endpoint.
type VersionInfo struct {
	NewVersion      string          `json:"new_version,omitempty"`
	Announcement    string          `json:"announcement,omitempty"`
	AnnouncementURL string          `json:"announcement_url,omitempty"`
	CanAutoUpdate   aghalg.NullBool `json:"can_autoupdate,omitempty"`
}

type versionResponse struct {
	VersionInfo
	Disabled bool `json:"disabled"`
}

// maxPrivilegedPort is the maximum port number.  This only applies to Unix, as
// on Windows, [aghnet.CanBindPrivilegedPorts] always returns `true`, `nil`.
const maxPrivilegedPort = 1024

// setAllowedToAutoUpdate sets CanAutoUpdate to true if ADBlock-Private-DNS is actually
// allowed to perform an automatic update by the OS.  l and tlsMgr must not be
// nil.
func (vr *versionResponse) setAllowedToAutoUpdate(
	ctx context.Context,
	l *slog.Logger,
	tlsMgr *tlsManager,
) (err error) {
	if vr.CanAutoUpdate != aghalg.NBTrue {
		return nil
	}

	canUpdate := true
	if tlsConfUsesPrivilegedPorts(tlsMgr.extendedTLSConfig()) ||
		config.HTTPConfig.Address.Port() < maxPrivilegedPort ||
		config.DNS.Port < maxPrivilegedPort {
		canUpdate, err = aghnet.CanBindPrivilegedPorts(ctx, l)
		if err != nil {
			return fmt.Errorf("checking ability to bind privileged ports: %w", err)
		}
	}

	vr.CanAutoUpdate = aghalg.BoolToNullBool(canUpdate)

	return nil
}

// tlsConfUsesPrivilegedPorts returns true if the provided TLS configuration
// indicates that privileged ports are used.  c must be valid
func tlsConfUsesPrivilegedPorts(c *tlsConfigSettings) (ok bool) {
	return c.Enabled && (c.PortHTTPS < maxPrivilegedPort ||
		c.PortDNSOverTLS < maxPrivilegedPort ||
		c.PortDNSOverQUIC < maxPrivilegedPort)
}


