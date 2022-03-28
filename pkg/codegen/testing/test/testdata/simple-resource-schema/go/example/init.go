// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package example

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "bar::BarResource":
		r = &BarResource{}
	case "example::OtherResource":
		r = &OtherResource{}
	case "example::Resource":
		r = &Resource{}
	case "example::TypeUses":
		r = &TypeUses{}
	case "foo::FooResource":
		r = &FooResource{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:example" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"example",
		"",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"example",
		&pkg{version},
	)
}
