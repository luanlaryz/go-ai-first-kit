package ui

import (
	"context"

	"github.com/charmbracelet/huh"
)

func ConfirmPersist(ctx context.Context) (bool, error) {
	value := false
	err := huh.NewConfirm().
		Title("Deseja persistir este relatorio em markdown?").
		Affirmative("sim").
		Negative("nao").
		Value(&value).
		Run()
	if ctx.Err() != nil {
		return false, ctx.Err()
	}
	return value, err
}
