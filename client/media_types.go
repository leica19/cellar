// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": Application Media Types
//
// Command:
// $ goagen
// --design=cellar/design
// --out=$(GOPATH)/src/cellar
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// A bottle of wine (default view)
//
// Identifier: application/vnd.goa.example.bottle+json; view=default
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique bottle ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Name of wine
	Name string `form:"name" json:"name" yaml:"name" xml:"name"`
}

// Validate validates the GoaExampleBottle media type instance.
func (mt *GoaExampleBottle) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	return
}

// DecodeGoaExampleBottle decodes the GoaExampleBottle instance encoded in resp body.
func (c *Client) DecodeGoaExampleBottle(resp *http.Response) (*GoaExampleBottle, error) {
	var decoded GoaExampleBottle
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
