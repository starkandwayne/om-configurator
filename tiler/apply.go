package tiler

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

func (c *Tiler) Apply(t Template) error {
	db, err := newTemplateRenderer(t, c.templateStore).evaluate(nil)
	if err != nil {
		return err
	}

	var deployment Deployment
	if err = yaml.Unmarshal(db, &deployment); err != nil {
		return err
	}

	if err = deployment.Validate(c.templateStore, t.Vars); err != nil {
		return err
	}

	err = c.client.ConfigureAuthentication()
	if err != nil {
		return err
	}

	err = c.configureDirector(deployment.Director, t.Vars)
	if err != nil {
		return err
	}

	for _, tile := range deployment.Tiles {
		err = c.downloadAndUploadProduct(tile.PivnetMeta)
		if err != nil {
			return err
		}

		err = c.client.StageProduct(StageProductArgs{
			ProductName:    tile.OpsmanMeta.Name,
			ProductVersion: tile.OpsmanMeta.Version,
		})
		if err != nil {
			return err
		}

		err = c.configureProduct(tile, t.Vars)
		if err != nil {
			return err
		}
	}

	err = c.client.ApplyChanges()
	if err != nil {
		return err
	}

	return nil
}

func (c *Tiler) downloadAndUploadProduct(p PivnetMeta) error {
	dir, err := ioutil.TempDir("", p.Slug)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	glob := p.Glob
	if glob == "" {
		glob = "*.pivotal"
	}

	err = c.client.DownloadProduct(DownloadProductArgs{
		OutputDirectory:      dir,
		PivnetProductSlug:    p.Slug,
		PivnetProductVersion: p.Version,
		PivnetProductGlob:    glob,
		StemcellIaas:         p.StemcellIaas,
	})
	if err != nil {
		return err
	}

	tile, err := findFileInDir(dir, "*.pivotal")
	if err != nil {
		return err
	}

	if err = c.client.UploadProduct(tile); err != nil {
		return err
	}

	stemcell, err := findFileInDir(dir, "*.tgz")
	if err != nil {
		return err
	}

	return c.client.UploadStemcell(stemcell)
}

func (c *Tiler) configureProduct(t Tile, gv map[string]interface{}) error {
	tpl, err := newTemplateRenderer(t.ToTemplate(), c.templateStore).evaluate(gv)
	if err != nil {
		return err
	}

	return c.client.ConfigureProduct(tpl)
}

func (c *Tiler) configureDirector(d Director, gv map[string]interface{}) error {
	tpl, err := newTemplateRenderer(d.ToTemplate(), c.templateStore).evaluate(gv)
	if err != nil {
		return err
	}

	return c.client.ConfigureDirector(tpl)
}

func findFileInDir(dir, glob string) (string, error) {
	files, err := filepath.Glob(filepath.Join(dir, glob))
	if err != nil {
		return "", err
	}
	if len(files) != 1 {
		return "", fmt.Errorf("no file found for %s in %s", glob, dir)
	}
	return files[0], nil
}