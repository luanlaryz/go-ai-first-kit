package cli

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/inventa-co/go-ai-first-kit/internal/templatecatalog"
	"github.com/spf13/cobra"
)

type templateListOptions struct {
	template  string
	asJSON    bool
	asTree    bool
	filesOnly bool
	dirsOnly  bool
}

func newTemplateListCmd(_ *globalOptions) *cobra.Command {
	opts := &templateListOptions{}
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List the contents of the embedded template",
		Example: "  gakit template list\n" +
			"  gakit template list --tree\n" +
			"  gakit template list --json\n" +
			"  gakit template list --files-only\n" +
			"  gakit template list --dirs-only",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			if opts.filesOnly && opts.dirsOnly {
				return fmt.Errorf("--files-only e --dirs-only nao podem ser usados juntos")
			}

			inv, err := templatecatalog.List(cmd.Context(), opts.template)
			if err != nil {
				return err
			}
			inv.Entries = filterEntries(inv.Entries, opts.filesOnly, opts.dirsOnly)

			out := cmd.OutOrStdout()
			switch {
			case opts.asJSON:
				data, err := json.MarshalIndent(inv, "", "  ")
				if err != nil {
					return err
				}
				fmt.Fprintln(out, string(data))
			case opts.asTree:
				renderTemplateTree(out, inv)
			default:
				renderTemplateList(out, inv)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&opts.template, "template", "default", "template embarcado a inspecionar")
	cmd.Flags().BoolVar(&opts.asJSON, "json", false, "imprime o inventario em JSON indentado")
	cmd.Flags().BoolVar(&opts.asTree, "tree", false, "imprime o inventario como arvore")
	cmd.Flags().BoolVar(&opts.filesOnly, "files-only", false, "mostra apenas arquivos")
	cmd.Flags().BoolVar(&opts.dirsOnly, "dirs-only", false, "mostra apenas diretorios")
	return cmd
}

func filterEntries(entries []templatecatalog.Entry, filesOnly, dirsOnly bool) []templatecatalog.Entry {
	if !filesOnly && !dirsOnly {
		return entries
	}
	out := make([]templatecatalog.Entry, 0, len(entries))
	for _, e := range entries {
		if filesOnly && e.Type != templatecatalog.EntryTypeFile {
			continue
		}
		if dirsOnly && e.Type != templatecatalog.EntryTypeDir {
			continue
		}
		out = append(out, e)
	}
	return out
}

// renderTemplateList prints the top-level files and directories of the template
// grouped into "Directories:" and "Files:" sections.
func renderTemplateList(out io.Writer, inv templatecatalog.Inventory) {
	fmt.Fprintf(out, "Template: %s\n", inv.Name)

	var dirs, files []string
	for _, e := range inv.Entries {
		if strings.Contains(e.Path, "/") {
			continue
		}
		if e.Type == templatecatalog.EntryTypeDir {
			dirs = append(dirs, e.Path)
		} else {
			files = append(files, e.Path)
		}
	}

	if len(dirs) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Directories:")
		for _, d := range dirs {
			fmt.Fprintf(out, "  %s/\n", d)
		}
	}
	if len(files) > 0 {
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Files:")
		for _, f := range files {
			fmt.Fprintf(out, "  %s\n", f)
		}
	}
}

// renderTemplateTree prints a simple indentation-based tree of all entries.
// Entries are reordered so that a directory's children appear immediately after
// it (path separators sort before any other character), which keeps the tree
// hierarchically consistent even when a file and a directory share a prefix.
func renderTemplateTree(out io.Writer, inv templatecatalog.Inventory) {
	fmt.Fprintln(out, inv.Name)

	entries := make([]templatecatalog.Entry, len(inv.Entries))
	copy(entries, inv.Entries)
	sort.Slice(entries, func(i, j int) bool {
		return treeSortKey(entries[i].Path) < treeSortKey(entries[j].Path)
	})

	for _, e := range entries {
		depth := strings.Count(e.Path, "/")
		indent := strings.Repeat("  ", depth+1)
		name := e.Path
		if idx := strings.LastIndex(e.Path, "/"); idx >= 0 {
			name = e.Path[idx+1:]
		}
		if e.Type == templatecatalog.EntryTypeDir {
			fmt.Fprintf(out, "%s%s/\n", indent, name)
		} else {
			fmt.Fprintf(out, "%s%s\n", indent, name)
		}
	}
}

// treeSortKey makes the path separator sort before any other character so that
// nested entries are grouped under their parent directory.
func treeSortKey(p string) string {
	return strings.ReplaceAll(p, "/", "\x00")
}
