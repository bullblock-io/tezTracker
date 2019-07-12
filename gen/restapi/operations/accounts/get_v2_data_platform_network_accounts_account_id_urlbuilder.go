// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// GetV2DataPlatformNetworkAccountsAccountIDURL generates an URL for the get v2 data platform network accounts account ID operation
type GetV2DataPlatformNetworkAccountsAccountIDURL struct {
	AccountID string
	Network   string
	Platform  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) WithBasePath(bp string) *GetV2DataPlatformNetworkAccountsAccountIDURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/data/{platform}/{network}/accounts/{accountId}"

	accountID := o.AccountID
	if accountID != "" {
		_path = strings.Replace(_path, "{accountId}", accountID, -1)
	} else {
		return nil, errors.New("accountId is required on GetV2DataPlatformNetworkAccountsAccountIDURL")
	}

	network := o.Network
	if network != "" {
		_path = strings.Replace(_path, "{network}", network, -1)
	} else {
		return nil, errors.New("network is required on GetV2DataPlatformNetworkAccountsAccountIDURL")
	}

	platform := o.Platform
	if platform != "" {
		_path = strings.Replace(_path, "{platform}", platform, -1)
	} else {
		return nil, errors.New("platform is required on GetV2DataPlatformNetworkAccountsAccountIDURL")
	}

	_basePath := o._basePath
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetV2DataPlatformNetworkAccountsAccountIDURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetV2DataPlatformNetworkAccountsAccountIDURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetV2DataPlatformNetworkAccountsAccountIDURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
