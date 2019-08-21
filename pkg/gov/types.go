package gov

import (
"encoding/json"
)

/// ----------- Response -----------

type Response struct {
	Help    string          `json:"help"`
	Success bool            `json:"success"`
	Result  json.RawMessage `json:"result"`
}

/// ----------- Shared -----------

//
//
//

/// ----------- Package -----------

type PackageResource struct {
	PackageID       string `json:"package_id"`
	DataStoreActive bool   `json:"datastore_active"`
	ID              string `json:"id"`
	Size            int    `json:"size"`
	FileHashSum     string `json:"file_hash_sum"`
	State           string `json:"state"`
	Hash            string `json:"hash"`
	Description     string `json:"description"`
	Format          string `json:"format"`
	LastModified    string `json:"last_modified"`
	URLType         string `json:"url_type"`
	MIMEType        string `json:"mimetype"`
	Name            string `json:"name"`
	Created         string `json:"created"`
	URL             string `json:"url"`
	Position        int    `json:"position"`
	RevisionID      string `json:"revision_id"`
}

type PackageTag struct {
	State       string `json:"state"`
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

type PackageGroup struct {
	DisplayName     string `json:"display_name"`
	Description     string `json:"description"`
	ImageDisplayURL string `json:"image_display_url"`
	Title           string `json:"title"`
	ID              string `json:"id"`
	Name            string `json:"name"`
}

type Package struct {
	LicenseTitle                   string            `json:"license_title"`
	Maintainer                     string            `json:"maintainer"`
	TagString                      string            `json:"tag_string"`
	PurposeOfCollectingInformation string            `json:"purpose_of_collecting_information"`
	Private                        bool              `json:"private"`
	MaintainerEmail                string            `json:"maintainer_email"`
	NumTags                        int               `json:"num_tags"`
	UpdateFrequency                string            `json:"update_frequency"`
	ID                             string            `json:"id"`
	MetadataCreated                string            `json:"metadata_created"`
	MetadataModified               string            `json:"metadata_modified"`
	Author                         string            `json:"author"`
	AuthorEmail                    string            `json:"author_email"`
	State                          string            `json:"state"`
	Version                        string            `json:"version"`
	IsDataPackage                  string            `json:"is_datapackage"`
	CreatorUserID                  string            `json:"creator_user_id"`
	Type                           string            `json:"type"`
	Resources                      []PackageResource `json:"resources"`
	NumResources                   int               `json:"num_resources"`
	Tags                           []PackageTag      `json:"tags"`
	Language                       string            `json:"language"`
	Groups                         []PackageGroup    `json:"groups"`
	LicenseID                      string            `json:"license_id"`
	Name                           string            `json:"name"`
	IsOpen                         bool              `json:"isopen"`
	Notes                          string            `json:"notes"`
	OwnerOrg                       string            `json:"owner_org"`
	LicenseURL                     string            `json:"license_url"`
	Title                          string            `json:"title"`
	RevisionID                     string            `json:"revision_id"`
}

/// ----------- Resource -----------

type ResourceRevision struct {
	MIMEType        string `json:"mimetype"`
	Name            string `json:"name"`
	Format          string `json:"format"`
	URL             string `json:"url"`
	FileHashSum     string `json:"file_hash_sum"`
	ResourceCreated string `json:"resource_created"`
	Size            int    `json:"size"`
}

type Resource struct {
	PackageID         string             `json:"package_id"`
	DataStoreActive   bool               `json:"datastore_active"`
	ID                string             `json:"id"`
	Size              int                `json:"size"`
	FileHashSum       string             `json:"file_hash_sum"`
	State             string             `json:"state"`
	Hash              string             `json:"hash"`
	Description       string             `json:"description"`
	Format            string             `json:"format"`
	LastModified      string             `json:"last_modified"`
	ResourceRevisions []ResourceRevision `json:"resource_revisions"`
	URLType           string             `json:"url_type"`
	MIMEType          string             `json:"mimetype"`
	Name              string             `json:"nam e"`
	Created           string             `json:"created"`
	URL               string             `json:"url"`
	Position          int                `json:"position"`
	RevisionID        string             `json:"revision_id"`
}

/// ----------- Helpers -----------
