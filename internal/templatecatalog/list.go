package templatecatalog

import (
	"context"
	"io/fs"
	"sort"
	"strings"

	kit "github.com/inventa-co/go-ai-first-kit"
)

// List walks the embedded template identified by templateName and returns its
// inventory of files and directories sorted by path. An empty name resolves to
// the default template; any other unknown name returns an error.
func List(ctx context.Context, templateName string) (Inventory, error) {
	name, err := Resolve(templateName)
	if err != nil {
		return Inventory{}, err
	}

	inv := Inventory{Name: name, Root: kit.TemplateRoot}
	walkErr := fs.WalkDir(kit.TemplateFS, kit.TemplateRoot, func(p string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		if p == kit.TemplateRoot {
			return nil
		}

		rel := strings.TrimPrefix(p, kit.TemplateRoot+"/")
		if entry.IsDir() {
			inv.Entries = append(inv.Entries, Entry{Path: rel, Type: EntryTypeDir})
			return nil
		}

		e := Entry{Path: rel, Type: EntryTypeFile}
		if info, infoErr := entry.Info(); infoErr == nil {
			e.Size = info.Size()
		}
		inv.Entries = append(inv.Entries, e)
		return nil
	})
	if walkErr != nil {
		return Inventory{}, walkErr
	}

	sort.Slice(inv.Entries, func(i, j int) bool {
		return inv.Entries[i].Path < inv.Entries[j].Path
	})
	return inv, nil
}
