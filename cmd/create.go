package cmd

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

var createCmd = &cobra.Command{}
