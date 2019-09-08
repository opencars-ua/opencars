package gov

import (
	"encoding/json"
)

// Result is high-level representation for response structure.
type Result struct {
	Help    string          `json:"help"`
	Success bool            `json:"success"`
	Result  json.RawMessage `json:"result"`
}

// Resource is a data attached to package.
type Resource struct {
	PackageID       string `json:"package_id"`
	DataStoreActive bool   `json:"datastore_active"`
	ID              string `json:"id"`
	Size            int    `json:"size"`
	FileHashSum     string `json:"file_hash_sum"`
	State           string `json:"state"`
	Hash            string `json:"hash"`
	Description     string `json:"description"`
	Format          string `json:"format"`
	LastModified    Time   `json:"last_modified"`
	URLType         string `json:"url_type"`
	MIMEType        string `json:"mimetype"`
	Name            string `json:"name"`
	Created         string `json:"created"`
	URL             string `json:"url"`
	Position        int    `json:"position"`
	RevisionID      string `json:"revision_id"`
}

// Tag represents some short names applied to the package.
type Tag struct {
	State       string `json:"state"`
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// Group is a sphere associated with package.
type Group struct {
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	ImageDisplayURL string `json:"image_display_url"`
	Title           string `json:"title"`
	ID              string `json:"id"`
	Name            string `json:"name"`
}

// Package is a representation of data package on data.gov.ua portal.
type Package struct {
	LicenseTitle                   string     `json:"license_title"`
	Maintainer                     string     `json:"maintainer"`
	TagString                      string     `json:"tag_string"`
	PurposeOfCollectingInformation string     `json:"purpose_of_collecting_information"`
	Private                        bool       `json:"private"`
	MaintainerEmail                string     `json:"maintainer_email"`
	NumTags                        int        `json:"num_tags"`
	UpdateFrequency                string     `json:"update_frequency"`
	ID                             string     `json:"id"`
	MetadataCreated                string     `json:"metadata_created"`
	MetadataModified               string     `json:"metadata_modified"`
	Author                         string     `json:"author"`
	AuthorEmail                    string     `json:"author_email"`
	State                          string     `json:"state"`
	Version                        string     `json:"version"`
	IsDataPackage                  string     `json:"is_datapackage"`
	CreatorUserID                  string     `json:"creator_user_id"`
	Type                           string     `json:"type"`
	Resources                      []Resource `json:"resources"`
	NumResources                   int        `json:"num_resources"`
	Tags                           []Tag      `json:"tags"`
	Language                       string     `json:"language"`
	Groups                         []Group    `json:"groups"`
	LicenseID                      string     `json:"license_id"`
	Name                           string     `json:"name"`
	IsOpen                         bool       `json:"isopen"`
	Notes                          string     `json:"notes"`
	OwnerOrg                       string     `json:"owner_org"`
	LicenseURL                     string     `json:"license_url"`
	Title                          string     `json:"title"`
	RevisionID                     string     `json:"revision_id"`
}
