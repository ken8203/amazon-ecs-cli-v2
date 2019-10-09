// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package template provides usage templates to render help menus.
package template

import (
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func isInGroup(cmd *cobra.Command, group string) bool {
	return cmd.Annotations["group"] == group
}

func h1(text string) string {
	var s strings.Builder
	color.New(color.Bold, color.Underline).Fprintf(&s, text)
	return s.String()
}

func h2(text string) string {
	var s strings.Builder
	color.New(color.Bold).Fprintf(&s, text)
	return s.String()
}

func code(text string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.HasPrefix(strings.TrimSpace(line), "$") {
			// code sample
			lines[i] = color.HiBlackString(line)
		}
	}
	return strings.Join(lines, "\n")
}

func mkSlice(args ...interface{}) []interface{} {
	return args
}
